<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.7

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.7** was hardened automatically. 16 finding(s) were identified and resolved across 2 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Rule (a): The 'install task' run block interpolates `${{ inputs.task-version }}` directly into shell commands three times without routing through an env var. An attacker-controlled value could inject arbitrary shell commands. Offending lines include: `if ! task --version | grep -q "Task version: v${{ inputs.task-version }}";`, `major_version=$(echo "${{ inputs.task-version }}" | sed -E ...)`, and `go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}`.

Locations:

- `action.yml:107`

### script-injection (severity: high)

Rule (a): The git config run block interpolates `${{ inputs.github-token-for-downloading-private-go-modules }}` directly into a shell command: `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`. A newline or shell metacharacter in the token value could break out of the git config command.

Locations:

- `action.yml:120`

### script-injection (severity: high)

Rule (a): The 'Build binary' run block interpolates `${{ inputs.release-build-tags }}` and `${{ github.ref_name }}` directly into shell commands: `if [ -n "${{ inputs.release-build-tags }}" ]`, `-tags "${{ inputs.release-build-tags }}"`, and `-ldflags="-X 'main.Version=${{ github.ref_name }}'"`. These values flow through YAML template substitution before the shell sees them, enabling command injection.

Locations:

- `action.yml:248`

### script-injection (severity: high)

Rule (a): The 'Compute asset name' run block interpolates multiple untrusted expressions directly into shell commands: `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"`, `if [ -n "${{ inputs.release-build-tags }}" ]`, and `ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"`.

Locations:

- `action.yml:262`

### github-env-injection (severity: high)

The 'Compute asset name' step writes a value derived from multiple untrusted inputs (`${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, `${{ inputs.release-build-tags }}`) to `$GITHUB_OUTPUT` via `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT` without first sanitizing with `printf '%s' ... | tr -d '\n\r'`. A newline character in any of these values could inject additional key=value pairs into GITHUB_OUTPUT.

Locations:

- `action.yml:268`

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

**Fixes applied:** script-injection, static-inline-injection, github-env-injection

**Notes:**

Fixed all script injection and github-env-injection findings in action.yml:

1. 'install task' step: Moved ${{ inputs.task-version }} to env var TASK_VERSION; updated run script to use ${TASK_VERSION}.

2. git config step: Moved ${{ inputs.github-token-for-downloading-private-go-modules }} to env var GITHUB_TOKEN_FOR_PRIVATE_MODULES; updated run script to use ${GITHUB_TOKEN_FOR_PRIVATE_MODULES}.

3. 'Build binary' step: Added RELEASE_BUILD_TAGS and REF_NAME to existing env block; updated run script to use ${RELEASE_BUILD_TAGS} and ${REF_NAME} instead of inline ${{ }} expressions.

4. 'Compute asset name' step: Added env block with RELEASE_APPLICATION_NAME, REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS; updated run script to use env vars; added sanitization via 'printf | tr -d newlines' before writing to $GITHUB_OUTPUT to prevent newline injection.

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed two script-injection findings in action.yml:
1. Line 127 (install task step): Added double-quotes around the `go install` argument `"github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"` to prevent word splitting and glob expansion of the unquoted `${major_version}` and `${TASK_VERSION}` variables.
2. Line 133 (git config step): Added double-quotes around the git config key argument `"url.https://${GITHUB_TOKEN_FOR_PRIVATE_MODULES}@github.com/.insteadOf"` to prevent word splitting and glob expansion of the unquoted `${GITHUB_TOKEN_FOR_PRIVATE_MODULES}` variable. Both variables were already properly sourced from the step's `env:` block.

