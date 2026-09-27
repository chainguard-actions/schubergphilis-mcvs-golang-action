<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.2

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.2** was hardened automatically. 13 finding(s) were identified and resolved across 3 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Rule (a) violation: Multiple `run:` blocks in action.yml directly interpolate `${{ ... }}` expressions inside shell command strings, enabling script injection. (1) The 'install task' step uses `${{ inputs.task-version }}` directly in grep, echo, and `go install` commands (e.g. `grep -q "Task version: v${{ inputs.task-version }}"`). (2) An unnamed step uses `${{ inputs.github-token-for-downloading-private-go-modules }}` directly inside a `git config` URL. (3) The 'Build binary' step uses `${{ inputs.release-build-tags }}` and `${{ github.ref_name }}` directly in `go build` flags. (4) The 'Compute asset name' step uses `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` directly in shell variable assignments. All of these values are workflow-controllable and must be passed via `env:` and then double-quoted in the shell, never interpolated directly.

Locations:

- `action.yml:108`
- `action.yml:114`
- `action.yml:232`
- `action.yml:248`

### github-env-injection (severity: high)

The 'Compute asset name' step writes `asset_name=${ASSET_NAME}` to `$GITHUB_OUTPUT`, where `ASSET_NAME` is constructed directly from `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` via direct shell interpolation. No sanitization (`printf '%s' ... | tr -d '\n\r'`) is applied before the write. An attacker-controlled input containing newlines could inject arbitrary key-value pairs into the GitHub output context.

Locations:

- `action.yml:248`

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

Fixed all script injection and GitHub env injection findings in action.yml:
1. 'install task' step: moved ${{ inputs.task-version }} into env: TASK_VERSION; shell now uses ${TASK_VERSION} in grep, echo, and go install commands.
2. Unnamed git config step: moved ${{ inputs.github-token-for-downloading-private-go-modules }} into env: GH_TOKEN_PRIVATE_MODULES; shell uses ${GH_TOKEN_PRIVATE_MODULES} in the git config URL.
3. 'Build binary' step: moved ${{ inputs.release-build-tags }} into env: RELEASE_BUILD_TAGS and ${{ github.ref_name }} into env: GITHUB_REF_NAME; shell uses ${RELEASE_BUILD_TAGS} and ${GITHUB_REF_NAME}.
4. 'Compute asset name' step: moved all 5 expressions (${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, ${{ inputs.release-build-tags }}) into env: block; each value is sanitized with printf '%s' ... | tr -d '\n\r' before being used in the GITHUB_OUTPUT write, preventing newline injection.

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed two script-injection findings in hardened/action/action.yml:

1. Line 128 ('install task' step): Wrapped the go install module path in double quotes — `go install "github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"` — so that shell metacharacters in the caller-controlled TASK_VERSION input cannot be interpreted by the shell.

2. Line 244 ('Build binary' step): Added sanitization of GITHUB_REF_NAME before embedding it inside single quotes in the -ldflags argument. A new `safe_ref_name` variable strips single quotes and newlines via `printf '%s' "${GITHUB_REF_NAME}" | tr -d "'\n\r"`, preventing a ref name like `foo'$(evil)'` from breaking out of the single-quote context.

### Iteration 3

**Fixes applied:** script-injection

**Notes:**

Fixed the unquoted shell expansion in the git config step at action.yml line 133. Wrapped the `url.https://${GH_TOKEN_PRIVATE_MODULES}@github.com/.insteadOf` argument in double quotes so that any shell metacharacters in the token value are not interpreted by the shell before being passed to `git config`. The token is already safely sourced via the step's `env:` block.

