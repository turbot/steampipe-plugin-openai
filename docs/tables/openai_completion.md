---
title: "Steampipe Table: openai_completion - Query OpenAI Completions using SQL"
description: "Allows users to query Completions in OpenAI, specifically the generated text and the associated model details."
---

# Table: openai_completion - Query OpenAI Completions using SQL

OpenAI Completions is a service within the OpenAI API that allows you to generate text. It provides a way to interact with the API, enabling users to input a prompt and receive a generated text completion in response. OpenAI Completions is typically used for tasks such as drafting emails, writing code, answering questions, creating conversational agents, tutoring, translating languages, simulating characters for video games, and more.

## Table Usage Guide

The `openai_completion` table provides insights into text completions within OpenAI. As a data scientist or AI developer, explore completion-specific details through this table, including generated text, model used, and other related information. Utilize it to understand the performance of different models, analyze the generated text, and gain insights into the behavior of the OpenAI API.

**Important Notes**

- A `prompt` or `settings -> 'prompt'` where qualifier is required for all queries.

## Examples

### Basic example with default settings
Discover the segments that yield completion results for a specific prompt in an OpenAI model. This is particularly useful for generating creative content, such as taglines for businesses.

```sql+postgres
select
  completion
from
  openai_completion
where
  prompt = 'Write a tagline for an ice cream shop.';
```

```sql+sqlite
select
  completion
from
  openai_completion
where
  prompt = 'Write a tagline for an ice cream shop.';
```

### Completion with specific settings
Explore the completion of a specific task within the OpenAI platform by specifying certain settings. This can be useful to understand the output of a particular model under defined parameters, providing insights into its performance and usefulness for specific tasks.
`settings` is a JSONB object that accepts any of the [completion API request
parameters](https://beta.openai.com/docs/api-reference/completions/create).


```sql+postgres
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

```sql+sqlite
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
Analyze the settings to understand the completion status of specific tasks in OpenAI. This is particularly useful when you want to monitor the progress of tasks such as writing a tagline for an ice cream shop.
The `prompt` column takes precedence, but you can also provide prompt text
through `settings` if easier.


```sql+postgres
select
  completion
from
  openai_completion
where
  settings = '{"prompt": "Write a tagline for an ice cream shop."}';
```

```sql+sqlite
select
  completion
from
  openai_completion
where
  settings = '{"prompt": "Write a tagline for an ice cream shop."}';
```

### Code analysis with change of model and a stop sequence
Explore the quality of code snippets by analyzing their completion status. This query is useful to understand the efficiency of different code models and identify potential areas of improvement.

```sql+postgres
select
  completion
from
  openai_completion
where
  settings = '{
    "model": "code-davinci-002",
    "max_tokens": 64,
    "temperature": 0,
    "top_p": 1.0,
    "frequency_penalty": 0.0,
    "presence_penalty": 1,
    "stop": ["\"\"\""]
  }'
  and prompt = '
class Log:
    def __init__(self, path):
        dirname = os.path.dirname(path)
        os.makedirs(dirname, exist_ok=True)
        f = open(path, "a+")

        # Check that the file is newline-terminated
        size = os.path.getsize(path)
        if size > 0:
            f.seek(size - 1)
            end = f.read(1)
            if end != "\n":
                f.write("\n")
        self.f = f
        self.path = path

    def log(self, event):
        event["_event_id"] = str(uuid.uuid4())
        json.dump(event, self.f)
        self.f.write("\n")

    def state(self):
        state = {"complete": set(), "last": None}
        for line in open(self.path):
            event = json.loads(line)
            if event["type"] == "submit" and event["success"]:
                state["complete"].add(event["id"])
                state["last"] = event
        return state

"""
Here''s what the above class is doing:
1.';
```

```sql+sqlite
select
  completion
from
  openai_completion
where
  settings = '{
    "model": "code-davinci-002",
    "max_tokens": 64,
    "temperature": 0,
    "top_p": 1.0,
    "frequency_penalty": 0.0,
    "presence_penalty": 1,
    "stop": ["\"\"\""]
  }'
  and prompt = '
class Log:
    def __init__(self, path):
        dirname = os.path.dirname(path)
        os.makedirs(dirname, exist_ok=True)
        f = open(path, "a+")

        # Check that the file is newline-terminated
        size = os.path.getsize(path)
        if size > 0:
            f.seek(size - 1)
            end = f.read(1)
            if end != "\n":
                f.write("\n")
        self.f = f
        self.path = path

    def log(self, event):
        event["_event_id"] = str(uuid.uuid4())
        json.dump(event, self.f)
        self.f.write("\n")

    def state(self):
        state = {"complete": set(), "last": None}
        for line in open(self.path):
            event = json.loads(line)
            if event["type"] == "submit" and event["success"]:
                state["complete"].add(event["id"])
                state["last"] = event
        return state

"""
Here''s what the above class is doing:
1.';
```