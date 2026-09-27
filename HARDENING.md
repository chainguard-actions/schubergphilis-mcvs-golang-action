<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.6

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.6** was hardened automatically. 13 finding(s) were identified and resolved across 1 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Multiple `run:` blocks directly interpolate `${{ ... }}` expressions inside shell commands (rule a), allowing an attacker to inject arbitrary shell code via controlled inputs or GitHub context values.

1. **"install task" step** (line ~107): `${{ inputs.task-version }}` is interpolated directly into a `grep -q`, an `echo`, and a `go install` URL — e.g. `go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}`.

2. **git config step** (line ~117): `${{ inputs.github-token-for-downloading-private-go-modules }}` is embedded directly in a shell URL string: `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`.

3. **"Build binary" step** (line ~248): `${{ inputs.release-build-tags }}` appears unquoted inside `[ -n "${{ inputs.release-build-tags }}" ]` and in `-tags "${{ inputs.release-build-tags }}"`. `${{ github.ref_name }}` is interpolated into `-ldflags="-X 'main.Version=${{ github.ref_name }}'"` — `github.ref_name` is attacker-controlled via branch/tag names.

4. **"Compute asset name" step** (line ~261): `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` are all interpolated directly into shell variable assignments and a `[ -n ... ]` test.

Locations:

- `action.yml:107`
- `action.yml:117`
- `action.yml:248`
- `action.yml:261`

### github-env-injection (severity: high)

The **"Compute asset name"** step writes `ASSET_NAME` to `$GITHUB_OUTPUT` without sanitization. The value is constructed by directly interpolating `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` into a shell variable, then writing it with `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT`. Because `github.ref_name` and all `inputs.*` values are attacker-controllable and none are passed through `printf '%s' ... | tr -d '\n\r'` before the write, a newline embedded in any of these values can inject arbitrary key=value pairs into the GitHub output environment.

Locations:

- `action.yml:270`

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

Fixed all script-injection and github-env-injection findings in action.yml:
1. 'install task' step: moved ${{ inputs.task-version }} to env: TASK_VERSION; replaced all 3 inline occurrences with ${TASK_VERSION}.
2. git config step: moved ${{ inputs.github-token-for-downloading-private-go-modules }} to env: PRIVATE_GO_TOKEN; replaced inline occurrence with ${PRIVATE_GO_TOKEN}.
3. 'Build binary' step: added RELEASE_BUILD_TAGS and REF_NAME to existing env: block; replaced ${{ inputs.release-build-tags }} and ${{ github.ref_name }} in run: block.
4. 'Compute asset name' step: added full env: block with RELEASE_APPLICATION_NAME, REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS; replaced all inline expressions; added printf '%s' ... | tr -d '\n\r' sanitization for all values before constructing ASSET_NAME and writing to $GITHUB_OUTPUT to prevent newline injection.

