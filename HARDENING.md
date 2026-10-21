<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.13.4

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.13.4** was hardened automatically. 1 finding(s) were identified and resolved across 1 iteration(s).

## Findings Fixed

### github-env-injection (severity: high)

The step 'Compute asset name' (id compute_asset_name) writes `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT` without cleaning the value first. ASSET_NAME is built from env vars that come straight from untrusted inputs: inputs.release-application-name, github.ref_name, inputs.release-os, inputs.release-architecture and inputs.release-build-tags. No `printf '%s' ... | tr -d '\n\r'` step strips newlines before the write, so a value with a newline in it can add extra step outputs. That output is then passed to svenstaro/upload-release-action as asset_name. $GITHUB_OUTPUT is also not quoted.

Locations:

- `action.yml:403`

## Iteration Notes

### Iteration 1

**Fixes applied:** github-env-injection

**Notes:**

In action.yml, the 'Compute asset name' step now removes CR and LF from ASSET_NAME with printf '%s' | tr -d '\n\r' before it writes asset_name to $GITHUB_OUTPUT. $GITHUB_OUTPUT is now double-quoted.

