<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.8

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.8** was hardened automatically. 13 finding(s) were identified and resolved across 3 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Rule (a) violation: Multiple `${{ ... }}` expressions are interpolated directly inside `run:` shell command strings, allowing an attacker-controlled value to execute arbitrary shell commands.

1. 'install task' step (lines ~121-123): `${{ inputs.task-version }}` is interpolated three times directly in the shell script:
   - `grep -q "Task version: v${{ inputs.task-version }}"`
   - `major_version=$(echo "${{ inputs.task-version }}" | sed -E ...)`
   - `go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}`

2. Unnamed git-config step (line ~129): `${{ inputs.github-token-for-downloading-private-go-modules }}` is interpolated directly:
   - `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`

3. 'Build binary' step (lines ~228-233): `${{ inputs.release-build-tags }}` and `${{ github.ref_name }}` are interpolated directly:
   - `if [ -n "${{ inputs.release-build-tags }}" ]`
   - `-tags "${{ inputs.release-build-tags }}"`
   - `-ldflags="-X 'main.Version=${{ github.ref_name }}'"`

4. 'Compute asset name' step (lines ~240-244): `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` are all interpolated directly:
   - `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"`
   - `ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"`

Fix: Move all these values into `env:` variables and reference them as quoted shell variables (e.g., `"$TASK_VERSION"`) inside the `run:` block.

Locations:

- `action.yml:121`
- `action.yml:122`
- `action.yml:123`
- `action.yml:129`
- `action.yml:228`
- `action.yml:230`
- `action.yml:233`
- `action.yml:240`
- `action.yml:244`

### github-env-injection (severity: high)

The 'Compute asset name' step writes `ASSET_NAME` to `$GITHUB_OUTPUT` without sanitization. The value of `ASSET_NAME` is constructed by directly interpolating `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` into a shell variable, then writing it with `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT`. An attacker-controlled input containing newlines could inject additional key=value pairs into `$GITHUB_OUTPUT`, potentially overwriting other outputs. The required sanitization step (`printf '%s' "$VAR" | tr -d '\n\r'`) is absent before the write.

Locations:

- `action.yml:247`

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

Fixed all script injection findings in action.yml:
1. 'install task' step: moved ${{ inputs.task-version }} to TASK_VERSION env var
2. git-config step: moved ${{ inputs.github-token-for-downloading-private-go-modules }} to GH_TOKEN_PRIVATE_MODULES env var
3. 'Build binary' step: moved ${{ inputs.release-build-tags }} to RELEASE_BUILD_TAGS and ${{ github.ref_name }} to REF_NAME env vars
4. 'Compute asset name' step: moved all 5 inline expressions to env vars (RELEASE_APPLICATION_NAME, REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS)
5. Fixed github-env-injection in 'Compute asset name' step by sanitizing ASSET_NAME with `printf '%s' "${ASSET_NAME}" | tr -d '\n\r'` before writing to $GITHUB_OUTPUT

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed script injection in action.yml line 126: added double quotes around the `go install` argument `"github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"`. Both `${major_version}` (derived from TASK_VERSION via sed) and `${TASK_VERSION}` (from user-supplied `inputs.task-version`) are now quoted, preventing shell metacharacters in the caller-supplied input from being interpreted by the shell.

### Iteration 3

**Fixes applied:** script-injection

**Notes:**

Fixed the script injection vulnerability in action.yml at line 129. The git config command used an unquoted shell expansion `${GH_TOKEN_PRIVATE_MODULES}` in the URL argument. Added double quotes around the URL argument: `git config --global "url.https://${GH_TOKEN_PRIVATE_MODULES}@github.com/.insteadOf" https://github.com/`. This prevents any shell metacharacters in the token value from being interpreted by the shell.

