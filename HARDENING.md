<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.11.21

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `1`

Action **schubergphilis--mcvs-golang-action/v3.11.21** was hardened automatically. 14 finding(s) were identified and resolved across 1 iteration(s).

## Findings Fixed

### unpinned-uses (severity: high)

action.yml contains two unpinned `uses:` references that use mutable tags instead of full 40-character commit SHAs:
- `actions/setup-go@v6.4.0` (tag reference)
- `svenstaro/upload-release-action@2.11.5` (version tag reference)
These can be silently updated by the upstream repository, enabling supply-chain attacks. The third reference (`anchore/scan-action@e1165082ffb1fe366ebaf02d8526e7c4989ea9d2`) is correctly pinned to a SHA.

Locations:

- `action.yml:75`
- `action.yml:244`

### script-injection (severity: high)

Multiple `run:` blocks in action.yml interpolate `${{ ... }}` expressions directly into shell command strings (sub-rule a), allowing an attacker-controlled value to inject arbitrary shell commands:

1. **"install task" step** (lines ~101-103): `${{ inputs.task-version }}` is interpolated directly into a grep pattern, an echo argument, and a `go install` URL:
   - `grep -q "Task version: v${{ inputs.task-version }}"`
   - `echo "${{ inputs.task-version }}" | sed -E ...`
   - `go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}`

2. **git config step** (line ~108): `${{ inputs.github-token-for-downloading-private-go-modules }}` is interpolated directly into a git config URL:
   - `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/...`

3. **"Build binary" step** (lines ~222-229): `${{ inputs.release-build-tags }}` and `${{ github.ref_name }}` are interpolated directly into shell conditionals and compiler flags:
   - `if [ -n "${{ inputs.release-build-tags }}" ]`
   - `-tags "${{ inputs.release-build-tags }}"`
   - `-ldflags="-X 'main.Version=${{ github.ref_name }}'"`

4. **"Compute asset name" step** (lines ~234-239): Multiple inputs and github context values are interpolated directly into shell variable assignments:
   - `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"`
   - `ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"`

Locations:

- `action.yml:101`
- `action.yml:108`
- `action.yml:222`
- `action.yml:234`

### github-env-injection (severity: high)

The "Compute asset name" step writes `ASSET_NAME` to `$GITHUB_OUTPUT` without sanitization. `ASSET_NAME` is derived entirely from untrusted inputs (`${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, `${{ inputs.release-build-tags }}`). A newline character embedded in any of these values could inject additional key=value pairs into `$GITHUB_OUTPUT`, potentially overwriting other step outputs. The required sanitization (`printf '%s' "$VAR" | tr -d '\n\r'`) is absent before the write:
  `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT`

Locations:

- `action.yml:240`

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

**Fixes applied:** unpinned-uses, script-injection, github-env-injection, static-inline-injection

**Notes:**

Fixed all findings in action.yml:
1. Pinned actions/setup-go@v6.4.0 to SHA 4a3601121dd01d1626a1e23e37211e3254c1c06c
2. Pinned svenstaro/upload-release-action@2.11.5 to SHA 29e53e917877a24fad85510ded594ab3c9ca12de
3. 'install task' step: moved ${{ inputs.task-version }} to TASK_VERSION env var, replaced all inline expressions with ${TASK_VERSION}
4. git config step: moved ${{ inputs.github-token-for-downloading-private-go-modules }} to GITHUB_TOKEN_FOR_PRIVATE_MODULES env var
5. 'Build binary' step: moved ${{ inputs.release-build-tags }} and ${{ github.ref_name }} to RELEASE_BUILD_TAGS and REF_NAME env vars
6. 'Compute asset name' step: moved all five ${{ }} expressions to env vars (RELEASE_APPLICATION_NAME, REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS) and added sanitization with `printf '%s' | tr -d '\n\r'` before writing to $GITHUB_OUTPUT

