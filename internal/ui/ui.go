package ui

import (
    "context"
    "fmt"
    "strings"

    "github.com/charmbracelet/bubbles/spinner"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/glamour"
    "github.com/charmbracelet/lipgloss"
    "github.com/muesli/reflow/wordwrap"
    "mycliai/internal/ai"
)

type UI struct {
    input     string
    cursor    bool
    messages  []Message
    aiClient  ai.Client
    spinner   spinner.Model
    loading   bool
    renderer  *glamour.TermRenderer
    width     int
}

type Message struct {
    Role    string
    Content string
}

func New(aiClient ai.Client) *UI {
    s := spinner.New()
    s.Spinner = spinner.Dot
    s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))

    renderer, _ := glamour.NewTermRenderer(
        glamour.WithAutoStyle(),
        glamour.WithWordWrap(0), // Will be set dynamically based on terminal width
    )

    return &UI{
        aiClient:  aiClient,
        spinner:   s,
        renderer:  renderer,
        messages:  []Message{{Role: "system", Content: "Welcome to MyCliAI!\nAsk me to write a haiku or anything else!\n"}},
        cursor:    true,
        width:     80, // Default width
    }
}

func (ui *UI) Init() tea.Cmd {
    return tea.Batch(
        ui.spinner.Tick,
        tea.EnterAltScreen,
    )
}

func (ui *UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd

    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.Type {
        case tea.KeyCtrlC:
            return ui, tea.Quit
        case tea.KeyEnter:
            if ui.input != "" && !ui.loading {
                ui.loading = true
                ui.messages = append(ui.messages, Message{Role: "user", Content: ui.input})
                ui.input = ""
                return ui, tea.Batch(
                    ui.spinner.Tick,
                    ui.getAIResponse,
                )
            }
        case tea.KeyBackspace:
            if len(ui.input) > 0 {
                ui.input = ui.input[:len(ui.input)-1]
            }
        default:
            if !ui.loading && msg.String() != "" {
                ui.input += msg.String()
            }
        }

    case tea.WindowSizeMsg:
        ui.width = msg.Width
        // Update glamour renderer with new width
        ui.renderer, _ = glamour.NewTermRenderer(
            glamour.WithAutoStyle(),
            glamour.WithWordWrap(msg.Width-2), // Leave some margin
        )

    case AIResponseMsg:
        ui.loading = false
        ui.messages = append(ui.messages, Message{Role: "assistant", Content: string(msg)})
        return ui, nil

    case spinner.TickMsg:
        if ui.loading {
            var spinCmd tea.Cmd
            ui.spinner, spinCmd = ui.spinner.Update(msg)
            return ui, spinCmd
        }
    }

    return ui, cmd
}

func (ui *UI) View() string {
    var sb strings.Builder

    // Render messages with proper wrapping
    for _, msg := range ui.messages {
        switch msg.Role {
        case "user":
            prefix := lipgloss.NewStyle().Foreground(lipgloss.Color("5")).Render("You: ")
            wrapped := wordwrap.String(msg.Content, ui.width-len([]rune(prefix)))
            sb.WriteString(prefix + wrapped + "\n")
        case "assistant":
            prefix := lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Render("AI: ")
            rendered, err := ui.renderer.Render(msg.Content)
            if err != nil {
                rendered = wordwrap.String(msg.Content, ui.width-len([]rune(prefix)))
            }
            sb.WriteString(prefix + rendered + "\n")
        default:
            wrapped := wordwrap.String(msg.Content, ui.width)
            sb.WriteString(wrapped)
        }
    }

    // Show spinner when loading
    if ui.loading {
        sb.WriteString(ui.spinner.View())
        sb.WriteString(" Thinking...\n")
    }

    // Input prompt with cursor
    prompt := "> "
    inputWidth := ui.width - len(prompt)
    wrapped := wordwrap.String(ui.input, inputWidth)
    
    sb.WriteString("\n" + prompt + wrapped)
    if ui.cursor {
        sb.WriteString("█")
    }
    
    sb.WriteString("\n\n")
    sb.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render("Press Ctrl+C to quit"))

    return sb.String()
}

// AI response handling
type AIResponseMsg string

func (ui *UI) getAIResponse() tea.Msg {
    response, err := ui.aiClient.Complete(context.Background(), []ai.Message{{
        Role:    "user",
        Content: ui.messages[len(ui.messages)-1].Content,
    }})
    
    if err != nil {
        return AIResponseMsg(fmt.Sprintf("Error: %v", err))
    }
    
    return AIResponseMsg(response)
}

func (ui *UI) Start() error {
    p := tea.NewProgram(
        ui,
        tea.WithAltScreen(),
    )
    _, err := p.Run()
    return err
}
