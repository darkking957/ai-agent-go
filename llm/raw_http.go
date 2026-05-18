package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// 请求体结构（和 OpenAI 规范完全一致）
type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"` // "system" / "user" / "assistant"
	Content string `json:"content"`
}

// 响应体结构（只取我们需要的字段）
type ChatResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func ChatRaw(userMsg string) (string, error) {
	apiKey := os.Getenv("DEEPSEEK_API_KEY")
	url := "https://api.deepseek.com/v1/chat/completions"

	// 1. 构造请求体
	reqBody := ChatRequest{
		Model: "deepseek-chat",
		Messages: []Message{
			{Role: "system", Content: "你是一个有帮助的助手。"},
			{Role: "user", Content: userMsg},
		},
	}

	// 2. 序列化为 JSON
	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("marshal error: %w", err)
	}

	// 3. 构造 HTTP 请求
	req, err := http.NewRequest("POST", url, bytes.NewReader(bodyBytes))
	if err != nil {
		return "", fmt.Errorf("create request error: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// 4. 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("http error: %w", err)
	}
	defer resp.Body.Close()

	// 5. 读取响应
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read body error: %w", err)
	}

	// 6. 解析 JSON
	var chatResp ChatResponse
	if err := json.Unmarshal(respBytes, &chatResp); err != nil {
		return "", fmt.Errorf("unmarshal error: %w, body: %s", err, respBytes)
	}

	if len(chatResp.Choices) == 0 {
		return "", fmt.Errorf("no choices returned, raw: %s", respBytes)
	}

	return chatResp.Choices[0].Message.Content, nil
}
