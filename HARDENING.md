<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.11

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.11** was hardened automatically. 13 finding(s) were identified and resolved across 2 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Multiple run: blocks in action.yml directly interpolate ${{ inputs.* }} and ${{ github.* }} expressions inside shell command strings (rule a). This allows an attacker who controls the inputs to inject arbitrary shell commands.

1. 'install task' step: `grep -q "Task version: v${{ inputs.task-version }}"`, `echo "${{ inputs.task-version }}"`, and `go install ...@v${{ inputs.task-version }}` — inputs.task-version is interpolated directly into shell.

2. git config step: `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/...` — token interpolated directly into a shell URL string.

3. 'Build binary' step: `if [ -n "${{ inputs.release-build-tags }}" ]`, `-tags "${{ inputs.release-build-tags }}"`, and `-ldflags="-X 'main.Version=${{ github.ref_name }}'"` — inputs and github context interpolated directly into shell.

4. 'Compute asset name' step: `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"` and `if [ -n "${{ inputs.release-build-tags }}" ]` — multiple inputs and github context interpolated directly into shell.

Locations:

- `action.yml:108`
- `action.yml:116`
- `action.yml:121`
- `action.yml:165`
- `action.yml:232`
- `action.yml:237`
- `action.yml:248`
- `action.yml:258`
- `action.yml:261`
- `action.yml:264`

### github-env-injection (severity: high)

The 'Compute asset name' run: step writes values derived from untrusted inputs (${{ inputs.release-application-name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, ${{ inputs.release-build-tags }}) and ${{ github.ref_name }} to $GITHUB_OUTPUT without the required sanitization step (printf '%s' ... | tr -d '\n\r'). An attacker-controlled input containing newlines could inject arbitrary key=value pairs into GITHUB_OUTPUT, potentially overwriting subsequent step outputs.

Offending line: `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT` where ASSET_NAME is constructed from unsanitized ${{ inputs.* }} and ${{ github.ref_name }} values.

Locations:

- `action.yml:267`

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

1. 'install task' step: Moved ${{ inputs.task-version }} to env: block as TASK_VERSION; replaced all three inline interpolations with ${TASK_VERSION}.

2. git config step: Moved ${{ inputs.github-token-for-downloading-private-go-modules }} to env: block as GITHUB_TOKEN_FOR_PRIVATE_MODULES; replaced inline interpolation in the git config URL.

3. 'Build binary' step: Added RELEASE_BUILD_TAGS and RELEASE_REF_NAME to the existing env: block; replaced ${{ inputs.release-build-tags }} and ${{ github.ref_name }} inline interpolations with env var references.

4. 'Compute asset name' step: Added full env: block for all five injected values; sanitized each with printf '%s' ... | tr -d '\n\r' before use to prevent GITHUB_OUTPUT injection; replaced all inline ${{ }} interpolations with sanitized shell variables.

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed the script injection vulnerability in the 'install task' step of action.yml (line 128). The `go install` command's module path argument was unquoted, allowing attacker-controlled `task-version` input to inject shell metacharacters via the `${major_version}` and `${TASK_VERSION}` expansions. Fixed by wrapping the entire argument in double quotes: `go install "github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"`.

