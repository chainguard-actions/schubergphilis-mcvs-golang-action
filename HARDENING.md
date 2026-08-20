<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.5

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.5** was hardened automatically. 16 finding(s) were identified and resolved across 2 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Sub-rule (a): The 'install task' run block directly interpolates `${{ inputs.task-version }}` into shell commands. An attacker-controlled value for `inputs.task-version` is substituted by the YAML template engine before the shell parses the command, enabling command injection. Offending lines include: `grep -q "Task version: v${{ inputs.task-version }}"`, `echo "${{ inputs.task-version }}" | sed -E ...`, and `go install ...@v${{ inputs.task-version }}`.

Locations:

- `action.yml:109`

### script-injection (severity: high)

Sub-rule (a): The git config step directly interpolates `${{ inputs.github-token-for-downloading-private-go-modules }}` into a shell command: `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`. A caller-controlled token value is substituted before the shell parses the command, enabling injection.

Locations:

- `action.yml:117`

### script-injection (severity: high)

Sub-rule (a): The 'Build binary' run block directly interpolates multiple `${{ ... }}` expressions into shell commands: `if [ -n "${{ inputs.release-build-tags }}" ]`, `-tags "${{ inputs.release-build-tags }}"`, and `-ldflags="-X 'main.Version=${{ github.ref_name }}'"`. These values are substituted by the YAML template engine before the shell parses the command, enabling command injection via attacker-controlled inputs.

Locations:

- `action.yml:220`

### script-injection (severity: high)

Sub-rule (a): The 'Compute asset name' run block directly interpolates multiple `${{ ... }}` expressions into shell commands: `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"` and `ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"`. These values are substituted by the YAML template engine before the shell parses the command, enabling command injection.

Locations:

- `action.yml:233`

### github-env-injection (severity: high)

The 'Compute asset name' step writes `${ASSET_NAME}` to `$GITHUB_OUTPUT` without sanitization. `ASSET_NAME` is derived entirely from untrusted inputs interpolated directly in the run block: `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}`. A newline character in any of these values would allow injection of arbitrary key-value pairs into GITHUB_OUTPUT. The required sanitization (`printf '%s' ... | tr -d '\n\r'`) is absent.

Locations:

- `action.yml:238`

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

1. 'install task' step: Added env block with TASK_VERSION=${{ inputs.task-version }} and replaced all 3 inline ${{ inputs.task-version }} occurrences in the run block with ${TASK_VERSION}.

2. git config step (private modules): Added env block with GH_TOKEN_PRIVATE_MODULES=${{ inputs.github-token-for-downloading-private-go-modules }} and replaced the inline token interpolation in the run block with ${GH_TOKEN_PRIVATE_MODULES}.

3. 'Build binary' step: Added RELEASE_BUILD_TAGS and RELEASE_REF_NAME to the existing env block, replaced all inline ${{ inputs.release-build-tags }} and ${{ github.ref_name }} in the run block with the env vars.

4. 'Compute asset name' step: Added a full env block with all 5 values (RELEASE_APPLICATION_NAME, RELEASE_REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS), replaced all inline ${{ }} expressions in the run block with env vars, and added printf '%s' ... | tr -d '\n\r' sanitization for each value before writing to $GITHUB_OUTPUT to prevent newline injection.

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed three script-injection findings in hardened/action/action.yml:
1. Line 113 ('install task' step): Added double quotes around the `go install` argument: `go install "github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"` — both TASK_VERSION and major_version are now properly double-quoted.
2. Line 127 (git config step): Added double quotes around the git config key argument: `git config --global "url.https://${GH_TOKEN_PRIVATE_MODULES}@github.com/.insteadOf" https://github.com/` — GH_TOKEN_PRIVATE_MODULES is now properly double-quoted.
3. Line 233 ('Build binary' step): Extracted the ldflags value into a separate variable with proper double-quoting: `ldflags="-X 'main.Version=${RELEASE_REF_NAME}'"` then `go build -ldflags="${ldflags}" ...` — RELEASE_REF_NAME is now properly double-quoted during expansion.

