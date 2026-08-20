package main

import (
	"context"
	"log"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/schema"
)

// 查询订单的业务逻辑
type OrderArgs struct{
	OrderID string `json:"order_id" jsonschema:"required" jsonschema_description:"订单号，例如ORD-00111"`
}

type Order struct {
	OrderID string `json:"order_id"`
	Info string `json:"info"`
	Status string `json:"status"`
}

func QueryOrder(ctx context.Context, args OrderArgs) (*Order, error) {
	// mock查询DB
	order := &Order{
		OrderID: args.OrderID,
		Info: "汤臣一品大平层一套（2000平）",
	}
	return order, nil
}

// 然后封装成一个agent可以使用的工具
// 基于已有的函数创建工具
func createTool1() tool.InvokableTool {
	// 自己拼接参数
	queryOrderTool := utils.NewTool(&schema.ToolInfo{
		Name: "query_order",
		Desc: "查询当前用户的订单，包含商品信息，订单状态，金额，运单还等信息",
		ParamsOneOf: schema.NewParamsOneOfByParams(
			map[string]*schema.ParameterInfo{
				"order_id": &schema.ParameterInfo{
					Type: schema.String,
					Required: true,
				},
			},
		),
	}, QueryOrder)
	return queryOrderTool
}


func createTool2() tool.InvokableTool {
	queryOrderTool, err := utils.InferTool(
		"query_order",
		"查询当前用户的订单，包含商品信息，订单状态，金额，运单还等信息",
		QueryOrder,
	)
	if err!=nil {
		log.Fatal("creatre tool failed")
	}

	return queryOrderTool
	
}