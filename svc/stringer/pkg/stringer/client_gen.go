package stringer

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

// Client ...
type Client struct {
	svc           *lambda.Lambda
	stage         string
	RetryCount    int
	RetryDuration time.Duration
}

// New Client
func New(stage string) (*Client, error) {
	return nil, errors.New("not implemented")
}

// Count calculates number of characters in string. For empty string returns an error.
// FIXME: add rate limiting, timeouts, ...
func (client *Client) Count(ctx context.Context, input string) (int, error) {

	// marshal data
	payload, err := json.Marshal(input)
	if err != nil {
		return 0, err
	}

	req := client.svc.InvokeRequest(&lambda.InvokeInput{
		FunctionName: aws.String("stringer-count-dev"), // service-function-stage.
		Payload:      payload,
	})

	req.SetContext(ctx)
	req.Retryable = aws.Bool(true)
	req.RetryCount = client.RetryCount
	req.RetryDelay = client.RetryDuration

	output, err := req.Send()
	if err != nil {
		return 0, err
	}

	// FIXME: use function error
	// output.FunctionError

	var result int
	if err = json.Unmarshal(output.Payload, &result); err != nil {
		return 0, err
	}

	return result, nil
}
