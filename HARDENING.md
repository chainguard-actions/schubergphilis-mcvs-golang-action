<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.8

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.8** was hardened automatically. 16 finding(s) were identified and resolved across 2 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Sub-rule (a): The 'install task' step directly interpolates `${{ inputs.task-version }}` inside the `run:` shell script in three places: (1) `grep -q "Task version: v${{ inputs.task-version }}"`, (2) `$(echo "${{ inputs.task-version }}" | sed -E ...)`, and (3) `go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}`. An attacker-controlled `task-version` input can inject shell metacharacters.

Locations:

- `action.yml:123`

### script-injection (severity: high)

Sub-rule (a): An unnamed step directly interpolates `${{ inputs.github-token-for-downloading-private-go-modules }}` inside the `run:` shell command: `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`. A caller-supplied token value containing shell metacharacters or newlines could alter the command.

Locations:

- `action.yml:133`

### script-injection (severity: high)

Sub-rule (a): The 'Build binary' step directly interpolates `${{ inputs.release-build-tags }}` in `if [ -n "${{ inputs.release-build-tags }}" ]` and `-tags "${{ inputs.release-build-tags }}"`, and interpolates `${{ github.ref_name }}` in `-ldflags="-X 'main.Version=${{ github.ref_name }}'"`. These are direct expression interpolations inside a `run:` shell script, allowing injection of shell metacharacters.

Locations:

- `action.yml:228`

### script-injection (severity: high)

Sub-rule (a): The 'Compute asset name' step directly interpolates multiple untrusted expressions inside the `run:` shell script: `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"` and `ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"`. All of these are direct expression interpolations that allow shell metacharacter injection.

Locations:

- `action.yml:243`

### github-env-injection (severity: high)

The 'Compute asset name' step constructs `ASSET_NAME` from untrusted inputs (`${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, `${{ inputs.release-build-tags }}`) directly in the shell, then writes it to `$GITHUB_OUTPUT` via `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT` without the required sanitization step (`printf '%s' ... | tr -d '\n\r'`). A newline embedded in any of these inputs could inject arbitrary output variables.

Locations:

- `action.yml:249`

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
1. 'install task' step: moved ${{ inputs.task-version }} to env: TASK_VERSION; replaced all 3 inline occurrences with ${TASK_VERSION}.
2. Unnamed git config step: moved ${{ inputs.github-token-for-downloading-private-go-modules }} to env: GITHUB_TOKEN_FOR_PRIVATE_MODULES; replaced inline expression with ${GITHUB_TOKEN_FOR_PRIVATE_MODULES}.
3. 'Build binary' step: added RELEASE_BUILD_TAGS and REF_NAME to env: block; replaced ${{ inputs.release-build-tags }} and ${{ github.ref_name }} in run: block.
4. 'Compute asset name' step: added RELEASE_APPLICATION_NAME, REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS to env: block; sanitized each value with printf '%s' ... | tr -d '\n\r' before constructing ASSET_NAME and writing to $GITHUB_OUTPUT (fixing github-env-injection); also quoted $GITHUB_OUTPUT.

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed two script injection vulnerabilities in action.yml:
1. Line 128 ('install task' step): Wrapped the go install module path in double quotes — `go install "github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"` — so that TASK_VERSION and major_version expansions are properly quoted against shell metacharacter injection.
2. Line 136 (private-modules git-config step): Wrapped the git config URL argument in double quotes — `git config --global "url.https://${GITHUB_TOKEN_FOR_PRIVATE_MODULES}@github.com/.insteadOf" https://github.com/` — so that GITHUB_TOKEN_FOR_PRIVATE_MODULES is properly quoted against shell metacharacter injection.

