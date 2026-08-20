package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudwego/eino-ext/components/document/loader/file"
	"github.com/cloudwego/eino-ext/components/document/transformer/splitter/markdown"
	"github.com/cloudwego/eino/components/document"
	"github.com/cloudwego/eino/schema"
)

func main() {
	ctx := context.Background()
	docs := loadFile(ctx, "docs.md")
	transformerDemo(ctx, docs)
}
func transformerDemo(ctx context.Context, doc []*schema.Document) []*schema.Document {
	transformerDemo, err := markdown.NewHeaderSplitter(ctx, &markdown.HeaderConfig{
		Headers: map[string]string{
			"##": "h2",
			"###" : "h3",
		},
	})
	if err != nil {
		log.Fatal("transformer failed")
	}
	transformerDocs, err := transformerDemo.Transform(ctx, doc)
	if err != nil {
		log.Fatal("transformerDocs failed")
	}
	for _, doc := range transformerDocs {
		fmt.Println(doc)
	}
	return transformerDocs
}


func loadFile(ctx context.Context, src string) []*schema.Document{
	loader, err :=file.NewFileLoader(ctx, &file.FileLoaderConfig{})
	if err != nil {
		log.Fatal("loader failed")
	}
	docs, err := loader.Load(ctx, document.Source{
		URI: src,
	})
	if err != nil {
		log.Fatal("docs failed")
	}
	for _, doc:= range docs {
		fmt.Println(doc)
	}

	return docs
}