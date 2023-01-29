# Table: openai_completion

Generate completions for a given text prompt.

Notes:
* A `prompt` or `settings -> 'prompt'` where qualifier is required for all queries.

## Examples

### Basic example with default settings

```sql
select
  completion
from
  openai_completion
where
  prompt = 'Write a tagline for an ice cream shop.';
```

### Completion with specific settings

`settings` is a JSONB object that accepts any of the [completion API request
parameters](https://beta.openai.com/docs/api-reference/completions/create).

```sql
select
  completion
from
  openai_completion
where
  settings = '{
    "model": "text-davinci-003",
    "max_tokens": 60,
    "temperature": 0.7,
    "top_p": 1.0,
    "frequency_penalty": 0.0,
    "presence_penalty": 1
  }'
  and prompt = 'A neutron star is the collapsed core of a massive supergiant star,
which had a total mass of between 10 and 25 solar masses, possibly more if the
star was especially metal-rich.[1] Neutron stars are the smallest and densest
stellar objects, excluding black holes and hypothetical white holes, quark
stars, and strange stars.[2] Neutron stars have a radius on the order of 10
kilometres (6.2 mi) and a mass of about 1.4 solar masses.[3] They result from
the supernova explosion of a massive star, combined with gravitational
collapse, that compresses the core past white dwarf star density to that of
atomic nuclei.

TL;DR

';
```

### Prompt through settings

The `prompt` column takes precedence, but you can also provide prompt text
through `settings` if easier.

```sql
select
  completion
from
  openai_completion
where
  settings = '{"prompt": "Write a tagline for an ice cream shop."}';
```