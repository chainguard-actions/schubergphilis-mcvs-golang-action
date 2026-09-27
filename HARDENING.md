<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.10

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.10** was hardened automatically. 13 finding(s) were identified and resolved across 2 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Rule (a) violation: Multiple ${{ }} expressions are interpolated directly inside run: shell command strings in action.yml, allowing an attacker-controlled value to be executed as shell code.

1. 'install task' step: `${{ inputs.task-version }}` is embedded directly in three shell commands — in a grep pattern, a sed pipeline, and a `go install` URL. A malicious task-version value (e.g. containing `;`, `$(...)`, or backticks) would be executed by the shell.
   Offending lines:
   - `if ! task --version | grep -q "Task version: v${{ inputs.task-version }}"; then`
   - `major_version=$(echo "${{ inputs.task-version }}" | sed -E 's/^([0-9]+).*/\1/')`
   - `go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}`

2. Unnamed git-config step: `${{ inputs.github-token-for-downloading-private-go-modules }}` is interpolated directly into a `git config` URL. A token value containing shell metacharacters would be executed.
   Offending line:
   - `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`

3. 'Build binary' step: `${{ inputs.release-build-tags }}` and `${{ github.ref_name }}` are interpolated directly into shell commands inside the run: block.
   Offending lines:
   - `if [ -n "${{ inputs.release-build-tags }}" ]; then`
   - `-tags "${{ inputs.release-build-tags }}"`
   - `-ldflags="-X 'main.Version=${{ github.ref_name }}'"`

4. 'Compute asset name' step: `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` are all interpolated directly into shell variable assignments and a conditional.
   Offending lines:
   - `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"`
   - `if [ -n "${{ inputs.release-build-tags }}" ]; then`
   - `ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"`

Fix: Move all ${{ }} values into env: variables and reference them as quoted shell variables (e.g. "$TASK_VERSION") inside the run: block.

Locations:

- `action.yml:124`
- `action.yml:125`
- `action.yml:126`
- `action.yml:132`
- `action.yml:232`
- `action.yml:234`
- `action.yml:237`
- `action.yml:243`
- `action.yml:244`
- `action.yml:246`

### github-env-injection (severity: high)

The 'Compute asset name' step writes a value derived from multiple untrusted inputs (${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, ${{ inputs.release-build-tags }}) to $GITHUB_OUTPUT without sanitization. The shell variable ASSET_NAME is constructed from these ${{ }} expressions (which are interpolated directly into the run: block) and then written with `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT`. No `printf '%s' ... | tr -d '\n\r'` sanitization step is applied before the write. A newline character injected via any of these inputs could allow an attacker to inject arbitrary key=value pairs into GITHUB_OUTPUT, potentially overwriting outputs consumed by downstream steps.

Offending line: `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT`

Fix:
```bash
safe_name=$(printf '%s' "${ASSET_NAME}" | tr -d '\n\r')
echo "asset_name=${safe_name}" >> "$GITHUB_OUTPUT"
```

Locations:

- `action.yml:250`

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
1. 'install task' step: moved ${{ inputs.task-version }} to env var TASK_VERSION, referenced as ${TASK_VERSION} in run block
2. git config step: moved ${{ inputs.github-token-for-downloading-private-go-modules }} to env var GH_TOKEN_PRIVATE_MODULES, referenced as ${GH_TOKEN_PRIVATE_MODULES} in run block
3. 'Build binary' step: moved ${{ inputs.release-build-tags }} to RELEASE_BUILD_TAGS and ${{ github.ref_name }} to REF_NAME env vars, referenced as shell variables in run block
4. 'Compute asset name' step: moved all 5 expressions (${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, ${{ inputs.release-build-tags }}) to env vars; added printf '%s' "${ASSET_NAME}" | tr -d '\n\r' sanitization before writing to $GITHUB_OUTPUT to prevent newline injection

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed the script injection finding in the 'install task' step of action.yml (line 128). The `go install` command argument `github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}` was unquoted, allowing an attacker-controlled `task-version` input containing shell metacharacters to cause unexpected behavior. The fix wraps the entire argument in double quotes: `go install "github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"`.

