<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.7

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.7** was hardened automatically. 13 finding(s) were identified and resolved across 2 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Multiple run: blocks in action.yml directly interpolate ${{ inputs.* }} and ${{ github.* }} expressions into shell command strings, violating sub-rule (a). This allows an attacker who controls those inputs to inject arbitrary shell commands.

1. 'install task' step: `grep -q "Task version: v${{ inputs.task-version }}"`, `echo "${{ inputs.task-version }}"`, and `go install ...@v${{ inputs.task-version }}` — inputs.task-version is interpolated directly into the shell.

2. git config step: `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/...` — the token input is interpolated directly into the shell command.

3. 'Build binary' step: `if [ -n "${{ inputs.release-build-tags }}" ]`, `-tags "${{ inputs.release-build-tags }}"`, and `-ldflags="-X 'main.Version=${{ github.ref_name }}'"` — inputs and github context interpolated directly into shell.

4. 'Compute asset name' step: `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"` and `ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"` — multiple inputs and github context interpolated directly into shell.

Locations:

- `action.yml:113`
- `action.yml:114`
- `action.yml:115`
- `action.yml:120`
- `action.yml:253`
- `action.yml:255`
- `action.yml:259`
- `action.yml:265`
- `action.yml:267`

### github-env-injection (severity: high)

The 'Compute asset name' step writes ASSET_NAME to $GITHUB_OUTPUT without sanitization. ASSET_NAME is constructed directly from ${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, and ${{ inputs.release-build-tags }} — all untrusted inputs — interpolated directly into the shell variable. The write `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT` is not preceded by the required sanitization step (`printf '%s' ... | tr -d '\n\r'`), enabling newline injection into the GitHub output environment file.

Locations:

- `action.yml:271`

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

1. 'install task' step: Moved ${{ inputs.task-version }} to env: block as TASK_VERSION; updated run: to use ${TASK_VERSION}.

2. git config step: Moved ${{ inputs.github-token-for-downloading-private-go-modules }} to env: block as GITHUB_TOKEN_FOR_PRIVATE_MODULES; updated run: to use ${GITHUB_TOKEN_FOR_PRIVATE_MODULES}.

3. 'Build binary' step: Added RELEASE_BUILD_TAGS and RELEASE_VERSION to existing env: block; updated run: to use ${RELEASE_BUILD_TAGS} and ${RELEASE_VERSION} instead of inline ${{ }} expressions.

4. 'Compute asset name' step: Added env: block with RELEASE_APPLICATION_NAME, RELEASE_REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS; updated run: to use env vars; added sanitization via `printf '%s' "${ASSET_NAME}" | tr -d '\n\r'` before writing to $GITHUB_OUTPUT to prevent newline injection.

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed two script injection vulnerabilities in action.yml:
1. 'install task' step (line 113): Quoted the go install argument string so that TASK_VERSION and derived major_version variables cannot inject shell metacharacters: changed `go install github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}` to `go install "github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"`.
2. 'git config' step (line 117): Quoted the URL portion containing the token to prevent shell metacharacter injection: changed `url.https://${GITHUB_TOKEN_FOR_PRIVATE_MODULES}@github.com/.insteadOf` to `url."https://${GITHUB_TOKEN_FOR_PRIVATE_MODULES}@github.com/".insteadOf` (which also follows standard git config URL quoting syntax).

