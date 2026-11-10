---
# lazysql-w7e1
title: Printable keys leak from WHERE filter to shortcuts
status: completed
type: bug
priority: high
created_at: 2026-11-10T10:02:33Z
updated_at: 2026-11-10T10:02:33Z
---

Typing "=" into the WHERE filter widened the tree; "+", "-", "[", "]", "{", "}", "X", "<", ">" had similar side effects. The home and right-wrapper key handlers skip their shortcuts only when the tab's filtering/editing flag is set, and that flag is set only by the "/" shortcut. Focusing the field with the mouse (mouse support is on) leaves the flag false, so printable keys are treated as commands.

- [x] Pass printable keys through when a tview InputField or TextArea has focus
- [x] Apply to both homeInputCapture and rightWrapperInputCapture
- [x] Add regression test and run suite

## Summary of Changes

Both key handlers return rune events untouched whenever keyboard focus is inside a text field, independent of the filtering/editing flags. Shortcuts behave as before when the table or tree is focused. Regression test TestPrintableKeysReachFocusedTextInput. go test ./... passes.
