# Tool 与 ToolMode

Tool 是模型可选择调用的外部能力。

ToolNode 是 eino 框架指定的 Tool 执行器，无论是在 Graph 内还是 Agent 中，Tool 的执行都要通过 ToolNode。

工具内容较多，推荐查看官方文档。

调用 Tool 需要两部分：1. 大模型根据 Tool 的功能和参数需求构造调用参数（首先得有工具，其次得让大模型知道它有这些工具）；2. 实际调用工具。

```
遇事不决写注释
1. 定义一个工具
2. 告诉大模型我有这么一个tool
3. 大模型分析并决定是否需要调用 Tool，以及调用什么 Tool、参数是什么
4. 解析大模型回复的结果；如果需要工具调用，就调用工具
5. 把调用结果告诉大模型
6. 大模型根据用户的问题+工具执行结果生成回答
```

接口有基础型和增强型

给出基础型工具调用demo

实现工具的方式有四种：1. 直接实现接口；2. 用 `NewTool` 方法把一个函数封装成工具；3. 用 eino-ext 中提供的 Tool；4. 用 MCP 协议。

```
type InvokeFunc[T, D any] func(ctx context.Context, input T) (output D, err error)

```

这里给函数定义了一个类型别名，包含两个泛型类型。如果你的函数满足这样的输入和输出，就可以定义成一个 `InvokeFunc`。回头看我们定义的 `QueryOrder` 函数正好满足这样的条件，所以可以把它放到 `NewTool` 的第二个参数里。

`CreateTool1` 方法比较繁琐。注意到定义 `queryArg` 时已经打好了详细的 JSON tag，可以使用 `InferTool`。
