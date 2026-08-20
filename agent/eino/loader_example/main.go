package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudwego/eino-ext/components/document/loader/file"
	urlloader "github.com/cloudwego/eino-ext/components/document/loader/url"
	"github.com/cloudwego/eino/components/document"
	
)

func main() {
	//loadFile(context.Background(), "docs.md")
	loadURL(context.Background(),"https://www.cloudwego.io/zh/docs/eino/")
}

func loadFile(ctx context.Context, src string) {
	loader, err := file.NewFileLoader(ctx, &file.FileLoaderConfig{})

	if err != nil {
		log.Fatal("loader failed")
	}
	docs, err := loader.Load(ctx, document.Source{
		URI: src,
	})
	if err != nil {
		log.Fatal("load failed")
	}
	for _, doc := range docs {
		fmt.Println(doc)
	}
}

func loadURL(ctx context.Context, src string) {
	loader, err := urlloader.NewLoader(ctx, &urlloader.LoaderConfig{})

	if err != nil {
		log.Fatal("loader failed")
	}
	docs, err := loader.Load(ctx, document.Source{
		URI: src,
	})
	if err != nil {
		log.Fatal("load failed")
	}
	for _, doc := range docs {
		fmt.Println(doc)
	}
}
