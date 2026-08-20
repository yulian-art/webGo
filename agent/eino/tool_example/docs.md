# tool与toolMode

tool是模型可选择调用的外部能力

toolNode是eino框架指定的tool执行器，无论是Graph内还是Agent中，tool的执行都要通过toolNode

工具内容较多，推荐看官方文档

调用tool需要两部：1 大模型根据tool的功能和参数需求构造调用参数（首先得有工具，其次得让大模型知道他有这些工具）2 实际调用工具

```
遇事不决写注释
1. 定义一个工具
2. 告诉大模型我有这么一个tool
3. 大模型分析决策出是否需要调用tool，以及调用什么tool， 参数是什么
4. 解析大模型回复的结果， 如果需要工具调用那就调用工具
5. 把调用结果告诉大模型
6. 大模型根据用户的问题+工具执行结果生成回答
```

接口有基础型和增强型

给出基础型工具调用demo

实现工具的方式有四种1. 直接实现接口。2. 用newTool方法，把一个函数封装成一个工具。3.用eino-ext中提供的tool。4. 用mcp协议 

```
type InvokeFunc[T, D any] func(ctx context.Context, input T) (output D, err error)

```

这里是函数定义了一个类型别名，两个泛型的类型，如果你的函数满足这样的输入和输出，那么都可以定义成一个invokeFunc，回头看我们定义的QueryOrder函数正好满足这样的条件，所以可以把它放到NewTool的第二个参数里

createTool1方法比较繁琐，注意到定义queryArg是已经打好一个详细的jsonTag，可以使用InferTool