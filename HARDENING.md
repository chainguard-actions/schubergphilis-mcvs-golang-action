<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.9

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.9** was hardened automatically. 13 finding(s) were identified and resolved across 3 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Rule (a) violation: Multiple `${{ }}` expressions are interpolated directly inside `run:` shell command strings in action.yml, allowing script injection.

**Step: "install task"** — `${{ inputs.task-version }}` is interpolated three times directly in the shell script:
```
if ! task --version | grep -q "Task version: v${{ inputs.task-version }}"; then
  major_version=$(echo "${{ inputs.task-version }}" | sed -E 's/^([0-9]+).*/\1/')
  go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}
```
A calling workflow can supply a crafted `task-version` input containing shell metacharacters.

**Step: unnamed git config step** — `${{ inputs.github-token-for-downloading-private-go-modules }}` is interpolated directly in the shell command:
```
git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/
```

**Step: "Build binary"** — `${{ inputs.release-build-tags }}` and `${{ github.ref_name }}` are interpolated directly in the shell script:
```
if [ -n "${{ inputs.release-build-tags }}" ]; then
  go build -tags "${{ inputs.release-build-tags }}" ...
else
  go build -ldflags="-X 'main.Version=${{ github.ref_name }}'" ...
```

**Step: "Compute asset name"** — `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` are all interpolated directly in the shell script:
```
ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"
if [ -n "${{ inputs.release-build-tags }}" ]; then
  ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"
```
All of these allow an attacker to inject arbitrary shell commands via crafted input values.

Locations:

- `action.yml:108`
- `action.yml:118`
- `action.yml:249`
- `action.yml:264`

### github-env-injection (severity: high)

The "Compute asset name" step builds the `ASSET_NAME` shell variable from multiple unsanitized `${{ }}` expressions (`inputs.release-application-name`, `github.ref_name`, `inputs.release-os`, `inputs.release-architecture`, `inputs.release-build-tags`) and then writes it directly to `$GITHUB_OUTPUT` without applying the required sanitization step (`printf '%s' ... | tr -d '\n\r'`):
```
ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"
if [ -n "${{ inputs.release-build-tags }}" ]; then
  ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"
else
  ASSET_NAME="${ASSET_NAME_BASE}"
fi
echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT
```
An attacker can inject newlines into any of these input values to write arbitrary key-value pairs into `$GITHUB_OUTPUT`, potentially overwriting other step outputs and influencing downstream job behaviour.

Locations:

- `action.yml:264`

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

Fixed all script injection findings in action.yml:
1. 'install task' step: moved ${{ inputs.task-version }} to env: TASK_VERSION, replaced 3 inline uses with ${TASK_VERSION}
2. git config step: moved ${{ inputs.github-token-for-downloading-private-go-modules }} to env: GITHUB_TOKEN_FOR_PRIVATE_MODULES, replaced inline use
3. 'Build binary' step: added RELEASE_BUILD_TAGS and REF_NAME to env: block, replaced ${{ inputs.release-build-tags }} and ${{ github.ref_name }} inline uses
4. 'Compute asset name' step: moved all 5 expressions (release-application-name, github.ref_name, release-os, release-architecture, release-build-tags) to env: block, sanitized each with printf '%s' ... | tr -d '\n\r' before use, and quoted $GITHUB_OUTPUT to fix github-env-injection finding

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed the script injection in the 'install task' step of action.yml by double-quoting the go install argument: changed `go install github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}` to `go install "github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"`. This prevents shell word-splitting and glob expansion on the caller-controlled TASK_VERSION input value.

### Iteration 3

**Fixes applied:** script-injection

**Notes:**

Fixed the script injection vulnerability in the 'git config' step at action.yml line 117. Added double quotes around the URL argument `"url.https://${GITHUB_TOKEN_FOR_PRIVATE_MODULES}@github.com/.insteadOf"` to prevent shell metacharacters in the token value from breaking out of the URL context and injecting arbitrary shell commands. The variable is still expanded within the double quotes (correct behavior), but the argument is now treated as a single shell word.

