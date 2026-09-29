<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.11.23

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.11.23** was hardened automatically. 14 finding(s) were identified and resolved across 3 iteration(s).

## Findings Fixed

### unpinned-uses (severity: high)

The action uses `svenstaro/upload-release-action@2.11.5` which is pinned to a mutable version tag rather than an immutable 40-character commit SHA. This is vulnerable to supply-chain attacks if the tag is moved. The other two `uses:` references (actions/setup-go and anchore/scan-action) are correctly pinned to SHAs.

Locations:

- `action.yml:222`

### script-injection (severity: high)

Multiple `run:` blocks directly interpolate `${{ inputs.* }}` and `${{ github.* }}` expressions into shell command strings (sub-rule a). This allows an attacker to inject arbitrary shell commands via crafted input values.

1. 'install task' step: `${{ inputs.task-version }}` is interpolated directly into shell commands — e.g. `grep -q "Task version: v${{ inputs.task-version }}"`, `echo "${{ inputs.task-version }}"`, and `go install ...@v${{ inputs.task-version }}`.

2. git config step: `${{ inputs.github-token-for-downloading-private-go-modules }}` is interpolated directly into `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`.

3. 'Build binary' step: `${{ inputs.release-build-tags }}` and `${{ github.ref_name }}` are interpolated directly into shell commands — e.g. `if [ -n "${{ inputs.release-build-tags }}" ]`, `-tags "${{ inputs.release-build-tags }}"`, and `-ldflags="-X 'main.Version=${{ github.ref_name }}'"`. These should be passed via env vars and double-quoted.

4. 'Compute asset name' step: `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` are all interpolated directly into shell variable assignments and conditionals.

Locations:

- `action.yml:107`
- `action.yml:120`
- `action.yml:183`
- `action.yml:207`

### github-env-injection (severity: high)

The 'Compute asset name' step writes `${ASSET_NAME}` to `$GITHUB_OUTPUT` without sanitization. `ASSET_NAME` is constructed by directly interpolating `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` into the shell script. An attacker-controlled newline in any of these values could inject arbitrary key=value pairs into `$GITHUB_OUTPUT`. The required sanitization step (`printf '%s' "$VAR" | tr -d '\n\r'`) is absent before the `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT` write.

Locations:

- `action.yml:207`

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

Fixed all findings in action.yml:
1. Pinned svenstaro/upload-release-action@2.11.5 to full SHA 29e53e917877a24fad85510ded594ab3c9ca12de.
2. 'install task' step: moved ${{ inputs.task-version }} to TASK_VERSION env var; replaced all 3 inline occurrences.
3. git config step: moved ${{ inputs.github-token-for-downloading-private-go-modules }} to GITHUB_TOKEN_FOR_PRIVATE_MODULES env var.
4. 'Build binary' step: moved ${{ inputs.release-build-tags }} to RELEASE_BUILD_TAGS and ${{ github.ref_name }} to GITHUB_REF_NAME env vars; replaced all inline occurrences.
5. 'Compute asset name' step: moved all 5 inline expressions (${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, ${{ inputs.release-build-tags }}) to env vars; added sanitization via 'printf | tr -d \n\r' before writing asset_name to $GITHUB_OUTPUT.

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed the script injection vulnerability in the 'install task' step of action.yml. The `go install` command at line 117 used unquoted `${major_version}` and `${TASK_VERSION}` variables. Fixed by double-quoting the entire module path argument: `go install "github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"`. This prevents attacker-controlled `task-version` input containing shell metacharacters from causing command injection.

### Iteration 1

**Fixes applied:** script-injection

**Notes:**

Fixed script injection vulnerability in action.yml at the git-config step (line 134). The git config command arguments containing the GITHUB_TOKEN_FOR_PRIVATE_MODULES variable were unquoted, allowing shell metacharacters in the token value to be interpreted as shell commands. Fixed by wrapping both arguments in double-quotes: `git config --global "url.https://${GITHUB_TOKEN_FOR_PRIVATE_MODULES}@github.com/.insteadOf" "https://github.com/"`

