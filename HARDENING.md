<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.9

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.9** was hardened automatically. 16 finding(s) were identified and resolved across 4 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Rule (a): The 'install task' run: block directly interpolates ${{ inputs.task-version }} into shell commands without routing through an env: variable. This allows an attacker-controlled input to inject arbitrary shell commands. Offending lines: `if ! task --version | grep -q "Task version: v${{ inputs.task-version }}";`, `major_version=$(echo "${{ inputs.task-version }}" | ...)`, and `go install ...@v${{ inputs.task-version }}`.

Locations:

- `action.yml:124`

### script-injection (severity: high)

Rule (a): An unnamed run: block directly interpolates ${{ inputs.github-token-for-downloading-private-go-modules }} into a git config shell command: `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`. This allows injection of arbitrary characters (including shell metacharacters) into the command.

Locations:

- `action.yml:132`

### script-injection (severity: high)

Rule (a): The 'Build binary' run: block directly interpolates ${{ inputs.release-build-tags }} and ${{ github.ref_name }} into shell commands: `if [ -n "${{ inputs.release-build-tags }}" ]`, `-tags "${{ inputs.release-build-tags }}"`, and `-ldflags="-X 'main.Version=${{ github.ref_name }}'"`. These expressions are substituted before the shell parses the script, enabling script injection.

Locations:

- `action.yml:228`

### script-injection (severity: high)

Rule (a): The 'Compute asset name' run: block directly interpolates multiple ${{ inputs.* }} and ${{ github.ref_name }} expressions into shell variable assignments: `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"` and `if [ -n "${{ inputs.release-build-tags }}" ]` / `ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"`.

Locations:

- `action.yml:244`

### github-env-injection (severity: high)

The 'Compute asset name' run: block writes a value derived from multiple untrusted inputs (${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, ${{ inputs.release-build-tags }}) to $GITHUB_OUTPUT via `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT` without applying the required sanitization step (`printf '%s' ... | tr -d '\n\r'`). A newline character in any of these inputs could inject arbitrary key=value pairs into the GitHub output context.

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

Fixed all script injection and github-env-injection findings in action.yml:

1. 'install task' step: Moved ${{ inputs.task-version }} to env var TASK_VERSION; replaced all 3 inline occurrences in run: block.

2. Unnamed git config step: Moved ${{ inputs.github-token-for-downloading-private-go-modules }} to env var GITHUB_TOKEN_FOR_PRIVATE_MODULES; replaced inline occurrence in run: block.

3. 'Build binary' step: Added RELEASE_BUILD_TAGS and REF_NAME to existing env: block; replaced ${{ inputs.release-build-tags }} and ${{ github.ref_name }} in run: block.

4. 'Compute asset name' step: Added env: block with RELEASE_APPLICATION_NAME, REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS; replaced all inline expressions in run: block; added printf/tr sanitization for each value before writing to $GITHUB_OUTPUT to prevent newline injection; also quoted $GITHUB_OUTPUT reference.

### Iteration 2

**Fixes applied:** unsafe-shell

**Notes:**

Fixed two unsafe curl-pipe-to-shell patterns in hardened/action/build/task.yml:
1. helm-install (line ~246): Changed `curl ... | bash -s -- --version {{.HELM_VERSION}}` to download the script to a mktemp file first, then execute `bash "$HELM_INSTALL_SCRIPT" --version {{.HELM_VERSION}}` (dropped '--' which was the shell's stdin option terminator, not a script argument), then clean up.
2. golangci-lint-install (line ~261): Changed `curl ... | sh -s -- -b {{.GOBIN}} {{.GOLANGCI_LINT_VERSION}}` to download the script to a mktemp file first, then execute `sh "$GOLANGCI_INSTALL_SCRIPT" -b {{.GOBIN}} {{.GOLANGCI_LINT_VERSION}}` (dropped '--' per the same reasoning), then clean up.

### Iteration 3

**Fixes applied:** script-injection

**Notes:**

Fixed the unquoted variable expansion in the 'install task' step's go install command. Changed `go install github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}` to `go install "github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"`. Double-quoting the argument prevents word splitting and glob expansion from attacker-controlled values in the task-version input, while still allowing the necessary variable substitution.

### Iteration 4

**Fixes applied:** script-injection

**Notes:**

Fixed the unquoted variable expansion in the git config command at action.yml line 119. Wrapped the URL argument containing ${GITHUB_TOKEN_FOR_PRIVATE_MODULES} in double quotes: changed `git config --global url.https://${GITHUB_TOKEN_FOR_PRIVATE_MODULES}@github.com/.insteadOf https://github.com/` to `git config --global "url.https://${GITHUB_TOKEN_FOR_PRIVATE_MODULES}@github.com/.insteadOf" https://github.com/`. This ensures shell metacharacters in the token value cannot alter the command structure.

