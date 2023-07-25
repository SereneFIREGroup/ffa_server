package chatGPT

import (
	"context"
	"errors"

	jujuErrors "github.com/juju/errors"
	"github.com/sashabaranov/go-openai"
	"github.com/serenefiregroup/ffa_server/pkg/config"
	"github.com/serenefiregroup/ffa_server/pkg/log"
)

func Call(prompt string) (string, error) {
	client := openai.NewClient(config.String("openai_key", ""))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	err = parseError(err)
	if err != nil {
		return "", jujuErrors.Trace(err)
	}
	return resp.Choices[0].Message.Content, nil
}

func parseError(err error) error {
	e := &openai.APIError{}
	if errors.As(err, &e) {
		switch e.HTTPStatusCode {
		case 401:
			// invalid auth or key (do not retry)
			log.Error("invalid auth or key, err: %+v", err)
		case 429:
			// rate limiting or engine overload (wait and retry)
			log.Error("rate limiting or engine overload, err: %+v", err)
		case 500:
			// openai server error (retry)
			log.Error("openai server error, err: %+v", err)
		default:
			// unhandled
			log.Error("unhandled error, err: %+v", err)
		}
		return jujuErrors.Trace(err)
	}
	return nil
}
