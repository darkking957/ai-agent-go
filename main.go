package main

import (
	"fmt"
	"log"

	"github.com/darkking957/ai-agent-go/llm"
)

func main() {
	// //part.1
	// reply, err := llm.ChatRaw("用一句话解释什么是大语言模型")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("LLM 回复:", reply)
	// //part.2
	// if err := llm.ChatStream("写一首关于 Go 语言的短诗"); err != nil {
	// 	log.Fatal(err)
	// }
	//part.3
	analysis, err := llm.AnalyzeText("今天天气真好，心情愉快，适合出去踏青！")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("情感: %s\n关键词: %v\n摘要: %s\n置信度: %.2f\n",
		analysis.Sentiment, analysis.Keywords, analysis.Summary, analysis.Confidence)
}
