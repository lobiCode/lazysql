---
# lazysql-d3qa
title: Show rows affected after editor DML
status: completed
type: bug
priority: high
created_at: 2026-11-10T09:51:10Z
updated_at: 2026-11-10T09:51:10Z
---

Running INSERT/UPDATE/DELETE from the SQL editor gave no visible confirmation. The driver message ("1 rows affected") was written to the results-info panel, but the success handler immediately switched the editor pages to the empty results grid, hiding the panel. Regressed in 4f546ab when focus handling was changed to mirror the SELECT branch.

- [x] Show the results-info panel on DML success while keeping focus on the table
- [x] Append the elapsed time to the message
- [x] Switch back to the grid when the next query starts so the running status is visible
- [x] Add regression test and run suite

## Summary of Changes

Successful editor DML now displays "<driver message> (<elapsed>)" in the results-info panel. The next query execution restores the grid page. Regression test TestEditorDMLShowsAffectedRowsPanel covers both transitions. go test ./... passes.
