# ai-agent-go

> 用 Go 从零构建 AI Agent —— 从原始 HTTP 到流式 SDK，逐步深入。

## 目标

- 理解 LLM API 的底层通信原理
- 掌握 Go 调用 Chat Completion 的多种姿势
- 逐步演进为具备工具调用能力的 AI Agent

## 技术栈

- **语言**: Go 1.22+
- **API**: DeepSeek (OpenAI 兼容)
- **SDK**: openai-go

## 项目结构
ai-agent-go/
├── 01-raw-http/        # 用 net/http 手写请求
├── 02-streaming/       # openai-go SDK + 流式输出
└── README.md

## 快速开始

```bash
export DEEPSEEK_API_KEY=your_key_here
go run ./01-raw-http/
```

## 学习笔记

- [ ] Anthropic — Building Effective Agents (前两节)
- [ ] 理解 streaming SSE 协议
- [ ] Tool Calling 原理
