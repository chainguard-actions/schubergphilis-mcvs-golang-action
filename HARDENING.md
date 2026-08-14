<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.2

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.2** was hardened automatically. 16 finding(s) were identified and resolved across 2 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Sub-rule (a): The 'install task' run: block directly interpolates ${{ inputs.task-version }} inside shell commands. The expression is embedded in a grep pattern string, an echo pipeline, and a go install URL — all before the shell processes the command. An attacker-controlled value for `task-version` can inject shell metacharacters. Offending lines: `grep -q "Task version: v${{ inputs.task-version }}"`, `echo "${{ inputs.task-version }}"`, and `go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}`.

Locations:

- `action.yml:114`

### script-injection (severity: high)

Sub-rule (a): A run: block directly interpolates ${{ inputs.github-token-for-downloading-private-go-modules }} inside a git config shell command: `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`. The token value is expanded by YAML template substitution before the shell sees it, allowing injection of shell metacharacters.

Locations:

- `action.yml:127`

### script-injection (severity: high)

Sub-rule (a): The 'Build binary' run: block directly interpolates ${{ inputs.release-build-tags }} and ${{ github.ref_name }} inside shell commands passed to `go build`. Offending lines: `if [ -n "${{ inputs.release-build-tags }}" ]`, `-tags "${{ inputs.release-build-tags }}"`, and `-ldflags="-X 'main.Version=${{ github.ref_name }}'"`. Attacker-controlled input values are expanded by YAML before the shell processes them.

Locations:

- `action.yml:243`

### script-injection (severity: high)

Sub-rule (a): The 'Compute asset name' run: block directly interpolates multiple ${{ inputs.* }} and ${{ github.ref_name }} expressions inside shell variable assignments: `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"` and `if [ -n "${{ inputs.release-build-tags }}" ]`. These values are YAML-expanded before the shell sees them, enabling shell injection.

Locations:

- `action.yml:257`

### github-env-injection (severity: high)

The 'Compute asset name' run: block writes a value derived from untrusted inputs (${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, ${{ inputs.release-build-tags }}) to $GITHUB_OUTPUT via `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT` without first sanitizing with `printf '%s' ... | tr -d '\n\r'`. A newline embedded in any of these values could inject additional key=value pairs into the output file.

Locations:

- `action.yml:264`

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

Fixed all script injection and github-env-injection findings in action.yml:
1. 'install task' step: moved ${{ inputs.task-version }} to env: block as TASK_VERSION; shell script now uses ${TASK_VERSION}.
2. git config step: moved ${{ inputs.github-token-for-downloading-private-go-modules }} to env: block as GH_PRIVATE_TOKEN; shell script now uses ${GH_PRIVATE_TOKEN}.
3. 'Build binary' step: moved ${{ inputs.release-build-tags }} and ${{ github.ref_name }} to env: block as RELEASE_BUILD_TAGS and REF_NAME; shell script now uses these env vars.
4. 'Compute asset name' step: moved all five ${{ inputs.* }} and ${{ github.ref_name }} expressions to env: block; shell script uses env vars; added printf/tr sanitization before writing to $GITHUB_OUTPUT to prevent newline injection; quoted $GITHUB_OUTPUT.

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed the script injection vulnerability in action.yml at the 'install task' / private-modules step. The git config command argument `url.https://${GH_PRIVATE_TOKEN}@github.com/.insteadOf` was unquoted, allowing a token value containing shell metacharacters to cause word splitting or command injection. Added double quotes around the argument: `git config --global "url.https://${GH_PRIVATE_TOKEN}@github.com/.insteadOf" https://github.com/`.

