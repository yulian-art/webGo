package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
)
var (
	de = 2048
	ti = time.Second * 30
)
func main() {
	embeddingDemo(context.Background())
}

func embeddingDemo(ctx context.Context){
	embedder, err := ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
		APIKey: os.Getenv("KEY"),
		Model: "doubao-embedding-vision-250615",
		BaseURL: "https://ark.cn-beijing.volces.com/api/v3",
		Dimensions: &de,
		Timeout: &ti,
		APIType: new(ark.APITypeMultiModal),
	})
	if err != nil {
		log.Fatal("embedder failed")
	}
	IDs, err := embedder.EmbedStrings(ctx, []string{"hello", "how are you"})
	if err != nil {
		log.Fatal("ids failed")
	}
	for _, id := range IDs {
		fmt.Println(id)
	}
}