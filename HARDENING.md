<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.4

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.4** was hardened automatically. 16 finding(s) were identified and resolved across 2 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Rule (a) violation: The 'install task' step directly interpolates ${{ inputs.task-version }} inside the run: shell script in three places — inside a grep -q string, inside an echo/sed pipeline, and as part of a go install URL. An attacker-controlled input value is substituted into the shell command before the shell ever sees it, enabling arbitrary command injection.

Offending lines:
  if ! task --version | grep -q "Task version: v${{ inputs.task-version }}"; then
  major_version=$(echo "${{ inputs.task-version }}" | sed -E 's/^([0-9]+).*/\1/')
  go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}

Locations:

- `action.yml:107`

### script-injection (severity: high)

Rule (a) violation: An unnamed step directly interpolates ${{ inputs.github-token-for-downloading-private-go-modules }} inside a git config run: command. This embeds the raw input value into the shell command string before the shell parses it, allowing injection of arbitrary git config arguments or shell metacharacters.

Offending line:
  git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/

Locations:

- `action.yml:118`

### script-injection (severity: high)

Rule (a) violation: The 'Build binary' step directly interpolates ${{ inputs.release-build-tags }} and ${{ github.ref_name }} inside the run: shell script. These expressions are substituted into the shell command string before the shell parses it, enabling arbitrary command injection via crafted input values.

Offending lines:
  if [ -n "${{ inputs.release-build-tags }}" ]; then
    -tags "${{ inputs.release-build-tags }}" \
    -ldflags="-X 'main.Version=${{ github.ref_name }}'" \

Locations:

- `action.yml:232`

### script-injection (severity: high)

Rule (a) violation: The 'Compute asset name' step directly interpolates ${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, and ${{ inputs.release-build-tags }} inside the run: shell script. These expressions are substituted into the shell command string before the shell parses it, enabling arbitrary command injection.

Offending lines:
  ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"
  if [ -n "${{ inputs.release-build-tags }}" ]; then
    ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"

Locations:

- `action.yml:245`

### github-env-injection (severity: high)

The 'Compute asset name' step constructs the ASSET_NAME shell variable from values directly interpolated from ${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, and ${{ inputs.release-build-tags }}, then writes it to $GITHUB_OUTPUT without any sanitization (no 'printf "%s" ... | tr -d '\n\r'' step). A newline character injected via any of these inputs could poison subsequent GITHUB_OUTPUT entries.

Offending line:
  echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT

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

**Fixes applied:** script-injection, static-inline-injection, github-env-injection

**Notes:**

Fixed all script injection findings in action.yml:
1. 'install task' step: moved ${{ inputs.task-version }} to env: TASK_VERSION; replaced 3 inline occurrences with ${TASK_VERSION}.
2. git config step: moved ${{ inputs.github-token-for-downloading-private-go-modules }} to env: GITHUB_TOKEN_PRIVATE_MODULES; replaced inline occurrence with ${GITHUB_TOKEN_PRIVATE_MODULES}.
3. 'Build binary' step: added RELEASE_BUILD_TAGS and RELEASE_REF_NAME to env: block; replaced ${{ inputs.release-build-tags }} and ${{ github.ref_name }} in run: script.
4. 'Compute asset name' step: moved all 5 expressions (release-application-name, github.ref_name, release-os, release-architecture, release-build-tags) to env: block; replaced inline occurrences in run: script; added printf '%s' ... | tr -d '\n\r' sanitization before writing to $GITHUB_OUTPUT to prevent newline injection.

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed all three script-injection findings in hardened/action/action.yml:
1. Line 120 ('install task' step): Double-quoted the go install argument string so ${major_version} and ${TASK_VERSION} are protected: `go install "github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"`.
2. Line 126 (private modules step): Double-quoted the git config key argument containing ${GITHUB_TOKEN_PRIVATE_MODULES}: `git config --global "url.https://${GITHUB_TOKEN_PRIVATE_MODULES}@github.com/.insteadOf" https://github.com/`.
3. Line 240 ('Build binary' step): Added sanitization of RELEASE_REF_NAME before use in ldflags: `safe_ref_name=$(printf '%s' "${RELEASE_REF_NAME}" | tr -cd 'a-zA-Z0-9._-')` and then used `${safe_ref_name}` in the ldflags string, stripping any characters that could break the ldflags syntax or escape the surrounding quotes.

