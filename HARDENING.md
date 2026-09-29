<!-- markdownlint-disable -->

# Hardening Report: schubergphilis--mcvs-golang-action/v3.11.21

> This file was generated automatically by the hardening agent.

**Policy SHA:** `d636be7e43ef829af6e853da6b3c7566db9f72fe`

**Test Policy SHA:** `843adf9e4b8f85d0c08b27b9d0b09dd094b54702`

**Harden Agent Version:** `2`

Action **schubergphilis--mcvs-golang-action/v3.11.21** was hardened automatically. 14 finding(s) were identified and resolved across 2 iteration(s).

## Findings Fixed

### script-injection (severity: high)

Sub-rule (a): Multiple ${{ }} expressions are interpolated directly inside run: shell command strings in action.yml, allowing an attacker-controlled value to be executed as shell code.

1. 'install task' step: `${{ inputs.task-version }}` appears three times inside the run: block — in a grep string, an echo/sed pipeline, and a `go install` URL. A malicious task-version value (e.g. containing `;`, `$(...)`, or backticks) would be executed by the shell.

2. Unnamed git-config step: `${{ inputs.github-token-for-downloading-private-go-modules }}` is interpolated directly into `git config --global url.https://<TOKEN>@github.com/...`. A crafted token value could inject shell metacharacters.

3. 'Build binary' step: `${{ inputs.release-build-tags }}` is interpolated directly into `if [ -n "${{ inputs.release-build-tags }}" ]` and `-tags "${{ inputs.release-build-tags }}"`, and `${{ github.ref_name }}` is interpolated into `-ldflags="-X 'main.Version=${{ github.ref_name }}'"`. Both allow shell injection.

4. 'Compute asset name' step: `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` are all interpolated directly into shell variable assignments and a conditional test inside the run: block.

Locations:

- `action.yml:118`
- `action.yml:119`
- `action.yml:120`
- `action.yml:126`
- `action.yml:222`
- `action.yml:224`
- `action.yml:227`
- `action.yml:234`
- `action.yml:235`
- `action.yml:236`
- `action.yml:237`
- `action.yml:238`

### github-env-injection (severity: high)

The 'Compute asset name' step writes the value of `${ASSET_NAME}` to `$GITHUB_OUTPUT` without sanitization. `ASSET_NAME` is constructed directly from `${{ inputs.release-application-name }}`, `${{ github.ref_name }}`, `${{ inputs.release-os }}`, `${{ inputs.release-architecture }}`, and `${{ inputs.release-build-tags }}` — all of which are untrusted inputs. A value containing a newline character could inject arbitrary key=value pairs into the GitHub output context (GITHUB_OUTPUT injection). The required sanitization step (`printf '%s' "$VAR" | tr -d '\n\r'`) is absent before the `echo "asset_name=${ASSET_NAME}" >> $GITHUB_OUTPUT` write.

Locations:

- `action.yml:241`

### unpinned-uses (severity: high)

Two `uses:` references in action.yml are pinned to mutable version tags rather than immutable full 40-character commit SHAs, making the action vulnerable to supply-chain attacks if the upstream tag is moved or the repository is compromised:

1. `uses: actions/setup-go@v6.4.0` — pinned to a version tag, not a SHA.
2. `uses: svenstaro/upload-release-action@2.11.5` — pinned to a version tag, not a SHA.

Note: `uses: anchore/scan-action@e1165082ffb1fe366ebaf02d8526e7c4989ea9d2` is correctly pinned to a full SHA and is safe.

Locations:

- `action.yml:87`
- `action.yml:258`

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

Fixed all findings in action.yml:
1. Pinned actions/setup-go@v6.4.0 to SHA 4a3601121dd01d1626a1e23e37211e3254c1c06c
2. Pinned svenstaro/upload-release-action@2.11.5 to SHA 29e53e917877a24fad85510ded594ab3c9ca12de
3. Moved ${{ inputs.task-version }} to TASK_VERSION env var in 'install task' step
4. Moved ${{ inputs.github-token-for-downloading-private-go-modules }} to GITHUB_TOKEN_FOR_PRIVATE_MODULES env var in git config step
5. Moved ${{ inputs.release-build-tags }} to RELEASE_BUILD_TAGS and ${{ github.ref_name }} to REF_NAME env vars in 'Build binary' step
6. Moved all five ${{ }} expressions in 'Compute asset name' step to env vars (RELEASE_APPLICATION_NAME, REF_NAME, RELEASE_OS, RELEASE_ARCHITECTURE, RELEASE_BUILD_TAGS)
7. Added newline sanitization (tr -d '\n\r') before writing asset_name to GITHUB_OUTPUT to prevent env injection

### Iteration 2

**Fixes applied:** script-injection, script-injection

**Notes:**

Fixed two unquoted shell variable expansion issues in action.yml:
1. Line ~111 ('install task' step): Quoted the go install URL argument: `go install "github.com/go-task/task/v${major_version}/cmd/task@v${TASK_VERSION}"` to prevent shell metacharacter injection via inputs.task-version.
2. Line ~124 (git-config step): Quoted the git config URL key argument: `git config --global "url.https://${GITHUB_TOKEN_FOR_PRIVATE_MODULES}@github.com/.insteadOf" https://github.com/` to prevent shell metacharacter injection via inputs.github-token-for-downloading-private-go-modules.

