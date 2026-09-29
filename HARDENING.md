<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.5

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.5** was hardened automatically. 13 finding(s) were identified and resolved across 2 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Multiple `run:` blocks in action.yml directly interpolate `${{ inputs.* }}` and `${{ github.* }}` expressions into shell commands (sub-rule a), allowing an attacker who controls those inputs to inject arbitrary shell commands.

1. "install task" step (lines 109–111): `${{ inputs.task-version }}` is interpolated directly into grep, echo, and `go install` shell commands — e.g. `grep -q "Task version: v${{ inputs.task-version }}"` and `go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}`.

2. git config step (line 117): `${{ inputs.github-token-for-downloading-private-go-modules }}` is interpolated directly into a `git config` URL — `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/...`.

3. "Build binary" step (lines 258–264): `${{ inputs.release-build-tags }}` is interpolated into an `if [ -n "..." ]` test and a `-tags` flag; `${{ github.ref_name }}` is interpolated into `-ldflags`.

4. "Compute asset name" step (lines 271–273): `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` are all interpolated directly into shell variable assignments.

All of these should be moved to `env:` blocks and referenced as quoted `"$ENV_VAR"` shell variables.

Locations:

- `action.yml:109`
- `action.yml:110`
- `action.yml:111`
- `action.yml:117`
- `action.yml:258`
- `action.yml:260`
- `action.yml:264`
- `action.yml:271`
- `action.yml:272`
- `action.yml:273`

### github-env-injection (severity: high)

The "Compute asset name" step writes a value derived from multiple untrusted inputs (`${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, `${{ inputs.release-build-tags }}`) to `$GITHUB_OUTPUT` without sanitization. The shell variable `ASSET_NAME` is built by directly interpolating these expressions and then written with `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT`. No `printf '%s' ... | tr -d '\n\r'` sanitization step is applied before the write, so a newline injected into any of those inputs could poison the output file and set arbitrary output variables.

Locations:

- `action.yml:271`
- `action.yml:276`

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

1. 'install task' step: Moved `${{ inputs.task-version }}` to env: block as TASK_VERSION; updated run script to use ${TASK_VERSION}.

2. git config step: Moved `${{ inputs.github-token-for-downloading-private-go-modules }}` to env: block as GITHUB_TOKEN_FOR_PRIVATE_MODULES; updated run script to use ${GITHUB_TOKEN_FOR_PRIVATE_MODULES}.

3. 'Build binary' step: Added RELEASE_BUILD_TAGS and GITHUB_REF_NAME to the existing env: block; updated run script to use ${RELEASE_BUILD_TAGS} and ${GITHUB_REF_NAME} instead of inline expressions.

4. 'Compute asset name' step: Added env: block with all five expressions (RELEASE_APPLICATION_NAME, GITHUB_REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS); added printf/tr sanitization for each value before writing to $GITHUB_OUTPUT to prevent newline injection.

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed two script-injection findings in hardened/action/action.yml:
1. Line 113 (install task step): Wrapped the entire go install URL in double quotes so both `${major_version}` and `${TASK_VERSION}` are protected: `go install "github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"`.
2. Line 118 (git config step): Double-quoted the git config key argument containing `${GITHUB_TOKEN_FOR_PRIVATE_MODULES}`: `git config --global "url.https://${GITHUB_TOKEN_FOR_PRIVATE_MODULES}@github.com/.insteadOf" https://github.com/`.

