<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.5

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.5** was hardened automatically. 16 finding(s) were identified and resolved across 2 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Rule (a): Direct ${{ }} expression interpolation inside run: shell commands. The 'install task' step interpolates ${{ inputs.task-version }} directly into shell commands (grep string, echo pipeline, and go install URL). An attacker-controlled task-version value could inject arbitrary shell commands. Offending lines: `grep -q "Task version: v${{ inputs.task-version }}"`, `echo "${{ inputs.task-version }}"`, and `go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}`.

Locations:

- `action.yml:108`

### script-injection (severity: high)

Rule (a): Direct ${{ }} expression interpolation inside a run: shell command. The anonymous git-config step embeds ${{ inputs.github-token-for-downloading-private-go-modules }} directly into a git config URL string: `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`. A newline or shell metacharacter in the token value could inject additional git config commands or shell commands.

Locations:

- `action.yml:124`

### script-injection (severity: high)

Rule (a): Direct ${{ }} expression interpolation inside run: shell commands in the 'Build binary' step. The expressions ${{ inputs.release-build-tags }} and ${{ github.ref_name }} are interpolated directly into go build command arguments: `-tags "${{ inputs.release-build-tags }}"` and `-ldflags="-X 'main.Version=${{ github.ref_name }}'"`. Attacker-controlled values could inject arbitrary build flags or shell metacharacters.

Locations:

- `action.yml:220`

### script-injection (severity: high)

Rule (a): Direct ${{ }} expression interpolation inside run: shell commands in the 'Compute asset name' step. The expressions ${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, and ${{ inputs.release-build-tags }} are all interpolated directly into shell variable assignments and a conditional test: `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"` and `if [ -n "${{ inputs.release-build-tags }}" ]`. Attacker-controlled values could inject shell metacharacters.

Locations:

- `action.yml:235`

### github-env-injection (severity: high)

The 'Compute asset name' step writes ASSET_NAME to $GITHUB_OUTPUT without sanitization. ASSET_NAME is derived from multiple untrusted inputs interpolated directly in the shell: ${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, and ${{ inputs.release-build-tags }}. The write `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT` is not preceded by the required sanitization step (`printf '%s' ... | tr -d '\n\r'`). A newline character in any of these values could inject arbitrary key=value pairs into GITHUB_OUTPUT.

Locations:

- `action.yml:243`

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

Fixed all script-injection and github-env-injection findings in action.yml:

1. 'install task' step: Moved ${{ inputs.task-version }} to env: block as TASK_VERSION; replaced 3 inline uses in run: block.

2. git config step (anonymous): Moved ${{ inputs.github-token-for-downloading-private-go-modules }} to env: block as GITHUB_TOKEN_FOR_PRIVATE_MODULES; replaced inline use in run: block.

3. 'Build binary' step: Added RELEASE_BUILD_TAGS and RELEASE_REF_NAME to existing env: block; replaced ${{ inputs.release-build-tags }} and ${{ github.ref_name }} inline uses in run: block.

4. 'Compute asset name' step: Added new env: block with RELEASE_APPLICATION_NAME, RELEASE_REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS; replaced all 5 inline ${{ }} uses in run: block; added sanitization (printf | tr -d '\n\r') before writing to $GITHUB_OUTPUT to fix github-env-injection.

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed two script injection findings in hardened/action/action.yml:
1. Line 128 (install task step): Added double-quotes around the go install argument `"github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"` to prevent shell metacharacter injection from the TASK_VERSION input-derived variables.
2. Line 134 (git config step): Added double-quotes around the git config URL key argument `"url.https://${GITHUB_TOKEN_FOR_PRIVATE_MODULES}@github.com/.insteadOf"` to prevent shell metacharacter injection from the GITHUB_TOKEN_FOR_PRIVATE_MODULES input-derived variable.

