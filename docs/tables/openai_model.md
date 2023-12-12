---
title: "Steampipe Table: openai_model - Query OpenAI Models using SQL"
description: "Allows users to query OpenAI Models specifically details about each model including its id, name, and status, providing insights into model configurations and availability."
---

# Table: openai_model - Query OCI OpenAI Models using SQL

OpenAI Models are a set of pre-trained machine learning models that can be used for a variety of AI tasks. These models can be used for tasks such as text generation, translation, summarization, and more. The models are trained on a large amount of data and can be fine-tuned to perform specific tasks.

## Table Usage Guide

The `openai_model` table provides insights into OpenAI Models within OCI. As a data scientist or AI engineer, explore model-specific details through this table, including model ids, names, and statuses. Utilize it to uncover information about models, such as their configurations, training details, and availability status.

## Examples

### List models
Explore the various models available in your OpenAI database. This can be useful in identifying and managing the resources in your OpenAI environment.

```sql+postgres
select
  id
from
  openai_model;
```

```sql+sqlite
select
  id
from
  openai_model;
```

### Models that allow fine tuning

```sql+postgres
select
  id
from
  openai_model
where
  (permission -> 0 -> 'allow_fine_tuning')::bool;
```

```sql+sqlite
Error: The corresponding SQLite query is unavailable.
```