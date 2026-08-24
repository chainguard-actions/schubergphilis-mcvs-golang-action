<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.6

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.6** was hardened automatically. 13 finding(s) were identified and resolved across 3 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Multiple `run:` blocks in action.yml directly interpolate `${{ inputs.* }}` and `${{ github.* }}` expressions inside shell command strings (sub-rule a), enabling script injection. An attacker-controlled value in any of these inputs can inject arbitrary shell commands.

1. 'install task' step: `if ! task --version | grep -q "Task version: v${{ inputs.task-version }}"; then` and `go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}`
2. git config step: `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`
3. 'Build binary' step: `if [ -n "${{ inputs.release-build-tags }}" ]` and `-tags "${{ inputs.release-build-tags }}"` and `-ldflags="-X 'main.Version=${{ github.ref_name }}'"` 
4. 'Compute asset name' step: `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"` and `if [ -n "${{ inputs.release-build-tags }}" ]` and `ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"`

Fix: move all `${{ inputs.* }}` and `${{ github.* }}` values into `env:` variables and reference them as quoted shell variables (e.g. `"$INPUT_TASK_VERSION"`).

Locations:

- `action.yml:113`
- `action.yml:114`
- `action.yml:115`
- `action.yml:120`
- `action.yml:289`
- `action.yml:293`
- `action.yml:296`
- `action.yml:305`
- `action.yml:308`
- `action.yml:311`
- `action.yml:313`

### github-env-injection (severity: high)

The 'Compute asset name' `run:` block writes `ASSET_NAME` to `$GITHUB_OUTPUT` without sanitization. `ASSET_NAME` is derived directly from `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` — all untrusted inputs. A newline character in any of these values can inject arbitrary environment variables or outputs into the runner. The required sanitization step (`printf '%s' "$VAR" | tr -d '\n\r'`) is absent before the `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT` write.

Fix: sanitize each component before writing, e.g.:
```
safe_name=$(printf '%s' "${ASSET_NAME}" | tr -d '\n\r')
echo "asset_name=${safe_name}" >> "$GITHUB_OUTPUT"
```

Locations:

- `action.yml:305`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.task-version }}" appears directly in run: block of step "install task"; move to env: map

Locations:

- `action.yml:124`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.task-version }}" appears directly in run: block of step "install task"; move to env: map

Locations:

- `action.yml:125`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.task-version }}" appears directly in run: block of step "install task"; move to env: map

Locations:

- `action.yml:126`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.github-token-for-downloading-private-go-modules }}" appears directly in run: block of step ""; move to env: map

Locations:

- `action.yml:132`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-build-tags }}" appears directly in run: block of step "Build binary"; move to env: map

Locations:

- `action.yml:293`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-build-tags }}" appears directly in run: block of step "Build binary"; move to env: map

Locations:

- `action.yml:295`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-application-name }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:310`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-os }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:310`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-architecture }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:310`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-build-tags }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:311`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-build-tags }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:312`

## Iteration Notes

### Iteration 1

**Fixes applied:** script-injection, github-env-injection, static-inline-injection

**Notes:**

Fixed all script injection and github-env-injection findings in action.yml:

1. 'install task' step: Added `env: TASK_VERSION: ${{ inputs.task-version }}` and replaced all `${{ inputs.task-version }}` references in the run block with `${TASK_VERSION}`.

2. git config step: Added `env: GITHUB_TOKEN_FOR_PRIVATE_MODULES: ${{ inputs.github-token-for-downloading-private-go-modules }}` and replaced the inline expression with `${GITHUB_TOKEN_FOR_PRIVATE_MODULES}`.

3. 'Build binary' step: Added `RELEASE_BUILD_TAGS: ${{ inputs.release-build-tags }}` and `GITHUB_REF_NAME: ${{ github.ref_name }}` to the existing env block, and replaced all inline expressions in the run block with the corresponding shell variables.

4. 'Compute asset name' step: Added a new env block with `RELEASE_APPLICATION_NAME`, `GITHUB_REF_NAME`, `RELEASE_OS`, `RELEASE_ARCHITECTURE`, and `RELEASE_BUILD_TAGS` env vars. Replaced all inline expressions in the run block with shell variables. Added sanitization (`safe_name=$(printf '%s' "${ASSET_NAME}" | tr -d '\n\r')`) before writing to `$GITHUB_OUTPUT`, and quoted `$GITHUB_OUTPUT` properly.

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed the script injection vulnerability in the 'install task' step of action.yml (line 115). The `go install` command now wraps the entire URL string in double quotes: `go install "github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"`. This ensures both `${major_version}` (derived from `inputs.task-version`) and `${TASK_VERSION}` (sourced from `inputs.task-version`) are properly quoted, preventing shell metacharacters in the input from being interpreted as shell commands.

### Iteration 1

**Fixes applied:** script-injection

**Notes:**

Fixed script injection vulnerability in action.yml at line 134. The git config argument `url.https://${GITHUB_TOKEN_FOR_PRIVATE_MODULES}@github.com/.insteadOf` was unquoted, allowing shell metacharacters in the token value to cause command injection. Fixed by wrapping the argument in double quotes: `"url.https://${GITHUB_TOKEN_FOR_PRIVATE_MODULES}@github.com/.insteadOf"`.

