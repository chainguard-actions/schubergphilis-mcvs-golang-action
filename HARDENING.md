<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.12.16

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.12.16** was hardened automatically. 12 finding(s) were identified and resolved across 2 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Sub-rule (a): The unnamed `run:` step directly interpolates `${{ inputs.github-token-for-downloading-private-go-modules }}` inside a shell command string: `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`. Any value containing shell metacharacters supplied by the calling workflow will be executed by the shell before quoting can protect it.

Locations:

- `action.yml:163`

### script-injection (severity: high)

Sub-rule (a): The 'Build binary' run: block directly interpolates multiple `${{ inputs.* }}` and `${{ github.* }}` expressions inside shell commands: `if [ -n "${{ inputs.release-build-tags }}" ]`, `-tags "${{ inputs.release-build-tags }}"`, and `-ldflags="-X 'main.Version=${{ github.ref_name }}'"`. These template substitutions happen before the shell parses the command, allowing an attacker-controlled value to inject arbitrary shell commands.

Locations:

- `action.yml:243`

### script-injection (severity: high)

Sub-rule (a): The 'Compute asset name' run: block directly interpolates `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` inside shell command strings, e.g. `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"`. These are expanded by the YAML template engine before the shell sees them, enabling shell command injection.

Locations:

- `action.yml:258`

### github-env-injection (severity: high)

The 'Compute asset name' step constructs `ASSET_NAME` from values derived from `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` — all untrusted inputs — and then writes the result directly to `$GITHUB_OUTPUT` via `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT` without any sanitization (`printf '%s' ... | tr -d '\n\r'`). A newline embedded in any of these values can inject arbitrary key=value pairs into the output file.

Locations:

- `action.yml:265`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.github-token-for-downloading-private-go-modules }}" appears directly in run: block of step ""; move to env: map

Locations:

- `action.yml:179`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-build-tags }}" appears directly in run: block of step "Build binary"; move to env: map

Locations:

- `action.yml:342`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-build-tags }}" appears directly in run: block of step "Build binary"; move to env: map

Locations:

- `action.yml:344`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-application-name }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:359`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-os }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:359`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-architecture }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:359`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-build-tags }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:360`

### static-inline-injection (severity: high)

shell injection: expression "${{ inputs.release-build-tags }}" appears directly in run: block of step "Compute asset name"; move to env: map

Locations:

- `action.yml:361`

## Iteration Notes

### Iteration 1

**Fixes applied:** script-injection, static-inline-injection, github-env-injection

**Notes:**

Fixed three groups of findings in action.yml:

1. Unnamed git config step (line ~163): Moved `${{ inputs.github-token-for-downloading-private-go-modules }}` to env block as `GH_TOKEN_PRIVATE_MODULES`; shell command now uses `${GH_TOKEN_PRIVATE_MODULES}`.

2. 'Build binary' step (lines ~243-257): Moved `${{ inputs.release-build-tags }}` and `${{ github.ref_name }}` to env block as `RELEASE_BUILD_TAGS` and `RELEASE_REF_NAME`; run block references env vars.

3. 'Compute asset name' step (lines ~258-265): Moved all five inline expressions to env block; added `printf '%s' ... | tr -d '\n\r'` sanitization for each value before use; the GITHUB_OUTPUT write now uses sanitized values to prevent newline injection.

### Iteration 1

**Fixes applied:** script-injection

**Notes:**

Fixed script injection vulnerability in action.yml at line 163. The git config URL argument containing ${GH_TOKEN_PRIVATE_MODULES} was unquoted, allowing shell metacharacters in the token value to be interpreted as shell commands. Fixed by wrapping the URL argument in double quotes: `git config --global "url.https://${GH_TOKEN_PRIVATE_MODULES}@github.com/.insteadOf" https://github.com/`

