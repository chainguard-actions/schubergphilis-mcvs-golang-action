<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.13.6

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.13.6** was hardened automatically. 1 finding(s) were identified and resolved across 1 iteration(s).

## Findings Fixed

### github-env-injection (severity: high)

The 'Compute asset name' step builds ASSET_NAME from untrusted sources: APPLICATION_NAME (inputs.release-application-name), REF_NAME (github.ref_name), RELEASE_OS (inputs.release-os), RELEASE_ARCHITECTURE (inputs.release-architecture) and RELEASE_BUILD_TAGS (inputs.release-build-tags). It then writes the value straight to the output file with `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT`. Nothing strips newlines first with `printf '%s' ... | tr -d '\n\r'`, so a newline in any of these values could inject extra step outputs. Fix: `safe=$(printf '%s' "$ASSET_NAME" | tr -d '\n\r'); echo "asset_name=$safe" >> "$GITHUB_OUTPUT"`.

Locations:

- `action.yml:386`
- `action.yml:369`

## Iteration Notes

### Iteration 1

**Fixes applied:** github-env-injection

**Notes:**

In action.yml, the 'Compute asset name' step now strips newlines and carriage returns from ASSET_NAME (printf '%s' | tr -d '\n\r') before it writes asset_name to GITHUB_OUTPUT. It also quotes "$GITHUB_OUTPUT". The untrusted inputs were already passed through env: in the step.

