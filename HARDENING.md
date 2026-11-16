<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.13.8

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.13.8** was hardened automatically. 1 finding(s) were identified and resolved across 1 iteration(s).

## Findings Fixed

### github-env-injection (severity: high)

Step 'Compute asset name' writes a value built from untrusted inputs to $GITHUB_OUTPUT without the required sanitization (printf '%s' ... | tr -d '\n\r'). ASSET_NAME is put together from APPLICATION_NAME (inputs.release-application-name), REF_NAME (github.ref_name), RELEASE_OS (inputs.release-os), RELEASE_ARCHITECTURE (inputs.release-architecture) and RELEASE_BUILD_TAGS (inputs.release-build-tags). It is then written with: echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT. A newline in any of these values could inject extra step outputs.

Locations:

- `action.yml:386`

## Iteration Notes

### Iteration 1

**Fixes applied:** github-env-injection

**Notes:**

In the 'Compute asset name' step, ASSET_NAME now has newlines and carriage returns stripped with printf '%s' | tr -d '\n\r' before it is written to $GITHUB_OUTPUT. $GITHUB_OUTPUT is now quoted.

