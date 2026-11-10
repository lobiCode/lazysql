---
# lazysql-t4rm
title: Remove telemetry entirely
status: completed
type: task
priority: normal
created_at: 2026-11-10T09:42:23Z
updated_at: 2026-11-10T09:42:23Z
---

This fork must never contact a telemetry collector, not even as an opt-out option. Remove the telemetry client, consent flow, usage counters, the Cloudflare collector, and every build/CI hook that could inject an endpoint.

- [x] Delete internal/telemetry, components/telemetry.go and telemetry/collector
- [x] Strip usage counters from components and the --no-telemetry flag from main
- [x] Remove TELEMETRY_ENDPOINT wiring from GoReleaser and workflows, update docs
- [x] Build, vet and run the suite

## Summary of Changes

The binary has no telemetry code path left: no HTTP client for usage data, no consent prompt, no telemetry.toml, no ldflags injection. CI no longer tests the collector. The release updater (GitHub releases check) is unchanged.
