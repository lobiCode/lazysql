---
# lazysql-u8vn
title: Remove in-app updater
status: completed
type: task
priority: normal
created_at: 2026-11-10T09:46:58Z
updated_at: 2026-11-10T09:46:58Z
---

The updater checked and installed releases from the upstream jorgerojas26/lazysql repository. This fork does not publish releases through it and should make no outbound calls on startup, so remove the feature instead of repointing it.

- [x] Delete internal/updater and components/updates.go with tests
- [x] Remove the CheckForUpdates command, F10 binding, OnUpdateRequest hook and DisableUpdateCheck setting
- [x] Update README and GoReleaser comment, tidy go.mod
- [x] Build, vet and run the suite

## Summary of Changes

No GitHub API call remains in the binary. The connection picker footer and the home border version label drawn by the update UI are gone; `lazysql --version` still reports the build version. golang.org/x/mod is no longer a dependency.
