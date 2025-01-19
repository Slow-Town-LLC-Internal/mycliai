package ui

import (
    "context"
    "fmt"
    "strings"
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
        if isSafetyError(err) {
            return AIResponseMsg("Error: blocked: candidate: FinishReasonSafety")
        }
        return AIResponseMsg(fmt.Sprintf("Error: %v", err))
    }
    
    return AIResponseMsg(response)
}

func isSafetyError(err error) bool {
    return err != nil && (
        strings.Contains(err.Error(), "safety") ||
        strings.Contains(err.Error(), "blocked") ||
        strings.Contains(err.Error(), "FinishReason"))
}
