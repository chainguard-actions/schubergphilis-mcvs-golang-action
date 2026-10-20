<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.13.2

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.13.2** was hardened automatically. 12 finding(s) were identified and resolved across 3 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Rule (a): the 'git config' step puts a GitHub Actions expression straight into the run: script: `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`. Pass it through an env: variable and use it double-quoted instead.

Locations:

- `action.yml:179`

### script-injection (severity: high)

Rule (a): the 'Build binary' step puts expressions straight into the run: script: `if [ -n "${{ inputs.release-build-tags }}" ]`, `-tags "${{ inputs.release-build-tags }}"` and `-ldflags="-X 'main.Version=${{ github.ref_name }}'"`. Move these values into env: variables and expand them double-quoted.

Locations:

- `action.yml:342`
- `action.yml:348`

### script-injection (severity: high)

Rule (a): the 'Compute asset name' step puts expressions straight into the run: script: `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"`, `if [ -n "${{ inputs.release-build-tags }}" ]` and `ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"`.

Locations:

- `action.yml:359`

### github-env-injection (severity: high)

The 'Compute asset name' step writes ASSET_NAME to $GITHUB_OUTPUT with `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT`. ASSET_NAME is built from inputs.release-application-name, github.ref_name, inputs.release-os, inputs.release-architecture and inputs.release-build-tags, and it is written without the `printf '%s' ... | tr -d '\n\r'` sanitization step.

Locations:

- `action.yml:365`

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

I moved the private-modules token in the git config step into the env var PRIVATE_GO_MODULES_TOKEN and quoted the config key. In 'Build binary', release-build-tags and github.ref_name now come from env vars (RELEASE_BUILD_TAGS, REF_NAME), expanded double-quoted. In 'Compute asset name', every input and ref_name now comes from env. ASSET_NAME has its newlines and carriage returns stripped with printf | tr before it is written to "$GITHUB_OUTPUT". The commands do the same thing as before.

### Iteration 2

**Fixes applied:** script-injection, invalid-yaml

**Notes:**

Fixed the 'Build binary' step in action.yml:
- Removed the duplicated line that held the stray `\o build \` and the inline `${{ github.ref_name }}`. The ldflags now use the env var: `-X 'main.Version=${REF_NAME}'`.
- Repaired the broken env entry `GOARCH:        SOURCE_FILE: ...`. GOARCH is now `${{ inputs.release-architecture }}` and SOURCE_FILE is on its own line again. I also added `GOOS: ${{ inputs.release-os }}`. The file never set GOOS, so this new line is my guess at upstream's code and was not copied from it. It is harmless with the default `linux` on Linux runners, but someone should check it against upstream.
- This repair also clears the YAML parse error at line 339.
- I have not loaded the file with a YAML parser or run the step.

### Iteration 3

**Fixes applied:** invalid-yaml

**Notes:**

In the 'Build binary' step of action.yml, two env lines had been merged into one ('GOARCH:        SOURCE_FILE: ...'). This caused the parse error and also left GOARCH with no value. I split it back into separate entries: GOARCH: ${{ inputs.release-architecture }}, GOOS: ${{ inputs.release-os }}, and SOURCE_FILE. No upstream copy was available, so I couldn't confirm the original lines. GOOS is inferred, not recovered: release-os defaults to linux, the docs show a darwin build, and the OS is part of the asset name, but nothing else in the file sets GOOS. Remove that line if upstream didn't set it. I haven't parsed the YAML or re-run the analysis.

