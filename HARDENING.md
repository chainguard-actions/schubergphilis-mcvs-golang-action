<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.1

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.1** was hardened automatically. 13 finding(s) were identified and resolved across 2 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Rule (a): Multiple `run:` blocks in action.yml directly interpolate `${{ ... }}` expressions inside shell commands, enabling script injection. 

1. "install task" step: `${{ inputs.task-version }}` is interpolated directly in three shell commands:
   - `grep -q "Task version: v${{ inputs.task-version }}"`
   - `major_version=$(echo "${{ inputs.task-version }}" | sed ...)`
   - `go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}`

2. git config step: `${{ inputs.github-token-for-downloading-private-go-modules }}` is interpolated directly in:
   - `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`

3. "Build binary" step: `${{ inputs.release-build-tags }}` and `${{ github.ref_name }}` are interpolated directly in shell commands:
   - `if [ -n "${{ inputs.release-build-tags }}" ]`
   - `-tags "${{ inputs.release-build-tags }}"`
   - `-ldflags="-X 'main.Version=${{ github.ref_name }}'"`

4. "Compute asset name" step: `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` are all interpolated directly in shell commands:
   - `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"`
   - `if [ -n "${{ inputs.release-build-tags }}" ]`
   - `ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"`

All of these allow an attacker-controlled value to be injected as shell code before the shell ever sees it.

Locations:

- `action.yml:118`
- `action.yml:124`
- `action.yml:222`
- `action.yml:234`

### github-env-injection (severity: high)

The "Compute asset name" step writes `ASSET_NAME` to `$GITHUB_OUTPUT` without sanitization. The value is constructed by directly interpolating `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` into a shell variable, then writing it with `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT`. Any of these inputs could contain newline characters that inject additional key=value pairs into GITHUB_OUTPUT. The required sanitization step (`printf '%s' "$VAR" | tr -d '\n\r'`) is absent before the write.

Locations:

- `action.yml:241`

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

1. 'install task' step: Moved ${{ inputs.task-version }} to env var TASK_VERSION; replaced all 3 inline expressions in run: block.

2. git config step: Moved ${{ inputs.github-token-for-downloading-private-go-modules }} to env var GH_TOKEN_PRIVATE_MODULES; replaced inline expression in run: block.

3. 'Build binary' step: Added RELEASE_BUILD_TAGS and RELEASE_REF_NAME to existing env: block; replaced ${{ inputs.release-build-tags }} and ${{ github.ref_name }} in run: block.

4. 'Compute asset name' step: Added env: block with RELEASE_APPLICATION_NAME, RELEASE_REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS; replaced all 5 inline expressions in run: block; sanitized GITHUB_OUTPUT write with printf '%s' | tr -d '\n\r' to prevent newline injection.

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed three script injection findings in hardened/action/action.yml:
1. Line 121 (install task step): Double-quoted the go install URL argument containing ${major_version} and ${TASK_VERSION} to prevent word splitting.
2. Line 128 (private-modules git config step): Double-quoted the git config URL argument containing ${GH_TOKEN_PRIVATE_MODULES} to prevent word splitting.
3. Line 261 (Build binary step): Refactored the ldflags construction to assign to a local variable `version_ldflags` without inner single quotes (which could be broken by a ref name containing single quotes), then passed it double-quoted as `-ldflags="${version_ldflags}"`.

