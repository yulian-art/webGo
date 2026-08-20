package main

import (
	"context"
	"fmt"
	"log"

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