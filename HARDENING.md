<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.10

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.10** was hardened automatically. 16 finding(s) were identified and resolved across 1 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Sub-rule (a): The 'install task' run: block directly interpolates ${{ inputs.task-version }} into shell commands without routing through an env: variable. The expression is embedded in a grep string, an echo/sed pipeline, and a go install command. An attacker-controlled task-version value could inject arbitrary shell commands.

Offending lines:
  if ! task --version | grep -q "Task version: v${{ inputs.task-version }}"; then
  major_version=$(echo "${{ inputs.task-version }}" | sed -E 's/^([0-9]+).*/\1/')
  go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}

Locations:

- `action.yml:129`

### script-injection (severity: high)

Sub-rule (a): An anonymous run: block directly interpolates ${{ inputs.github-token-for-downloading-private-go-modules }} into a git config URL shell command. A newline or shell metacharacter in the token value could break out of the URL context and inject arbitrary git config directives or shell commands.

Offending line:
  git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/

Locations:

- `action.yml:136`

### script-injection (severity: high)

Sub-rule (a): The 'Build binary' run: block directly interpolates ${{ inputs.release-build-tags }} and ${{ github.ref_name }} into shell commands. These expressions appear inside [ -n "..." ] tests, -tags flags, and -ldflags values. An attacker-controlled value could inject arbitrary shell commands or compiler flags.

Offending lines:
  if [ -n "${{ inputs.release-build-tags }}" ]; then
  -tags "${{ inputs.release-build-tags }}" \
  -ldflags="-X 'main.Version=${{ github.ref_name }}'" \

Locations:

- `action.yml:231`

### script-injection (severity: high)

Sub-rule (a): The 'Compute asset name' run: block directly interpolates multiple ${{ inputs.* }} and ${{ github.ref_name }} expressions into shell variable assignments. An attacker-controlled value containing shell metacharacters could break out of the double-quoted string context.

Offending lines:
  ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"
  if [ -n "${{ inputs.release-build-tags }}" ]; then
  ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"

Locations:

- `action.yml:248`

### github-env-injection (severity: high)

The 'Compute asset name' run: block constructs ASSET_NAME from multiple untrusted inputs (${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, ${{ inputs.release-build-tags }}) and writes the result directly to $GITHUB_OUTPUT without sanitization (no 'printf "%s" ... | tr -d "\n\r"' step). A newline character in any of these inputs could inject additional key=value pairs into the GitHub output file, allowing an attacker to set arbitrary step outputs.

Offending line:
  echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT

Locations:

- `action.yml:253`

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

1. 'install task' step: Moved ${{ inputs.task-version }} to env var TASK_VERSION; replaced all 3 inline occurrences in the run block.

2. Anonymous git config step: Moved ${{ inputs.github-token-for-downloading-private-go-modules }} to env var GITHUB_TOKEN_FOR_PRIVATE_MODULES; replaced inline occurrence in the git config URL.

3. 'Build binary' step: Added RELEASE_BUILD_TAGS and RELEASE_REF_NAME to the existing env: block; replaced ${{ inputs.release-build-tags }} and ${{ github.ref_name }} in the run block.

4. 'Compute asset name' step: Added full env: block with RELEASE_APPLICATION_NAME, RELEASE_REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS; replaced all inline expressions; added printf/tr -d '\n\r' sanitization for each value before writing to $GITHUB_OUTPUT to prevent newline injection.

