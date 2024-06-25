package openai

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpenAiFile(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "openai_file",
		Description: "Files uploaded to OpenAI for fine-tuning.",
		List: &plugin.ListConfig{
			Hydrate: listFile,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getFile,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the file, e.g. davinci."},
			{Name: "file_name", Type: proto.ColumnType_STRING, Description: "Name of the file."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CreatedAt").Transform(transform.UnixToTimestamp), Description: "Timestamp of when the file was created."},
			{Name: "bytes", Type: proto.ColumnType_INT, Description: "Size of the file in bytes"},
			// Other columns
			{Name: "object", Type: proto.ColumnType_STRING, Description: "The type of the uploaded document, e.g. file."},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "Organization that owns the file, e.g. openai."},
			{Name: "purpose", Type: proto.ColumnType_STRING, Description: "The intended purpose of the uploaded documents."},
		}),
	}
}

func listFile(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("openai_file.listFile", "connection_error", err)
		return nil, err
	}
	resp, err := conn.ListFiles(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("openai_file.listFile", "query_error", err)
		return nil, err
	}
	for _, i := range resp.Files {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getFile(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("openai_file.getFile", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["id"].GetStringValue()
	item, err := conn.GetFile(ctx, id)
	if err != nil {
		plugin.Logger(ctx).Error("openai_file.getFile", "query_error", err, "id", id)
		return nil, err
	}
	return item, err
}
