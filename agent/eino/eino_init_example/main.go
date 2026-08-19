package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

func main() {
	ctx := context.Background()
	key := os.Getenv("KEY")
	modelname := os.Getenv("MODEL")
	baseURL := os.Getenv("URL")
	fmt.Printf("KEY=%s, MODEL=%s, URL=%s\n", key, modelname, baseURL)
	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		BaseURL: baseURL,
		Model:   modelname,
		APIKey:  key,
	})
	if err != nil {
		log.Fatal("create chat model failed")
	}
	chatTemplate := prompt.FromMessages(
		schema.FString,
		schema.SystemMessage("你是go语言编程助手"),
		schema.UserMessage("{question}"),
	)
	msg, err := chatTemplate.Format(ctx, map[string]any{"question": "后端组最美组长是谁"})
	if err != nil {
		log.Fatalf("format failed")
	}
	reply, err := chatModel.Generate(ctx, msg, model.WithTemperature(0.1))
	if err != nil {
		log.Fatalf("reply failed")
	}
	fmt.Println("reply: %#v\n", reply)
}
