<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.10

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.10** was hardened automatically. 16 finding(s) were identified and resolved across 2 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Rule (a) violation: The 'install task' run: block directly interpolates ${{ inputs.task-version }} into shell commands without routing through an env: variable. Offending lines: `if ! task --version | grep -q "Task version: v${{ inputs.task-version }}";`, `major_version=$(echo "${{ inputs.task-version }}" | sed -E ...)`, and `go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}`. An attacker-controlled task-version value could inject arbitrary shell commands.

Locations:

- `action.yml:108`

### script-injection (severity: high)

Rule (a) violation: An unnamed run: block directly interpolates ${{ inputs.github-token-for-downloading-private-go-modules }} into a git config shell command: `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`. A malicious token value containing shell metacharacters could inject arbitrary commands.

Locations:

- `action.yml:115`

### script-injection (severity: high)

Rule (a) violation: The 'Build binary' run: block directly interpolates ${{ inputs.release-build-tags }} and ${{ github.ref_name }} into shell commands: `if [ -n "${{ inputs.release-build-tags }}" ]`, `-tags "${{ inputs.release-build-tags }}"`, and `-ldflags="-X 'main.Version=${{ github.ref_name }}'"`. Attacker-controlled input values could inject arbitrary shell commands or compiler flags.

Locations:

- `action.yml:220`

### script-injection (severity: high)

Rule (a) violation: The 'Compute asset name' run: block directly interpolates multiple ${{ inputs.* }} and ${{ github.ref_name }} expressions into shell commands: `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"`, `if [ -n "${{ inputs.release-build-tags }}" ]`, and `ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"`. Any of these attacker-controlled values could inject arbitrary shell commands.

Locations:

- `action.yml:232`

### github-env-injection (severity: high)

The 'Compute asset name' run: block writes ASSET_NAME to $GITHUB_OUTPUT without sanitization. ASSET_NAME is constructed by directly interpolating ${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, and ${{ inputs.release-build-tags }} into the shell script. The final write `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT` is not preceded by the required sanitization step (`printf '%s' ... | tr -d '\n\r'`). A newline injected via any of these inputs could poison GITHUB_OUTPUT with arbitrary key-value pairs.

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

1. 'install task' step: Moved ${{ inputs.task-version }} to env var TASK_VERSION; replaced all three inline occurrences in run: block.

2. Unnamed git config step: Moved ${{ inputs.github-token-for-downloading-private-go-modules }} to env var GITHUB_TOKEN_FOR_PRIVATE_MODULES; replaced inline occurrence in run: block.

3. 'Build binary' step: Added RELEASE_BUILD_TAGS and REF_NAME to existing env: block; replaced ${{ inputs.release-build-tags }} and ${{ github.ref_name }} in run: block.

4. 'Compute asset name' step: Added full env: block with RELEASE_APPLICATION_NAME, REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS; replaced all inline expressions; added printf/tr sanitization for each value before writing to GITHUB_OUTPUT to prevent newline injection; quoted $GITHUB_OUTPUT.

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed script injection vulnerability in the 'install task' step at action.yml line 120. The `go install` URL argument was unquoted, allowing shell metacharacters in `inputs.task-version` to achieve command injection. Fixed by wrapping the entire URL in double quotes: `go install "github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"`.

