package ai

import (
    "context"

    "github.com/google/generative-ai-go/genai"
    "google.golang.org/api/option"
)

type GeminiClient struct {
    client *genai.Client
    model  *genai.GenerativeModel
}

func NewGeminiClient(apiKey string) *GeminiClient {
    ctx := context.Background()
    client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
    if err != nil {
        return nil
    }

    return &GeminiClient{
        client: client,
        model:  client.GenerativeModel("gemini-pro"),
    }
}

func (c *GeminiClient) Complete(ctx context.Context, messages []Message) (string, error) {
    chat := c.model.StartChat()
    
    // For now, just use the last message for the basic functionality
    lastMsg := messages[len(messages)-1]
    response, err := chat.SendMessage(ctx, genai.Text(lastMsg.Content))
    if err != nil {
        return "", err
    }

    // Extract the text content from the response
    if len(response.Candidates) > 0 && len(response.Candidates[0].Content.Parts) > 0 {
        if textPart, ok := response.Candidates[0].Content.Parts[0].(genai.Text); ok {
            return string(textPart), nil
        }
    }
    
    return "", nil
}

func (c *GeminiClient) Close() error {
    if c.client != nil {
        c.client.Close()
    }
    return nil
}
