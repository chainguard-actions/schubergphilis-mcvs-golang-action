<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.11.22

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.11.22** was hardened automatically. 17 finding(s) were identified and resolved across 3 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Rule (a): The 'install task' run: block directly interpolates ${{ inputs.task-version }} into shell commands. An attacker-controlled value for this input can inject arbitrary shell commands. Offending lines: `if ! task --version | grep -q "Task version: v${{ inputs.task-version }}"`, `major_version=$(echo "${{ inputs.task-version }}" | sed -E ...)`, and `go install github.com/go-task/task/v${major_version}/cmd/task@v${{ inputs.task-version }}`

Locations:

- `action.yml:101`

### script-injection (severity: high)

Rule (a): An unnamed run: block directly interpolates ${{ inputs.github-token-for-downloading-private-go-modules }} into a git config shell command. An attacker-controlled token value can inject arbitrary shell commands. Offending line: `git config --global url.https://${{ inputs.github-token-for-downloading-private-go-modules }}@github.com/.insteadOf https://github.com/`

Locations:

- `action.yml:108`

### script-injection (severity: high)

Rule (a): The 'Build binary' run: block directly interpolates ${{ inputs.release-build-tags }} and ${{ github.ref_name }} into shell commands. These values flow through YAML template substitution before the shell sees them, enabling command injection. Offending lines: `if [ -n "${{ inputs.release-build-tags }}" ]`, `-tags "${{ inputs.release-build-tags }}"`, and `-ldflags="-X 'main.Version=${{ github.ref_name }}'"`.

Locations:

- `action.yml:185`

### script-injection (severity: high)

Rule (a): The 'Compute asset name' run: block directly interpolates ${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, and ${{ inputs.release-build-tags }} into shell variable assignments. An attacker-controlled input can inject arbitrary shell commands. Offending line: `ASSET_NAME_BASE="${{ inputs.release-application-name }}-${{ github.ref_name }}-${{ inputs.release-os }}-${{ inputs.release-architecture }}"`

Locations:

- `action.yml:200`

### github-env-injection (severity: high)

The 'Compute asset name' run: block writes ASSET_NAME to $GITHUB_OUTPUT without sanitization. ASSET_NAME is assembled directly from ${{ inputs.release-application-name }}, ${{ github.ref_name }}, ${{ inputs.release-os }}, ${{ inputs.release-architecture }}, and ${{ inputs.release-build-tags }} via direct expression interpolation in the shell script. No `printf '%s' ... | tr -d '\n\r'` sanitization is applied before the write: `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT`

Locations:

- `action.yml:207`

### unpinned-uses (severity: high)

Two uses: references are pinned to mutable version tags rather than immutable 40-character commit SHAs, making the action vulnerable to supply-chain attacks if those tags are moved: (1) `uses: actions/setup-go@v6.4.0` — should be pinned to a full SHA commit; (2) `uses: svenstaro/upload-release-action@2.11.5` — should be pinned to a full SHA commit. Note: anchore/scan-action is correctly SHA-pinned.

Locations:

- `action.yml:80`
- `action.yml:218`

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

**Fixes applied:** unpinned-uses, script-injection, static-inline-injection, github-env-injection

**Notes:**

Fixed all 17 findings in hardened/action/action.yml:

1. Pinned actions/setup-go@v6.4.0 → @4a3601121dd01d1626a1e23e37211e3254c1c06c # v6.4.0
2. Pinned svenstaro/upload-release-action@2.11.5 → @29e53e917877a24fad85510ded594ab3c9ca12de # 2.11.5
3. 'install task' step: moved ${{ inputs.task-version }} to env var TASK_VERSION
4. git config step: moved ${{ inputs.github-token-for-downloading-private-go-modules }} to env var GITHUB_TOKEN_FOR_PRIVATE_MODULES
5. 'Build binary' step: moved ${{ inputs.release-build-tags }} and ${{ github.ref_name }} to env vars RELEASE_BUILD_TAGS and REF_NAME
6. 'Compute asset name' step: moved all five inline expressions to env vars (RELEASE_APPLICATION_NAME, REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS)
7. 'Compute asset name' step: sanitized ASSET_NAME with printf | tr -d '\n\r' before writing to $GITHUB_OUTPUT to prevent env injection

### Iteration 2

**Fixes applied:** script-injection

**Notes:**

Fixed the script injection vulnerability in the 'install task' step of action.yml (line 119). The `go install` command now double-quotes the URL argument containing `${major_version}` and `${TASK_VERSION}` expansions: `go install "github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"`. This prevents shell metacharacters in the attacker-controllable `task-version` input from being interpreted as shell commands.

### Iteration 3

**Fixes applied:** script-injection

**Notes:**

Fixed script injection in action.yml at line 131. The URL string containing ${GITHUB_TOKEN_FOR_PRIVATE_MODULES} was unquoted, allowing shell metacharacters in the caller-controlled input to be interpreted. Fixed by wrapping the URL in double quotes: `url."https://${GITHUB_TOKEN_FOR_PRIVATE_MODULES}@github.com/".insteadOf` — this prevents shell word-splitting and command injection while maintaining valid git config key syntax.

