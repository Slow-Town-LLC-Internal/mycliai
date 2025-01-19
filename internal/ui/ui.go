package ui

import (
    "context"
    "github.com/charmbracelet/bubbletea"
    "mycliai/internal/ai"
)

type UI struct {
    input    string
    messages []string
    aiClient ai.Client
}

func New(aiClient ai.Client) *UI {
    return &UI{
        aiClient: aiClient,
        messages: []string{"Welcome to MyCliAI!\nAsk me to write a haiku or anything else!\n"},
    }
}

func (ui *UI) Init() tea.Cmd {
    return nil
}

func (ui *UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.Type {
        case tea.KeyCtrlC:
            return ui, tea.Quit
        case tea.KeyEnter:
            if ui.input != "" {
                // Add user message to history
                ui.messages = append(ui.messages, "You: "+ui.input+"\n")
                
                // Get AI response
                response, err := ui.aiClient.Complete(context.Background(), []ai.Message{{
                    Role:    "user",
                    Content: ui.input,
                }})
                
                if err != nil {
                    ui.messages = append(ui.messages, "Error: "+err.Error()+"\n")
                } else {
                    ui.messages = append(ui.messages, "AI: "+response+"\n")
                }
                
                ui.input = ""
            }
            return ui, nil
        case tea.KeyBackspace:
            if len(ui.input) > 0 {
                ui.input = ui.input[:len(ui.input)-1]
            }
        default:
            if msg.String() != "" {
                ui.input += msg.String()
            }
        }
    }
    return ui, nil
}

func (ui *UI) View() string {
    // Combine all messages
    output := ""
    for _, msg := range ui.messages {
        output += msg
    }
    
    // Add current input
    output += "\n> " + ui.input
    
    return output + "\n\nPress Ctrl+C to quit\n"
}

func (ui *UI) Start() error {
    p := tea.NewProgram(ui)
    _, err := p.Run()
    return err
}
