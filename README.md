![image](https://hub.steampipe.io/images/plugins/turbot/openai-social-graphic.png)

# OpenAI Plugin for Steampipe

Use SQL to query models, completions and more from OpenAI.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/openai)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/openai/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-openai/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install openai
```

Configure your API key in `~/.steampipe/config/openai.spc`:

```hcl
connection "openai" {
  plugin  = "openai"
  api_key = "sk-CGG8G29a47ViRhvVsCGPT8BlbkFJBvFr65mZcMJWH8fayZO8"
}
```

Run steampipe:

```shell
steampipe query
```

Run a query:

```sql
select
  completion
from
  openai_completion
where
  prompt = 'Write a tagline for an ice cream shop.';
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-openai.git
cd steampipe-plugin-openai
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/openai.spc
```

Try it!

```
steampipe query
> .inspect openai
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-openai/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [OpenAI Plugin](https://github.com/turbot/steampipe-plugin-openai/labels/help%20wanted)
