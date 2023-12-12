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
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# OpenAI + Steampipe

[OpenAI](https://openai.com) is an Artificial Intelligence research and development company that provides APIs for general models.

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

Generate completions for a given text prompt in your OpenAI account:

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

### Credentials

| Item        | Description                                                                                                                                                                                                                                                                                 |
|-------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Credentials | OpenAI requires an [API Key](https://beta.openai.com/account/api-keys) for all requests.                                                                                                                                                                                 |
| Permissions | API Keys have the same permissions as the user who creates them, and if the user permissions change, the API key permissions also change.                                                                                                                                               |
| Radius      | Each connection represents a single OpenAI Installation.                                                                                                                                                                                                                                   |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/openai.spc`)<br />2. Credentials specified in environment variables, e.g., `OPENAI_API_KEY`. |

### Configuration

Installing the latest openai plugin will create a config file (`~/.steampipe/config/openai.spc`) with a single connection named `openai`:

```hcl
connection "openai" {
  plugin = "openai"

  # Get your API key at https://beta.openai.com/account/api-keys
  # This can also be set via the `OPENAI_API_KEY` environment variable.
  api_key = "sk-CGG8G29a47ViRhvVsCGPT8BlbkFJBvFr65mZcMJWH8fay..."
}
```

### Credentials from Environment Variables

The OpenAI plugin will use the standard OpenAI environment variables to obtain credentials **only if other arguments (`api_key`) are not specified** in the connection:

```sh
export OPENAI_API_KEY=sk-CGG8G29a47ViRhvVsCGPT8BlbkFJBvFr65mZcMJWH8fayZO8
```


