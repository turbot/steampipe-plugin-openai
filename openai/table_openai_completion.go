package openai

import (
	"context"
	"encoding/json"

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
				{Name: "prompt", Require: plugin.Optional},
				{Name: "settings", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "completion", Type: proto.ColumnType_STRING, Transform: transform.FromField("Completion"), Description: "Completions for a given text prompt."},
			{Name: "index", Type: proto.ColumnType_INT, Transform: transform.FromField("Index"), Description: "The index location of the result."},
			{Name: "finish_reason", Type: proto.ColumnType_STRING, Description: "The reason for the execution to be terminated."},
			{Name: "log_probs", Type: proto.ColumnType_JSON, Description: "Include the log probabilities on the logprobs most likely tokens, as well the chosen tokens."},
			{Name: "prompt", Type: proto.ColumnType_STRING, Transform: transform.FromQual("prompt"), Description: "The prompt to generate completions for, encoded as a string."},
			{Name: "settings", Type: proto.ColumnType_JSON, Transform: transform.FromQual("settings"), Description: "Settings is a JSONB object that accepts any of the completion API request parameters."},
		},
	}
}

type CompletionRequestQual struct {
	Model            *string        `json:"model"`
	Prompt           *string        `json:"prompt,omitempty"`
	Suffix           *string        `json:"suffix,omitempty"`
	MaxTokens        *int           `json:"max_tokens,omitempty"`
	Temperature      *float32       `json:"temperature,omitempty"`
	TopP             *float32       `json:"top_p,omitempty"`
	N                *int           `json:"n,omitempty"`
	Stream           *bool          `json:"stream,omitempty"`
	LogProbs         *int           `json:"logprobs,omitempty"`
	Stop             []string       `json:"stop,omitempty"`
	PresencePenalty  *float32       `json:"presence_penalty,omitempty"`
	FrequencyPenalty *float32       `json:"frequency_penalty,omitempty"`
	BestOf           *int           `json:"best_of,omitempty"`
	LogitBias        map[string]int `json:"logit_bias,omitempty"`
	User             *string        `json:"user,omitempty"`
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

	// these are the defaults before reading settings
	cr := openai.ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: d.EqualsQualString("prompt"),
			},
		},
		Temperature:      0.7,
		MaxTokens:        256,
		Stop:             []string{},
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}

	settingsString := d.EqualsQuals["settings"].GetJsonbValue()
	if settingsString != "" {
		// Overwrite any settings provided in the settings qual. If a field
		// is not passed in the settings, then default to the settings above.
		var crQual CompletionRequestQual
		err := json.Unmarshal([]byte(settingsString), &crQual)
		if err != nil {
			plugin.Logger(ctx).Error("openai_completion.listCompletion", "unmarshal_error", err)
			return nil, err
		}
		if crQual.Model != nil {
			cr.Model = *crQual.Model
		}
		if crQual.MaxTokens != nil {
			cr.MaxTokens = *crQual.MaxTokens
		}
		if crQual.Temperature != nil {
			cr.Temperature = *crQual.Temperature
		}
		if crQual.TopP != nil {
			cr.TopP = *crQual.TopP
		}
		if crQual.N != nil {
			cr.N = *crQual.N
		}
		if crQual.Stream != nil {
			cr.Stream = *crQual.Stream
		}
		if crQual.LogProbs != nil {
			cr.LogProbs = *crQual.LogProbs != 0
		}
		if crQual.Stop != nil {
			cr.Stop = crQual.Stop
		}
		if crQual.PresencePenalty != nil {
			cr.PresencePenalty = *crQual.PresencePenalty
		}
		if crQual.FrequencyPenalty != nil {
			cr.FrequencyPenalty = *crQual.FrequencyPenalty
		}
		if crQual.LogitBias != nil {
			cr.LogitBias = crQual.LogitBias
		}
		if crQual.User != nil {
			cr.User = *crQual.User
		}
	}

	if d.EqualsQuals["prompt"] == nil {
		// No prompt, so return zero rows
		return nil, nil
	}

	prompt := d.EqualsQualString("prompt")
	plugin.Logger(ctx).Debug("openai_completion.listCompletion", "prompt", prompt)

	resp, err := conn.CreateChatCompletion(ctx, cr)

	if err != nil {
		plugin.Logger(ctx).Error("openai_completion.listCompletion", "prompt", cr, "completion_error", err)
		return nil, err
	}
	plugin.Logger(ctx).Debug("openai_completion.listCompletion", "completion_response", resp)

	for _, i := range resp.Choices {
		row := CompletionRow{i.Message.Content, prompt}
		plugin.Logger(ctx).Debug("openai_completion.listCompletion", "row", row)
		d.StreamListItem(ctx, row)
	}

	return nil, nil
}
