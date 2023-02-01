---
organization: Turbot
category: ["ai"]
icon_url: "/images/plugins/turbot/openai.svg"
brand_color: "#000000"
display_name: "OpenAI"
short_name: "openai"
description: "Steampipe plugin to query models, completions and more from OpenAI."
og_description: "Query OpenAI with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/openai-social-graphic.png"
---

# OpenAI + Steampipe

[OpenAI](https://openai.com) is an Artificial Intelligence research and development company that provides APIs for general models.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List instances in your OpenAI account:

```sql
select
  completion
from
  openai_completion
where
  prompt = 'Write a tagline for an ice cream shop.';
```

```
+-------------------------------------------------------+
| completion                                            |
+-------------------------------------------------------+
| 1. Cool down with a scoop of our delicious ice cream! |
| 2. Indulge your cravings with creamy goodness!        |
| 3. Sweeten your day with a scoop of our tasty treats! |
+-------------------------------------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/openai/tables)**

## Get started

### Install

Download and install the latest OpenAI plugin:

```bash
steampipe plugin install openai
```

### Configuration

Installing the latest openai plugin will create a config file (`~/.steampipe/config/openai.spc`) with a single connection named `openai`:

```hcl
connection "openai" {
  plugin = "openai"

  # Get your API key at https://beta.openai.com/account/api-keys
  api_key = "sk-CGG8G29a47ViRhvVsCGPT8BlbkFJBvFr65mZcMJWH8fayZO8"
}
```

- `api_key` - API key to authenticate requests.

Environment variables are also available as an alternate configuration method:
* `OPENAI_API_KEY`

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-openai
- Community: [Slack Channel](https://steampipe.io/community/join)
