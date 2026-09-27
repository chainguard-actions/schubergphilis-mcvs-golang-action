<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.11.21

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.11.21** was hardened automatically. 14 finding(s) were identified and resolved across 2 iteration(s).

## Findings Fixed

### unpinned-uses (severity: high)

action.yml contains two `uses:` references pinned to mutable version tags instead of immutable 40-character commit SHAs, making the action vulnerable to supply-chain attacks if those tags are moved:
- `actions/setup-go@v6.4.0` (line 84)
- `svenstaro/upload-release-action@2.11.5` (line 257)

The third reference `anchore/scan-action@e1165082ffb1fe366ebaf02d8526e7c4989ea9d2` is correctly pinned to a SHA.

Locations:

- `action.yml:84`
- `action.yml:257`

### script-injection (severity: high)

Multiple `run:` blocks in action.yml directly interpolate `${{ inputs.* }}` and `${{ github.* }}` expressions inside shell command strings (sub-rule a). This allows an attacker who controls those inputs to inject arbitrary shell commands.

1. **"install task" step (line ~108):** `${{ inputs.task-version }}` is interpolated three times directly in the shell:
   - `grep -q "Task version: v${{ inputs.task-version }}"`
   - `major_version=$(echo "${{ inputs.task-version }}" | sed -E ...)`
   - `go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}`

2. **git config step (line ~115):** `${{ inputs.github-token-for-downloading-private-go-modules }}` is interpolated directly into a `git config` URL in the shell.

3. **"Build binary" step (line ~232):** `${{ inputs.release-build-tags }}` and `${{ github.ref_name }}` are interpolated directly inside the shell `run:` block:
   - `if [ -n "${{ inputs.release-build-tags }}" ]`
   - `-tags "${{ inputs.release-build-tags }}"`
   - `-ldflags="-X 'main.Version=${{ github.ref_name }}'"`

4. **"Compute asset name" step (line ~243):** `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` are all interpolated directly in the shell.

Locations:

- `action.yml:108`
- `action.yml:115`
- `action.yml:232`
- `action.yml:243`

### github-env-injection (severity: high)

The "Compute asset name" step in action.yml writes a value derived from untrusted inputs directly to `$GITHUB_OUTPUT` without sanitization. The shell variable `ASSET_NAME` is constructed by directly interpolating `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` — all of which are attacker-controllable. The value is then written with:

```
echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT
```

No `printf '%s' ... | tr -d '\n\r'` sanitization is applied before the write, allowing newline injection into the GitHub output file, which can be used to set arbitrary environment variables or outputs in subsequent steps.

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

**Fixes applied:** unpinned-uses, script-injection, github-env-injection, static-inline-injection

**Notes:**

Fixed all findings in hardened/action/action.yml:
1. Pinned actions/setup-go@v6.4.0 to SHA 4a3601121dd01d1626a1e23e37211e3254c1c06c.
2. Pinned svenstaro/upload-release-action@2.11.5 to SHA 29e53e917877a24fad85510ded594ab3c9ca12de.
3. Moved ${{ inputs.task-version }} to env: TASK_VERSION in the 'install task' step.
4. Moved ${{ inputs.github-token-for-downloading-private-go-modules }} to env: GH_TOKEN_PRIVATE_MODULES in the git config step.
5. Moved ${{ inputs.release-build-tags }} and ${{ github.ref_name }} to env: block in the 'Build binary' step.
6. Moved all ${{ inputs.* }} and ${{ github.* }} expressions to env: block in the 'Compute asset name' step, and sanitized the ASSET_NAME with tr -d '\n\r' before writing to $GITHUB_OUTPUT.

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed two script-injection findings in hardened/action/action.yml:
1. Line 115 ('install task' step): Added double quotes around the go install package path argument: `go install "github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"`. This prevents shell metacharacters in TASK_VERSION (sourced from inputs.task-version) or the derived major_version from being interpreted as shell commands.
2. Line 127 (git config step): Added double quotes around the git config key argument: `git config --global "url.https://${GH_TOKEN_PRIVATE_MODULES}@github.com/.insteadOf" https://github.com/`. This prevents shell metacharacters in GH_TOKEN_PRIVATE_MODULES (sourced from inputs.github-token-for-downloading-private-go-modules) from being interpreted as shell commands. The second argument (https://github.com/) is a literal string and does not need quoting.

