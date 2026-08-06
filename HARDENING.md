<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.1

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.1** was hardened automatically. 13 finding(s) were identified and resolved across 2 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Rule (a): Multiple `${{ }}` expressions are interpolated directly inside `run:` shell command strings in action.yml, allowing script injection.

1. "install task" step: `${{ inputs.task-version }}` is interpolated directly into shell commands:
   - `if ! task --version | grep -q "Task version: v${{ inputs.task-version }}"`
   - `major_version=$(echo "${{ inputs.task-version }}" | sed -E 's/^([0-9]+).*/\1/')`
   - `go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}`

2. Unnamed git config step: `${{ inputs.github-token-for-downloading-private-go-modules }}` is interpolated directly into a shell command:
   - `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`

3. "Build binary" step: `${{ inputs.release-build-tags }}` and `${{ github.ref_name }}` are interpolated directly into shell commands:
   - `if [ -n "${{ inputs.release-build-tags }}" ]`
   - `-tags "${{ inputs.release-build-tags }}"`
   - `-ldflags="-X 'main.Version=${{ github.ref_name }}'"`

4. "Compute asset name" step: Multiple expressions interpolated directly into shell:
   - `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"`
   - `ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"`

All of these should be moved to `env:` variables and referenced as `"$VAR"` in the shell script.

Locations:

- `action.yml:118`
- `action.yml:119`
- `action.yml:120`
- `action.yml:125`
- `action.yml:249`
- `action.yml:251`
- `action.yml:253`
- `action.yml:260`
- `action.yml:262`

### github-env-injection (severity: high)

The "Compute asset name" step writes `${ASSET_NAME}` to `$GITHUB_OUTPUT` without sanitization. The value of `ASSET_NAME` is constructed by directly interpolating untrusted `${{ inputs.* }}` (release-application-name, release-os, release-architecture, release-build-tags) and `${{ github.ref_name }}` expressions into the shell variable. A newline character in any of these values could inject arbitrary key-value pairs into GITHUB_OUTPUT. The required sanitization step (`printf '%s' "$VAR" | tr -d '\n\r'`) is not applied before the write:

```
ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-..."
...
echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT
```

Locations:

- `action.yml:265`

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
1. 'install task' step: Added env: block with TASK_VERSION=${{ inputs.task-version }} and replaced all 3 inline ${{ inputs.task-version }} uses with ${TASK_VERSION}.
2. Unnamed git config step: Added env: block with GITHUB_TOKEN_FOR_PRIVATE_MODULES=${{ inputs.github-token-for-downloading-private-go-modules }} and replaced the inline use.
3. 'Build binary' step: Added RELEASE_BUILD_TAGS and REF_NAME to the existing env: block, replaced all inline ${{ inputs.release-build-tags }} and ${{ github.ref_name }} uses.
4. 'Compute asset name' step: Added full env: block with RELEASE_APPLICATION_NAME, REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS; replaced all inline expressions; added printf/tr sanitization before writing to $GITHUB_OUTPUT to prevent newline injection.

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed the script injection vulnerability in the 'install task' step of action.yml. The `go install` command at line 113 now wraps the entire module path argument in double quotes: `go install "github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"`. This prevents shell metacharacters in the `task-version` input (already correctly placed in the `env:` block as `TASK_VERSION`) from being interpreted as shell commands. Both `${major_version}` (derived from TASK_VERSION via sed) and `${TASK_VERSION}` are now safely enclosed within a double-quoted string.

