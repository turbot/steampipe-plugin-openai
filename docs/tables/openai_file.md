---
title: "Steampipe Table: openai_file - Query OpenAI Files using SQL"
description: "Allows users to query Files in OpenAI, specifically the file ID, filename, purpose and status, providing insights into the file data used for training models."
---

# Table: openai_file - Query OpenAI Files using SQL

OpenAI File is a resource within OpenAI that represents the data files used for training models. It provides detailed information about each file including its ID, name, purpose, and status. OpenAI File is vital for understanding the data used in model training and for keeping track of file statuses.

## Table Usage Guide

The `openai_file` table provides insights into the data files within OpenAI used for training models. As a data scientist or AI engineer, explore file-specific details through this table, including file ID, filename, purpose, and status. Utilize it to understand the data used in model training, monitor the status of files, and manage your resources effectively.

## Examples

### List files
Explore which files are taking up the most space in your system. This is particularly useful for managing storage and identifying large files that may no longer be necessary.

```sql+postgres
select
  id,
  file_name,
  bytes
from
  openai_file;
```

```sql+sqlite
select
  id,
  file_name,
  bytes
from
  openai_file;
```