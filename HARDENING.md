<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.11

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.11** was hardened automatically. 13 finding(s) were identified and resolved across 2 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Multiple `run:` blocks in action.yml directly interpolate `${{ ... }}` expressions inside shell command strings (sub-rule a), allowing an attacker-controlled value to be parsed as shell syntax before the shell ever sees it.

1. **"install task" step** (line ~119): `${{ inputs.task-version }}` is interpolated three times directly in the shell:
   - `grep -q "Task version: v${{ inputs.task-version }}"`
   - `major_version=$(echo "${{ inputs.task-version }}" | sed -E ...)`
   - `go install ...@v${{ inputs.task-version }}`

2. **git config step** (line ~126): `${{ inputs.github-token-for-downloading-private-go-modules }}` is interpolated directly in the shell URL:
   - `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/...`

3. **"Build binary" step** (lines ~224-232): `${{ inputs.release-build-tags }}` and `${{ github.ref_name }}` are interpolated directly in shell commands:
   - `if [ -n "${{ inputs.release-build-tags }}" ]`
   - `-tags "${{ inputs.release-build-tags }}"`
   - `-ldflags="-X 'main.Version=${{ github.ref_name }}'"`

4. **"Compute asset name" step** (lines ~240-247): Multiple expressions interpolated directly in shell:
   - `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"`
   - `if [ -n "${{ inputs.release-build-tags }}" ]`
   - `ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"`

All of these should be moved to `env:` variables and referenced as `"$VAR"` in the shell.

Locations:

- `action.yml:119`
- `action.yml:126`
- `action.yml:224`
- `action.yml:240`

### github-env-injection (severity: high)

The **"Compute asset name"** step writes to `$GITHUB_OUTPUT` using a value (`ASSET_NAME`) that is built directly from `${{ inputs.* }}` and `${{ github.ref_name }}` expressions interpolated in the shell, without any sanitization (`printf '%s' ... | tr -d '\n\r'`) before the write:

```yaml
run: |
  ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"
  if [ -n "${{ inputs.release-build-tags }}" ]; then
    ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"
  else
    ASSET_NAME="${ASSET_NAME_BASE}"
  fi
  echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT
```

An attacker-controlled input (e.g. `inputs.release-application-name` or `inputs.release-build-tags`) containing newline characters could inject arbitrary key-value pairs into `$GITHUB_OUTPUT`, potentially overwriting other outputs or injecting environment variables consumed by downstream steps. The sanitization step `safe=$(printf '%s' "$VAR" | tr -d '\n\r')` must be applied to each untrusted value before the write.

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

Fixed all script injection and github-env-injection findings in action.yml:
1. 'install task' step: Added env block with TASK_VERSION=${{ inputs.task-version }} and replaced all three inline ${{ inputs.task-version }} occurrences with ${TASK_VERSION}.
2. git config step: Added env block with GH_TOKEN_PRIVATE_MODULES=${{ inputs.github-token-for-downloading-private-go-modules }} and replaced the inline token in the URL.
3. 'Build binary' step: Added RELEASE_BUILD_TAGS=${{ inputs.release-build-tags }} and REF_NAME=${{ github.ref_name }} to the existing env block, replaced all inline expressions in the run: script.
4. 'Compute asset name' step: Added env block with all five expressions (RELEASE_APPLICATION_NAME, REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS), replaced all inline expressions, and added printf '%s' ... | tr -d '\n\r' sanitization for each value before building ASSET_NAME and writing to $GITHUB_OUTPUT.

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed two script injection vulnerabilities in action.yml:
1. Line 128 (install task step): Quoted the go install module path argument: `go install "github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"`. This prevents shell metacharacters in the caller-controlled `task-version` input from being interpreted by the shell.
2. Line 134 (git-config step): Quoted the URL argument in the git config command: `git config --global url."https://${GH_TOKEN_PRIVATE_MODULES}@github.com/".insteadOf https://github.com/`. This prevents shell metacharacters in the caller-controlled `github-token-for-downloading-private-go-modules` input from being interpreted by the shell.

