# Table: openai_model

List all models available in the OpenAI account.

## Examples

### List models

```sql
select
  id
from
  openai_model
```

### Models that allow fine tuning

```
select
  id
from
  openai_model
where
  (permission -> 0 -> 'allow_fine_tuning')::bool
```
