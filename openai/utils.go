package openai

import (
	"context"
	"errors"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*openai.Client, error) {
	conn, err := connectCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	return conn.(*openai.Client), nil
}

var connectCached = plugin.HydrateFunc(connectUncached).Memoize()

func connectUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (any, error) {

	var conn *openai.Client

	plugin.Logger(ctx).Debug("openai connectUncached")

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

	conn = openai.NewClient(apiKey)

	return conn, nil
}

func isNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "status code: 404")
}
