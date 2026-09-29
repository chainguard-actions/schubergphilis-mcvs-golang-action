<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.6

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.6** was hardened automatically. 13 finding(s) were identified and resolved across 2 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Rule (a): Multiple run: blocks in action.yml directly interpolate ${{ }} expressions inside shell command strings, enabling script injection. (1) The 'install task' step interpolates ${{ inputs.task-version }} directly in grep, echo, and go install commands. (2) An unnamed step interpolates ${{ inputs.github-token-for-downloading-private-go-modules }} directly in a git config URL. (3) The 'Build binary' step interpolates ${{ inputs.release-build-tags }} in a -tags flag and ${{ github.ref_name }} in -ldflags. (4) The 'Compute asset name' step interpolates ${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, and ${{ inputs.release-build-tags }} directly in shell variable assignments. All of these allow an attacker-controlled value to be parsed by the shell before quoting can protect it.

Locations:

- `action.yml:108`
- `action.yml:122`
- `action.yml:239`
- `action.yml:256`

### github-env-injection (severity: high)

The 'Compute asset name' step constructs ASSET_NAME_BASE and ASSET_NAME from unsanitized ${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, and ${{ inputs.release-build-tags }} expressions, then writes the result to $GITHUB_OUTPUT via `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT` without applying the required sanitization step (`printf '%s' ... | tr -d '\n\r'`). A newline injected into any of these inputs could allow an attacker to inject arbitrary key=value pairs into the GitHub output environment.

Locations:

- `action.yml:262`

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
1. 'install task' step: moved ${{ inputs.task-version }} to env: block as TASK_VERSION, updated all references in the run: block.
2. Unnamed git config step: moved ${{ inputs.github-token-for-downloading-private-go-modules }} to env: block as GITHUB_TOKEN_FOR_PRIVATE_MODULES, updated the git config URL to use the env var.
3. 'Build binary' step: added RELEASE_BUILD_TAGS and RELEASE_REF_NAME to the existing env: block, replaced all inline ${{ }} expressions in the run: block with env var references.
4. 'Compute asset name' step: added a full env: block with RELEASE_APPLICATION_NAME, RELEASE_REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS; replaced all inline ${{ }} expressions in the run: block; added sanitization (printf '%s' ... | tr -d '\n\r') before writing to $GITHUB_OUTPUT; quoted $GITHUB_OUTPUT.

### Iteration 1

**Fixes applied:** script-injection

**Notes:**

Fixed two script-injection findings in hardened/action/action.yml:

1. **Line 128 (install task step)**: Added a validation check before using `TASK_VERSION` in the `go install` command. The check `[[ "${TASK_VERSION}" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]` ensures the value is a valid semver string (digits and dots only), rejecting any value containing `$(...)` or other shell metacharacters.

2. **Line 306 (Build binary step)**: Added a validation check before using `RELEASE_REF_NAME` in the `-ldflags` string. The check `[[ "${RELEASE_REF_NAME}" =~ ^[a-zA-Z0-9._/+-]+$ ]]` ensures the value only contains characters valid in git ref names (alphanumeric, dots, underscores, forward slashes, hyphens, plus signs), rejecting any value containing `$(...)` or other shell metacharacters.

Both validations fail fast with a clear error message if the input is invalid, preventing any malicious command substitution from executing.

