package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ModelType string

var (
	ModelTypeGPT3_5TURBO      ModelType = "gpt-3.5-turbo"
	ModelTypeGPT3_5TURBO_0301 ModelType = "gpt-3.5-turbo-0301"
)

type completionRequest struct {
	Model    ModelType     `json:"model"`
	Messages []chatMessage `json:"messages"`
}

type chatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type completionResponse struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Model   ModelType `json:"model"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Choices []struct {
		Message      chatMessage `json:"message"`
		FinishReason string      `json:"finish_reason,omitempty"`
		Index        int         `json:"index,omitempty"`
	} `json:"choices"`
}

type Reply struct {
	Role    string
	Message string
}

type CompletionClient interface {
	Say(role string, message string) (*Reply, error)
}

type openAICompletionClient struct {
	apiKey string
}

func (c *openAICompletionClient) Say(role, message string) (*Reply, error) {
	url := "https://api.openai.com/v1/chat/completions"

	msg := completionRequest{
		Model: ModelTypeGPT3_5TURBO,
		Messages: []chatMessage{
			{Role: role, Content: message},
		},
	}

	rawMsg, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(rawMsg))
	req.Header.Add("Authorization", "Bearer "+c.apiKey)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	compResp := completionResponse{}
	err = json.Unmarshal(body, &compResp)
	if err != nil {
		return nil, err
	}

	// TODO
	fmt.Println(string(body))

	return &Reply{
		Message: compResp.Choices[0].Message.Content,
		Role:    compResp.Choices[0].Message.Role,
	}, nil
}

func NewClient(apiKey string) CompletionClient {
	return &openAICompletionClient{apiKey: apiKey}
}
