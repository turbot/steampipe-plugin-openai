# Table: openai_file

List all files uploaded to the OpenAI account for fine-tuning.

## Examples

### List files

```sql
select
  id,
  file_name,
  bytes
from
  openai_file;
```
