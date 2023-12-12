## v0.3.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#](https://github.com/turbot/steampipe-plugin-openai/pull/))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#](https://github.com/turbot/steampipe-plugin-openai/pull/))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-openai/blob/main/docs/LICENSE). ([#](https://github.com/turbot/steampipe-plugin-openai/pull/))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#21](https://github.com/turbot/steampipe-plugin-openai/pull/21))

## v0.2.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#15](https://github.com/turbot/steampipe-plugin-openai/pull/15))

## v0.2.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#13](https://github.com/turbot/steampipe-plugin-openai/pull/13))
- Recompiled plugin with Go version `1.21`. ([#13](https://github.com/turbot/steampipe-plugin-openai/pull/13))

## v0.1.0 [2023-04-11]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which adds go-getter support to dynamic tables. ([#7](https://github.com/turbot/steampipe-plugin-openai/pull/7))

## v0.0.1 [2023-02-01]

_What's new?_

- New tables added
  - [openai_completion](https://hub.steampipe.io/plugins/turbot/openai/tables/openai_completion)
  - [openai_file](https://hub.steampipe.io/plugins/turbot/openai/tables/openai_file)
  - [openai_model](https://hub.steampipe.io/plugins/turbot/openai/tables/openai_model)
