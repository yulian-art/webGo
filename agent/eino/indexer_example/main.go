package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino-ext/components/indexer/milvus2"
	"github.com/cloudwego/eino/schema"
	"github.com/milvus-io/milvus/client/v2/milvusclient"
)

func main() {
	indexerDemo(context.Background())
}

var (
	de = 2048
	ti = time.Second * 30
)

func indexerDemo(ctx context.Context) {
	embedder, err := ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
		APIKey:     os.Getenv("KEY"),
		Model:      "doubao-embedding-vision-250615",
		BaseURL:    "https://ark.cn-beijing.volces.com/api/v3",
		Dimensions: &de,
		Timeout:    &ti,
		APIType:    new(ark.APITypeMultiModal),
	})
	if err != nil {
		log.Fatal("embedder failed")
	}

	indexer, err := milvus2.NewIndexer(ctx, &milvus2.IndexerConfig{
		ClientConfig: &milvusclient.ClientConfig{
			Address: "localhost:19530",
		},
		Collection: "julien_demo",
		Vector: &milvus2.VectorConfig{
			Dimension:    int64(de),
			MetricType:   milvus2.COSINE,
			IndexBuilder: milvus2.NewHNSWIndexBuilder().WithM(16).WithEfConstruction(200),
		},

		Embedding: embedder,
	})

	docs := []*schema.Document{
		{
			ID:      "doc1",
			Content: "EINO is a framework for building AI applications",
		},
		{
			ID:      "doc2",
			Content: "好累啊好累啊",
		},
		{
			ID:      "doc3",
			Content: "不想上学，因为不想见到神人",
		},
		{
			ID:      "doc4",
			Content: "我超威，那些神人有客观看问题的角度吗，阅读理解及格了吗",
		},
	}
	ids, err := indexer.Store(ctx, docs)
	if err != nil {
		log.Fatal("ids failed")
	}
	fmt.Println(ids)
}
