<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.11.23

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.11.23** was hardened automatically. 17 finding(s) were identified and resolved across 3 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Sub-rule (a): The 'install task' run: block directly interpolates ${{ inputs.task-version }} inside shell commands without routing through an env: variable. This allows an attacker-controlled input to inject arbitrary shell commands. Offending lines:
  `if ! task --version | grep -q "Task version: v${{ inputs.task-version }}"`
  `major_version=$(echo "${{ inputs.task-version }}" | sed -E 's/^([0-9]+).*/\1/')`
  `go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}`

Locations:

- `action.yml:101`

### script-injection (severity: high)

Sub-rule (a): An unnamed run: block directly interpolates ${{ inputs.github-token-for-downloading-private-go-modules }} inside a git config shell command. A caller-controlled token value containing shell metacharacters could inject arbitrary commands. Offending line:
  `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`

Locations:

- `action.yml:113`

### script-injection (severity: high)

Sub-rule (a): The 'Build binary' run: block directly interpolates ${{ inputs.release-build-tags }} and ${{ github.ref_name }} inside shell commands. These expressions are expanded by the YAML template engine before the shell sees them, allowing injection of arbitrary shell commands. Offending lines:
  `if [ -n "${{ inputs.release-build-tags }}" ]`
  `-tags "${{ inputs.release-build-tags }}"`
  `-ldflags="-X 'main.Version=${{ github.ref_name }}'"` 

Locations:

- `action.yml:222`

### script-injection (severity: high)

Sub-rule (a): The 'Compute asset name' run: block directly interpolates ${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, and ${{ inputs.release-build-tags }} inside shell commands. These expressions are expanded before the shell sees them, allowing injection of arbitrary shell commands. Offending lines:
  `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"`
  `ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"`

Locations:

- `action.yml:237`

### github-env-injection (severity: high)

The 'Compute asset name' run: block constructs ASSET_NAME from ${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, and ${{ inputs.release-build-tags }} — all untrusted inputs — and writes the result directly to $GITHUB_OUTPUT without the required sanitization step (printf '%s' ... | tr -d '\n\r'). A newline embedded in any of these values could inject arbitrary key=value pairs into GITHUB_OUTPUT, poisoning subsequent steps. Offending line:
  `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT`

Locations:

- `action.yml:244`

### unpinned-uses (severity: high)

The 'Upload binaries to release' step uses svenstaro/upload-release-action@2.11.5, which is pinned to a mutable version tag rather than an immutable full 40-character commit SHA. If the tag is moved or the repository is compromised, the action could execute arbitrary malicious code. It should be pinned to a full SHA, e.g. svenstaro/upload-release-action@<40-char-sha> # 2.11.5.

Locations:

- `action.yml:252`

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

**Fixes applied:** script-injection, github-env-injection, unpinned-uses, static-inline-injection

**Notes:**

Fixed all 17 findings in action.yml:
1. 'install task' step: moved ${{ inputs.task-version }} to TASK_VERSION env var; shell now uses ${TASK_VERSION} throughout.
2. Unnamed git config step: moved ${{ inputs.github-token-for-downloading-private-go-modules }} to GH_TOKEN_PRIVATE_MODULES env var; shell uses ${GH_TOKEN_PRIVATE_MODULES}.
3. 'Build binary' step: moved ${{ inputs.release-build-tags }} to RELEASE_BUILD_TAGS and ${{ github.ref_name }} to REF_NAME env vars; shell uses env vars.
4. 'Compute asset name' step: moved all 5 expressions to env vars (RELEASE_APPLICATION_NAME, REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS); each sanitized with printf '%s' ... | tr -d '\n\r' before use; GITHUB_OUTPUT written with printf to prevent newline injection.
5. 'Upload binaries to release': pinned svenstaro/upload-release-action from mutable tag @2.11.5 to immutable SHA @29e53e917877a24fad85510ded594ab3c9ca12de # 2.11.5.

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed the unquoted shell variable expansion in the 'install task' step at action.yml line 128. The `go install` argument was changed from `github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}` (unquoted) to `"github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"` (double-quoted), ensuring that both `${major_version}` and `${TASK_VERSION}` are properly quoted and cannot be exploited via shell metacharacters in the caller-controlled `task-version` input.

### Iteration 3

**Fixes applied:** script-injection

**Notes:**

Fixed the script injection vulnerability in action.yml at line 133. The git config command's URL argument containing ${GH_TOKEN_PRIVATE_MODULES} was not double-quoted, allowing shell metacharacters in the token value to be interpreted. Fixed by wrapping the URL in double quotes: `url."https://${GH_TOKEN_PRIVATE_MODULES}@github.com/".insteadOf` instead of `url.https://${GH_TOKEN_PRIVATE_MODULES}@github.com/.insteadOf`. The token is already correctly sourced from the env block rather than directly from ${{ inputs.* }}, so only the quoting fix was needed.

