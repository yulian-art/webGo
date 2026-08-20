package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/cloudwego/eino-ext/components/embedding/ark"
    "github.com/cloudwego/eino-ext/components/retriever/milvus2"
    "github.com/cloudwego/eino-ext/components/retriever/milvus2/search_mode"
    // 不需要 indexer 包，除非用到 search_mode 等，但我们可以用 milvus2 包内的常量
    "github.com/milvus-io/milvus/client/v2/milvusclient"
)

func main() {
    retrieverDemo(context.Background())
}

var (
    de = 2048
    ti = time.Second * 30
)

func retrieverDemo(ctx context.Context) {
    emb, err := ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
        APIKey:     os.Getenv("KEY"),
        Model:      "doubao-embedding-vision-250615",
        BaseURL:    "https://ark.cn-beijing.volces.com/api/v3",
        Dimensions: &de,
        Timeout:    &ti,
        APIType:    new(ark.APITypeMultiModal),
    })
    if err != nil {
        log.Fatal("embedder failed: ", err)
    }

    // 创建 Milvus Retriever
    r, err := milvus2.NewRetriever(ctx, &milvus2.RetrieverConfig{
        ClientConfig: &milvusclient.ClientConfig{
            Address: "localhost:19530",
        },
        Collection: "julien_demo",
        TopK:       3,
        // 使用 milvus2 包内定义的 SearchMode 常量，例如 SearchModeApproximate
        SearchMode: search_mode.NewApproximate(milvus2.COSINE), // 可能需要指定 MetricType，看实际定义
        Embedding:  emb,
        // 如果 SearchMode 需要额外参数，可以用其他方式，比如使用 WithSearchMode 选项，但这里 RetrieverConfig 可能支持
    })
    if err != nil {
        log.Fatal("retriever failed: ", err)
    }

    // 使用 Retrieve 方法，而不是转换
    docs, err := r.Retrieve(ctx, "为什么不想上学？")
    if err != nil {
        log.Fatal("docs failed: ", err)
    }

    for _, doc := range docs {
        // doc 可能实现了 Score() 方法，但注意接口可能不同，用 .Score 或类似
        fmt.Printf("document:%v\n  score:%v \n", doc, doc.Score())
    }
}