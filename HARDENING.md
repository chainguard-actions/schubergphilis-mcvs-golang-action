<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.13.7

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.13.7** was hardened automatically. 1 finding(s) were identified and resolved across 1 iteration(s).

## Findings Fixed

### github-env-injection (severity: high)

In the 'Compute asset name' step, ASSET_NAME is built from APPLICATION_NAME (inputs.release-application-name), REF_NAME (github.ref_name), RELEASE_OS, RELEASE_ARCHITECTURE and RELEASE_BUILD_TAGS (all inputs). The line `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT` then writes it without first running `printf '%s' ... | tr -d '\n\r'`. A value containing a newline could inject extra step outputs. Fix: `safe=$(printf '%s' "$ASSET_NAME" | tr -d '\n\r'); echo "asset_name=$safe" >> "$GITHUB_OUTPUT"`.

Locations:

- `action.yml:386`

## Iteration Notes

### Iteration 1

**Fixes applied:** github-env-injection

**Notes:**

In the 'Compute asset name' step, the value is now stripped of newline and carriage-return characters before it is written to $GITHUB_OUTPUT, and $GITHUB_OUTPUT is now quoted.

