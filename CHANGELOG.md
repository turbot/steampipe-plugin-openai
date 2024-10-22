## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#41](https://github.com/turbot/steampipe-plugin-openai/pull/41))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#41](https://github.com/turbot/steampipe-plugin-openai/pull/41))

## v0.4.0 [2024-07-26]

_Enhancements_

- The `org_id` column has now been assigned as a connection key column across all the tables which facilitates more precise and efficient querying across multiple OpenAI connections. ([#28](https://github.com/turbot/steampipe-plugin-openai/pull/28))
- Updated the plugin to use `sashabaranov/go-openai` SDK package instead of `sashabaranov/go-gpt3`. ([#27](https://github.com/turbot/steampipe-plugin-openai/pull/27))
- The Plugin and the Steampipe Anywhere binaries are now built with the `netgo` package. ([#35](https://github.com/turbot/steampipe-plugin-openai/pull/35))
- Added the `version` flag to the plugin's Export tool. ([#65](https://github.com/turbot/steampipe-export/pull/65))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.10.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v5100-2024-04-10) that adds support for connection key columns. ([#28](https://github.com/turbot/steampipe-plugin-openai/pull/28))

## v0.3.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#21](https://github.com/turbot/steampipe-plugin-openai/pull/21))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#21](https://github.com/turbot/steampipe-plugin-openai/pull/21))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-openai/blob/main/docs/LICENSE). ([#21](https://github.com/turbot/steampipe-plugin-openai/pull/21))

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
