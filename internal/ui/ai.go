package ui

import (
    "context"
    "fmt"
    tea "github.com/charmbracelet/bubbletea"
    "mycliai/internal/ai"
)

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
