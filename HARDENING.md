<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.13.5

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.13.5** was hardened automatically. 1 finding(s) were identified and resolved across 1 iteration(s).

## Findings Fixed

### github-env-injection (severity: high)

In the 'Compute asset name' step, ASSET_NAME is written to $GITHUB_OUTPUT without sanitization: `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT`. ASSET_NAME is built from env vars that come from untrusted sources: inputs.release-application-name, github.ref_name, inputs.release-os, inputs.release-architecture and inputs.release-build-tags. A value with a newline in it could add extra outputs. Clean the value first with `safe=$(printf '%s' "$ASSET_NAME" | tr -d '\n\r')`, then write it as `echo "asset_name=$safe" >> "$GITHUB_OUTPUT"`.

Locations:

- `action.yml:386`

## Iteration Notes

### Iteration 1

**Fixes applied:** github-env-injection

**Notes:**

In the 'Compute asset name' step of action.yml, ASSET_NAME now has newline and carriage-return characters stripped (tr -d '\n\r') before it is written. $GITHUB_OUTPUT is now quoted. No other changes.

