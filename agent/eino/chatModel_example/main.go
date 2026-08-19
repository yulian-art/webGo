package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/schema"
)

func main() {
	ctx := context.Background()
	chatModelDemo(ctx)
}

func chatModelDemo(ctx context.Context){
	//model.BaseChatModel
	cm, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("KEY"),
		BaseURL: os.Getenv("URL"),
		Model: os.Getenv("MODEL"),
	})
	fmt.Println(cm)
	if err != nil {
		log.Fatal("client failed")
	}

	//schema.Message
	message := []*schema.Message{
		schema.SystemMessage("你是go语言技术僵尸"),
		schema.UserMessage("请用两句话介绍eino框架"),
	}

	msg , _:= cm.Generate(ctx, message)
	fmt.Println(msg)

}