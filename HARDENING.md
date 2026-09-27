<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.8

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.8** was hardened automatically. 16 finding(s) were identified and resolved across 3 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Rule (a): The 'install task' step directly interpolates ${{ inputs.task-version }} inside the run: shell script in three places: in a grep string, in an echo/sed pipeline, and in a go install command. An attacker-controlled input value is substituted into the shell command before the shell parses it, enabling command injection.

Offending lines:
  if ! task --version | grep -q "Task version: v${{ inputs.task-version }}"; then
  major_version=$(echo "${{ inputs.task-version }}" | sed -E 's/^([0-9]+).*/\1/')
  go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}

Locations:

- `action.yml:124`
- `action.yml:125`
- `action.yml:126`

### script-injection (severity: high)

Rule (a): An unnamed step directly interpolates ${{ inputs.github-token-for-downloading-private-go-modules }} inside a git config URL in the run: shell script. A caller-controlled token value is substituted into the shell command before the shell parses it, enabling injection of arbitrary git config arguments or URL components.

Offending line:
  git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/

Locations:

- `action.yml:132`

### script-injection (severity: high)

Rule (a): The 'Build binary' step directly interpolates ${{ inputs.release-build-tags }} into a -tags flag and ${{ github.ref_name }} into a -ldflags argument inside the run: shell script. These caller-controlled values are substituted into shell commands before the shell parses them, enabling command injection.

Offending lines:
  if [ -n "${{ inputs.release-build-tags }}" ]; then
  -tags "${{ inputs.release-build-tags }}" \
  -ldflags="-X 'main.Version=${{ github.ref_name }}'" \

Locations:

- `action.yml:205`
- `action.yml:207`
- `action.yml:211`

### script-injection (severity: high)

Rule (a): The 'Compute asset name' step directly interpolates ${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, and ${{ inputs.release-build-tags }} inside the run: shell script. These caller-controlled values are substituted into shell commands before the shell parses them, enabling command injection.

Offending lines:
  ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"
  if [ -n "${{ inputs.release-build-tags }}" ]; then
  ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"

Locations:

- `action.yml:218`
- `action.yml:219`
- `action.yml:220`

### github-env-injection (severity: high)

The 'Compute asset name' step writes ASSET_NAME (derived from caller-controlled inputs ${{ inputs.release-application-name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, ${{ inputs.release-build-tags }}, and ${{ github.ref_name }}) to $GITHUB_OUTPUT without any sanitization (printf '%s' ... | tr -d '\n\r'). A newline character in any of these values could inject arbitrary key=value pairs into the GitHub output environment.

Offending line:
  echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT

Locations:

- `action.yml:224`

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

Fixed all script-injection, static-inline-injection, and github-env-injection findings in hardened/action/action.yml:
1. 'install task' step: Moved ${{ inputs.task-version }} to env var TASK_VERSION; run: block now uses ${TASK_VERSION}.
2. git config step: Moved ${{ inputs.github-token-for-downloading-private-go-modules }} to env var GITHUB_TOKEN_FOR_PRIVATE_MODULES; run: block now uses ${GITHUB_TOKEN_FOR_PRIVATE_MODULES}.
3. 'Build binary' step: Added RELEASE_BUILD_TAGS and RELEASE_REF_NAME env vars; run: block uses them instead of inline expressions.
4. 'Compute asset name' step: Added five env vars for all inline expressions; run: block sanitizes each with printf '%s' ... | tr -d '\n\r' before use, and writes to $GITHUB_OUTPUT using the sanitized values. Also quoted $GITHUB_OUTPUT reference.

### Iteration 2

**Fixes applied:** unsafe-shell

**Notes:**

Fixed two unsafe curl-pipe-to-shell patterns in hardened/action/build/task.yml:
1. helm-install task: Changed from `curl ... | bash -s -- --version {{.HELM_VERSION}}` to downloading the script to a temp file first with `curl -o "$HELM_INSTALL_SCRIPT"`, then executing `bash "$HELM_INSTALL_SCRIPT" --version {{.HELM_VERSION}}`. Dropped the `--` (it was the shell's stdin option terminator, not a script argument).
2. golangci-lint-install task: Changed from `curl ... | sh -s -- -b {{.GOBIN}} {{.GOLANGCI_LINT_VERSION}}` to downloading the script to a temp file first with `curl -o "$GOLANGCI_LINT_INSTALL_SCRIPT"`, then executing `sh "$GOLANGCI_LINT_INSTALL_SCRIPT" -b {{.GOBIN}} {{.GOLANGCI_LINT_VERSION}}`. Dropped the `--` (same reason). Both temp files are cleaned up with `rm -f` after execution.

### Iteration 3

**Fixes applied:** script-injection

**Notes:**

Fixed all three script-injection findings in hardened/action/action.yml:
1. Line 124 (install task step): Double-quoted TASK_VERSION in the go install command: `go install "github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"`
2. Line 130 (git config step): Double-quoted the URL containing GITHUB_TOKEN_FOR_PRIVATE_MODULES: `git config --global "url.https://${GITHUB_TOKEN_FOR_PRIVATE_MODULES}@github.com/.insteadOf" https://github.com/`
3. Line 228 (Build binary step): Double-quoted RELEASE_REF_NAME in the ldflags argument: `-ldflags="-X main.Version=\"${RELEASE_REF_NAME}\""` - changed from single-quote wrapping (which doesn't prevent injection when inside a double-quoted string) to proper double-quoting of the variable.

