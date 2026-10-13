<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.13.1

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.13.1** was hardened automatically. 10 finding(s) were identified and resolved across 2 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Multiple `${{ }}` expressions are interpolated directly inside `run:` shell command strings, violating sub-rule (a). This allows a caller to inject arbitrary shell commands via input values or GitHub context.

1. `git config` step: `${{ inputs.github-token-for-downloading-private-go-modules }}` is embedded directly in the shell command: `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`

2. 'Build binary' step: `${{ inputs.release-build-tags }}` appears in `if [ -n "${{ inputs.release-build-tags }}" ]` and `-tags "${{ inputs.release-build-tags }}"`, and `${{ github.ref_name }}` appears in `-ldflags="-X 'main.Version=${{ github.ref_name }}'"` — all directly inside the `run:` script.

3. 'Compute asset name' step: `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` are all interpolated directly inside the `run:` script shell string.

Locations:

- `action.yml:168`
- `action.yml:295`
- `action.yml:310`

### github-env-injection (severity: high)

The 'Compute asset name' step writes `ASSET_NAME` to `$GITHUB_OUTPUT` without sanitization. The variable is constructed directly from `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` — all interpolated inline in the shell script. A newline character in any of these values would allow injection of arbitrary key=value pairs into `$GITHUB_OUTPUT`. The required sanitization step (`printf '%s' ... | tr -d '\n\r'`) is absent before the `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT` write.

Locations:

- `action.yml:316`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.github-token-for-downloading-private-go-modules }}" appears directly in run: block of step ""; move to env: map

Locations:

- `action.yml:179`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-build-tags }}" appears directly in run: block of step "Build binary"; move to env: map

Locations:

- `action.yml:342`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-build-tags }}" appears directly in run: block of step "Build binary"; move to env: map

Locations:

- `action.yml:344`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-application-name }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:359`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-os }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:359`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-architecture }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:359`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-build-tags }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:360`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-build-tags }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:361`

## Iteration Notes

### Iteration 1

**Fixes applied:** script-injection, github-env-injection, static-inline-injection

**Notes:**

Fixed three locations in action.yml:
1. Git config step: moved `${{ inputs.github-token-for-downloading-private-go-modules }}` to env var `GH_TOKEN_PRIVATE_MODULES` and referenced it as `${GH_TOKEN_PRIVATE_MODULES}` in the run block.
2. Build binary step: moved `${{ inputs.release-build-tags }}` and `${{ github.ref_name }}` to env vars `RELEASE_BUILD_TAGS` and `RELEASE_REF_NAME` in the existing env block, and updated the run block to use shell variable references.
3. Compute asset name step: moved all five inline expressions to a new env block (`RELEASE_APPLICATION_NAME`, `RELEASE_REF_NAME`, `RELEASE_OS`, `RELEASE_ARCHITECTURE`, `RELEASE_BUILD_TAGS`), updated the run block to use shell variables, and added `printf '%s' "${ASSET_NAME}" | tr -d '\n\r'` sanitization before writing to `$GITHUB_OUTPUT` to prevent newline injection.

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed the unquoted variable expansion in the git config command at action.yml line 180. Changed `git config --global url.https://${GH_TOKEN_PRIVATE_MODULES}@github.com/.insteadOf https://github.com/` to `git config --global "url.https://${GH_TOKEN_PRIVATE_MODULES}@github.com/.insteadOf" https://github.com/`. The double-quoting of the argument containing the variable prevents shell word-splitting and metacharacter interpretation if the token value contains shell special characters. The variable was already correctly placed in the step's `env:` block.

