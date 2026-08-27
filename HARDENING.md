<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.7

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.7** was hardened automatically. 13 finding(s) were identified and resolved across 2 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Multiple `${{ ... }}` expressions from `inputs.*` and `github.*` contexts are interpolated directly inside `run:` shell command strings in action.yml, violating rule (a). This allows an attacker who controls those inputs to inject arbitrary shell commands.

(1) "install task" step: `${{ inputs.task-version }}` is interpolated directly into shell commands — used inside a `grep -q` string, an `echo` pipeline, and a `go install` URL. Example offending line: `if ! task --version | grep -q "Task version: v${{ inputs.task-version }}"`

(2) Unnamed git-config step: `${{ inputs.github-token-for-downloading-private-go-modules }}` is interpolated directly into a `git config` URL: `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`

(3) "Build binary" step: `${{ inputs.release-build-tags }}` and `${{ github.ref_name }}` are interpolated directly into shell commands: `-tags "${{ inputs.release-build-tags }}"` and `-ldflags="-X 'main.Version=${{ github.ref_name }}'"`. An attacker-controlled `release-build-tags` value can break out of the quoted string.

(4) "Compute asset name" step: `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` are all interpolated directly into shell variable assignments before being written to `$GITHUB_OUTPUT`.

Locations:

- `action.yml:117`
- `action.yml:130`
- `action.yml:220`
- `action.yml:233`

### github-env-injection (severity: high)

The "Compute asset name" step in action.yml writes a value derived from multiple unsanitized `inputs.*` and `github.*` expressions to `$GITHUB_OUTPUT` without the required sanitization step (`printf '%s' ... | tr -d '\n\r'`). The variable `ASSET_NAME` is constructed by directly interpolating `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` into a shell string, then written with `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT`. A newline character embedded in any of these inputs could inject additional key=value pairs into the GitHub output environment, potentially overwriting other step outputs.

Locations:

- `action.yml:233`

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

Fixed all shell injection and github-env-injection findings in action.yml:

1. 'install task' step: Moved ${{ inputs.task-version }} to env block as TASK_VERSION; replaced all 3 inline occurrences in run: block with ${TASK_VERSION}.

2. git config step (unnamed): Moved ${{ inputs.github-token-for-downloading-private-go-modules }} to env block as GH_TOKEN_PRIVATE_MODULES; replaced inline occurrence in run: block.

3. 'Build binary' step: Added RELEASE_BUILD_TAGS and RELEASE_REF_NAME to existing env block; replaced ${{ inputs.release-build-tags }} and ${{ github.ref_name }} in run: block with env vars.

4. 'Compute asset name' step: Added full env block with RELEASE_APPLICATION_NAME, RELEASE_REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS; sanitized each value with printf '%s' ... | tr -d '\n\r' before constructing ASSET_NAME; fixed $GITHUB_OUTPUT quoting. This addresses both script-injection and github-env-injection findings.

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed two script injection vulnerabilities in hardened/action/action.yml:
1. Line 122 ('install task' step): Wrapped the `go install` argument in double quotes so that `${major_version}` and `${TASK_VERSION}` (derived from `inputs.task-version`) are properly quoted: `go install "github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"`.
2. Line 128 (git-config step): Wrapped the URL argument containing `${GH_TOKEN_PRIVATE_MODULES}` in double quotes so the variable expansion is properly quoted: `git config --global "url.https://${GH_TOKEN_PRIVATE_MODULES}@github.com/.insteadOf" https://github.com/`.

