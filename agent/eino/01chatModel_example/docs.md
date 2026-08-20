# chatModel模型

## 前言

eino 有官方文档，虽然它槽点很多，但是信息较全。建议大致看一下概述，然后实践一下快速开始，最好写一些小项目（但是里面的顺序好乱，应该把它定位成一个 eino 的百科全书）。

如果你是开发小白，只看文档很多概念或巧思无法 get 到，非常正常，因为我也 get 不到。

## 正文

ChatModel 是个非常基础的组件，是与大模型交互的门户，以 `schema.Message` 的切片作为输入。无论下面连接哪个模型，都通过 ChatModel 与大模型进行交互，大模型返回的任何信息都会抽象成 Message 发出。

事先预置好了一些消息类型，让我们快速构造消息。

```
遇事不决写注释
1. 创建大模型的 client：`model.BaseChatModel` --> `ark.ChatModel`
2. 准备消息（schema.Message）
3. 发送消息（generate stream）
4. 打印输出
```

ChatModel 有 `Stream` 和 `Generate` 两个方法，说明它实现了 `BaseChatModel` 这个接口。`BaseChatModel` 是统一接口，`ark.ChatModel` 是下游不同厂商的具体实现。

写完函数别忘了在 `main` 里面调用……我超，运行三次一个输出都没有，我还以为环境变量又没配进来。
