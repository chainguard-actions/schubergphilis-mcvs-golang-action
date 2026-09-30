<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.13

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.13** was hardened automatically. 16 finding(s) were identified and resolved across 3 iteration(s).

## Findings Fixed

### script-injection (severity: high)

The 'install task' run: block directly interpolates ${{ inputs.task-version }} into shell commands three times (in a grep pattern, an echo/sed pipeline, and a go install URL), violating rule (a). An attacker controlling the `task-version` input can inject arbitrary shell commands.

Offending lines:
  if ! task --version | grep -q "Task version: v${{ inputs.task-version }}"; then
  major_version=$(echo "${{ inputs.task-version }}" | sed -E 's/^([0-9]+).*/\1/')
  go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}

Locations:

- `action.yml:124`

### script-injection (severity: high)

An anonymous run: block directly interpolates ${{ inputs.github-token-for-downloading-private-go-modules }} into a git config shell command, violating rule (a). An attacker controlling this input can inject arbitrary shell commands or exfiltrate the token value.

Offending line:
  git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/

Locations:

- `action.yml:132`

### script-injection (severity: high)

The 'Build binary' run: block directly interpolates ${{ inputs.release-build-tags }} (twice) and ${{ github.ref_name }} into shell commands, violating rule (a). An attacker controlling these inputs can inject arbitrary shell commands.

Offending lines:
  if [ -n "${{ inputs.release-build-tags }}" ]; then
  -tags "${{ inputs.release-build-tags }}" \
  -ldflags="-X 'main.Version=${{ github.ref_name }}'" \

Locations:

- `action.yml:228`

### script-injection (severity: high)

The 'Compute asset name' run: block directly interpolates ${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, and ${{ inputs.release-build-tags }} (twice) into shell commands, violating rule (a). An attacker controlling these inputs can inject arbitrary shell commands.

Offending lines:
  ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"
  if [ -n "${{ inputs.release-build-tags }}" ]; then
  ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"

Locations:

- `action.yml:243`

### github-env-injection (severity: high)

The 'Compute asset name' run: block writes `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT` where ASSET_NAME is constructed by directly interpolating ${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, and ${{ inputs.release-build-tags }} into the shell script without any sanitization (no `printf '%s' ... | tr -d '\n\r'` step). A newline injected via any of these inputs can poison GITHUB_OUTPUT with attacker-controlled key-value pairs.

Locations:

- `action.yml:249`

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

- `action.yml:295`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-build-tags }}" appears directly in run: block of step "Build binary"; move to env: map

Locations:

- `action.yml:297`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-application-name }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:312`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-os }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:312`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-architecture }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:312`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-build-tags }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:313`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-build-tags }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:314`

## Iteration Notes

### Iteration 1

**Fixes applied:** script-injection, static-inline-injection, github-env-injection

**Notes:**

Fixed all 16 findings across 4 steps in action.yml:

1. 'install task' step: Moved ${{ inputs.task-version }} to env var TASK_VERSION and replaced all 3 inline interpolations with ${TASK_VERSION}.

2. Anonymous git config step: Moved ${{ inputs.github-token-for-downloading-private-go-modules }} to env var GH_TOKEN_PRIVATE_MODULES and replaced the inline interpolation.

3. 'Build binary' step: Moved ${{ inputs.release-build-tags }} and ${{ github.ref_name }} to env vars RELEASE_BUILD_TAGS and RELEASE_REF_NAME, replacing both inline interpolations.

4. 'Compute asset name' step: Moved all 5 expressions (${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, ${{ inputs.release-build-tags }}) to env vars. Each value is sanitized with printf '%s' ... | tr -d '\n\r' before use, preventing newline injection into GITHUB_OUTPUT (fixing the github-env-injection finding). The GITHUB_OUTPUT write now uses quoted "$GITHUB_OUTPUT".

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed two script injection vulnerabilities in hardened/action/action.yml:
1. Line 128: Double-quoted ${TASK_VERSION} in `go install github.com/go-task/task/v${major_version}/cmd/task@v"${TASK_VERSION}"` to prevent shell metacharacter injection via the task-version input.
2. Line 134: Double-quoted the URL segment containing ${GH_TOKEN_PRIVATE_MODULES} in the git config command: `url."https://${GH_TOKEN_PRIVATE_MODULES}@github.com/".insteadOf` to prevent shell metacharacter injection via the github-token-for-downloading-private-go-modules input.

### Iteration 3

**Fixes applied:** script-injection

**Notes:**

Fixed unquoted `${major_version}` variable in the 'install task' step's `go install` command. Changed `v${major_version}/cmd/task` to `v"${major_version}"/cmd/task` so the shell variable is double-quoted, preventing word-splitting and glob expansion from attacker-controlled input routed through `inputs.task-version` → `TASK_VERSION` env var → `major_version`.

