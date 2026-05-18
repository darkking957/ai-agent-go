package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/openai/openai-go/v3"
)

type TextAnalysis struct {
	Sentiment  string   `json:"sentiment"`
	Keywords   []string `json:"keywords`
	Summary    string   `json:"summary"`
	Confidence float64  `json:"confidence"`
}

func AnalyzeText(text string) (*TextAnalysis, error) {
	client := NewDeepSeekClient()

	systemPrompt := `你是一个文本分析器。
分析用户输入的文本，只返回以下 JSON 格式，不要有任何其他文字：
{
  "sentiment": "positive|negative|neutral",
  "keywords": ["关键词1", "关键词2"],
  "summary": "一句话总结",
  "confidence": 0.95
}
注意：只输出 JSON，不要有 markdown 代码块，不要有任何解释。`

	resp, err := client.Chat.Completions.New(
		context.Background(),
		openai.ChatCompletionNewParams{
			Model: "deepseek-chat",
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage(systemPrompt),
				openai.UserMessage(text),
			},
			Temperature: openai.Float(0.1),
		},
	)

	if err != nil {
		return nil, err
	}

	rawJSON := resp.Choices[0].Message.Content
	fmt.Printf("LLM原始输出：", rawJSON)

	var result TextAnalysis
	if err := json.Unmarshal([]byte(rawJSON), &result); err != nil {
		// LLM 偶尔会在 JSON 外面包 ```json ... ```，需要清洗
		cleaned := cleanJSONFences(rawJSON)
		if err2 := json.Unmarshal([]byte(cleaned), &result); err2 != nil {
			return nil, fmt.Errorf("parse JSON failed: %w\nraw: %s", err2, rawJSON)
		}
	}
	return &result, nil
}

// 清除 LLM 有时会加的 markdown 代码块标记
func cleanJSONFences(s string) string {
	// 简单处理：去掉 ```json 和 ``` 包裹
	if len(s) > 7 && s[:7] == "```json" {
		s = s[7:]
	}
	if len(s) > 3 && s[len(s)-3:] == "```" {
		s = s[:len(s)-3]
	}
	return strings.TrimSpace(s) // 需要 import "strings"
}
