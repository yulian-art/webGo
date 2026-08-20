package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

func main() {
	mockToolDemo(context.Background())
}

func mockToolDemo(ctx context.Context){
	// 告诉大模型
	queryOrderTool := createTool2()

	//创建toolsNode
	ToolsNode, err := compose.NewToolNode(
		ctx,
		&compose.ToolsNodeConfig{
			Tools: []tool.BaseTool{
				queryOrderTool,
			},
			ExecuteSequentially: true,//如果返回多个工具，必须有一条依赖关系链式调用，如果是false则是并行调用
		},
	)
	if err != nil {
		log.Fatal("ToolNode failed")
	}

	// 调用
	// mock大模型输出
	input := &schema.Message{
		Role: schema.Assistant,
		ToolCalls: []schema.ToolCall{
			{
				Function: schema.FunctionCall{
					Name: "query_order",
					Arguments: `{"order_id": "lyx-1234556"}`,
				},
			},
		},
	}
	toolMessages, err := ToolsNode.Invoke(ctx, input)
	if err != nil {
		log.Fatal("mock failed")
	}

	for _, msg := range toolMessages{
		fmt.Println(msg)
	}
}

func ranTool(ctx context.Context) {
	chatModel, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("KEY"),
		Model: os.Getenv("MODEL"),
		BaseURL: os.Getenv("URL"),
	})
	if err != nil {
		log.Fatal("create client failed")
	}
	queryOrderTool := createTool2()

	cma, err := adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{
		Name: "order_assistant",
		Description: "可以查询订单的智能助手",
		Instruction: "你是一个智能助手，用户查询订单时必须调用query_order工具，并且根据工具返回结果",
		Model: chatModel,
		ToolsConfig: adk.ToolsConfig{
			ToolsNodeConfig: compose.ToolsNodeConfig{
				Tools: []tool.BaseTool{
					queryOrderTool,
				},
			},
		},
		MaxIterations: 3,
	})

	if err != nil {
		log.Fatal("create chat model agent failed")
	}
	// 创建运行器
	runner := adk.NewRunner(ctx, adk.RunnerConfig{
		Agent: cma,
		EnableStreaming: true,
	})
	// 运行
	events:= runner.Run(
		ctx,
		[]*schema.Message{
			schema.UserMessage("帮我查一下我刚才下单的东西"),
		},
	)
	//遍历流式事件
	for{
		event, ok := events.Next()
		if !ok {
			break
		}
		if event == nil {
			fmt.Println("event == nil")
			return
		}
		if event.Err != nil {
			fmt.Println(event.Err)
			return
		}
		if event.Output != nil && event.Output.MessageOutput != nil {
			if err := printMessage(event.Output.MessageOutput); err != nil {
				fmt.Println("print message failed:", err)
				return
			}
		}
	}
}


func printMessage(outPut *adk.MessageVariant) error {
	// 流式模式下会先消费整个流，非流式模式下直接返回完整消息
	message, err := outPut.GetMessage()
	if err != nil {
		return err
	}
	if message == nil {
		return nil
	}

	switch outPut.Role {
	case schema.Assistant:
		// 大模型生成的回答
		if message.Content != "" {
			fmt.Printf("assistant: %s\n", message.Content)
		}
	case schema.Tool:
		// 工具执行结果
		if message.Content != "" {
			fmt.Printf("tool(%s): %s\n", outPut.ToolName, message.Content)
		}
	}
	return nil
}