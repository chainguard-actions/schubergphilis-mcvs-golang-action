<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.1

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.1** was hardened automatically. 16 finding(s) were identified and resolved across 3 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Sub-rule (a): The 'install task' run: block directly interpolates ${{ inputs.task-version }} into shell commands without routing through an env: variable. This allows an attacker-controlled input to inject arbitrary shell commands. Offending lines:
  - `if ! task --version | grep -q "Task version: v${{ inputs.task-version }}"; then`
  - `major_version=$(echo "${{ inputs.task-version }}" | sed -E 's/^([0-9]+).*/\1/')`
  - `go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}`

Locations:

- `action.yml:110`

### script-injection (severity: high)

Sub-rule (a): The git config run: block directly interpolates ${{ inputs.github-token-for-downloading-private-go-modules }} into a shell command. This embeds the token value directly into the command string before the shell sees it, enabling both script injection and token exposure in process listings. Offending line:
  - `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`

Locations:

- `action.yml:117`

### script-injection (severity: high)

Sub-rule (a): The 'Build binary' run: block directly interpolates ${{ inputs.release-build-tags }} and ${{ github.ref_name }} into shell commands. An attacker-controlled input value can inject arbitrary shell commands or compiler flags. Offending lines:
  - `if [ -n "${{ inputs.release-build-tags }}" ]; then`
  - `-tags "${{ inputs.release-build-tags }}"`
  - `-ldflags="-X 'main.Version=${{ github.ref_name }}'"` 

Locations:

- `action.yml:228`

### script-injection (severity: high)

Sub-rule (a): The 'Compute asset name' run: block directly interpolates multiple ${{ inputs.* }} and ${{ github.ref_name }} expressions into shell variable assignments. Attacker-controlled inputs can inject arbitrary shell commands. Offending lines:
  - `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"`
  - `if [ -n "${{ inputs.release-build-tags }}" ]; then`
  - `ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"`

Locations:

- `action.yml:243`

### github-env-injection (severity: high)

The 'Compute asset name' run: block writes ASSET_NAME to $GITHUB_OUTPUT without sanitization. ASSET_NAME is constructed from unsanitized ${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, and ${{ inputs.release-build-tags }} expressions. A newline character in any of these values could inject additional key=value pairs into GITHUB_OUTPUT, allowing an attacker to set arbitrary step outputs. The required sanitization step (`printf '%s' "$VAR" | tr -d '\n\r'`) is absent before the write. Offending line:
  - `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT`

Locations:

- `action.yml:251`

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
1. 'install task' step: moved ${{ inputs.task-version }} to env: TASK_VERSION, replaced all inline occurrences with ${TASK_VERSION}.
2. git config step: moved ${{ inputs.github-token-for-downloading-private-go-modules }} to env: GH_TOKEN_PRIVATE_MODULES, replaced inline occurrence with ${GH_TOKEN_PRIVATE_MODULES}.
3. 'Build binary' step: added RELEASE_BUILD_TAGS and REF_NAME to existing env: block, replaced ${{ inputs.release-build-tags }} and ${{ github.ref_name }} with env vars.
4. 'Compute asset name' step: added env: block with RELEASE_APPLICATION_NAME, REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS; replaced all inline ${{ }} expressions with env vars; added sanitization (printf '%s' | tr -d '\n\r') before writing asset_name to $GITHUB_OUTPUT.

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed script injection in the 'install task' step of action.yml (line 107). Added strict semver validation (^[0-9]+\.[0-9]+\.[0-9]+$) for TASK_VERSION before use, so any value containing shell metacharacters is rejected with an error and exit 1. Also double-quoted the go install argument string to prevent word splitting and globbing on the expanded variables.

### Iteration 1

**Fixes applied:** script-injection

**Notes:**

Fixed the unquoted shell variable expansion in the private-modules step of action.yml. The git config command's URL key containing `${GH_TOKEN_PRIVATE_MODULES}` was wrapped in double quotes: changed `git config --global url.https://${GH_TOKEN_PRIVATE_MODULES}@github.com/.insteadOf https://github.com/` to `git config --global "url.https://${GH_TOKEN_PRIVATE_MODULES}@github.com/.insteadOf" https://github.com/`. This prevents attacker-controlled token values containing shell metacharacters from breaking out of the URL context and injecting arbitrary shell commands.

