<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.11.23

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.11.23** was hardened automatically. 14 finding(s) were identified and resolved across 2 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Multiple ${{ }} expressions are directly interpolated inside run: shell command strings in action.yml, violating sub-rule (a). This allows an attacker who controls the calling workflow's inputs or github context to inject arbitrary shell commands.

1. 'install task' step: `${{ inputs.task-version }}` is interpolated 3 times directly in the run: script — in a grep pattern, an echo pipeline, and a go install URL.
2. 'git config' step (unnamed): `${{ inputs.github-token-for-downloading-private-go-modules }}` is interpolated directly into a git config URL in the run: script.
3. 'Build binary' step: `${{ inputs.release-build-tags }}` is used in an `if [ -n "${{ inputs.release-build-tags }}" ]` test and as a `-tags` argument; `${{ github.ref_name }}` is interpolated into `-ldflags`.
4. 'Compute asset name' step: `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` are all interpolated directly into shell variable assignments and an `if [ -n ... ]` test.

All of these should be moved to env: variables and then referenced as quoted shell variables (e.g., "$VAR").

Locations:

- `action.yml:117`
- `action.yml:118`
- `action.yml:119`
- `action.yml:125`
- `action.yml:243`
- `action.yml:244`
- `action.yml:247`
- `action.yml:258`
- `action.yml:260`
- `action.yml:262`

### github-env-injection (severity: high)

The 'Compute asset name' step writes the value of ASSET_NAME to $GITHUB_OUTPUT without sanitization. ASSET_NAME is constructed by directly interpolating ${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, and ${{ inputs.release-build-tags }} into a shell variable. Because these values come from untrusted inputs and the github context, a newline character embedded in any of them could inject additional key=value pairs into $GITHUB_OUTPUT. The required sanitization step (`printf '%s' "$VAR" | tr -d '\n\r'`) is absent before the `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT` write.

Locations:

- `action.yml:264`

### unpinned-uses (severity: high)

The 'Upload binaries to release' step uses `svenstaro/upload-release-action@2.11.5`, which is pinned to a mutable version tag rather than an immutable 40-character commit SHA. If the tag is moved or the repository is compromised, the action could execute arbitrary code in the runner. It should be pinned to a full SHA, e.g. `svenstaro/upload-release-action@<40-char-sha> # 2.11.5`.

Locations:

- `action.yml:272`

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

**Fixes applied:** script-injection, static-inline-injection, github-env-injection, unpinned-uses

**Notes:**

Fixed all findings in action.yml:
1. 'install task' step: moved ${{ inputs.task-version }} to env: TASK_VERSION, referenced as ${TASK_VERSION} in run: block.
2. git config step: moved ${{ inputs.github-token-for-downloading-private-go-modules }} to env: GITHUB_TOKEN_FOR_PRIVATE_MODULES, referenced as ${GITHUB_TOKEN_FOR_PRIVATE_MODULES} in run: block.
3. 'Build binary' step: added RELEASE_BUILD_TAGS and RELEASE_REF_NAME to env: block, replaced all inline ${{ }} expressions in run: with shell variable references.
4. 'Compute asset name' step: moved all five inline expressions to env: block (RELEASE_APPLICATION_NAME, RELEASE_REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS), replaced inline expressions in run: with shell variables, and added sanitization (printf '%s' | tr -d '\n\r') before writing to $GITHUB_OUTPUT.
5. Pinned svenstaro/upload-release-action@2.11.5 to immutable SHA @29e53e917877a24fad85510ded594ab3c9ca12de with tag preserved as comment.

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed script injection in the 'install task' step of action.yml (line 110). The `go install` command's URL argument was unquoted, allowing attacker-controlled `task-version` input (exposed as `TASK_VERSION`) and the derived `major_version` variable to inject shell metacharacters. Fixed by wrapping the entire URL in double quotes: `go install "github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"`.

