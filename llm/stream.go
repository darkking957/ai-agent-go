package llm

import (
	"context"
	"fmt"
	"os"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func NewDeepSeekClient() openai.Client {
	return openai.NewClient(
		option.WithAPIKey(os.Getenv("DEEPSEEK_API_KEY")),
		option.WithBaseURL("https://api.deepseek.com/v1/"),
	)
}

func ChatStream(userMsg string) error {
	client := NewDeepSeekClient()

	stream := client.Chat.Completions.NewStreaming(
		context.Background(),
		openai.ChatCompletionNewParams{
			Model: "deepseek-chat",
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage("你是一个有帮助的助手。"),
				openai.UserMessage(userMsg),
			},
		},
	)
	defer stream.Close()

	fmt.Print("LLM 流式输出：")

	for stream.Next() {
		chunk := stream.Current()
		if len(chunk.Choices) > 0 {
			fmt.Print(chunk.Choices[0].Delta.Content)
		}
	}
	fmt.Println()

	if err := stream.Err(); err != nil {
		return fmt.Errorf("stream Error:%w", err)
	}

	return nil
}
