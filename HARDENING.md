<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.13.3

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.13.3** was hardened automatically. 12 finding(s) were identified and resolved across 1 iteration(s).

## Findings Fixed

### script-injection (severity: high)

(a) Direct expression interpolation in a run: block: `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`. The input should be passed through an env: variable and quoted in the script.

Locations:

- `action.yml:185`

### script-injection (severity: high)

(a) Direct expression interpolation in the 'Build binary' run: block: `if [ -n "${{ inputs.release-build-tags }}" ]`, `-tags "${{ inputs.release-build-tags }}"`, and `-ldflags="-X 'main.Version=${{ github.ref_name }}'"`.

Locations:

- `action.yml:364`

### script-injection (severity: high)

(a) Direct expression interpolation in the 'Compute asset name' run: block: `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"` and `${{ inputs.release-build-tags }}`.

Locations:

- `action.yml:380`

### github-env-injection (severity: high)

The 'Compute asset name' step builds ASSET_NAME from inputs.* and github.ref_name and writes it with `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT`. There is no `printf '%s' ... | tr -d '\n\r'` sanitization step before the write.

Locations:

- `action.yml:386`

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

**Fixes applied:** script-injection, github-env-injection, static-inline-injection

**Notes:**

In action.yml, the private-modules token for git config now comes from an env: variable (PRIVATE_GO_MODULES_TOKEN) and the git config key is quoted. In the 'Build binary' step, release-build-tags and github.ref_name are now env vars (RELEASE_BUILD_TAGS, REF_NAME) and stay quoted the way upstream had them. In the 'Compute asset name' step, all five inputs are now env vars, and asset_name passes through printf | tr -d '\n\r' before it is written to the quoted "$GITHUB_OUTPUT".

