<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.14.0

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.14.0** was hardened automatically. 1 finding(s) were identified and resolved across 1 iteration(s).

## Findings Fixed

### github-env-injection (severity: high)

The 'Compute asset name' step builds ASSET_NAME from values the caller controls: inputs.release-application-name, inputs.release-os, inputs.release-architecture and inputs.release-build-tags, plus github.ref_name. These reach the step through env vars (APPLICATION_NAME, REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS). The step then writes the value to $GITHUB_OUTPUT with `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT` without stripping newlines first. A newline in any of these values could add extra step outputs. Fix: run `safe=$(printf '%s' "$ASSET_NAME" | tr -d '\n\r')` before the write, then `echo "asset_name=$safe" >> "$GITHUB_OUTPUT"`.

Locations:

- `action.yml:517`

## Iteration Notes

### Iteration 1

**Fixes applied:** github-env-injection

**Notes:**

In action.yml, the 'Compute asset name' step now removes newline and carriage-return characters from ASSET_NAME before writing it to GITHUB_OUTPUT. It uses printf piped to tr, and the GITHUB_OUTPUT path is now quoted. No other changes.

