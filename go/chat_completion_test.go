package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompletions(t *testing.T) {
	apiKey := os.Getenv("OPENAI_APIKEY")

	cli := NewClient(apiKey)

	reply, err := cli.Say("user", "show me python example of http request")

	assert.Nil(t, err)

	fmt.Println(reply)
}

func TestParseCompletionResponse(t *testing.T) {
	fixture := `{
  "id": "chatcmpl-6p9XYPYSTTRi0xEviKjjilqrWU2Ve",
  "object": "chat.completion",
  "created": 1677649420,
  "model": "gpt-3.5-turbo",
  "usage": {"prompt_tokens": 56, "completion_tokens": 31, "total_tokens": 87},
  "choices": [
    {
      "message": {
        "role": "assistant",
        "content": "The 2020 World Series was played in Arlington, Texas at the Globe Life Field, which was the new home stadium for the Texas Rangers."},
        "finish_reason": "stop",
        "index": 0
    }
  ]
}
`

	d := completionResponse{}
	err := json.Unmarshal([]byte(fixture), &d)

	assert.Nil(t, err)

	assert.Equal(t, "chatcmpl-6p9XYPYSTTRi0xEviKjjilqrWU2Ve", d.ID)
	assert.Equal(t, "chat.completion", d.Object)
	assert.Equal(t, 1677649420, d.Created)
	assert.Equal(t, ModelTypeGPT3_5TURBO, d.Model)

	assert.Equal(t, 56, d.Usage.PromptTokens)
	assert.Equal(t, 31, d.Usage.CompletionTokens)
	assert.Equal(t, 87, d.Usage.TotalTokens)

	assert.Equal(t, "stop", d.Choices[0].FinishReason)
	assert.Equal(t, 0, d.Choices[0].Index)
	assert.Equal(t, "assistant", d.Choices[0].Message.Role)
}
