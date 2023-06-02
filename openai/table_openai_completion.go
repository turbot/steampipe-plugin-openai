package openai

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpenAiCompletion(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openai_completion",
		Description: "Completions available in OpenAI.",
		List: &plugin.ListConfig{
			Hydrate: listCompletion,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "prompt", Require: plugin.Required},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "completion", Type: proto.ColumnType_STRING, Description: "Completion for a given text prompt."},
			{Name: "prompt", Type: proto.ColumnType_STRING, Transform: transform.FromQual("prompt"), Description: "The prompt to generate completions for, encoded as a string."},
		},
	}
}

type CompletionRequestQual struct {
	Model  *string `json:"model"`
	Prompt *string `json:"prompt,omitempty"`
}

type CompletionRow struct {
	Completion string
	Prompt     string
}

func listCompletion(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("openai_completion.listCompletion", "connection_error", err)
		return nil, err
	}
	plugin.Logger(ctx).Debug("openai_completion.listCompletion", "conn", conn)

	prompt := d.EqualsQualString("prompt")
	plugin.Logger(ctx).Debug("openai_completion.listCompletion", "prompt", prompt)

	resp, err := conn.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: d.EqualsQualString("prompt"),
				},
			},
		},
	)

	if err != nil {
		plugin.Logger(ctx).Error("openai_completion.listCompletion", "completion_error", err)
		return nil, err
	}

	plugin.Logger(ctx).Debug("openai_completion.listCompletion", "completion_response", resp)

	row := CompletionRow{resp.Choices[0].Message.Content, prompt}

	plugin.Logger(ctx).Debug("openai_completion.listCompletion", "row", row)

	d.StreamListItem(ctx, row)

	return nil, nil
}
