package openai

import (
	"context"
	"errors"
	"os"
	"strings"

	gpt3 "github.com/sashabaranov/go-gpt3"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*gpt3.Client, error) {
	conn, err := connectCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	return conn.(*gpt3.Client), nil
}

var connectCached = plugin.HydrateFunc(connectUncached).Memoize()

func connectUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (any, error) {

	var conn *gpt3.Client

	// Default to the env var settings
	apiKey := os.Getenv("OPENAI_API_KEY")

	// Prefer config settings
	openaiConfig := GetConfig(d.Connection)
	if openaiConfig.APIKey != nil {
		apiKey = *openaiConfig.APIKey
	}

	// Error if the minimum config is not set
	if apiKey == "" {
		return conn, errors.New("api_key must be configured")
	}

	conn = gpt3.NewClient(apiKey)

	return conn, nil
}

func isNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "status code: 404")
}
