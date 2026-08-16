package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	baseURL = "https://qpi.deepseek.con"
	model   = "deepseek-v4-pro"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model   string    `json:"model"`
	Message []Message `json:"message"`
	Stream  string    `json:"stream"`
}

type ChatResponce struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`

	Usage struct {
		PromoptTokens     int `json:"promopt_tokens"`
		CompletionsTokens int `json:"completions_tokens"`
		totalTokens       int `json:"total_tokens"`
	} `json:"usage"`
}

func main() {
	apiKey := os.Getenv("LLM_API_KEY")
	if apiKey == "" {
		log.Fatal("do not have api key")
	} 

	question := "say what is agent"
	if len(os.Args) >1 {
		question = strings.Join(os.Args[1:]," ")
	}

	payload := ChatRequest{
		Model: model,
		Message: []Message{
			{Role: "user", Content: question},
		},
		Stream: "false",
	}
	body, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		baseURL + "/chat/completeions",
		bytes.NewReader(body),
	)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content_Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 && resp.StatusCode >= 300 {
		raw , _ := io.ReadAll(io.LimitReader(resp.Body, 8<<10))
		log.Fatalf("return %s: %s", resp.Status, raw)
	}

	var result ChatResponce
	if err := json.NewDecoder(resp.Body).Decode(&result); err!=nil {
		log.Fatal(err)
	}

	fmt.Println(result.Choices[0].Message.Content)
}