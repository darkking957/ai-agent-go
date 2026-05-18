## 结构化 JSON 输出
- 关键：System Prompt(Go语言的text/template) 明确 schema + 低温度 (0.1)
- 坑：LLM 偶尔包 ```json，需要清洗
- 用途：所有需要程序解析 LLM 输出的场景（Tool calling 的前身）

## Temperature 实验结论
- 0.0-0.2：JSON/代码/事实型输出
- 0.7：通用对话
- 1.0+：创意写作/头脑风暴

## Few-shot 结论
- 例子让 LLM 更理解输出格式和边界案例
- 3-5 个例子通常比 1 个好很多
- 例子要覆盖边界情况，不只是简单正例