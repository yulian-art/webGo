package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

func main() {
	ctx := context.Background()
	chatTemplateDemo(ctx)
}
func chatTemplateDemo(ctx context.Context){
	cm, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("KEY"),
		BaseURL: os.Getenv("URL"),
		Model: os.Getenv("MODEL"),
	})
	if err != nil {
		log.Fatal("create model failed")
	}

	//prompt.ChatTemplate
	tpl := prompt.FromMessages(
		schema.FString,
		schema.SystemMessage(
			"你是{brand}客服。只依据已知资料回答。",
				),
			schema.MessagesPlaceholder("history", true),//包含历史信息
			schema.UserMessage(
				"资料：\n{context}\n\n问题：{question}",
			),
	)

	retrievedContext := `
	《julien售后手册》
	## 有问题包解答
	1. 有质量的问题
	2. 查过ai后的问题
	3. julien会的问题（很重要）
	`

	userQuestion := "julien是什么"
	msg, err := tpl.Format(ctx, map[string]any{
		"brand": "julien",

		"context": retrievedContext, //搜到的资料
		"question": userQuestion,

	})

	msg1, _:=cm.Generate(ctx, msg)
	fmt.Println(msg1)
}