# chatModel模型

## 前言

eino有官方文档，虽然它槽点很多但是信息较全，建议大致看一下概述，然后实践一下快速开始，最好写一些小项目（但是里面的顺序好乱，应该把它定位成一个eino的百科全书）

如果你是开发小白，只看文档很多概念或巧思无法get到，非常正常，因为我也get不到

## 正文

chatModel是个非常基础的组件，与大模型进行交互的门户，以schema.Message的切片作为输入，无论下面连接哪个模型，都通过chatModel与大模型进行交互，大模型返回的任何信息抽象成Message发出

事先预制好了一些消息类型让我们去快速构造消息

```
遇事不决写注释
1. 创建大模型的client model.BaseChatModel --> ark.ChatModel
2. 准备消息（schema.Message）
3. 发送消息（generate stream）
4. 打印输出
```

chatModel有stream和generate两个方法，说明它实现了baseChatModel这个接口。baseChatModel是统一接口，ark.chatModel是下游不同厂商具体的实现

写完函数别忘了在main里面调用。。。我超威运行三次一个输出都没有，我还以为环境变量又没配进来