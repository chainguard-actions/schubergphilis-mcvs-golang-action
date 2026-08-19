<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.11.22

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.11.22** was hardened automatically. 14 finding(s) were identified and resolved across 3 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Multiple run: blocks in action.yml directly interpolate ${{ }} expressions inside shell commands, violating rule (a). (1) 'install task' step: `${{ inputs.task-version }}` is interpolated directly into grep and go install shell commands. (2) Unnamed git config step: `${{ inputs.github-token-for-downloading-private-go-modules }}` is interpolated directly into a git config URL. (3) 'Build binary' step: `${{ inputs.release-build-tags }}` and `${{ github.ref_name }}` are interpolated directly into go build flags inside the run: script. (4) 'Compute asset name' step: `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` are all interpolated directly into shell variable assignments. Any of these inputs can contain shell metacharacters enabling command injection.

Locations:

- `action.yml:108`
- `action.yml:121`
- `action.yml:233`
- `action.yml:248`

### github-env-injection (severity: high)

The 'Compute asset name' step constructs ASSET_NAME from untrusted inputs (${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, ${{ inputs.release-build-tags }}) and writes the result to $GITHUB_OUTPUT via `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT` without first sanitizing with `printf '%s' ... | tr -d '\n\r'`. A newline embedded in any of these values could inject arbitrary key-value pairs into the GitHub output context.

Locations:

- `action.yml:255`

### unpinned-uses (severity: high)

Multiple action references use mutable tags instead of immutable 40-character commit SHAs, making the action vulnerable to supply-chain attacks if the tag is moved. Failing references: action.yml: `actions/setup-go@v6.4.0`, `svenstaro/upload-release-action@2.11.5`. .github/workflows/action-sorted-inputs.yml: `actions/checkout@v6.0.3`. .github/workflows/general.yml: `actions/checkout@v6.0.3`, `schubergphilis/mcvs-general-action@v0.5.8`. .github/workflows/golang.yml: `actions/checkout@v6.0.3`. .github/workflows/gomod-go-version-updater.yml: `schubergphilis/gomod-go-version-updater-action@v0.3.5`. .github/workflows/mcvs-pr-validation.yml: `actions/checkout@v6.0.3`, `schubergphilis/mcvs-pr-validation-action@v0.2.2`. .github/workflows/package-version-updater.yml: `actions/checkout@v6.0.3`. .github/workflows/stale.yml: `actions/stale@v10.3.0`. .github/workflows/taskfile-sorted-units.yml: `actions/checkout@v6.0.3`. .github/workflows/taskfile-without-empty-lines.yml: `actions/checkout@v6.0.3`.

Locations:

- `action.yml:88`
- `action.yml:270`
- `.github/workflows/action-sorted-inputs.yml:13`
- `.github/workflows/general.yml:17`
- `.github/workflows/general.yml:18`
- `.github/workflows/golang.yml:16`
- `.github/workflows/gomod-go-version-updater.yml:12`
- `.github/workflows/mcvs-pr-validation.yml:16`
- `.github/workflows/mcvs-pr-validation.yml:17`
- `.github/workflows/package-version-updater.yml:12`
- `.github/workflows/stale.yml:10`
- `.github/workflows/taskfile-sorted-units.yml:10`
- `.github/workflows/taskfile-without-empty-lines.yml:10`

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

**Fixes applied:** script-injection, static-inline-injection, github-env-injection, unpinned-uses

**Notes:**

Fixed all findings in action.yml and .github/workflows/ files:

1. script-injection / static-inline-injection (action.yml):
   - 'install task' step: moved inputs.task-version to env: TASK_VERSION
   - git config step: moved inputs.github-token-for-downloading-private-go-modules to env: GITHUB_TOKEN_FOR_PRIVATE_MODULES
   - 'Build binary' step: moved inputs.release-build-tags and github.ref_name to env block
   - 'Compute asset name' step: moved all 5 inline expressions to env block

2. github-env-injection (action.yml line 255):
   - 'Compute asset name' step: sanitized ASSET_NAME with `printf '%s' ... | tr -d '\n\r'` before writing to $GITHUB_OUTPUT; also quoted $GITHUB_OUTPUT

3. unpinned-uses:
   - action.yml: actions/setup-go@v6.4.0 → SHA 4a3601121dd01d1626a1e23e37211e3254c1c06c
   - action.yml: svenstaro/upload-release-action@2.11.5 → SHA 29e53e917877a24fad85510ded594ab3c9ca12de
   - All 7 workflow files: actions/checkout@v6.0.3 → SHA df4cb1c069e1874edd31b4311f1884172cec0e10
   - general.yml: schubergphilis/mcvs-general-action@v0.5.8 → SHA f52c4433add29d8eff9036bf37b5b69a2c4cf28b
   - gomod-go-version-updater.yml: schubergphilis/gomod-go-version-updater-action@v0.3.5 → SHA c12bc6c5e4ae85e82ee346115d1657e43121c195
   - mcvs-pr-validation.yml: schubergphilis/mcvs-pr-validation-action@v0.2.2 → SHA 6f153c01301ab2f92336781173b7f3a796de5f3e
   - stale.yml: actions/stale@v10.3.0 → SHA eb5cf3af3ac0a1aa4c9c45633dd1ae542a27a899

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed the 'install task' step in action.yml: added double quotes around the `go install` argument `"github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"` to prevent shell metacharacter injection from the attacker-controlled `inputs.task-version` input.

### Iteration 3

**Fixes applied:** script-injection

**Notes:**

Fixed the unquoted shell variable expansion in the git config command at action.yml line 121. Changed `git config --global url.https://${GITHUB_TOKEN_FOR_PRIVATE_MODULES}@github.com/.insteadOf https://github.com/` to `git config --global "url.https://${GITHUB_TOKEN_FOR_PRIVATE_MODULES}@github.com/.insteadOf" https://github.com/`. The double quotes around the git config key prevent shell metacharacters in the token value from being interpreted as shell commands while still allowing variable expansion.

