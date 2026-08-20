<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.4

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.4** was hardened automatically. 16 finding(s) were identified and resolved across 2 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Sub-rule (a): The 'install task' step in action.yml directly interpolates ${{ inputs.task-version }} inside a run: shell script. This allows an attacker-controlled input value to be injected into the shell command before the shell ever parses it, enabling arbitrary command execution. Offending lines include: `if ! task --version | grep -q "Task version: v${{ inputs.task-version }}";`, `major_version=$(echo "${{ inputs.task-version }}" | sed -E ...)`, and `go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}`.

Locations:

- `action.yml:108`

### script-injection (severity: high)

Sub-rule (a): An anonymous run: step directly interpolates ${{ inputs.github-token-for-downloading-private-go-modules }} inside a shell command: `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`. A newline or shell metacharacter in the input value could break out of the git config argument and execute arbitrary commands.

Locations:

- `action.yml:115`

### script-injection (severity: high)

Sub-rule (a): The 'Build binary' step directly interpolates multiple ${{ inputs.* }} and ${{ github.* }} expressions inside a run: shell script. Offending lines include: `if [ -n "${{ inputs.release-build-tags }}" ]`, `-tags "${{ inputs.release-build-tags }}"`, and `-ldflags="-X 'main.Version=${{ github.ref_name }}'"`. Attacker-controlled values for inputs.release-build-tags or a crafted github.ref_name (e.g. via a tag push) can inject arbitrary shell commands.

Locations:

- `action.yml:233`

### script-injection (severity: high)

Sub-rule (a): The 'Compute asset name' step directly interpolates multiple ${{ inputs.* }} and ${{ github.ref_name }} expressions inside a run: shell script. Offending lines include: `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"` and `ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"`.

Locations:

- `action.yml:249`

### github-env-injection (severity: high)

The 'Compute asset name' step builds ASSET_NAME from unsanitized ${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, and ${{ inputs.release-build-tags }} expressions (all interpolated directly into the shell), then writes the result to $GITHUB_OUTPUT via `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT` without applying the required sanitization step (`printf '%s' ... | tr -d '\n\r'`). A newline character embedded in any of these values could inject arbitrary key=value pairs into the GITHUB_OUTPUT file, poisoning subsequent steps.

Locations:

- `action.yml:257`

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

Fixed all script-injection and github-env-injection findings in hardened/action/action.yml:
1. 'install task' step: moved ${{ inputs.task-version }} to env: block as TASK_VERSION; replaced all three inline uses in run: with ${TASK_VERSION}.
2. Anonymous git config step: moved ${{ inputs.github-token-for-downloading-private-go-modules }} to env: block as GITHUB_TOKEN_FOR_PRIVATE_MODULES; replaced inline use in run: with ${GITHUB_TOKEN_FOR_PRIVATE_MODULES}.
3. 'Build binary' step: added RELEASE_BUILD_TAGS and RELEASE_REF_NAME to existing env: block; replaced ${{ inputs.release-build-tags }} and ${{ github.ref_name }} in run: with env var references.
4. 'Compute asset name' step: added full env: block with RELEASE_APPLICATION_NAME, RELEASE_REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS; replaced all inline ${{ }} in run: with env vars; added printf '%s' ... | tr -d '\n\r' sanitization for each value before writing to $GITHUB_OUTPUT to prevent newline injection; also properly quoted $GITHUB_OUTPUT.

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed the script injection vulnerability in the 'install task' step at action.yml line 113. The `go install` command now wraps the entire module path argument in double quotes: `go install "github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"`. This prevents attacker-controlled values in `inputs.task-version` (mapped to `TASK_VERSION` env var) containing shell metacharacters from breaking out of the argument and executing arbitrary commands.

