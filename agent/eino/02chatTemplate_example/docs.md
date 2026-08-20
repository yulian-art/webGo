# ChatTemplate

ChatTemplate 和 ChatModel 是交互密切的两个组件，用来把一个变量的 map 迅速格式化成 ChatModel 使用的消息列表，作用是为大模型调用准备上下文。

调用 `Format` 方法，然后返回一个消息列表。

通常会把一个系统提示词定义成一个模板。要注意模板里的变量要和代码里的变量对应上，对应不上就会报运行时错误。

替换语法有三种，请自行查阅资料。

创建模板可以用 `prompt` 包下的 `FromMessages` 方法快速创建。

然后调用 `Format` 方法格式化一个消息列表。

```
遇事不决写注释
1. 准备 ChatModel
2. 准备消息
   上次是直接声明一个 `schema.Message` 指针数组，然后传入 `Generate`。
   缺点是：每次都是写死的，但是实际业务场景会根据输入变化。
   所以这次用模板增加灵活性。
3. ……
```

`history` 要求必须是会话历史的格式：`[]*schema.Message`。

后期提示：在编排场景下通常把 ChatTemplate 作为 ChatModel 的前置节点。

我们通常在调用大模型之前把上下文准备好。
