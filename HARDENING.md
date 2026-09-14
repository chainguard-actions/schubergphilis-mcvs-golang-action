<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.12

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.12** was hardened automatically. 13 finding(s) were identified and resolved across 2 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Rule (a): Multiple `run:` blocks in action.yml directly interpolate `${{ inputs.* }}` and `${{ github.* }}` expressions inside shell commands, enabling script injection.

1. "install task" step (lines ~116-118): `${{ inputs.task-version }}` is interpolated directly into grep, echo, and `go install` shell commands:
   - `grep -q "Task version: v${{ inputs.task-version }}"`
   - `echo "${{ inputs.task-version }}" | sed -E ...`
   - `go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}`

2. git config step (line ~123): `${{ inputs.github-token-for-downloading-private-go-modules }}` is interpolated directly into a `git config` URL, allowing a caller to inject arbitrary git config values:
   - `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`

3. "Build binary" step (lines ~218-222): `${{ inputs.release-build-tags }}` and `${{ github.ref_name }}` are interpolated directly into `go build` shell commands:
   - `if [ -n "${{ inputs.release-build-tags }}" ]`
   - `-tags "${{ inputs.release-build-tags }}"`
   - `-ldflags="-X 'main.Version=${{ github.ref_name }}'"`

4. "Compute asset name" step (lines ~232-237): Multiple inputs and github context values are interpolated directly into shell variable assignments:
   - `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"`
   - `if [ -n "${{ inputs.release-build-tags }}" ]`
   - `ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"`

All of these should be moved to `env:` variables and referenced as `"$ENV_VAR"` in the shell.

Locations:

- `action.yml:116`
- `action.yml:123`
- `action.yml:218`
- `action.yml:232`

### github-env-injection (severity: high)

The "Compute asset name" step builds the `ASSET_NAME` shell variable from untrusted inputs (`${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, `${{ inputs.release-build-tags }}`) via direct `${{ }}` template interpolation, then writes the result to `$GITHUB_OUTPUT` without the required sanitization step (`printf '%s' ... | tr -d '\n\r'`). An attacker-controlled input containing newlines could inject arbitrary key=value pairs into the GitHub output environment.

Offending line: `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT`

Fix: sanitize each input before use, e.g.:
```
safe_name=$(printf '%s' "$RELEASE_APPLICATION_NAME" | tr -d '\n\r')
safe_ref=$(printf '%s' "$GITHUB_REF_NAME" | tr -d '\n\r')
...
echo "asset_name=${ASSET_NAME}" >> "$GITHUB_OUTPUT"
```

Locations:

- `action.yml:237`

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

- `action.yml:295`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-build-tags }}" appears directly in run: block of step "Build binary"; move to env: map

Locations:

- `action.yml:297`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-application-name }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:312`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-os }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:312`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-architecture }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:312`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-build-tags }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:313`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-build-tags }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:314`

## Iteration Notes

### Iteration 1

**Fixes applied:** script-injection, github-env-injection, static-inline-injection

**Notes:**

Fixed all script injection and github-env-injection findings in action.yml:
1. 'install task' step: moved ${{ inputs.task-version }} to env: TASK_VERSION, referenced as ${TASK_VERSION} in shell.
2. git config step: moved ${{ inputs.github-token-for-downloading-private-go-modules }} to env: GITHUB_TOKEN_FOR_PRIVATE_MODULES, referenced as ${GITHUB_TOKEN_FOR_PRIVATE_MODULES} in shell.
3. 'Build binary' step: moved ${{ inputs.release-build-tags }} to env: RELEASE_BUILD_TAGS and ${{ github.ref_name }} to env: GITHUB_REF_NAME, both referenced as env vars in shell.
4. 'Compute asset name' step: moved all ${{ inputs.* }} and ${{ github.* }} expressions to env: block, sanitized each with printf '%s' ... | tr -d '\n\r' before building ASSET_NAME, and quoted $GITHUB_OUTPUT in the echo command to prevent newline injection.

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed two script-injection findings in action.yml:

1. 'install task' step (line ~107): Wrapped the `go install` URL in double quotes so that `${major_version}` and `${TASK_VERSION}` expansions are properly quoted: `go install "github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"`.

2. 'Build binary' step (line ~247): Replaced the problematic `-ldflags="-X 'main.Version=${GITHUB_REF_NAME}'"` pattern (where single quotes inside double quotes provide no protection against word-splitting) with a two-step approach: first assign `ldflags="-X main.Version=${GITHUB_REF_NAME}"` (bash variable assignment RHS is not subject to word-splitting), then pass `-ldflags="${ldflags}"` as a properly double-quoted argument.

