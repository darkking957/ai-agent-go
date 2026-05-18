package main

import (
	"fmt"
	"log"

	"github.com/darkking957/ai-agent-go/llm"
)

func main() {
	reply, err := llm.ChatRaw("用一句话解释什么是大语言模型")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("LLM 回复:", reply)
	if err := llm.ChatStream("写一首关于 Go 语言的短诗"); err != nil {
		log.Fatal(err)
	}
}
