package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("Missing API Key")
	}

	ctx := context.Background()

	client := gpt3.NewClient(apiKey)
	resp, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:    []string{"The first thing you should know about golang is"},
		MaxTokens: gpt3.IntPtr(30),
		Echo:      true,
		Stop:      []string{"."},
	})

	for _, data := range resp.Choices {
		fmt.Println(data.Text)
	}

}
