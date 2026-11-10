---
# lazysql-s9k2
title: Schema node selected as database in PostgreSQL tree
status: completed
type: bug
priority: high
created_at: 2026-11-10T09:10:42Z
updated_at: 2026-11-10T09:10:42Z
---

With PostgreSQL (and MSSQL), schema nodes in the tree carried a bare `schema` reference. GetTreeNodeData classifies single-segment references as databases, so pressing Enter on a schema set the selected database to the schema name. The SQL editor then opened a connection to that "database" and every query failed with `pq: database "<schema>" does not exist` until a table was selected.

- [x] Qualify schema node references with their database
- [x] Add NodeTypeSchema so expanding a schema keeps the parent database selected
- [x] Add regression tests and run suite

## Summary of Changes

Schema nodes and their section nodes now use `database.schema[.section]` references. GetTreeNodeData returns NodeTypeSchema for two-segment references on schema-aware drivers, and selecting one expands it while selecting the parent database. Regression tests cover both tree build paths. go test ./... passes.
