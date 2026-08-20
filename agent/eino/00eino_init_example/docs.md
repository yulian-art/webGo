# eino介绍

## 框架

### Schema

定义大模型的一些 message。因为要接入各种厂商，每个厂商都有一套自己的消息结果，所以统一抽象出来。

```
message document streamBuilder streamWriter ToolInfo
```



### Callbacks

加入一些控制流程，也是偏底层的设计。

```
handler inject trigger
```

### components

组件，封装一些原子能力。

### compose

做一些编排。

### adk

agent 开发的一个工具包，开箱即用。

## 框架内部

整体定义一套规范，基于这个规范去开发。

写一个最简单的大模型调用 demo。

```
遇事不决先写注释
要调用大模型，先要经过哪几步？
1. 创建大模型的 client（通过 API key 和 base URL 连接具体的大模型厂商）【ChatModel】
2. 准备要发送的消息【msg】
3. 发送请求【reply】
4. 接收消息（解析消息）【fmt】
```

这里注意，最好把 key 配置进环境变量 `.env` 里面，然后：

```
source /xxx/xxx/.env
验证输出
echo $KEY
echo $URL
echo $MODEL
```

Windows 下可以设一个临时变量：

```
$env:URL = "https://ark.cn-beijing.volces.com/api/v3"
$env:KEY = "ark-xxx"
$env:MODEL = "doubao-seed-2-0-mini-260428"
```

在写代码的时候只需记住首字母（push）。

剩下的靠感觉，根据插件提示往下写，而不是死记硬背。不确定要填的类型时，按 Ctrl 并右键查看源码，一直追溯到它最终的类型。

定位最终类型后一般继续往下读，下方会定义此变量的使用方法，一般是常量或者一个返回该变量的方法。使用它们时注意所在的包名即可。

## 扩展

```
接收信息可以用流式输出，写好后可以积极交流。
```









