# ChatTemplate

chatTemplate和chatModel是交互密切的两个组件，用来把一个变量的map迅速格式化成一个chatModel使用的消息列表，作用是为大模型调用准备上下文

调用format方法然后返回一个消息列表

通常会把一个系统提示词定义成一个模板，要注意模板里的变量要和代码里的变量对应上，对应不上就会报运行时错误

替换语法有三种自行查阅资料

创建模板用prompt包下的FromMessage方法快速创建一个模板

然后调用format方法去格式化一个消息列表

```
遇事不决写注释
1. 准备chatModel
2. 准备消息
 上次是直接声明一个schema.Message指针数组，然后传入generate
 缺点是：每次都是写死的，但是实际业务场景是根据输入变化的
 所以这次用模板曾姐灵活性
```

history要求必须是会话历史的格式：[]*schema.Message