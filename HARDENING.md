<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.9

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.9** was hardened automatically. 16 finding(s) were identified and resolved across 1 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Sub-rule (a): The 'install task' run: block directly interpolates ${{ inputs.task-version }} into shell commands. An attacker-controlled value for this input can inject arbitrary shell commands. Offending lines:
  `if ! task --version | grep -q "Task version: v${{ inputs.task-version }}"; then`
  `major_version=$(echo "${{ inputs.task-version }}" | sed -E 's/^([0-9]+).*/\1/')`
  `go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}`
Fix: move inputs.task-version into an env: var and reference it as a quoted shell variable.

Locations:

- `action.yml:120`

### script-injection (severity: high)

Sub-rule (a): An anonymous run: block directly interpolates ${{ inputs.github-token-for-downloading-private-go-modules }} into a git config command. A newline or shell metacharacter in the token value can break out of the URL context and inject arbitrary git config directives or shell commands. Offending line:
  `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`
Fix: pass the token via an env: var and reference it as a quoted shell variable.

Locations:

- `action.yml:127`

### script-injection (severity: high)

Sub-rule (a): The 'Build binary' run: block directly interpolates ${{ inputs.release-build-tags }} and ${{ github.ref_name }} into shell commands passed to `go build`. An attacker-controlled tag name or build-tags value can inject arbitrary shell commands or compiler flags. Offending lines:
  `if [ -n "${{ inputs.release-build-tags }}" ]; then`
  `-tags "${{ inputs.release-build-tags }}"`
  `-ldflags="-X 'main.Version=${{ github.ref_name }}'"`
Fix: move all inputs and github context values into env: vars and reference them as quoted shell variables.

Locations:

- `action.yml:233`

### script-injection (severity: high)

Sub-rule (a): The 'Compute asset name' run: block directly interpolates ${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, and ${{ inputs.release-build-tags }} into shell variable assignments. An attacker-controlled value in any of these inputs can inject shell metacharacters. Offending lines:
  `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"`
  `if [ -n "${{ inputs.release-build-tags }}" ]; then`
  `ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"`
Fix: move all context values into env: vars and reference them as quoted shell variables.

Locations:

- `action.yml:248`

### github-env-injection (severity: high)

The 'Compute asset name' run: block constructs ASSET_NAME_BASE and ASSET_NAME directly from ${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, and ${{ inputs.release-build-tags }} — all untrusted inputs — and then writes the result to $GITHUB_OUTPUT without any newline-stripping sanitization (`printf '%s' ... | tr -d '\n\r'`). A value containing a newline can inject arbitrary key=value pairs into GITHUB_OUTPUT, potentially overwriting outputs consumed by downstream steps. Offending line:
  `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT`

Locations:

- `action.yml:255`

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
1. 'install task' step: Added env: block with TASK_VERSION=${{ inputs.task-version }}, replaced all 3 inline ${{ inputs.task-version }} references with ${TASK_VERSION} shell variable.
2. Anonymous git config step: Added env: block with GH_TOKEN_PRIVATE_MODULES=${{ inputs.github-token-for-downloading-private-go-modules }}, replaced inline expression with ${GH_TOKEN_PRIVATE_MODULES} shell variable.
3. 'Build binary' step: Added RELEASE_BUILD_TAGS=${{ inputs.release-build-tags }} and REF_NAME=${{ github.ref_name }} to existing env: block, replaced all inline expressions with shell variables.
4. 'Compute asset name' step: Added env: block with RELEASE_APPLICATION_NAME, REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS, replaced all inline expressions with shell variables.
5. 'Compute asset name' step (github-env-injection): Added newline sanitization using printf '%s' | tr -d '\n\r' before writing asset_name to $GITHUB_OUTPUT, and quoted $GITHUB_OUTPUT.

