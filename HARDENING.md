<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.2

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.2** was hardened automatically. 16 finding(s) were identified and resolved across 2 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Rule (a) violation — the 'install task' run: block directly interpolates ${{ inputs.task-version }} into shell commands without going through an env: variable. An attacker-controlled value is substituted by the YAML template engine before the shell ever sees the string, enabling command injection. Offending lines:
  `if ! task --version | grep -q "Task version: v${{ inputs.task-version }}"`
  `major_version=$(echo "${{ inputs.task-version }}" | sed -E 's/^([0-9]+).*/\1/')`
  `go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}`

Locations:

- `action.yml:119`

### script-injection (severity: high)

Rule (a) violation — an unnamed run: block directly interpolates ${{ inputs.github-token-for-downloading-private-go-modules }} into a git config URL shell command. The token value is substituted by the YAML template engine before the shell parses the command, enabling injection of arbitrary git config arguments or shell metacharacters. Offending line:
  `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`

Locations:

- `action.yml:125`

### script-injection (severity: high)

Rule (a) violation — the 'Build binary' run: block directly interpolates ${{ inputs.release-build-tags }} into a -tags flag and ${{ github.ref_name }} into a -ldflags string inside shell commands. These values are substituted by the YAML template engine before the shell parses the command, enabling injection of arbitrary shell metacharacters or compiler flags. Offending lines:
  `if [ -n "${{ inputs.release-build-tags }}" ]`
  `-tags "${{ inputs.release-build-tags }}"`
  `-ldflags="-X 'main.Version=${{ github.ref_name }}'"` 

Locations:

- `action.yml:232`

### script-injection (severity: high)

Rule (a) violation — the 'Compute asset name' run: block directly interpolates ${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, and ${{ inputs.release-build-tags }} into shell variable assignments. These values are substituted by the YAML template engine before the shell parses the script, enabling injection of arbitrary shell metacharacters. Offending lines:
  `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"`
  `if [ -n "${{ inputs.release-build-tags }}" ]`
  `ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"`

Locations:

- `action.yml:245`

### github-env-injection (severity: high)

The 'Compute asset name' run: block writes `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT` where ASSET_NAME is constructed by directly interpolating multiple untrusted inputs (${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, ${{ inputs.release-build-tags }}) into the shell script without sanitization. No `printf '%s' ... | tr -d '\n\r'` step is applied before the write, so a newline embedded in any of these values could inject additional key=value pairs into GITHUB_OUTPUT.

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

Fixed all script-injection and github-env-injection findings in action.yml:
1. 'install task' step: moved ${{ inputs.task-version }} to TASK_VERSION env var; shell script now uses ${TASK_VERSION} throughout.
2. git config step (private modules): moved ${{ inputs.github-token-for-downloading-private-go-modules }} to GITHUB_TOKEN_FOR_PRIVATE_MODULES env var.
3. 'Build binary' step: added RELEASE_BUILD_TAGS=${{ inputs.release-build-tags }} and REF_NAME=${{ github.ref_name }} to the existing env: block; shell script uses ${RELEASE_BUILD_TAGS} and ${REF_NAME}.
4. 'Compute asset name' step: added env: block with RELEASE_APPLICATION_NAME, REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS; each value is sanitized with 'printf "%s" | tr -d "\n\r"' before use, and the GITHUB_OUTPUT write uses the sanitized values to prevent newline injection.

### Iteration 2

**Fixes applied:** script-injection, script-injection

**Notes:**

Fixed two script-injection findings in hardened/action/action.yml:
1. Line 107 (install task step): Added double-quotes around the `go install` argument — `"github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"` — so that both `${major_version}` and `${TASK_VERSION}` are properly quoted and cannot be exploited via shell metacharacters in the caller-supplied `task-version` input.
2. Line 112 (git config step): Added double-quotes around the URL key argument — `"url.https://${GITHUB_TOKEN_FOR_PRIVATE_MODULES}@github.com/.insteadOf"` — so that the `${GITHUB_TOKEN_FOR_PRIVATE_MODULES}` variable (sourced from the caller-controlled `github-token-for-downloading-private-go-modules` input) cannot inject shell metacharacters.

