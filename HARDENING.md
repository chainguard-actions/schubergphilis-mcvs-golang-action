<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.11.22

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.11.22** was hardened automatically. 14 finding(s) were identified and resolved across 3 iteration(s).

## Findings Fixed

### unpinned-uses (severity: high)

Two `uses:` references in action.yml are pinned to mutable version tags rather than immutable 40-character commit SHAs, making the action vulnerable to supply-chain attacks if those tags are moved:
- `actions/setup-go@v6.4.0` (tag, not a SHA)
- `svenstaro/upload-release-action@2.11.5` (version tag, not a SHA)

The third reference (`anchore/scan-action@e1165082ffb1fe366ebaf02d8526e7c4989ea9d2`) is correctly pinned to a SHA.

Locations:

- `action.yml:79`
- `action.yml:233`

### script-injection (severity: high)

Multiple `run:` blocks in action.yml interpolate GitHub Actions expressions (`${{ ... }}`) directly into shell command strings, enabling script injection. An attacker who controls these input values can inject arbitrary shell commands.

**Sub-rule (a) violations — direct expression interpolation in run: blocks:**

1. `install task` step: `${{ inputs.task-version }}` is interpolated directly in the shell:
   ```
   if ! task --version | grep -q "Task version: v${{ inputs.task-version }}"; then
     major_version=$(echo "${{ inputs.task-version }}" | sed -E ...)
     go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}
   ```

2. `git config` step: `${{ inputs.github-token-for-downloading-private-go-modules }}` is interpolated directly in the shell:
   ```
   git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/
   ```

3. `Build binary` step: `${{ inputs.release-build-tags }}` and `${{ github.ref_name }}` are interpolated directly in the shell:
   ```
   if [ -n "${{ inputs.release-build-tags }}" ]; then
     go build -tags "${{ inputs.release-build-tags }}" ...
   else
     go build -ldflags="-X 'main.Version=${{ github.ref_name }}'" ...
   ```

4. `Compute asset name` step: `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` are all interpolated directly in the shell:
   ```
   ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"
   if [ -n "${{ inputs.release-build-tags }}" ]; then
     ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"
   ```

All these values should be passed via `env:` variables and then referenced as quoted shell variables (e.g. `"$INPUT_TASK_VERSION"`) instead.

Locations:

- `action.yml:101`
- `action.yml:116`
- `action.yml:196`
- `action.yml:213`

### github-env-injection (severity: high)

The `Compute asset name` step writes `${ASSET_NAME}` to `$GITHUB_OUTPUT` without sanitizing the value first. `ASSET_NAME` is constructed directly from unsanitized `${{ inputs.* }}` and `${{ github.ref_name }}` expressions interpolated into the shell. A newline character in any of these values could inject arbitrary key=value pairs into `$GITHUB_OUTPUT`, poisoning downstream steps.

Offending code:
```yaml
run: |
  ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"
  if [ -n "${{ inputs.release-build-tags }}" ]; then
    ASSET_NAME="${ASSET_NAME_BASE}-${{ inputs.release-build-tags }}"
  else
    ASSET_NAME="${ASSET_NAME_BASE}"
  fi
  echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT
```

The fix requires moving all inputs into `env:` variables, then sanitizing each with `printf '%s' "$VAR" | tr -d '\n\r'` before writing to `$GITHUB_OUTPUT`.

Locations:

- `action.yml:213`

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
1. Pinned actions/setup-go@v6.4.0 → @4a3601121dd01d1626a1e23e37211e3254c1c06c # v6.4.0
2. Pinned svenstaro/upload-release-action@2.11.5 → @29e53e917877a24fad85510ded594ab3c9ca12de # 2.11.5
3. 'install task' step: moved ${{ inputs.task-version }} to TASK_VERSION env var
4. 'git config' step: moved ${{ inputs.github-token-for-downloading-private-go-modules }} to GH_TOKEN_PRIVATE_MODULES env var
5. 'Build binary' step: moved ${{ inputs.release-build-tags }} to RELEASE_BUILD_TAGS and ${{ github.ref_name }} to REF_NAME env vars
6. 'Compute asset name' step: moved all ${{ }} expressions to env vars (RELEASE_APPLICATION_NAME, REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS) and sanitized each with printf '%s' | tr -d '\n\r' before writing to $GITHUB_OUTPUT

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed two script injection findings in hardened/action/action.yml:
1. Line ~110 ('install task' step): Added double quotes around ${TASK_VERSION} in the go install URL: `go install github.com/go-task/task/v${major_version}/cmd/task@v"${TASK_VERSION}"`
2. Line ~120 (private-modules step): Added double quotes around ${GH_TOKEN_PRIVATE_MODULES} in the git config command: `git config --global url.https://"${GH_TOKEN_PRIVATE_MODULES}"@github.com/.insteadOf https://github.com/`
Both variables were sourced from the env: block (mapped from inputs) but were used unquoted in shell commands, allowing attacker-controlled values with shell metacharacters to inject arbitrary commands.

### Iteration 3

**Fixes applied:** script-injection

**Notes:**

Fixed script injection in the 'install task' step of action.yml at line 128. The `go install` URL was unquoted, allowing shell metacharacters in `major_version` (derived from caller-controlled `inputs.task-version`) to be interpreted by the shell. Changed from `go install github.com/go-task/task/v${major_version}/cmd/task@v"${TASK_VERSION}"` to `go install "github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"`, wrapping the entire URL in double quotes to prevent shell interpretation of metacharacters.

