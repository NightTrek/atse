**User:**

<task>
use the atse tool to explain how authentication works using all the tools atse has avialble
</task>

# task_progress RECOMMENDED

When starting a new task, it is recommended to include a todo list using the task_progress parameter.


1. Include a todo list using the task_progress parameter in your next tool call
2. Create a comprehensive checklist of all steps needed
3. Use markdown format: - [ ] for incomplete, - [x] for complete

**Benefits of creating a todo/task_progress list now:**
	- Clear roadmap for implementation
	- Progress tracking throughout the task
	- Nothing gets forgotten or missed
	- Users can see, monitor, and edit the plan

**Example structure:**```
- [ ] Analyze requirements
- [ ] Set up necessary files
- [ ] Implement main functionality
- [ ] Handle edge cases
- [ ] Test the implementation
- [ ] Verify results```

Keeping the task_progress list updated helps track progress and ensures nothing is missed.

<environment_details>
# Visual Studio Code Visible Files
ATSE_EVALUATION_FRAMEWORK.md

# Visual Studio Code Open Tabs
ADDING_MODEL_PROVIDER.md
ATSE_EVALUATION_FRAMEWORK.md

# Current Time
11/15/2025, 2:01:13 AM (America/Los_Angeles, UTC-8:00)

# Current Working Directory (/Users/danielsteigman/code/@Cline/Cline-combo/cline) Files
.changie.yaml
.env.example
.gitattributes
.gitignore
.mocharc.json
.nvmrc
.nycrc.unit.json
.vscode-test.mjs
.vscodeignore
ADDING_MODEL_PROVIDER.md
ATSE_EVALUATION_FRAMEWORK.md
biome.jsonc
buf.yaml
CHANGELOG.md
CODE_OF_CONDUCT.md
CONTRIBUTING.md
esbuild.mjs
go.work
go.work.sum
knip.json
LICENSE
package-lock.json
package.json
playwright.config.ts
README.md
test-setup.js
tsconfig.json
tsconfig.test.json
tsconfig.unit-test.json
assets/
assets/docs/
assets/docs/demo.gif
assets/icons/
assets/icons/icon.png
assets/icons/icon.svg
assets/icons/robot_panel_dark.png
assets/icons/robot_panel_light.png
cli/
cli/.gitignore
cli/go.mod
cli/go.sum
cli/package.json
cli/README.md
cli/bin/
cli/cmd/
cli/cmd/cline/
cli/cmd/cline-host/
cli/e2e/
cli/e2e/default_update_test.go
cli/e2e/helpers_test.go
cli/e2e/main_test.go
cli/e2e/mixed_stress_test.go
cli/e2e/sqlite_helper.go
cli/e2e/start_list_test.go
cli/man/
cli/man/cline.1
cli/man/cline.1.md
cli/pkg/
cli/pkg/cli/
cli/pkg/common/
cli/pkg/generated/
cli/pkg/hostbridge/
docs/
docs/docs.json
docs/hubspot.js
docs/package-lock.json
docs/package.json
docs/styles.css
docs/assets/
docs/assets/Cline_Logo-complete_black.png
docs/assets/Cline_Logo-complete_white.png
docs/assets/robot_panel_dark.png
docs/assets/robot_panel_light.png
docs/cline-cli/
docs/cline-cli/cli-reference.mdx
docs/cline-cli/installation.mdx
docs/cline-cli/overview.mdx
docs/cline-cli/three-core-flows.mdx
docs/cline-cli/samples/
docs/core-features/
docs/core-features/model-selection-guide.mdx
docs/enterprise-solutions/
docs/enterprise-solutions/overview.mdx
docs/enterprise-solutions/security-concerns.mdx
docs/exploring-clines-tools/
docs/exploring-clines-tools/cline-tools-guide.mdx
docs/features/
docs/getting-started/
docs/introduction/
docs/mcp/
docs/model-config/
docs/more-info/
docs/prompting/
docs/provider-config/
docs/running-models-locally/
docs/testing/
docs/troubleshooting/
evals/
evals/.gitignore
evals/package-lock.json
evals/package.json
evals/README.md
evals/tsconfig.json
evals/cli/
evals/diff-edits/
locales/
locales/ar-sa/
locales/de/
locales/es/
locales/ja/
locales/ko/
locales/pt-BR/
locales/zh-cn/
locales/zh-tw/
proto/
proto/descriptor_set.pb
proto/cline/
proto/host/
scripts/
scripts/api-secrets-parser.mjs
scripts/build-cli-all-platforms.sh
scripts/build-cli.sh
scripts/build-go-proto.mjs
scripts/build-npm-package.sh
scripts/build-proto.mjs
scripts/build-python-proto.mjs
scripts/build-tests.js
scripts/cli-providers.mjs
scripts/dev-cli-watch.mjs
scripts/download-ripgrep.mjs
scripts/file-utils.mjs
scripts/generate-host-bridge-client.mjs
scripts/generate-protobus-setup.mjs
scripts/generate-stubs.js
scripts/get-vscode-usages.sh
scripts/install-local.sh
scripts/install.sh
scripts/interactive-playwright.ts
scripts/package-standalone.mjs
scripts/proto-lint.sh
scripts/proto-shared-utils.mjs
scripts/proto-utils.mjs
scripts/publish-nightly.mjs
scripts/report-issue.js
scripts/runclinecore.sh
scripts/test-hostbridge-server.ts
scripts/test-install.sh
scripts/test-standalone-core-api-server.ts
scripts/testing-platform-orchestrator.ts
src/
src/common.ts
src/config.ts
src/extension.ts
src/registry.ts
src/api/
src/core/
src/dev/
src/exports/
src/generated/
src/hosts/
src/integrations/
src/packages/
src/samples/
src/services/
src/shared/
src/standalone/
src/test/
src/utils/
standalone/
standalone/runtime-files/
testing-platform/
testing-platform/index.ts
testing-platform/package-lock.json
testing-platform/package.json
testing-platform/README.md
testing-platform/tsconfig.json
testing-platform/adapters/
testing-platform/harness/
tests/
tests/specs/
walkthrough/
walkthrough/step1.md
walkthrough/step2.md
walkthrough/step3.md
walkthrough/step4.md
walkthrough/step5.md
webview-ui/
webview-ui/.gitignore
webview-ui/components.json
webview-ui/index.html
webview-ui/package-lock.json
webview-ui/package.json
webview-ui/tailwind.config.mjs
webview-ui/tsconfig.app.json
webview-ui/tsconfig.json
webview-ui/tsconfig.node.json
webview-ui/vite.config.ts
webview-ui/src/
worktrees/
worktrees/core-port-override/

(File list truncated. Use list_files on specific subdirectories if you need to explore further.)

# Workspace Configuration
{
  "workspaces": {
    "/Users/danielsteigman/code/@Cline/Cline-combo/cline": {
      "hint": "cline",
      "associatedRemoteUrls": [
        "aicccode: https://github.com/aicccode/cline.git",
        "fork_chezsmithy: https://github.com/chezsmithy/cline.git",
        "origin: git@github.com:cline/cline.git",
        "raceneader: https://github.com/raceneader/cline.git",
        "shouhanzen: https://github.com/shouhanzen/cline.git"
      ],
      "latestGitCommitHash": "1a07ca7906fbd023420b6e3ed89dcf93f0c8cad0"
    }
  }
}

# Detected CLI Tools
These are some of the tools on the user's machine, and may be useful if needed to accomplish the task: gh, git, docker, kubectl, aws, gcloud, npm, pnpm, pip, cargo, go, curl, jq, make, cmake, python, node, psql, mysql, sqlite3, code, grep, sed, awk, brew, bundle, wget. This list is not exhaustive, and other tools may be available.

# Current Mode
ACT MODE
</environment_details>

---

**Assistant:**

[Unexpected content type]

I'll use the ATSE tool to analyze how authentication works in this codebase. Let me follow the recommended workflow: search → graph → extract → detailed analysis.
[Tool Use: execute_command]
Command: cd /Users/danielsteigman/code/@Cline/Cline-combo/cline && atse search "auth" ./src --limit 15 --benchmark
Requires_approval: false
Task_progress: - [ ] Step 1: Discovery - Search for authentication-related symbols
- [ ] Step 2: Understand scope - Build dependency graph
- [ ] Step 3: Extract code - Get full source code
- [ ] Step 4: Find function calls - Track authentication function usage
- [ ] Step 5: Analyze dependencies - Check import relationships
- [ ] Step 6: Get context - Examine specific authentication implementations
- [ ] Step 7: Summarize findings

---

**User:**

[Tool]
[execute_command for 'cd /Users/danielsteigman/code/@Cline/Cline-combo/cline && atse search "auth" ./src --limit 15 --benchmark'] Result:
Command executed.
Output:
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
⚡ Performance Benchmark
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Command:          search
Duration:         5.285s
Memory (Start):   0.2 MB
Memory (Peak):    10.5 MB
Memory (End):     10.5 MB
Files Processed:  821
Results Found:    15
Output Tokens:    252 (o200k_base)
Output Chars:     925
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Found 15 symbol(s):

AuthHandler (class) - hosts/external/AuthHandler.ts:16
AuthService (class) - services/auth/AuthService.ts:62
authenticate (method) - api/providers/sapaicore.ts:377
AuthServiceMock (class) - services/auth/AuthServiceMock.ts:11
authStateChanged (function) - controller/account/authStateChanged.ts:11
authenticatedRequest (method) - services/account/ClineAccountService.ts:49
authenticateMcpServer (function) - controller/mcp/authenticateMcpServer.ts:11
clearAuth (method) - auth/oca/OcaAuthService.ts:167
clearAuth (method) - oca/providers/OcaAuthProvider.ts:229
getAuthUrl (method) - oca/providers/OcaAuthProvider.ts:142
getAuthToken (method) - services/auth/AuthService.ts:114
getAuthToken (method) - auth/oca/OcaAuthService.ts:105
getAuthToken (method) - services/auth/AuthServiceMock.ts:40
handleDeauth (method) - services/auth/AuthService.ts:236
handleDeauth (method) - auth/oca/OcaAuthService.ts:155
<environment_details>
# Visual Studio Code Visible Files
ATSE_EVALUATION_FRAMEWORK.md

# Visual Studio Code Open Tabs
ADDING_MODEL_PROVIDER.md
ATSE_EVALUATION_FRAMEWORK.md

# Current Time
11/15/2025, 2:01:32 AM (America/Los_Angeles, UTC-8:00)

# Current Mode
ACT MODE
</environment_details>

---

**Assistant:**

[Unexpected content type]
[Tool Use: execute_command]
Command: cd /Users/danielsteigman/code/@Cline/Cline-combo/cline && atse graph "AuthService" ./src --mode bfs --depth 2 --verbose --benchmark
Requires_approval: false
Task_progress: - [x] Step 1: Discovery - Search for authentication-related symbols
- [ ] Step 2: Understand scope - Build dependency graph
- [ ] Step 3: Extract code - Get full source code
- [ ] Step 4: Find function calls - Track authentication function usage
- [ ] Step 5: Analyze dependencies - Check import relationships
- [ ] Step 6: Get context - Examine specific authentication implementations
- [ ] Step 7: Summarize findings

---

**User:**

[Tool]
[execute_command for 'cd /Users/danielsteigman/code/@Cline/Cline-combo/cline && atse graph "AuthService" ./src --mode bfs --depth 2 --verbose --benchmark'] Result:
Command executed.
Output:
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
⚡ Performance Benchmark
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Command:          graph
Duration:         74.931s
Memory (Start):   0.2 MB
Memory (Peak):    12.4 MB
Memory (End):     12.4 MB
Files Processed:  821
Results Found:    12
Output Tokens:    332 (o200k_base)
Output Chars:     1316
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Feature Graph for AuthService (bfs, depth=2):

Level 0 (Entry Point):
  AuthService (class) - services/auth/AuthService.ts:62
    → calls: openExternal, setWelcomeViewCompleted, getRequestRegistry

Level 1:
  openExternal (function) - src/utils/env.ts:41
    → calls: open
  setWelcomeViewCompleted (function) - controller/state/setWelcomeViewCompleted.ts:12
    → calls: clearOnboardingModelsCache
  getRequestRegistry (function) - core/controller/grpc-handler.ts:187
    → calls: recordRequest, withRecordingMiddleware, handleStreamingRequest, handleUnaryRequest, getHandler, postMessageToWebview, postMessageToWebview, postMessageToWebview, getHandler, postMessageToWebview, postMessageToWebview

Level 2:
  postMessageToWebview (method) - hosts/vscode/VscodeWebviewProvider.ts:187
  open (method) - integrations/editor/DiffViewProvider.ts:29
  clearOnboardingModelsCache (function) - controller/models/getClineOnboardingModels.ts:49
  recordRequest (function) - core/controller/grpc-handler.ts:41
  withRecordingMiddleware (function) - core/controller/grpc-handler.ts:22
  handleStreamingRequest (function) - core/controller/grpc-handler.ts:112
  handleUnaryRequest (function) - core/controller/grpc-handler.ts:74
  getHandler (function) - core/controller/grpc-handler.ts:191

Found 12 components across 7 files
<environment_details>
# Visual Studio Code Visible Files
ATSE_EVALUATION_FRAMEWORK.md

# Visual Studio Code Open Tabs
ADDING_MODEL_PROVIDER.md
ATSE_EVALUATION_FRAMEWORK.md

# Current Time
11/15/2025, 2:02:53 AM (America/Los_Angeles, UTC-8:00)

# Current Mode
ACT MODE
</environment_details>

---

**Assistant:**

[Unexpected content type]
[Tool Use: execute_command]
Command: cd /Users/danielsteigman/code/@Cline/Cline-combo/cline && atse extract "AuthService" ./src --depth 1 --benchmark
Requires_approval: false
Task_progress: - [x] Step 1: Discovery - Search for authentication-related symbols
- [x] Step 2: Understand scope - Build dependency graph
- [ ] Step 3: Extract code - Get full source code
- [ ] Step 4: Find function calls - Track authentication function usage
- [ ] Step 5: Analyze dependencies - Check import relationships
- [ ] Step 6: Get context - Examine specific authentication implementations
- [ ] Step 7: Summarize findings

---

**User:**

[Tool]
[execute_command for 'cd /Users/danielsteigman/code/@Cline/Cline-combo/cline && atse extract "AuthService" ./src --depth 1 --benchmark'] Result:
Command executed.
Output:
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
⚡ Performance Benchmark
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Command:          extract
Duration:         19.584s
Memory (Start):   0.2 MB
Memory (Peak):    13.1 MB
Memory (End):     13.1 MB
Files Processed:  821
Results Found:    4
Output Tokens:    2981 (o200k_base)
Output Chars:     12693
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
# Feature: AuthService (4 components, 4 files)

## Entry Point

### services/auth/AuthService.ts

```
class AuthService {
	protected static instance: AuthService | null = null
	protected _authenticated: boolean = false
	protected _clineAuthInfo: ClineAuthInfo | null = null
	protected _provider: IAuthProvider | null = null
	protected _activeAuthStatusUpdateHandlers = new Set<StreamingResponseHandler<AuthState>>()
	protected _handlerToController = new Map<StreamingResponseHandler<AuthState>, Controller>()
	protected _controller: Controller

	/**
	 * Creates an instance of AuthService.
	 * @param controller - Optional reference to the Controller instance.
	 */
	protected constructor(controller: Controller) {
		this._initProvider()
		this._controller = controller
	}

	/**
	 * Gets the singleton instance of AuthService.
	 * @param controller - Optional reference to the Controller instance.
	 * @returns The singleton instance of AuthService.
	 */
	public static getInstance(controller?: Controller): AuthService {
		if (!AuthService.instance) {
			if (!controller) {
				console.warn("Extension context was not provided to AuthService.getInstance, using default context")
				controller = {} as Controller
			}
			if (process.env.E2E_TEST) {
				// Use require instead of import to avoid circular dependency issues
				// eslint-disable-next-line @typescript-eslint/no-var-requires
				const { AuthServiceMock } = require("./AuthServiceMock")
				AuthService.instance = AuthServiceMock.getInstance(controller)
			} else {
				AuthService.instance = new AuthService(controller)
			}
		}
		if (controller !== undefined && AuthService.instance) {
			AuthService.instance.controller = controller
		}
		return AuthService.instance!
	}

	set controller(controller: Controller) {
		this._controller = controller
	}

	/**
	 * Returns the current authentication token with the appropriate prefix.
	 * Refreshing it if necessary.
	 */
	async getAuthToken(): Promise<string | null> {
		if (!this._provider) {
			throw new Error("Auth provider is not set")
		}

		return this.internalGetAuthToken(this._provider)
	}

	/**
	 * Gets the active organization ID from the authenticated user's info
	 * @returns The active organization ID, or null if no active organization exists
	 */
	getActiveOrganizationId(): string | null {
		if (!this._clineAuthInfo?.userInfo?.organizations) {
			return null
		}
		const activeOrg = this._clineAuthInfo.userInfo.organizations.find((org) => org.active)
		return activeOrg?.organizationId ?? null
	}

	/**
	 * Gets all organizations from the authenticated user's info
	 * @returns Array of organizations, or undefined if not available
	 */
	getUserOrganizations(): ClineAccountOrganization[] | undefined {
		return this._clineAuthInfo?.userInfo?.organizations
	}

	private async internalGetAuthToken(provider: IAuthProvider): Promise<string | null> {
		try {
			let clineAccountAuthToken = this._clineAuthInfo?.idToken
			if (!this._clineAuthInfo || !clineAccountAuthToken || this._clineAuthInfo.provider !== provider.name) {
				// Not authenticated
				return null
			}

			// Check if token has expired
			if (await provider.shouldRefreshIdToken(clineAccountAuthToken, this._clineAuthInfo.expiresAt)) {
				const updatedAuthInfo = await provider.retrieveClineAuthInfo(this._controller)
				if (updatedAuthInfo) {
					this._clineAuthInfo = updatedAuthInfo
					this._authenticated = true
					clineAccountAuthToken = updatedAuthInfo.idToken
				} else if (this.shouldClearAuthInfo(provider)) {
					this._clineAuthInfo = null
					this._authenticated = false
					telemetryService.captureAuthLoggedOut(this._provider?.name, LogoutReason.ERROR_RECOVERY)
				}
				await this.sendAuthStatusUpdate()
			}

			// IMPORTANT: Prefix with 'workos:' so backend can route verification to WorkOS provider
			return clineAccountAuthToken ? `workos:${clineAccountAuthToken}` : null
		} catch (error) {
			console.error("Error getting auth token:", error)
			return null
		}
	}

	private shouldClearAuthInfo(provider: IAuthProvider) {
		return this._clineAuthInfo?.provider === provider.name
	}

	protected _initProvider(): void {
		// Only ClineAuthProvider is supported going forward
		this._provider = new ClineAuthProvider()
	}

	/**
	 * Gets the provider name for the current authentication
	 * @returns The provider name (e.g., "cline", "firebase"), or null if not authenticated
	 */
	getProviderName(): string | null {
		return this._clineAuthInfo?.provider ?? null
	}

	getInfo(): AuthState {
		// TODO: this logic should be cleaner, but this will determine the authentication state for the webview -- if a user object is returned then the webview assumes authenticated, otherwise it assumes logged out (we previously returned a UserInfo object with empty fields, and this represented a broken logged in state)
		let user: any = null
		if (this._clineAuthInfo && this._authenticated) {
			const userInfo = this._clineAuthInfo.userInfo
			this._clineAuthInfo.userInfo.appBaseUrl = ClineEnv.config()?.appBaseUrl

			user = UserInfo.create({
				// TODO: create proto for new user info type
				uid: userInfo?.id,
				displayName: userInfo?.displayName,
				email: userInfo?.email,
				photoUrl: undefined,
				appBaseUrl: userInfo?.appBaseUrl,
			})
		}

		return AuthState.create({
			user,
		})
	}

	async createAuthRequest(strict = false): Promise<String> {
		// In strict mode, we do not open a new auth window if already authenticated
		if (strict && this._authenticated) {
			this.sendAuthStatusUpdate()
			return String.create({ value: "Already authenticated" })
		}

		if (!this._provider) {
			return String.create({
				value: "Authentication provider is not configured",
			})
		}

		const callbackHost = await HostProvider.get().getCallbackUrl()
		const callbackUrl = `${callbackHost}/auth`

		const authUrl = await this._provider.getAuthRequest(callbackUrl)
		const authUrlString = authUrl.toString()

		await openExternal(authUrlString)
		telemetryService.captureAuthStarted(this._provider.name)
		return String.create({ value: authUrlString })
	}

	async handleDeauth(reason: LogoutReason = LogoutReason.UNKNOWN): Promise<void> {
		if (!this._provider) {
			throw new Error("Auth provider is not set")
		}

		try {
			telemetryService.captureAuthLoggedOut(this._provider.name, reason)
			this._clineAuthInfo = null
			this._authenticated = false
			this.destroyTokens()
			this.sendAuthStatusUpdate()
		} catch (error) {
			console.error("Error signing out:", error)
			throw error
		}
	}

	async handleAuthCallback(authorizationCode: string, provider: string): Promise<void> {
		if (!this._provider) {
			throw new Error("Auth provider is not set")
		}

		try {
			this._clineAuthInfo = await this._provider.signIn(this._controller, authorizationCode, provider)
			this._authenticated = this._clineAuthInfo?.idToken !== undefined

			telemetryService.captureAuthSucceeded(this._provider.name)
			await setWelcomeViewCompleted(this._controller, { value: true })
		} catch (error) {
			console.error("Error signing in with custom token:", error)
			telemetryService.captureAuthFailed(this._provider.name)
			throw error
		} finally {
			await this.sendAuthStatusUpdate()
		}
	}

	/**
	 * @deprecated Use handleDeauth() instead. Storage clearing is now handled consistently within the auth domain.
	 * Clear the authentication token from the extension's storage.
	 * This is typically called when the user logs out.
	 */
	async clearAuthToken(): Promise<void> {
		this.destroyTokens()
	}

	/**
	 * Restores the authentication data from the extension's storage.
	 * This is typically called when the extension is activated.
	 */
	async restoreRefreshTokenAndRetrieveAuthInfo(): Promise<void> {
		if (!this._provider) {
			throw new Error("Auth provider is not set")
		}

		try {
			this._clineAuthInfo = await this.retrieveAuthInfo()
			if (this._clineAuthInfo) {
				this._authenticated = true
				await this.sendAuthStatusUpdate()
			} else {
				console.warn("No user found after restoring auth token")
				this._authenticated = false
				this._clineAuthInfo = null
				telemetryService.captureAuthLoggedOut(this._provider?.name, LogoutReason.ERROR_RECOVERY)
			}
		} catch (error) {
			console.error("Error restoring auth token:", error)
			this._authenticated = false
			this._clineAuthInfo = null
			telemetryService.captureAuthLoggedOut(this._provider?.name, LogoutReason.ERROR_RECOVERY)
			return
		}
	}

	private async retrieveAuthInfo(): Promise<ClineAuthInfo | null> {
		if (!this._provider) {
			throw new Error("Auth provider is not set")
		}

		return this._provider.retrieveClineAuthInfo(this._controller)
	}

	/**
	 * Subscribe to authStatusUpdate events
	 * @param controller The controller instance
	 * @param request The empty request
	 * @param responseStream The streaming response handler
	 * @param requestId The ID of the request (passed by the gRPC handler)
	 */
	async subscribeToAuthStatusUpdate(
		controller: Controller,
		_request: EmptyRequest,
		responseStream: StreamingResponseHandler<AuthState>,
		requestId?: string,
	): Promise<void> {
		// Add this subscription to the active subscriptions
		this._activeAuthStatusUpdateHandlers.add(responseStream)
		this._handlerToController.set(responseStream, controller)
		// Register cleanup when the connection is closed
		const cleanup = () => {
			this._activeAuthStatusUpdateHandlers.delete(responseStream)
			this._handlerToController.delete(responseStream)
		}
		// Register the cleanup function with the request registry if we have a requestId
		if (requestId) {
			getRequestRegistry().registerRequest(requestId, cleanup, { type: "authStatusUpdate_subscription" }, responseStream)
		}

		// Send the current authentication status immediately
		try {
			await this.sendAuthStatusUpdate()
		} catch (error) {
			console.error("Error sending initial auth status:", error)
			// Remove the subscription if there was an error
			this._activeAuthStatusUpdateHandlers.delete(responseStream)
			this._handlerToController.delete(responseStream)
		}
	}

	/**
	 * Send an authStatusUpdate event to all active subscribers
	 */
	async sendAuthStatusUpdate(): Promise<void> {
		// Compute once per broadcast
		const authInfo: AuthState = this.getInfo()
		const uniqueControllers = new Set<Controller>()

		// Send the event to all active subscribers
		const streamSends = Array.from(this._activeAuthStatusUpdateHandlers).map(async (responseStream) => {
			const controller = this._handlerToController.get(responseStream)
			if (controller) {
				uniqueControllers.add(controller)
			}
			try {
				await responseStream(
					authInfo,
					false, // Not the last message
				)
			} catch (error) {
				console.error("Error sending authStatusUpdate event:", error)
				// Remove the subscription if there was an error
				this._activeAuthStatusUpdateHandlers.delete(responseStream)
				this._handlerToController.delete(responseStream)
			}
		})

		await Promise.all(streamSends)
		// Identify the user in telemetry if available
		if (this._clineAuthInfo?.userInfo?.id) {
			telemetryService.identifyAccount(this._clineAuthInfo.userInfo)
			// Reset feature flags to ensure they are fetched for the new/logged in user
			featureFlagsService.reset(this._clineAuthInfo?.userInfo?.id)
		}
		// Poll feature flags to ensure they are up to date for all users
		await featureFlagsService.poll(this._clineAuthInfo?.userInfo?.id)

		// Update state in webviews once per unique controller
		await Promise.all(Array.from(uniqueControllers).map((c) => c.postStateToWebview()))
	}

	private destroyTokens() {
		this._controller.stateManager.setSecret("clineAccountId", undefined)
		this._controller.stateManager.setSecret("cline:clineAccountId", undefined)
	}
}
```

## Level 1

### controller/state/setWelcomeViewCompleted.ts

```
async function setWelcomeViewCompleted(controller: Controller, request: BooleanRequest): Promise<Empty> {
	try {
		// Update the global state to set welcomeViewCompleted to the requested value
		controller.stateManager.setGlobalState("welcomeViewCompleted", request.value)

		await controller.postStateToWebview()

		console.log(`Welcome view completed set to: ${request.value}`)
		return Empty.create({})
	} catch (error) {
		console.error("Failed to set welcome view completed:", error)
		throw error
	} finally {
		clearOnboardingModelsCache()
	}
}
```

### core/controller/grpc-handler.ts

```
function getRequestRegistry(): GrpcRequestRegistry {
	return requestRegistry
}
```

### src/utils/env.ts

```
async function openExternal(url: string): Promise<void> {
	console.log("Opening browser:", url)
	await open(url)
}
```

---
Total: 4 components across 4 files
<environment_details>
# Visual Studio Code Visible Files
ATSE_EVALUATION_FRAMEWORK.md

# Visual Studio Code Open Tabs
ADDING_MODEL_PROVIDER.md
ATSE_EVALUATION_FRAMEWORK.md

# Current Time
11/15/2025, 2:03:19 AM (America/Los_Angeles, UTC-8:00)

# Current Mode
ACT MODE
</environment_details>

---

**Assistant:**

[Unexpected content type]
[Tool Use: execute_command]
Command: cd /Users/danielsteigman/code/@Cline/Cline-combo/cline && atse find-fn "getAuthToken" ./src --limit 20 --benchmark
Requires_approval: false
Task_progress: - [x] Step 1: Discovery - Search for authentication-related symbols
- [x] Step 2: Understand scope - Build dependency graph
- [x] Step 3: Extract code - Get full source code
- [ ] Step 4: Find function calls - Track authentication function usage
- [ ] Step 5: Analyze dependencies - Check import relationships
- [ ] Step 6: Get context - Examine specific authentication implementations
- [ ] Step 7: Summarize findings

---

**User:**

[Tool]
[execute_command for 'cd /Users/danielsteigman/code/@Cline/Cline-combo/cline && atse find-fn "getAuthToken" ./src --limit 20 --benchmark'] Result:
Command executed.
Output:
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
⚡ Performance Benchmark
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Command:          find-fn
Duration:         2.803s
Memory (Start):   0.2 MB
Memory (Peak):    3.0 MB
Memory (End):     3.0 MB
Files Processed:  821
Results Found:    10450
Output Tokens:    363267 (o200k_base)
Output Chars:     1086744
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/api/providers/dify.ts:60:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/common.ts:44:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/common.ts:54:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/common.ts:57:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/common.ts:60:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/common.ts:63:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/common.ts:66:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/common.ts:73:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/common.ts:105:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/common.ts:113:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/common.ts:123:64 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/index.ts:448:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/index.ts:466:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:7:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:42:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:47:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:51:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:61:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:72:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:99:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:114:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:136:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:168:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:186:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:234:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:237:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:241:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:246:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:247:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:284:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:300:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:322:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:333:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:358:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:369:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:408:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:426:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:427:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:445:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:458:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:475:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:486:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:505:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:519:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:520:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:534:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:545:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:559:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:571:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:572:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:598:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:616:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:617:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:632:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:645:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:657:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:669:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:682:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:696:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:710:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:722:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/bedrock.test.ts:735:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/claude-code.test.ts:7:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/claude-code.test.ts:11:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/claude-code.test.ts:19:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/claude-code.test.ts:23:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/claude-code.test.ts:24:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/claude-code.test.ts:71:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/claude-code.test.ts:108:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/claude-code.test.ts:140:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/claude-code.test.ts:169:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/claude-code.test.ts:199:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/claude-code.test.ts:229:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/claude-code.test.ts:230:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/claude-code.test.ts:239:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:17:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:45:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:48:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:49:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:57:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:74:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:76:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:87:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:89:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:102:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:105:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:106:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:107:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:110:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:144:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:149:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:149:44 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:153:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:156:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:157:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:159:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:162:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:195:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:209:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:213:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:227:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:233:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:241:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:242:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/litellm.test.ts:243:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/ollama.test.ts:9:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/ollama.test.ts:13:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/ollama.test.ts:27:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/ollama.test.ts:37:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/ollama.test.ts:42:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/ollama.test.ts:43:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/ollama.test.ts:85:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/ollama.test.ts:101:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/ollama.test.ts:101:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/ollama.test.ts:136:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/ollama.test.ts:166:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/ollama.test.ts:183:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/sapaicore.test.ts:5:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/sapaicore.test.ts:8:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/sapaicore.test.ts:20:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/sapaicore.test.ts:24:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/sapaicore.test.ts:40:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/sapaicore.test.ts:70:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/sapaicore.test.ts:94:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/sapaicore.test.ts:95:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/sapaicore.test.ts:102:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/sapaicore.test.ts:112:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/__tests__/sapaicore.test.ts:113:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/aihubmix.ts:120:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/aihubmix.ts:244:72 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/anthropic.ts:46:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/anthropic.ts:90:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/anthropic.ts:139:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/asksage.ts:49:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/asksage.ts:74:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/baseten.ts:87:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/baseten.ts:100:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/baseten.ts:108:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/bedrock.ts:123:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/bedrock.ts:228:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/bedrock.ts:241:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/bedrock.ts:317:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/bedrock.ts:318:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/bedrock.ts:388:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/bedrock.ts:412:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/bedrock.ts:435:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/bedrock.ts:450:100 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/bedrock.ts:458:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/bedrock.ts:468:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/bedrock.ts:485:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/bedrock.ts:578:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/bedrock.ts:686:90 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/bedrock.ts:1071:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/bedrock.ts:1112:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/bedrock.ts:1206:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/bedrock.ts:1247:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/cerebras.ts:44:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/cerebras.ts:98:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/claude-code.ts:22:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/claude-code.ts:29:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/claude-code.ts:31:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/cline.ts:57:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/cline.ts:66:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/cline.ts:98:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/cline.ts:108:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/cline.ts:165:52 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/cline.ts:184:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/cline.ts:238:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/cline.ts:243:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/deepseek.ts:65:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/deepseek.ts:77:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/deepseek.ts:86:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/deepseek.ts:90:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/deepseek.ts:101:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/dify.ts:123:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/dify.ts:438:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/dify.ts:461:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/dify.ts:495:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/dify.ts:532:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/dify.ts:553:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/dify.ts:587:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/dify.ts:623:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/doubao.ts:52:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/doubao.ts:58:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/fireworks.ts:43:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/fireworks.ts:50:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/gemini.ts:108:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/gemini.ts:299:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/gemini.ts:484:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/groq.ts:125:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/groq.ts:194:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/groq.ts:206:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/groq.ts:220:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/hicap.ts:46:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/hicap.ts:53:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/huawei-cloud-maas.ts:64:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/huawei-cloud-maas.ts:70:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/huawei-cloud-maas.ts:79:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/huggingface.ts:57:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/huggingface.ts:71:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/huggingface.ts:79:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/huggingface.ts:89:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/litellm.ts:92:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/litellm.ts:108:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/litellm.ts:186:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/litellm.ts:189:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/litellm.ts:204:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/litellm.ts:210:41 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/lmstudio.ts:42:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/lmstudio.ts:47:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/lmstudio.ts:56:60 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/lmstudio.ts:57:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/lmstudio.ts:101:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/minimax.ts:48:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/mistral.ts:35:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/mistral.ts:50:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/mistral.ts:58:61 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/moonshot.ts:42:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/moonshot.ts:49:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/moonshot.ts:59:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/nebius.ts:41:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/nebius.ts:47:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/nebius.ts:48:52 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/nebius.ts:56:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/nousresearch.ts:39:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/nousresearch.ts:46:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/oca.ts:48:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/oca.ts:111:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/oca.ts:114:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/oca.ts:141:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/oca.ts:144:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/oca.ts:207:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/ollama.ts:50:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/ollama.ts:53:83 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/ollama.ts:59:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/ollama.ts:59:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/ollama.ts:68:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/ollama.ts:118:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/openai-native.ts:49:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/openai-native.ts:61:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/openai-native.ts:74:60 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/openai-native.ts:90:65 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/openai-native.ts:120:65 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/openai-native.ts:124:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/openai-native.ts:156:62 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/openai-native.ts:159:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/openai.ts:67:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/openai.ts:78:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/openai.ts:82:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/openai.ts:91:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/openai.ts:97:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/openai.ts:101:70 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/openai.ts:114:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/openrouter.ts:56:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/openrouter.ts:61:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/openrouter.ts:99:87 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/openrouter.ts:127:51 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/openrouter.ts:141:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/openrouter.ts:175:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/openrouter.ts:197:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/openrouter.ts:206:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/qwen-code.ts:48:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/qwen-code.ts:48:46 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/qwen-code.ts:75:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/qwen-code.ts:80:52 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/qwen-code.ts:97:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/qwen-code.ts:103:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/qwen-code.ts:125:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/qwen-code.ts:164:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/qwen-code.ts:172:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/qwen-code.ts:179:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/qwen-code.ts:190:47 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/qwen-code.ts:199:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/qwen.ts:83:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/qwen.ts:92:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/qwen.ts:107:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/qwen.ts:119:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/requesty.ts:46:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/requesty.ts:61:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/requesty.ts:68:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/requesty.ts:128:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/sambanova.ts:44:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/sambanova.ts:51:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/sambanova.ts:57:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/sambanova.ts:67:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/sapaicore.ts:141:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/sapaicore.ts:389:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/sapaicore.ts:416:55 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/sapaicore.ts:461:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/sapaicore.ts:647:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/sapaicore.ts:687:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/sapaicore.ts:691:91 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/sapaicore.ts:833:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/sapaicore.ts:873:97 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/sapaicore.ts:883:88 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/sapaicore.ts:1048:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/together.ts:44:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/together.ts:52:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/together.ts:56:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/together.ts:65:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/vercel-ai-gateway.ts:49:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/vercel-ai-gateway.ts:56:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/vertex.ts:69:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/vertex.ts:126:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/vertex.ts:152:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/vscode-lm.ts:368:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/vscode-lm.ts:384:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/vscode-lm.ts:482:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/vscode-lm.ts:507:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/vscode-lm.ts:557:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/xai.ts:46:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/xai.ts:62:60 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/xai.ts:66:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/xai.ts:86:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/zai.ts:78:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/zai.ts:84:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/providers/zai.ts:92:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/retry.test.ts:6:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/retry.test.ts:7:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/retry.test.ts:11:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/retry.test.ts:12:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/retry.test.ts:15:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/retry.test.ts:32:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/retry.test.ts:35:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/retry.test.ts:57:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/retry.test.ts:60:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/retry.test.ts:79:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/retry.test.ts:85:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/retry.test.ts:112:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/retry.test.ts:120:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/retry.test.ts:148:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/retry.test.ts:154:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/retry.test.ts:180:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/retry.test.ts:187:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/retry.test.ts:213:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/retry.test.ts:216:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/retry.ts:57:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/retry.ts:79:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/anthropic-format.ts:14:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/gemini-format.ts:51:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/o1-format.ts:206:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/o1-format.ts:264:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/o1-format.ts:293:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/o1-format.ts:319:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/o1-format.ts:369:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/openai-format.ts:188:54 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/openrouter-stream.ts:28:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/openrouter-stream.ts:139:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/openrouter-stream.ts:191:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vercel-ai-gateway-stream.ts:24:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vercel-ai-gateway-stream.ts:87:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:9:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:10:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:11:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:12:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:13:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:14:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:17:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:18:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:21:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:22:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:25:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:27:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:28:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:31:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:33:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:37:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:38:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:41:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:42:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:43:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:47:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:48:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:54:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:68:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:83:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:103:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:119:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:137:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:154:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:164:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:165:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:171:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:201:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.test.ts:205:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.ts:125:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.ts:162:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/api/transform/vscode-lm-format.ts:187:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:6:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:9:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:202:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:205:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:208:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:212:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:215:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:218:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:219:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:223:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:226:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:231:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:240:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:243:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:247:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:250:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:254:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:262:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:267:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:270:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:281:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:284:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:306:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:307:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:320:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:322:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:325:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:343:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:345:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:348:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:361:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:363:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:366:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:381:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.test.ts:383:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.ts:244:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.ts:283:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.ts:284:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.ts:285:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.ts:291:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.ts:298:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.ts:340:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.ts:345:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.ts:380:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.ts:586:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.ts:596:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.ts:604:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.ts:669:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff.ts:678:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:6:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:9:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:10:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:18:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:19:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:21:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:22:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:25:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:33:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:34:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:36:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:37:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:40:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:48:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:49:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:51:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:52:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:55:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:63:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:64:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:65:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:66:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:69:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:77:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:78:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:80:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:81:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:84:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:92:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:93:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:95:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:96:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:99:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:112:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:113:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:115:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:116:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:119:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:132:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:133:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:135:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/assistant-message/diff_edge_cases.test.ts:136:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/commands/reconstructTaskHistory.ts:42:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/commands/reconstructTaskHistory.ts:59:64 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/commands/reconstructTaskHistory.ts:76:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/commands/reconstructTaskHistory.ts:82:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/commands/reconstructTaskHistory.ts:87:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/commands/reconstructTaskHistory.ts:99:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/commands/reconstructTaskHistory.ts:108:61 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/commands/reconstructTaskHistory.ts:117:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/commands/reconstructTaskHistory.ts:124:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/commands/reconstructTaskHistory.ts:156:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/commands/reconstructTaskHistory.ts:162:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/commands/reconstructTaskHistory.ts:165:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/commands/reconstructTaskHistory.ts:170:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/ContextManager-legacy.ts:22:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/ContextManager.ts:71:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/ContextManager.ts:123:46 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/ContextManager.ts:165:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/ContextManager.ts:201:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/ContextManager.ts:441:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/ContextManager.ts:467:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/ContextManager.ts:477:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/ContextManager.ts:478:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/ContextManager.ts:492:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:5:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:26:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:29:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:33:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:34:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:37:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:40:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:41:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:44:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:47:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:48:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:50:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:54:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:57:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:58:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:63:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:64:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:67:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:68:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:73:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:77:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:80:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:81:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:84:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:87:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:88:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:95:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:98:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:103:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:106:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:110:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:111:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:114:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:117:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:118:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:123:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:124:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:125:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:126:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:129:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:130:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:135:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:136:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:137:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:138:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:141:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:142:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:147:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:148:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:149:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:150:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:152:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:153:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:154:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:157:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:186:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:190:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:191:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:195:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:196:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/__tests__/ContextManager.test.ts:197:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/context-error-handling.ts:5:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/context-error-handling.ts:6:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/context-error-handling.ts:7:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/context-error-handling.ts:8:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/context-error-handling.ts:15:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/context-error-handling.ts:29:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/context-error-handling.ts:45:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/context-error-handling.ts:66:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-management/context-error-handling.ts:68:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:13:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:26:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:54:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:60:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:64:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:68:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:69:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:72:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:75:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:78:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:79:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:80:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:81:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:82:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:85:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:89:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:93:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:101:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:104:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:105:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:106:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:107:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:108:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:111:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:118:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:119:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:120:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:121:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:122:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:125:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:132:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:133:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:134:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:135:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:139:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:142:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:160:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:163:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:167:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:168:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:171:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:175:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:178:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:181:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:196:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:199:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:203:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:206:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:224:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:227:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:231:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:234:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.test.ts:242:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.ts:47:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.ts:86:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.ts:109:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.ts:132:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.ts:133:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.ts:134:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.ts:158:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.ts:198:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/FileContextTracker.ts:283:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:8:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:16:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:28:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:32:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:47:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:48:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:51:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:57:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:61:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:62:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:63:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:64:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:71:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:97:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:103:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:106:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:114:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:125:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:157:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:158:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:164:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:168:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:169:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:170:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.test.ts:171:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.ts:11:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/context-tracking/ModelContextTracker.ts:35:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/cline-rules.ts:15:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/cline-rules.ts:16:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/cline-rules.ts:18:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/cline-rules.ts:19:41 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/cline-rules.ts:62:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/cline-rules.ts:63:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/cline-rules.ts:65:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/cline-rules.ts:70:41 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/cline-rules.ts:103:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/cline-rules.ts:104:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/cline-rules.ts:110:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/external-rules.ts:30:43 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/external-rules.ts:39:42 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/external-rules.ts:42:42 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/external-rules.ts:44:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/external-rules.ts:50:41 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/external-rules.ts:68:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/external-rules.ts:69:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/external-rules.ts:94:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/external-rules.ts:95:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/external-rules.ts:113:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/external-rules.ts:114:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/external-rules.ts:116:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/external-rules.ts:117:41 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/external-rules.ts:138:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/external-rules.ts:146:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/external-rules.ts:169:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/rule-helpers.ts:18:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/rule-helpers.ts:49:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/rule-helpers.ts:52:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/rule-helpers.ts:56:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/rule-helpers.ts:169:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/rule-helpers.ts:171:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/rule-helpers.ts:206:46 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/rule-helpers.ts:209:43 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/rule-helpers.ts:215:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/rule-helpers.ts:225:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/rule-helpers.ts:239:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/rule-helpers.ts:264:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/rule-helpers.ts:318:64 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/workflows.ts:19:44 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/workflows.ts:20:44 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/context/instructions/user-instructions/workflows.ts:25:38 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/account/openrouterAuthClicked.ts:13:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/account/requestyAuthClicked.ts:13:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/account/requestyAuthClicked.ts:21:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/account/setUserOrganization.ts:19:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/browser/discoverBrowser.ts:15:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/browser/discoverBrowser.ts:41:83 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/browser/testBrowserConnection.ts:21:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/browser/testBrowserConnection.ts:41:87 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/browser/testBrowserConnection.ts:57:82 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/checkpoints/checkpointRestore.ts:14:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/checkpoints/getCwdHash.ts:11:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/checkpoints/subscribeToCheckpoints.ts:66:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/checkpoints/subscribeToCheckpoints.ts:106:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/commands/addToCline.ts:16:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/commands/addToCline.ts:20:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/commands/addToCline.ts:24:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/commands/explainWithCline.ts:16:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/commands/fixWithCline.ts:9:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/commands/fixWithCline.ts:10:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/commands/improveWithCline.ts:16:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/dictation/startRecording.ts:61:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/dictation/startRecording.ts:63:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/dictation/startRecording.ts:136:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/dictation/startRecording.ts:138:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/dictation/startRecording.ts:152:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/dictation/startRecording.ts:155:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/dictation/transcribeAudio.ts:23:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/ifFileExistsRelativePath.test.ts:9:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/ifFileExistsRelativePath.test.ts:15:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/ifFileExistsRelativePath.test.ts:28:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/ifFileExistsRelativePath.test.ts:32:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/ifFileExistsRelativePath.test.ts:39:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/ifFileExistsRelativePath.test.ts:42:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/ifFileExistsRelativePath.test.ts:43:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/ifFileExistsRelativePath.test.ts:46:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/ifFileExistsRelativePath.test.ts:57:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/ifFileExistsRelativePath.test.ts:59:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/ifFileExistsRelativePath.test.ts:60:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/ifFileExistsRelativePath.test.ts:64:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/ifFileExistsRelativePath.test.ts:74:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/ifFileExistsRelativePath.test.ts:76:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/ifFileExistsRelativePath.test.ts:80:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/ifFileExistsRelativePath.test.ts:91:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/ifFileExistsRelativePath.test.ts:94:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/ifFileExistsRelativePath.test.ts:95:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/ifFileExistsRelativePath.test.ts:99:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/openFileRelativePath.test.ts:11:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/openFileRelativePath.test.ts:18:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/openFileRelativePath.test.ts:34:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/openFileRelativePath.test.ts:38:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/openFileRelativePath.test.ts:45:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/openFileRelativePath.test.ts:47:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/openFileRelativePath.test.ts:50:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/openFileRelativePath.test.ts:61:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/openFileRelativePath.test.ts:63:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/openFileRelativePath.test.ts:66:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/openFileRelativePath.test.ts:76:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/openFileRelativePath.test.ts:78:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/openFileRelativePath.test.ts:83:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/openFileRelativePath.test.ts:94:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/openFileRelativePath.test.ts:96:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/openFileRelativePath.test.ts:97:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/openFileRelativePath.test.ts:98:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/openFileRelativePath.test.ts:102:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/openFileRelativePath.test.ts:113:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/__tests__/openFileRelativePath.test.ts:115:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/copyToClipboard.ts:14:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/createRuleFile.ts:35:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/createRuleFile.ts:35:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/createRuleFile.ts:36:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/createRuleFile.ts:51:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/createRuleFile.ts:54:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/createRuleFile.ts:56:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/createRuleFile.ts:60:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/createRuleFile.ts:71:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/deleteRuleFile.ts:31:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/deleteRuleFile.ts:43:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/getRelativePaths.ts:18:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/getRelativePaths.ts:28:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/getRelativePaths.ts:36:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/ifFileExistsRelativePath.ts:14:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/openDiskConversationHistory.ts:16:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/openFile.ts:13:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/openFileRelativePath.ts:14:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/openFileRelativePath.ts:31:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/openFocusChainFile.ts:31:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/openFocusChainFile.ts:35:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/openFocusChainFile.ts:37:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/openImage.ts:13:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/openMention.ts:12:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/refreshRules.ts:17:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/refreshRules.ts:17:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/refreshRules.ts:18:48 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/refreshRules.ts:19:81 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/refreshRules.ts:23:64 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/searchCommits.ts:14:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/searchCommits.ts:20:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/searchFiles.ts:32:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/searchFiles.ts:41:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/searchFiles.ts:50:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/searchFiles.ts:59:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/searchFiles.ts:84:64 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/selectFiles.ts:13:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/file/toggleClineRule.ts:52:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:9:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:22:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:42:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:46:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:47:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:48:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:57:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:60:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:61:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:62:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:65:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:67:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:76:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:85:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:88:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:90:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:100:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:109:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:112:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:114:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:115:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:116:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:119:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:128:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:131:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:133:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:134:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:135:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:138:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:139:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:154:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:155:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:156:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:160:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:163:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:164:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:165:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:166:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:169:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:172:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:181:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:190:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:201:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:210:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:213:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:215:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:225:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:244:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:250:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:253:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:256:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:259:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:270:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:284:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:287:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:292:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:304:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:305:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:307:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:315:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:318:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:321:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:323:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:333:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:336:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:341:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:344:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:347:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:349:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:358:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:361:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:364:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:367:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:371:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:372:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:376:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:377:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:382:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:389:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:396:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:408:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:409:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.test.ts:412:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.ts:34:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.ts:57:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.ts:60:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.ts:63:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.ts:65:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.ts:81:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.ts:83:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.ts:85:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.ts:95:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.ts:98:52 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.ts:123:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.ts:136:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.ts:139:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.ts:146:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.ts:149:52 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-handler.ts:167:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.builder.test.ts:7:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.builder.test.ts:8:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.builder.test.ts:9:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.builder.test.ts:16:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.builder.test.ts:25:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.builder.test.ts:26:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.test.ts:6:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.test.ts:9:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.test.ts:16:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.test.ts:17:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.test.ts:79:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.test.ts:81:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.test.ts:94:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.test.ts:95:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.test.ts:101:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.test.ts:116:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.test.ts:119:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.test.ts:138:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.test.ts:147:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.test.ts:150:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.test.ts:175:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.test.ts:180:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.test.ts:181:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.test.ts:184:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.test.ts:209:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.test.ts:210:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/grpc-recorder.test.ts:211:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/log-file-handler.test.ts:5:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/log-file-handler.test.ts:8:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/log-file-handler.test.ts:10:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/log-file-handler.test.ts:13:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/log-file-handler.test.ts:14:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/log-file-handler.test.ts:16:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/test-hooks.test.ts:10:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/test-hooks.test.ts:16:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/test-hooks.test.ts:32:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/test-hooks.test.ts:36:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/test-hooks.test.ts:38:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/test-hooks.test.ts:45:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/__tests__/test-hooks.test.ts:50:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/grpc-recorder.builder.ts:61:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/grpc-recorder.builder.ts:66:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/grpc-recorder.builder.ts:80:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/grpc-recorder.builder.ts:110:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/grpc-recorder.ts:151:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/grpc-recorder.ts:201:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/grpc-recorder.ts:223:46 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/log-file-handler.ts:51:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/log-file-handler.ts:55:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/test-hooks.ts:14:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-recorder/test-hooks.ts:30:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-service.ts:97:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/grpc-service.ts:124:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:96:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:118:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:120:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:120:45 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:155:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:165:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:166:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:172:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:177:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:188:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:254:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:285:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:290:70 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:290:77 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:296:54 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:373:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:397:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:446:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:578:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:629:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:662:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:673:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:699:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:707:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:721:89 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:747:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:767:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:773:51 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:775:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:779:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:806:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:828:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:845:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:850:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:909:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/index.ts:912:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/mcp/addRemoteMcpServer.ts:25:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/mcp/deleteMcpServer.ts:18:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/mcp/downloadMcp.ts:40:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/mcp/downloadMcp.ts:78:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/mcp/getLatestMcpServers.ts:18:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/mcp/openMcpSettings.ts:14:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/mcp/restartMcpServer.ts:17:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/mcp/subscribeToMcpMarketplaceCatalog.ts:32:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/mcp/subscribeToMcpMarketplaceCatalog.ts:43:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/mcp/subscribeToMcpServers.ts:33:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/mcp/subscribeToMcpServers.ts:42:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/mcp/subscribeToMcpServers.ts:44:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/mcp/subscribeToMcpServers.ts:64:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/mcp/toggleMcpServer.ts:17:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/mcp/toggleToolAutoApprove.ts:19:41 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/mcp/updateMcpTimeout.ts:15:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/getAihubmixModels.ts:15:93 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/getClineOnboardingModels.ts:24:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/getLmStudioModels.ts:19:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/getOllamaModels.ts:20:58 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/getSapAiCoreModels.ts:37:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/getSapAiCoreModels.ts:70:54 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/getSapAiCoreModels.ts:121:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/getSapAiCoreModels.ts:122:56 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/getVsCodeLmModels.ts:17:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshBasetenModels.ts:18:47 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshBasetenModels.ts:54:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshBasetenModels.ts:62:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshBasetenModels.ts:74:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshBasetenModels.ts:75:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshBasetenModels.ts:78:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshBasetenModels.ts:113:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshBasetenModels.ts:162:47 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshBasetenModels.ts:163:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshBasetenModelsRpc.ts:17:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshBasetenModelsRpc.ts:18:55 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshGroqModels.ts:18:44 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshGroqModels.ts:56:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshGroqModels.ts:64:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshGroqModels.ts:74:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshGroqModels.ts:80:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshGroqModels.ts:120:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshGroqModels.ts:168:44 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshGroqModels.ts:169:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshGroqModelsRpc.ts:17:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshGroqModelsRpc.ts:18:55 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshHicapModels.ts:26:45 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshHicapModels.ts:37:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshHicapModels.ts:74:45 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshHicapModels.ts:75:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshHuggingFaceModels.ts:22:51 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshHuggingFaceModels.ts:30:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshHuggingFaceModels.ts:74:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshOcaModels.ts:23:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshOcaModels.ts:40:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshOcaModels.ts:43:60 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshOcaModels.ts:65:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshOcaModels.ts:66:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshOcaModels.ts:67:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshOcaModels.ts:68:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshOpenAiModels.ts:29:80 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshOpenRouterModels.ts:79:50 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshOpenRouterModels.ts:83:74 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshOpenRouterModels.ts:89:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshOpenRouterModels.ts:102:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshOpenRouterModels.ts:103:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshOpenRouterModels.ts:104:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshOpenRouterModels.ts:105:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshOpenRouterModels.ts:205:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshOpenRouterModels.ts:208:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshOpenRouterModels.ts:212:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshOpenRouterModels.ts:215:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshOpenRouterModels.ts:225:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshOpenRouterModels.ts:246:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshOpenRouterModels.ts:250:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshOpenRouterModelsRpc.ts:17:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshOpenRouterModelsRpc.ts:19:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshRequestyModels.ts:17:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshRequestyModels.ts:27:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshRequestyModels.ts:37:54 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshRequestyModels.ts:45:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshRequestyModels.ts:46:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshRequestyModels.ts:47:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshRequestyModels.ts:48:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshVercelAiGatewayModels.ts:16:55 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshVercelAiGatewayModels.ts:21:77 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshVercelAiGatewayModels.ts:27:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshVercelAiGatewayModels.ts:40:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshVercelAiGatewayModels.ts:41:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshVercelAiGatewayModels.ts:42:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshVercelAiGatewayModels.ts:43:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshVercelAiGatewayModels.ts:61:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshVercelAiGatewayModels.ts:74:55 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshVercelAiGatewayModels.ts:75:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshVercelAiGatewayModelsRpc.ts:17:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/refreshVercelAiGatewayModelsRpc.ts:18:55 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/subscribeToOpenRouterModels.ts:35:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/subscribeToOpenRouterModels.ts:47:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfiguration.ts:69:69 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfiguration.ts:107:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfiguration.ts:109:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfiguration.ts:116:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfiguration.ts:119:38 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfiguration.ts:121:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfiguration.ts:144:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationPartial.ts:34:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationPartial.ts:46:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:36:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:40:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:46:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:49:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:52:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:55:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:58:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:61:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:64:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:67:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:70:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:73:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:76:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:81:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:84:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:87:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:90:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:93:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:96:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:99:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:102:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:105:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:108:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:111:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/models/updateApiConfigurationProto.ts:121:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/captureOnboardingProgress.ts:15:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/checkCliInstallation.ts:12:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/getAvailableTerminalProfiles.ts:9:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/resetState.ts:22:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/resetState.ts:28:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/resetState.ts:42:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/resetState.ts:49:79 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/setWelcomeViewCompleted.ts:25:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/subscribeToState.ts:34:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/subscribeToState.ts:44:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/subscribeToState.ts:65:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateInfoBannerVersion.ts:11:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateModelBannerVersion.ts:11:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettings.test.ts:12:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettings.test.ts:16:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettings.test.ts:36:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettings.test.ts:46:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettings.test.ts:57:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettings.test.ts:65:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettings.test.ts:76:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettings.test.ts:84:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettings.test.ts:95:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettings.test.ts:111:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettings.test.ts:122:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettings.test.ts:130:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettings.test.ts:141:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettings.test.ts:149:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettings.test.ts:158:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettings.ts:33:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettings.ts:43:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettings.ts:46:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettings.ts:58:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettings.ts:141:69 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettings.ts:151:69 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettings.ts:165:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettings.ts:173:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettings.ts:179:68 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettingsCli.ts:50:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettingsCli.ts:105:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettingsCli.ts:110:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettingsCli.ts:119:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettingsCli.ts:124:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateSettingsCli.ts:134:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateTaskSettings.ts:93:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateTaskSettings.ts:98:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateTaskSettings.ts:107:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateTaskSettings.ts:112:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/state/updateTelemetrySetting.ts:13:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/task/deleteAllTaskHistory.ts:54:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/task/deleteAllTaskHistory.ts:95:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/task/deleteAllTaskHistory.ts:101:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/task/deleteAllTaskHistory.ts:107:148 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/task/deleteAllTaskHistory.ts:134:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/task/deleteTasksWithIds.ts:38:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/task/deleteTasksWithIds.ts:86:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/task/deleteTasksWithIds.ts:89:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/task/getTaskHistory.ts:17:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/task/getTaskHistory.ts:38:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/task/getTaskHistory.ts:45:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/task/getTotalTasksSize.ts:12:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/task/newTask.ts:69:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/task/newTask.ts:72:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/task/newTask.ts:78:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/task/newTask.ts:81:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/task/showTaskWithId.ts:26:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/task/showTaskWithId.ts:50:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/initializeWebview.ts:25:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/initializeWebview.ts:29:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/initializeWebview.ts:71:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/initializeWebview.ts:113:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/initializeWebview.ts:155:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/initializeWebview.ts:203:38 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/initializeWebview.ts:206:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/onDidShowAnnouncement.ts:15:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/openUrl.ts:14:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/subscribeToAccountButtonClicked.ts:33:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/subscribeToAccountButtonClicked.ts:45:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/subscribeToAddToInput.ts:34:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/subscribeToAddToInput.ts:49:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/subscribeToChatButtonClicked.ts:33:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/subscribeToChatButtonClicked.ts:45:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/subscribeToDidBecomeVisible.ts:33:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/subscribeToDidBecomeVisible.ts:45:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/subscribeToFocusChatInput.ts:31:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/subscribeToFocusChatInput.ts:43:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/subscribeToHistoryButtonClicked.ts:31:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/subscribeToHistoryButtonClicked.ts:43:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/subscribeToMcpButtonClicked.ts:33:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/subscribeToMcpButtonClicked.ts:45:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/subscribeToPartialMessage.ts:32:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/subscribeToPartialMessage.ts:44:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/subscribeToRelinquishControl.ts:31:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/subscribeToRelinquishControl.ts:43:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/subscribeToSettingsButtonClicked.ts:31:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/ui/subscribeToSettingsButtonClicked.ts:43:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/web/checkIsImageUrl.ts:16:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/web/fetchOpenGraphData.ts:17:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/web/fetchOpenGraphData.ts:20:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/controller/web/openInBrowser.ts:14:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/HookDiscoveryCache.ts:93:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/HookDiscoveryCache.ts:124:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/HookDiscoveryCache.ts:132:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/HookProcess.ts:64:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/HookProcess.ts:80:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/HookProcess.ts:93:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/HookProcess.ts:105:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/HookProcess.ts:106:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/HookProcess.ts:115:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/HookProcess.ts:118:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/HookProcess.ts:159:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/HookProcess.ts:171:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/HookProcess.ts:173:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/HookProcess.ts:183:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/HookProcess.ts:191:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/HookProcess.ts:199:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/HookProcess.ts:348:52 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/HookProcess.ts:369:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:11:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:14:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:29:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:30:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:49:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:63:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:64:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:81:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:82:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:92:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:110:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:118:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:135:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:145:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:162:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:167:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:186:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:191:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:210:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:219:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:236:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:237:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:246:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:266:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:267:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:290:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:316:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:334:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:335:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:352:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:363:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:380:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:384:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:390:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:394:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:400:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:409:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:419:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:435:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:443:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:451:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:464:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:472:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:485:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:493:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:502:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:517:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:525:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:534:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:548:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hook-factory.test.ts:557:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:6:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:10:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:15:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:25:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:26:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:27:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:36:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:37:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:41:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:42:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:46:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:47:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:52:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:56:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:57:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:65:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:66:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:70:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:71:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:75:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:76:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:83:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:84:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:92:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:93:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:94:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:97:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:105:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:106:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/hooks-utils.test.ts:107:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/setup.ts:36:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/setup.ts:76:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/setup.ts:78:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/setup.ts:81:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/setup.ts:84:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:5:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:18:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:26:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:27:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:28:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:31:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:33:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:37:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:39:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:43:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:45:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:49:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:51:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:56:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:58:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:62:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:64:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:69:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:71:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:76:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:78:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:83:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:85:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:89:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:91:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:95:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:97:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:101:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:103:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:107:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:109:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:113:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:115:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:119:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:121:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:126:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:128:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:133:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:134:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:135:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:138:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:140:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:144:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:146:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:150:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:152:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:156:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:158:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:163:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:165:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:170:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:172:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:177:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:179:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:184:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:186:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:190:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:192:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:196:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:198:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:202:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:204:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:208:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:210:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:215:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:216:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:217:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:219:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:223:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:224:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:226:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:230:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:231:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:233:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:237:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:238:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:240:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:244:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:245:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:247:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:252:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:253:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:254:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:269:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/shell-escape.test.ts:270:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:11:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:14:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:30:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:51:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:65:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:66:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:78:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:98:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:113:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:133:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:150:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:171:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:172:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:181:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:208:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:228:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:238:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:258:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:267:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:290:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:303:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:323:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:324:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:330:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:353:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:358:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:381:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:385:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:391:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:394:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:399:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:409:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:420:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:439:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:450:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:460:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:480:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:481:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:500:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:501:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:502:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:502:56 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:524:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:525:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:525:58 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:547:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:548:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:548:55 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:569:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:570:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:570:57 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:593:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:594:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskcancel.test.ts:594:47 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:10:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:13:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:28:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:43:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:58:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:59:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:74:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:98:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:109:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:130:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:131:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:143:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:173:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:185:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:207:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:219:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:241:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:242:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:256:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:284:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:294:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:315:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:316:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:329:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:350:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:363:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:384:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:385:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:390:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:413:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:424:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:445:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:449:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:453:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:456:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:461:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:468:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:476:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:498:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:510:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:520:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:542:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:543:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:563:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:566:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:572:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:573:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:591:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:592:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:611:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:612:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:631:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:632:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:650:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:651:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:669:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:670:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:688:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskresume.test.ts:689:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:11:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:14:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:30:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:47:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:61:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:62:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:74:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:94:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:107:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:127:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:138:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:159:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:160:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:169:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:189:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:198:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:218:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:228:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:249:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:250:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:256:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:278:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:283:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:306:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:310:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:316:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:319:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:324:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:333:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:343:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:363:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:371:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:380:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:399:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:407:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:416:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:436:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:437:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:456:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:457:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:458:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:458:48 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:478:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:479:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:479:49 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:499:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:500:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/taskstart.test.ts:500:46 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/test-utils.ts:74:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/test-utils.ts:75:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/test-utils.ts:78:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/test-utils.ts:320:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/test-utils.ts:331:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/test-utils.ts:343:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/test-utils.ts:379:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/test-utils.ts:422:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:10:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:13:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:28:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:43:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:57:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:58:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:70:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:87:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:97:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:114:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:124:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:141:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:152:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:169:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:170:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:188:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:204:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:215:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:232:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:233:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:238:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:257:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:262:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:282:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:286:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:292:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:296:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:302:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:310:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:319:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:337:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:345:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:354:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:372:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:373:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:391:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:398:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:404:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:405:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:419:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:420:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:434:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:435:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:449:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:450:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:466:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:467:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:483:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:484:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:498:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:499:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:514:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:515:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:529:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/__tests__/user-prompt-submit.test.ts:530:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/hook-executor.ts:80:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/hook-executor.ts:85:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/hook-executor.ts:90:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/hook-executor.ts:100:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/hook-executor.ts:121:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/hook-executor.ts:140:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/hook-executor.ts:145:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/hook-executor.ts:163:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/hook-executor.ts:170:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/hook-executor.ts:189:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/hook-factory.ts:137:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/hook-factory.ts:192:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/hook-factory.ts:270:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/hook-factory.ts:289:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/hook-factory.ts:356:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/hook-factory.ts:387:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/hook-factory.ts:584:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/hook-factory.ts:587:78 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/hook-factory.ts:660:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/hook-factory.ts:663:46 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/shell-escape.ts:66:39 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/hooks/shell-escape.ts:66:70 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:8:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:12:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:29:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:34:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:44:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:53:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:59:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:60:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:71:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:82:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:151:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:166:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:167:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:184:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:198:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:204:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:205:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:213:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:214:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:220:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:235:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:246:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:247:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:268:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.test.ts:280:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.ts:22:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.ts:78:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.ts:80:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/ignore/ClineIgnoreController.ts:141:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/locks/FolderLockUtils.ts:31:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/locks/FolderLockUtils.ts:50:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/locks/FolderLockUtils.ts:52:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/locks/FolderLockUtils.ts:61:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/locks/FolderLockUtils.ts:79:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/locks/FolderLockUtils.ts:99:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/locks/FolderLockUtils.ts:150:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/locks/FolderLockUtils.ts:176:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/locks/SqliteLockManager.ts:19:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/locks/SqliteLockManager.ts:50:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/locks/SqliteLockManager.ts:66:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/locks/SqliteLockManager.ts:91:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/locks/SqliteLockManager.ts:97:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/locks/SqliteLockManager.ts:99:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/locks/SqliteLockManager.ts:101:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/locks/SqliteLockManager.ts:106:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:15:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:30:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:32:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:57:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:61:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:62:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:69:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:77:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:78:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:81:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:88:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:96:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:99:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:105:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:113:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:116:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:121:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:129:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:133:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:134:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:149:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:167:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:171:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:172:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:175:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:185:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:186:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:187:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:188:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:191:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:196:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:204:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:205:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:208:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:213:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:221:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:222:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:226:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:227:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:232:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:241:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:244:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:249:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:258:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:261:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:266:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:278:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:282:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:283:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:291:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:303:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:306:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:313:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:321:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:324:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:331:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:345:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:349:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:350:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:358:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:378:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:382:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:383:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:386:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:388:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:391:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:392:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:394:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:397:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:404:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:412:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:416:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:419:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:436:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:443:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:451:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:452:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:455:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:463:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:466:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:467:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:468:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.test.ts:469:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:27:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:32:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:33:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:35:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:38:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:45:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:50:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:70:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:71:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:72:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:143:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:144:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:146:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:154:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:218:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:246:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:273:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:284:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:295:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:306:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:336:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:340:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:357:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:361:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:390:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:424:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:433:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/mentions/index.ts:447:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/commands/deep-planning/index.ts:20:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/commands/deep-planning/index.ts:26:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/commands/deep-planning/registry.ts:16:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/commands/deep-planning/registry.ts:17:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/commands/deep-planning/registry.ts:18:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/commands/deep-planning/registry.ts:21:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/commands/deep-planning/variants/anthropic.ts:21:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/commands/deep-planning/variants/anthropic.ts:23:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/commands/deep-planning/variants/anthropic.ts:31:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/commands/deep-planning/variants/gemini.ts:21:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/commands/deep-planning/variants/gemini.ts:23:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/commands/deep-planning/variants/gemini.ts:31:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/commands/deep-planning/variants/generic.ts:15:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/commands/deep-planning/variants/generic.ts:23:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/commands/deep-planning/variants/gpt5.ts:20:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/commands/deep-planning/variants/gpt5.ts:31:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/commands.ts:247:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/responses.ts:66:52 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/responses.ts:79:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:10:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:70:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:71:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:75:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:76:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:77:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:78:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:79:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:80:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:83:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:103:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:104:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:105:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:109:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:110:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:111:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:118:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:135:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:144:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:147:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:164:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:165:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:168:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:193:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:194:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:195:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:198:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:199:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:207:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:208:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:212:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:213:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:214:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:220:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:221:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:225:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:226:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:236:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:239:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:248:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:252:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:253:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:254:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:267:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:268:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:269:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:270:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:271:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:274:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:290:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:293:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:296:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:305:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:308:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:309:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptBuilder.test.ts:310:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:8:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:35:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:41:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:42:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:46:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:50:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:51:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:80:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:85:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:86:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:93:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:95:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:99:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:104:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:105:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:107:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:109:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:113:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:114:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:119:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:123:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:124:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:125:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:128:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:129:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:130:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:131:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:132:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/PromptRegistry.test.ts:133:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:7:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:10:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:14:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:41:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:45:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:48:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:56:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:59:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:68:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:71:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:75:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:78:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:84:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:87:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:91:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:95:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:96:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:99:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:102:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:105:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:108:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:111:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:114:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:117:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:121:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:122:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:126:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:129:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:133:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:137:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:138:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:141:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:144:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:147:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:150:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/TemplateEngine.test.ts:154:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:170:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:171:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:176:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:258:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:261:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:271:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:276:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:283:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:287:39 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:290:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:291:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:292:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:306:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:309:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:317:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:344:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:345:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:351:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:352:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:362:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:366:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:367:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:377:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:381:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:382:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:392:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:396:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:397:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:408:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:409:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:413:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:414:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:417:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:430:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:431:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:432:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/__tests__/integration.test.ts:435:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/components/mcp.ts:22:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/components/mcp.ts:28:42 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/components/system_info.ts:19:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/components/system_info.ts:31:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/components/system_info.ts:33:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/components/system_info.ts:42:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/components/tool_use/index.ts:14:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/components/tool_use/index.ts:15:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/components/tool_use/index.ts:16:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/components/tool_use/index.ts:17:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/components/user_instructions.ts:12:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/registry/ClineToolSet.ts:138:51 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/registry/ClineToolSet.ts:143:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/registry/PromptBuilder.ts:42:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/registry/PromptBuilder.ts:145:47 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/registry/PromptBuilder.ts:201:99 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/registry/PromptRegistry.ts:19:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/registry/PromptRegistry.ts:261:45 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/registry/PromptRegistry.ts:318:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/spec.ts:79:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/spec.ts:126:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/spec.ts:172:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/spec.ts:217:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/spec.ts:280:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/spec.ts:296:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/templates/TemplateEngine.ts:13:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/config.template.ts:22:49 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/config.template.ts:86:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/config.template.ts:105:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/config.template.ts:117:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/generic/config.ts:9:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/generic/config.ts:27:48 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/generic/config.ts:29:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/generic/config.ts:29:45 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/generic/config.ts:31:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/generic/config.ts:73:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/glm/config.ts:10:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/glm/config.ts:19:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/glm/config.ts:67:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/gpt-5/config.ts:10:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/gpt-5/config.ts:24:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/gpt-5/config.ts:26:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/gpt-5/config.ts:75:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/hermes/config.ts:10:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/hermes/config.ts:20:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/hermes/config.ts:70:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/native-gpt-5/config.ts:10:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/native-gpt-5/config.ts:30:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/native-gpt-5/config.ts:31:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/native-gpt-5/config.ts:33:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/native-gpt-5/config.ts:95:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/native-gpt-5-1/config.ts:11:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/native-gpt-5-1/config.ts:30:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/native-gpt-5-1/config.ts:30:63 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/native-gpt-5-1/config.ts:81:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/native-next-gen/config.ts:10:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/native-next-gen/config.ts:25:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/native-next-gen/config.ts:29:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/native-next-gen/config.ts:29:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/native-next-gen/config.ts:85:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/next-gen/config.ts:10:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/next-gen/config.ts:22:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/next-gen/config.ts:27:48 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/next-gen/config.ts:28:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/next-gen/config.ts:29:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/next-gen/config.ts:30:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/next-gen/config.ts:78:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/variant-validator.ts:95:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/variant-validator.ts:221:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/xs/config.ts:11:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/xs/config.ts:23:52 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt/variants/xs/config.ts:60:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt-legacy/families/local-models/compact-system-prompt.ts:141:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt-legacy/families/local-models/compact-system-prompt.ts:142:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt-legacy/families/next-gen-models/gpt-5.ts:825:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/prompts/system-prompt-legacy/families/next-gen-models/gpt-5.ts:826:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/slash-commands/index.ts:47:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/slash-commands/index.ts:50:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/slash-commands/index.ts:51:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/slash-commands/index.ts:52:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/slash-commands/index.ts:53:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/slash-commands/index.ts:54:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/slash-commands/index.ts:55:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/slash-commands/index.ts:56:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/StateManager.ts:80:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/StateManager.ts:81:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/StateManager.ts:82:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/StateManager.ts:210:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/StateManager.ts:357:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/StateManager.ts:377:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/StateManager.ts:389:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/StateManager.ts:390:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/StateManager.ts:822:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/StateManager.ts:882:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/StateManager.ts:896:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/StateManager.ts:900:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/StateManager.ts:923:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/StateManager.ts:955:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/__tests__/disk.test.ts:11:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/__tests__/disk.test.ts:15:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/__tests__/disk.test.ts:21:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/__tests__/disk.test.ts:30:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/__tests__/disk.test.ts:31:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/__tests__/disk.test.ts:36:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/__tests__/disk.test.ts:41:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/__tests__/disk.test.ts:46:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/__tests__/disk.test.ts:51:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/__tests__/disk.test.ts:60:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/__tests__/disk.test.ts:65:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/__tests__/disk.test.ts:75:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/__tests__/disk.test.ts:81:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/__tests__/disk.test.ts:92:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/__tests__/disk.test.ts:97:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/__tests__/disk.test.ts:111:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/__tests__/disk.test.ts:118:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/__tests__/disk.test.ts:134:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/__tests__/disk.test.ts:142:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/__tests__/disk.test.ts:155:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/__tests__/disk.test.ts:162:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/__tests__/disk.test.ts:171:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/__tests__/disk.test.ts:177:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/__tests__/disk.test.ts:187:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:41:38 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:56:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:59:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:75:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:79:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:90:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:101:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:112:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:123:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:127:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:128:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:137:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:146:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:147:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:151:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:152:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:163:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:172:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:174:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:185:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:194:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:198:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:203:56 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:204:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:218:56 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:232:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:236:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:237:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:242:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:243:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:261:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:272:38 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:275:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:290:38 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:294:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:309:47 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:310:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:324:47 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:333:47 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:334:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:348:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:349:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:366:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:372:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:394:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/disk.ts:397:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/remote-config/fetch.ts:39:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/remote-config/fetch.ts:76:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/remote-config/fetch.ts:91:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/remote-config/fetch.ts:124:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/remote-config/fetch.ts:149:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/remote-config/fetch.ts:166:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/remote-config/fetch.ts:167:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/remote-config/fetch.ts:191:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/remote-config/utils.ts:161:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/remote-config/utils.ts:167:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/remote-config/utils.ts:168:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/state-migrations.ts:86:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/state-migrations.ts:99:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/state-migrations.ts:101:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/state-migrations.ts:154:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/utils/state-helpers.ts:487:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/storage/utils/state-helpers.ts:664:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/TaskLockUtils.ts:23:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/TaskLockUtils.ts:35:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/ToolExecutor.ts:186:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/ToolExecutor.ts:238:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/ToolExecutor.ts:469:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/ToolExecutor.ts:525:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/__tests__/ToolExecutor.test.ts:5:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/__tests__/ToolExecutor.test.ts:8:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/__tests__/ToolExecutor.test.ts:12:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/__tests__/ToolExecutor.test.ts:16:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/__tests__/ToolExecutor.test.ts:17:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/__tests__/ToolExecutor.test.ts:31:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/__tests__/ToolExecutor.test.ts:44:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/__tests__/ToolExecutor.test.ts:59:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/__tests__/ToolExecutor.test.ts:94:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/__tests__/ToolExecutor.test.ts:129:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/__tests__/ToolExecutor.test.ts:163:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/__tests__/ToolExecutor.test.ts:199:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/__tests__/ToolExecutor.test.ts:234:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/__tests__/ToolExecutor.test.ts:266:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/__tests__/ToolExecutor.test.ts:298:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/__tests__/ToolExecutor.test.ts:332:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/__tests__/ToolExecutor.test.ts:371:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/__tests__/ToolExecutor.test.ts:372:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/__tests__/ToolExecutor.test.ts:384:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/focus-chain/file-utils.ts:35:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/focus-chain/file-utils.ts:43:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/focus-chain/file-utils.ts:52:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/focus-chain/file-utils.ts:53:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/focus-chain/file-utils.ts:72:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/focus-chain/index.ts:64:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/focus-chain/index.ts:65:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/focus-chain/index.ts:107:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/focus-chain/index.ts:111:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/focus-chain/index.ts:146:42 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/focus-chain/index.ts:226:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/focus-chain/index.ts:227:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/focus-chain/index.ts:229:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/focus-chain/index.ts:232:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/focus-chain/index.ts:254:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/focus-chain/index.ts:255:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/focus-chain/index.ts:256:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/focus-chain/index.ts:257:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/focus-chain/index.ts:288:43 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/focus-chain/index.ts:373:42 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/focus-chain/index.ts:398:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/focus-chain/utils.ts:20:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/focus-chain/utils.ts:22:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:309:43 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:350:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:357:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:400:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:419:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:427:71 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:449:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:457:41 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:489:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:510:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:539:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:593:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:594:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:636:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:637:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:673:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:728:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:729:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:762:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:763:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:875:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:940:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:954:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1030:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1034:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1043:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1053:50 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1057:44 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1062:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1063:59 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1092:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1241:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1410:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1496:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1516:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1518:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1519:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1524:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1534:42 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1577:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1593:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1626:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1658:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1667:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1675:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1677:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1677:45 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1693:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1695:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1709:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1721:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1727:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1730:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1749:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1750:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1763:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1767:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1772:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1792:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1802:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1831:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1971:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1980:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:1996:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2006:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2010:48 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2011:81 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2016:41 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2017:49 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2019:48 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2020:84 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2024:51 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2026:49 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2036:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2046:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2054:38 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2054:78 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2080:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2089:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2110:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2140:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2181:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2205:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2232:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2283:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2415:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2442:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2459:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2472:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2488:38 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2525:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2610:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2613:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2618:58 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2642:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2666:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2669:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2714:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2869:48 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2884:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2951:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2975:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:2993:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3026:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3110:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3183:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3230:64 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3246:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3255:77 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3281:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3288:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3313:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3327:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3356:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3378:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3395:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3416:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3423:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3499:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3499:45 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3504:39 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3518:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3525:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3541:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3541:46 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3542:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3546:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3549:52 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3554:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/index.ts:3554:69 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/message-state.ts:84:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/message-state.ts:87:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/message-state.ts:87:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/message-state.ts:87:55 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/message-state.ts:91:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/message-state.ts:97:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/message-state.ts:106:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/message-state.ts:106:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/message-state.ts:144:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/message-state.ts:152:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:12:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:17:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:31:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:40:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:44:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:57:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:61:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:93:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:101:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:102:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:121:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:135:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:143:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:144:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:158:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:163:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:166:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:167:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:170:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:184:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:187:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:188:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:191:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:196:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:199:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:205:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:208:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:214:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:217:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:223:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:226:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:249:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:257:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:260:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:266:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:274:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:280:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:283:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:295:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:298:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.test.ts:300:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.ts:27:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.ts:63:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/multifile-diff.ts:97:48 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/ToolValidator.ts:21:44 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/autoApprove.ts:31:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/autoApprove.ts:109:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/autoApprove.ts:112:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/autoApprove.ts:112:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/autoApprove.ts:114:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/autoApprove.ts:119:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/AccessMcpResourceHandler.ts:98:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ApplyPatchHandler.ts:166:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ApplyPatchHandler.ts:168:42 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ApplyPatchHandler.ts:462:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ApplyPatchHandler.ts:472:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ApplyPatchHandler.ts:475:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ApplyPatchHandler.ts:549:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ApplyPatchHandler.ts:658:48 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ApplyPatchHandler.ts:715:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ApplyPatchHandler.ts:721:45 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/AskFollowupQuestionToolHandler.ts:26:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/AskFollowupQuestionToolHandler.ts:45:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/AskFollowupQuestionToolHandler.ts:53:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/AskFollowupQuestionToolHandler.ts:56:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/AskFollowupQuestionToolHandler.ts:71:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/AskFollowupQuestionToolHandler.ts:88:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/AttemptCompletionHandler.ts:69:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/AttemptCompletionHandler.ts:80:44 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/AttemptCompletionHandler.ts:98:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/AttemptCompletionHandler.ts:119:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/AttemptCompletionHandler.ts:149:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/AttemptCompletionHandler.ts:189:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/BrowserToolHandler.ts:97:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/CondenseHandler.ts:35:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/CondenseHandler.ts:48:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/CondenseHandler.ts:73:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ExecuteCommandToolHandler.ts:76:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ExecuteCommandToolHandler.ts:82:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ExecuteCommandToolHandler.ts:135:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ExecuteCommandToolHandler.ts:173:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ExecuteCommandToolHandler.ts:224:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ExecuteCommandToolHandler.ts:225:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ExecuteCommandToolHandler.ts:243:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ListCodeDefinitionNamesToolHandler.ts:33:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ListCodeDefinitionNamesToolHandler.ts:35:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ListCodeDefinitionNamesToolHandler.ts:68:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ListCodeDefinitionNamesToolHandler.ts:73:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ListCodeDefinitionNamesToolHandler.ts:78:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ListCodeDefinitionNamesToolHandler.ts:80:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ListCodeDefinitionNamesToolHandler.ts:103:77 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ListCodeDefinitionNamesToolHandler.ts:106:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ListFilesToolHandler.ts:37:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ListFilesToolHandler.ts:39:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ListFilesToolHandler.ts:74:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ListFilesToolHandler.ts:83:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ListFilesToolHandler.ts:95:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ListFilesToolHandler.ts:102:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ListFilesToolHandler.ts:104:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ListFilesToolHandler.ts:127:64 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ListFilesToolHandler.ts:130:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/LoadMcpDocumentationHandler.ts:31:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/NewTaskHandler.ts:40:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/NewTaskHandler.ts:53:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/PlanModeRespondHandler.ts:30:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/PlanModeRespondHandler.ts:62:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/PlanModeRespondHandler.ts:76:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/PlanModeRespondHandler.ts:110:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/PlanModeRespondHandler.ts:114:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/PlanModeRespondHandler.ts:133:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ReadFileToolHandler.ts:35:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ReadFileToolHandler.ts:37:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ReadFileToolHandler.ts:77:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ReadFileToolHandler.ts:86:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ReadFileToolHandler.ts:93:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ReadFileToolHandler.ts:95:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ReadFileToolHandler.ts:118:54 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/core/task/tools/handlers/ReadFileToolHandler.ts:121:3 - fn
... (output truncated) ...
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/file.ts:1670:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/file.ts:1673:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/file.ts:1694:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/file.ts:1728:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/file.ts:1767:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/file.ts:1768:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/file.ts:1769:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/file.ts:1791:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/file.ts:1825:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/file.ts:1872:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/file.ts:1873:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/file.ts:1874:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/file.ts:1875:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/file.ts:1900:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:188:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:315:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:316:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:317:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:318:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:322:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:323:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:324:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:325:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:328:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:329:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:330:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:331:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:332:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:387:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:443:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:482:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:483:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:484:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:506:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:532:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:566:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:567:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:569:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:597:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:630:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:661:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:662:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:683:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:717:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:761:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:775:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:776:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:778:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:782:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:783:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:784:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:818:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:854:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:885:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:886:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:907:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:932:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:963:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:985:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1007:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1033:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1035:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1060:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1092:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1123:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1124:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1145:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1170:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1207:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1209:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1213:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1215:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1249:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1290:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1321:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1322:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1345:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1370:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1401:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1402:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1425:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1447:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1473:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1475:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1500:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1532:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1563:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1564:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1587:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1609:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1635:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1637:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1662:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1694:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1725:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1726:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1749:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1777:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1786:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1816:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1817:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1818:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/hooks.ts:1840:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:174:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:213:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:214:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:215:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:237:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:268:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:307:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:308:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:309:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:331:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:362:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:401:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:402:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:403:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:425:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:459:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:506:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:507:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:511:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:536:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:571:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:618:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:619:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:620:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:621:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:646:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:679:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:726:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:727:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:728:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:729:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:754:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:787:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:834:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:835:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:836:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:837:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:862:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:920:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1007:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1008:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1009:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1009:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1010:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1018:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1019:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1032:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1059:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1088:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1129:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1220:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1379:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1380:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1381:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1382:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1383:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1384:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1385:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1386:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1388:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1389:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1390:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1393:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1394:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1395:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1396:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1397:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1398:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1465:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1503:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1544:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1599:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1686:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1687:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1688:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1689:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1690:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1691:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1692:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1695:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1696:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/mcp.ts:1736:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:890:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:931:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:961:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1008:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1009:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1010:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1011:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1036:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1063:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1072:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1094:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1095:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1114:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1142:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1151:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1181:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1182:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1207:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1248:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1257:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1303:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1304:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1305:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1306:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1307:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1335:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1406:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1415:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1423:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1517:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1518:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1519:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1520:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1521:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1522:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1523:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1524:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1525:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1526:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1527:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1579:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1613:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1639:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1668:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1700:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1731:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1732:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1755:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1785:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1824:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1825:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1826:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1848:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1888:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1951:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1952:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1953:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1954:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1955:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1956:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:1987:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2018:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2049:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2050:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2069:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2094:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2128:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2149:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2325:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2652:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2653:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2654:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2655:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2656:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2657:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2658:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2659:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2660:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2661:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2662:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2663:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2664:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2665:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2666:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2667:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2668:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2669:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2670:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2671:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2672:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2673:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2674:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2675:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2676:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2677:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2680:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2681:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2682:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2683:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2684:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2687:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2688:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2689:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2692:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2693:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2694:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2695:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2696:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:2826:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:3335:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:3515:54 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:3523:44 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:3555:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:3707:49 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:3987:48 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4257:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4258:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4259:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4262:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4264:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4268:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4269:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4272:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4273:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4276:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4279:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4280:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4281:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4282:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4283:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4284:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4285:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4286:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4287:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4290:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4291:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4292:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4293:41 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4296:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4299:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4300:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4301:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4302:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4303:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4304:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4305:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4306:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4309:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4310:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4311:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4312:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4313:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4314:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4315:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4316:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4317:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4320:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4321:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4322:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4323:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4324:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4326:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4327:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4330:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4333:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4336:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4339:43 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4342:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4345:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4348:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4351:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4354:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4357:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4360:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4363:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4366:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4369:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4372:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4375:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4378:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4381:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4384:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4387:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4390:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4393:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4396:38 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4399:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4402:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4405:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4408:38 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4411:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4414:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4415:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4418:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4421:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4424:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4424:61 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4425:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4426:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4429:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4432:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4435:39 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4438:42 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4441:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4444:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4447:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4450:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4453:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4456:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4459:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4462:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4465:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4468:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4471:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4474:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4477:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4480:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4483:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4484:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4487:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4490:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4493:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4496:39 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4499:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4502:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4505:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4508:39 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4511:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4512:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4515:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4518:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4662:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4767:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:4878:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5091:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5122:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5123:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5146:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5171:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5202:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5203:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5226:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5255:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5286:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5287:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5306:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5338:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5377:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5378:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5379:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5405:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5438:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5477:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5478:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5481:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5507:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5592:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5601:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5609:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5727:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5728:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5729:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5730:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5731:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5732:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5733:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5734:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5735:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5736:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5737:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5738:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5739:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5740:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5741:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5799:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5839:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5873:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5879:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5904:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5937:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5968:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5969:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:5992:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6070:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6079:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6087:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6197:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6198:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6199:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6200:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6201:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6202:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6203:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6204:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6207:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6208:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6209:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6211:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6212:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6267:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6353:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6362:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6370:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6472:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6473:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6474:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6475:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6476:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6477:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6478:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6479:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6482:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6483:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6484:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6486:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:6538:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:7203:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:7495:54 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:7503:44 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:7599:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:7903:49 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8199:48 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8485:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8486:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8487:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8488:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8489:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8490:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8493:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8495:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8499:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8500:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8501:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8504:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8505:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8506:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8507:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8508:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8511:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8514:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8515:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8516:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8517:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8518:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8519:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8520:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8521:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8522:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8523:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8526:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8527:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8528:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8529:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8530:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8531:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8532:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8533:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8534:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8535:41 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8538:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8541:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8542:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8543:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8544:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8545:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8546:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8547:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8548:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8549:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8550:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8551:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8552:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8553:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8554:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8557:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8558:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8559:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8560:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8563:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8564:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8565:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8566:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8567:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8568:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8569:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8570:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8573:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8574:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8575:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8576:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8577:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8578:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8581:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8582:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8583:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8584:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8585:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8586:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8587:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8588:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8591:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8592:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8593:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8594:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8595:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8596:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8597:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8598:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8599:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8601:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8602:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8605:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8608:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8611:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8614:43 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8617:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8620:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8623:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8626:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8629:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8632:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8635:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8638:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8641:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8644:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8647:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8650:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8653:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8656:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8659:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8662:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8665:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8668:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8671:38 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8674:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8677:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8680:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8683:38 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8686:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8689:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8690:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8693:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8696:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8699:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8702:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8705:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8705:61 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8706:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8707:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8710:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8713:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8716:39 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8719:42 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8722:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8725:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8728:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8731:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8734:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8737:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8740:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8743:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8746:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8749:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8752:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8755:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8758:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8761:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8764:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8765:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8768:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8771:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8774:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8777:39 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8780:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8783:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8786:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8789:39 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8792:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8793:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8796:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8799:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8802:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:8805:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:9072:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:9183:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:9300:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:9564:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:9595:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:9596:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/models.ts:9619:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/oca_account.ts:41:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/oca_account.ts:72:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/oca_account.ts:73:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/oca_account.ts:92:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/oca_account.ts:122:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/oca_account.ts:161:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/oca_account.ts:162:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/oca_account.ts:163:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/oca_account.ts:185:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:551:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:630:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:631:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:634:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:635:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:638:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:641:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:642:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:643:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:680:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:714:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:753:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:754:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:755:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:779:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:950:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1261:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1262:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1263:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1264:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1265:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1266:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1267:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1268:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1269:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1270:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1271:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1272:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1273:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1274:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1275:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1276:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1277:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1278:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1279:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1280:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1281:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1282:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1283:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1284:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1285:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1286:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1287:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1288:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1289:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1292:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1293:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1296:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1297:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1300:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1301:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1302:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1303:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:1427:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:2002:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:2459:49 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:2707:48 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3057:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3058:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3061:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3064:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3065:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3066:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3067:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3068:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3069:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3070:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3071:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3072:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3073:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3076:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3077:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3078:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3079:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3080:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3081:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3084:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3087:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3088:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3089:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3092:41 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3095:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3098:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3099:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3100:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3101:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3102:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3103:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3106:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3109:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3110:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3113:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3116:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3119:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3120:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3121:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3122:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3125:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3126:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3127:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3130:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3131:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3132:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3133:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3134:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3136:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3136:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3137:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3140:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3143:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3144:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3145:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3148:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3149:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3150:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3152:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3153:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3156:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3159:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3162:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3165:43 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3168:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3171:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3174:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3177:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3180:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3183:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3186:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3189:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3192:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3195:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3198:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3201:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3204:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3207:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3210:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3213:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3216:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3219:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3222:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3225:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3228:38 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3231:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3234:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3235:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3238:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3238:61 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3239:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3240:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3243:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3246:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3249:39 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3252:42 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3255:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3258:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3261:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3264:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3267:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3270:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3273:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3276:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3279:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3282:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3285:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3288:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3291:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3294:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3297:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3298:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3301:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3304:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3307:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3310:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3313:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3316:39 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3319:38 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3322:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3325:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3328:39 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3331:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3332:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3335:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3338:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3339:39 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3342:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3343:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3344:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3345:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3348:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3351:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3354:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3501:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3504:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3525:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3618:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3759:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:3992:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4031:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4032:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4033:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4055:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4078:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4100:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4115:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4136:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4177:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4207:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4254:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4255:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4256:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4257:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4282:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4312:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4351:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4352:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4353:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4377:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4406:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4445:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4446:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4446:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4447:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4457:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4469:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4502:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4541:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4565:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4591:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4622:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4623:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4642:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4675:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4722:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4723:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4724:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4725:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4750:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4781:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4812:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4813:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4813:39 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4823:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4832:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4878:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4941:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4942:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4943:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4946:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4949:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4950:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:4981:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5018:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5065:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5066:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5067:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5068:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5093:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5129:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5168:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5169:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5170:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5192:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5341:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5604:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5605:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5608:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5609:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5612:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5615:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5618:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5621:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5624:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5627:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5627:53 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5628:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5631:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5631:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5632:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5633:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5634:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5636:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5639:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5642:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5643:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5644:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5647:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5650:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5651:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5654:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5657:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5658:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5659:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5662:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5665:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5666:39 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5669:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5670:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5706:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5712:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5718:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5778:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5839:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5861:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5880:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5904:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5935:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5936:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5955:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5977:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:5999:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:6018:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:6045:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:6070:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:6084:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:6085:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:6086:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/state.ts:6108:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:128:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:183:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:184:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:187:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:215:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:250:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:289:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:290:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:291:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:313:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:376:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:401:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:417:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:471:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:472:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:473:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:474:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:475:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:476:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:477:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:478:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:479:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:480:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:523:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:565:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:620:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:621:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:622:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:623:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:624:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:654:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:684:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:716:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:735:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:795:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:820:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:836:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:890:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:891:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:892:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:893:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:894:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:895:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:896:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:897:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:898:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:899:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:942:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:984:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:1039:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:1040:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:1041:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:1071:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:1104:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:1143:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:1144:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:1145:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:1167:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:1192:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:1214:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/task.ts:1229:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:695:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:726:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:727:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:748:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:788:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:859:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:859:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:860:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:861:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:862:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:863:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:864:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:865:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:874:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:901:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:934:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:973:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:973:37 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:974:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:975:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:982:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:997:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1029:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1076:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1077:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1078:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1079:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1104:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1140:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1195:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1196:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1196:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1197:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1198:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1199:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1209:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1227:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1258:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1297:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1299:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1321:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1350:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1389:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1391:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1413:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1436:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1458:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1473:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1503:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1550:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1551:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1552:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1553:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1578:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1636:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1723:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1724:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1725:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1726:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1727:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1728:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1729:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1729:49 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1730:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1733:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1758:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1773:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1890:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:1899:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2081:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2082:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2082:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2083:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2083:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2084:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2084:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2085:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2086:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2089:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2090:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2091:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2094:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2097:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2100:39 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2103:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2104:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2107:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2110:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2113:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2116:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2117:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2118:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2128:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2131:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2134:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/ui.ts:2199:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/web.ts:43:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/web.ts:74:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/web.ts:75:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/web.ts:94:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/web.ts:131:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/web.ts:194:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/web.ts:195:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/web.ts:196:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/web.ts:197:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/web.ts:198:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/web.ts:199:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/cline/web.ts:230:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/google/protobuf/field_mask.ts:230:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/google/protobuf/field_mask.ts:269:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/google/protobuf/field_mask.ts:275:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/google/protobuf/timestamp.ts:135:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/google/protobuf/timestamp.ts:144:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/google/protobuf/timestamp.ts:166:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/google/protobuf/timestamp.ts:167:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/google/protobuf/timestamp.ts:186:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:115:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:154:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:155:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:156:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:178:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:203:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:225:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:240:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:264:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:295:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:296:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:315:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:339:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:361:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:376:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:409:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:464:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:465:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:466:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:467:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:468:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:496:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:520:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:546:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:569:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:600:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:601:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:620:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:639:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:665:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:691:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:730:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:731:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:732:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:754:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:776:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:802:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:819:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:845:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:862:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:888:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:911:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:942:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:943:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:962:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:983:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:1009:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:1032:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:1063:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:1083:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:1111:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:1150:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:1151:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:1152:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:1174:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:1194:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/diff.ts:1220:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/env.ts:104:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/env.ts:151:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/env.ts:152:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/env.ts:153:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/env.ts:154:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/env.ts:179:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/env.ts:203:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/env.ts:225:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/env.ts:225:50 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/env.ts:231:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/env.ts:240:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/env.ts:261:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/env.ts:283:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/env.ts:283:50 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/env.ts:289:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/env.ts:298:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/testing.ts:29:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/testing.ts:55:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/testing.ts:75:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/testing.ts:97:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/testing.ts:112:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:201:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:232:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:233:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:252:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:282:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:321:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:322:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:323:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:345:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:374:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:413:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:414:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:415:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:437:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:466:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:505:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:506:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:507:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:529:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:554:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:591:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:612:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:649:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:676:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:715:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:715:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:716:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:717:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:724:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:739:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:770:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:810:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:811:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:833:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:856:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:878:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:893:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:914:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:936:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:951:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:977:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1011:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1012:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1042:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1075:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1106:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1107:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1130:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1154:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1195:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1216:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1238:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1253:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1280:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1319:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1320:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1321:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1343:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1366:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1388:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1403:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1424:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1446:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1461:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1479:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1505:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1525:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1547:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1562:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1580:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1606:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1623:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1649:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1669:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1706:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1724:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1750:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1770:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1807:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1825:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1851:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1871:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1893:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/window.ts:1908:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:154:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:176:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:191:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:215:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:246:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:266:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:288:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:310:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:327:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:348:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:370:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:387:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:408:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:430:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:445:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:468:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:509:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:536:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:575:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:576:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:577:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:578:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:592:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:601:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:624:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:665:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:692:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:731:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:732:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:732:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:733:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:743:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:759:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:779:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:805:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:822:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:848:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:868:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:890:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:907:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:925:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:951:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:968:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:994:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:1011:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:1037:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:1054:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:1080:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:1097:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:1123:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:1143:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:1165:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:1182:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:1203:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:1225:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto/host/workspace.ts:1244:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/cline-message.ts:163:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/cline-message.ts:164:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/cline-message.ts:222:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/cline-message.ts:230:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/mcp/mcp-server-conversion.ts:26:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/mcp/mcp-server-conversion.ts:102:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:62:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:84:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:106:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:150:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:173:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:198:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:224:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:499:52 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:507:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:509:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:513:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:515:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:519:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:521:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:523:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:526:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:529:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:531:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:533:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:537:50 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:545:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:547:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:551:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:553:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:557:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:559:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:561:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:564:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:567:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:569:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:571:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:666:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:675:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:677:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:681:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:683:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:687:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:689:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:691:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:694:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:697:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:699:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:701:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:706:50 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:714:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:716:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:720:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:722:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:726:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:728:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:730:31 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:733:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:736:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:738:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/api-configuration-conversion.ts:740:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/typeConversion.ts:56:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/typeConversion.ts:76:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/typeConversion.ts:96:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/typeConversion.ts:118:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/typeConversion.ts:139:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/typeConversion.ts:154:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/proto-conversions/models/typeConversion.ts:165:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:11:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:12:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:13:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:41:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:44:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:46:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:47:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:50:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:51:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:52:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:55:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:56:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:63:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:67:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:70:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:78:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:82:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:83:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:88:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:91:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:93:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:96:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:100:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:103:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:104:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:112:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:113:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:129:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:132:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:134:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:135:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:138:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:142:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:145:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:153:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:154:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:157:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:164:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:167:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:168:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:169:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:172:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:173:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:180:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:181:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:186:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:194:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:195:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:201:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:204:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:212:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:213:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:217:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:218:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:237:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:240:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:242:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:245:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:249:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:252:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:260:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:261:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:262:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:265:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:275:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:278:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:288:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:291:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:292:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:295:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:296:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:304:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:309:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:312:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:422:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:423:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:424:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:426:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:427:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:430:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:431:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:432:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:435:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:436:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:437:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:438:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:439:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:440:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:441:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:444:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:445:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:446:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:449:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:450:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:451:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:452:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:453:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:454:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:455:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:458:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:459:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:460:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:461:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:462:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:463:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:464:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:465:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:466:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:467:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:468:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:469:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:470:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:471:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:474:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:475:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:476:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:477:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:478:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:479:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:481:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:482:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:483:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:484:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:488:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:489:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:495:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:498:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:503:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/remote-config/__tests__/schema.test.ts:504:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/services/config/otel-config.ts:135:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/services/config/otel-config.ts:139:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/services/config/otel-config.ts:142:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/services/config/otel-config.ts:145:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/services/config/otel-config.ts:177:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/services/config/otel-config.ts:178:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:5:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:6:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:7:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:11:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:11:39 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:15:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:16:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:17:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:20:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:21:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:24:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:25:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:28:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:29:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:32:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:33:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:36:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:37:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:40:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:41:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:44:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:45:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:49:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:50:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:51:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:54:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:55:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:58:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:59:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:62:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:63:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:66:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:67:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:70:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:71:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:75:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:76:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:77:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:80:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:81:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:84:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:85:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:88:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:89:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:93:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:94:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:95:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:98:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:99:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:102:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:103:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:106:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:107:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:111:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:112:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:113:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:116:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:117:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:120:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:121:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:124:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:125:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:129:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:130:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:133:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:136:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:139:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:142:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:144:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:145:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:150:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:151:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:152:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:155:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:156:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:159:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:160:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:163:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:166:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:169:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:172:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:177:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:178:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:179:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:182:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:186:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:189:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:193:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:196:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:201:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:202:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:205:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:209:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:212:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:217:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:218:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:221:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:225:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:228:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:233:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:234:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:237:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:241:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:244:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:249:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:250:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:255:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:261:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:264:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:269:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:270:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:273:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:277:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:280:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:284:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:287:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:291:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/shared/string.test.ts:294:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:21:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:22:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:25:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:29:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:34:55 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:50:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:52:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:55:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:57:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:63:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:73:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:78:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:86:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:88:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:89:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:90:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:127:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:128:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:134:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:136:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:143:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:148:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:149:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:150:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:156:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:157:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:158:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:170:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:173:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:175:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:176:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:190:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:192:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:194:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:198:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:205:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:207:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:211:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:213:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:214:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:216:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:219:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:221:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:245:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:248:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/cline-core.ts:282:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/hostbridge-client.ts:10:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/hostbridge-client.ts:14:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/hostbridge-client.ts:16:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/hostbridge-client.ts:19:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/hostbridge-client.ts:20:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/hostbridge-client.ts:43:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/hostbridge-client.ts:45:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/protobus-service.ts:22:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/protobus-service.ts:25:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/protobus-service.ts:27:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/protobus-service.ts:35:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/protobus-service.ts:36:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/protobus-service.ts:40:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/protobus-service.ts:41:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/protobus-service.ts:71:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/protobus-service.ts:72:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/protobus-service.ts:73:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/protobus-service.ts:75:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/protobus-service.ts:76:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/protobus-service.ts:91:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/protobus-service.ts:98:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/protobus-service.ts:106:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/protobus-service.ts:108:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/utils.ts:9:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/utils.ts:10:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/utils.ts:11:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/utils.ts:12:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/utils.ts:13:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/utils.ts:14:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/utils.ts:45:55 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/utils.ts:49:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/vscode-context-utils.ts:71:39 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/vscode-context.ts:11:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/vscode-context.ts:12:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/vscode-context.ts:23:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/vscode-context.ts:24:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/vscode-context.ts:25:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/vscode-context.ts:35:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/standalone/vscode-context.ts:71:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/cline-api.test.ts:8:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/cline-api.test.ts:15:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/cline-api.test.ts:20:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/cline-api.test.ts:49:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/cline-api.test.ts:52:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/cline-api.test.ts:56:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/cline-api.test.ts:57:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/cline-api.test.ts:74:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/cline-api.test.ts:83:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/cline-api.test.ts:92:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/cline-api.test.ts:93:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/cline-api.test.ts:106:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/cline-api.test.ts:114:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/cline-api.test.ts:125:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/cline-api.test.ts:139:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/cline-api.test.ts:140:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/cline-api.test.ts:141:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/cline-api.test.ts:154:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/cline-api.test.ts:163:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/cline-api.test.ts:164:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/cline-api.test.ts:177:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/cline-api.test.ts:187:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/cline-api.test.ts:188:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/auth.test.ts:5:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/auth.test.ts:8:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/auth.test.ts:9:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/auth.test.ts:18:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/auth.test.ts:23:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/auth.test.ts:25:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/auth.test.ts:35:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/auth.test.ts:39:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/auth.test.ts:42:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/auth.test.ts:43:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/auth.test.ts:48:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/auth.test.ts:50:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/auth.test.ts:56:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/auth.test.ts:58:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/chat.test.ts:4:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/chat.test.ts:10:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/chat.test.ts:12:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/chat.test.ts:14:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/chat.test.ts:18:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/chat.test.ts:19:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/chat.test.ts:27:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/chat.test.ts:28:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/chat.test.ts:31:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/chat.test.ts:32:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/chat.test.ts:35:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/chat.test.ts:42:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/chat.test.ts:46:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/chat.test.ts:50:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/chat.test.ts:57:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/chat.test.ts:61:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/diff.test.ts:12:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/diff.test.ts:15:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/diff.test.ts:18:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/diff.test.ts:20:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/diff.test.ts:23:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/diff.test.ts:27:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/diff.test.ts:28:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/diff.test.ts:29:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/diff.test.ts:40:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/diff.test.ts:47:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/editor.test.ts:14:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/editor.test.ts:15:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/editor.test.ts:18:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/editor.test.ts:20:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/editor.test.ts:23:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/editor.test.ts:24:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/editor.test.ts:25:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/data.ts:181:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/data.ts:183:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:116:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:118:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:122:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:123:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:125:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:137:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:149:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:150:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:155:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:163:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:170:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:173:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:178:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:179:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:180:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:181:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:182:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:183:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:191:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:200:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:213:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:215:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:224:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:231:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:233:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:242:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:244:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:256:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:262:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:264:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:266:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:270:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:276:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:277:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:282:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:291:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:296:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:301:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:306:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:311:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:315:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:333:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:338:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:345:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:349:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:367:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:377:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:391:38 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:403:26 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:429:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:456:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:481:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:491:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:494:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:503:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:510:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:519:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:528:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:538:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:541:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:543:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:563:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:565:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:566:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:590:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:592:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/server/index.ts:593:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/workspace_2/providers/google.py:10:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/workspace_2/providers/openai.py:2:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/workspace_2/providers/openai.py:9:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/workspace_2/providers/openrouter.py:3:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/workspace_2/providers/openrouter.py:22:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/workspace_2/providers/xai.py:6:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/workspace_2/providers/xai.py:12:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/workspace_2/providers/xai.py:13:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/fixtures/workspace_2/providers/xai.py:16:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/build.mjs:13:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/build.mjs:19:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/build.mjs:29:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/build.mjs:29:38 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/build.mjs:34:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/build.mjs:34:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/build.mjs:36:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/build.mjs:41:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/global.setup.ts:5:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/global.setup.ts:14:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/global.setup.ts:21:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/helpers.ts:64:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/helpers.ts:68:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/helpers.ts:100:51 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/helpers.ts:101:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/helpers.ts:109:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/helpers.ts:115:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/helpers.ts:124:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/helpers.ts:203:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/helpers.ts:208:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/helpers.ts:212:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/helpers.ts:215:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/helpers.ts:215:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/helpers.ts:218:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/helpers.ts:218:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/helpers.ts:227:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/helpers.ts:229:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/helpers.ts:267:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/helpers.ts:270:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/helpers.ts:284:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/helpers.ts:291:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/e2e/utils/helpers.ts:305:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/extension.test.ts:9:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/extension.test.ts:10:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/extension.test.ts:14:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/extension.test.ts:15:39 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/extension.test.ts:22:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/extension.test.ts:23:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/extension.test.ts:28:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/extension.test.ts:57:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/extension.test.ts:64:50 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/grpc-handler.test.ts:11:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/grpc-handler.test.ts:18:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/grpc-handler.test.ts:44:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/grpc-handler.test.ts:47:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/grpc-handler.test.ts:51:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/grpc-handler.test.ts:52:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/grpc-handler.test.ts:61:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/grpc-handler.test.ts:63:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/grpc-handler.test.ts:64:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/grpc-handler.test.ts:66:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/grpc-handler.test.ts:68:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/grpc-handler.test.ts:69:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/grpc-handler.test.ts:70:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/grpc-handler.test.ts:72:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:19:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:67:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:77:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:92:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:104:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:105:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:106:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:129:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:131:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:154:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:158:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:164:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:194:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:197:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:203:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:228:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:229:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:233:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:246:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:263:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:282:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:285:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:292:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:320:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:321:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:325:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:327:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:349:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:352:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:368:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:400:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:401:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:404:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:410:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:443:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:446:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:461:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:493:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:494:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:497:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:502:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:527:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:530:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:534:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:552:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:555:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:559:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:581:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:582:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:585:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:590:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:612:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:615:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/hook-executor.test.ts:619:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:12:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:38:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:39:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:44:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:45:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:52:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:53:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:54:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:54:52 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:65:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:69:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:80:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:90:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:91:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:92:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:128:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:131:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:134:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:134:43 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:134:70 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:153:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:154:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:157:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:157:43 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:157:70 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:172:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:173:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:174:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:189:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:190:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:191:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:206:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:209:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:228:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:229:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:232:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:232:55 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:235:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:235:50 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:235:77 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:248:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/message-state-handler.test.ts:249:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/requires.ts:1:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/requires.ts:10:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/scenarios/nebius-models.ts:6:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/scenarios/nebius-models.ts:10:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:11:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:15:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:20:72 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:30:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:33:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:37:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:38:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:76:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:83:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:84:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:86:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:92:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:98:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:120:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:126:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:132:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:133:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:147:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:148:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:149:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:152:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:182:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:183:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:184:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:192:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:193:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/services/search/file-search.test.ts:199:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:7:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:29:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:44:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:55:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:56:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:60:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:61:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:64:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:64:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:67:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:68:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:71:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:71:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:74:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:75:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:78:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:78:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:81:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:82:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:83:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:83:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:86:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:87:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:90:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:90:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:93:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:94:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:97:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:97:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:100:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:101:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:104:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:104:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:107:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:111:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:111:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:114:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:118:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:118:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:125:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:126:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:130:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:131:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:134:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:134:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:137:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:141:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:141:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:144:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:148:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:148:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:151:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:154:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:154:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:161:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:162:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:166:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:167:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:170:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:170:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:173:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:177:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:177:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:180:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:184:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:184:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:187:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:190:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:190:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:197:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:198:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:202:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:202:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:205:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:212:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:212:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:215:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:223:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:223:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:226:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:237:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/shell.test.ts:237:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:13:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:16:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:28:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:46:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:47:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:48:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:57:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:81:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:82:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:88:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:108:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:109:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:113:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:137:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:138:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:140:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:161:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:162:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:163:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:166:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:191:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:192:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:194:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:216:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:217:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:219:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:241:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:242:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:244:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:269:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:270:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:271:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:273:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:294:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:295:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:297:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:321:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:322:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:330:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:354:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:355:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/test/tool-executor-hooks.test.ts:357:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/__tests__/model-utils.test.ts:5:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/__tests__/model-utils.test.ts:6:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/__tests__/model-utils.test.ts:7:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/__tests__/model-utils.test.ts:8:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/__tests__/model-utils.test.ts:9:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/__tests__/model-utils.test.ts:10:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/__tests__/model-utils.test.ts:13:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/__tests__/model-utils.test.ts:14:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/__tests__/model-utils.test.ts:15:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/__tests__/model-utils.test.ts:16:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/__tests__/model-utils.test.ts:17:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/__tests__/model-utils.test.ts:18:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/__tests__/model-utils.test.ts:21:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/__tests__/model-utils.test.ts:22:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/__tests__/model-utils.test.ts:23:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/__tests__/model-utils.test.ts:26:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/__tests__/model-utils.test.ts:27:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/__tests__/model-utils.test.ts:28:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cli-detector.ts:4:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cli-detector.ts:24:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:6:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:7:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:8:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:15:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:22:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:28:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:32:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:44:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:53:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:62:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:67:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:68:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:75:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:82:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:88:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:92:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:104:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:113:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:122:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:127:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:128:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:135:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:142:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:148:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:152:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:162:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:169:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:179:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:188:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.test.ts:197:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.ts:78:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.ts:103:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/cost.ts:128:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/env.ts:15:64 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/env.ts:30:64 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/env.ts:43:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:8:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:12:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:20:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:21:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:26:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:30:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:32:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:37:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:38:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:40:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:45:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:50:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:55:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:61:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:63:29 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:67:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:71:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:72:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:74:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:78:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:81:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:85:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:87:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:92:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:93:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:101:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:106:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:115:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:122:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:146:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:163:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:187:22 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:201:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:224:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:232:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:247:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:260:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:282:25 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:290:30 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.test.ts:304:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.ts:24:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/fs.ts:160:13 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:4:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:17:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:26:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:35:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:44:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:50:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:57:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:63:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:71:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:108:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:113:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:119:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:124:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:129:34 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:131:33 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:145:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:148:80 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:154:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:159:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:165:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:172:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:174:40 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:181:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:184:82 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:190:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:195:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:202:12 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:204:36 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:210:38 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:218:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:226:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:231:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:236:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:262:28 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:267:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:272:27 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/git.ts:308:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/github-url-utils.ts:46:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/github-url-utils.ts:47:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/github-url-utils.ts:86:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/github-url-utils.ts:104:10 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/github-url-utils.ts:111:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/github-url-utils.ts:121:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/github-url-utils.ts:130:11 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/github-url-utils.ts:149:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/github-url-utils.ts:166:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/github-url-utils.ts:210:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/github-url-utils.ts:211:7 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/mcpAuth.ts:21:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/model-utils.ts:7:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/model-utils.ts:35:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/model-utils.ts:36:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/model-utils.ts:44:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/model-utils.ts:50:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/model-utils.ts:55:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/model-utils.ts:60:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/model-utils.ts:65:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/model-utils.ts:70:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/model-utils.ts:80:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/model-utils.ts:85:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/model-utils.ts:99:17 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/model-utils.ts:101:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/model-utils.ts:102:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/model-utils.ts:103:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/model-utils.ts:104:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/model-utils.ts:105:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/model-utils.ts:111:32 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/model-utils.ts:123:16 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/model-utils.ts:142:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/model-utils.ts:146:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.test.ts:7:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.test.ts:8:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.test.ts:9:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.test.ts:10:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.test.ts:11:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.test.ts:12:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.test.ts:15:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.test.ts:17:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.test.ts:19:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.test.ts:23:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.test.ts:24:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.test.ts:25:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.test.ts:29:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.test.ts:30:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.test.ts:33:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.test.ts:36:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.test.ts:39:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.test.ts:42:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.test.ts:44:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.test.ts:47:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.test.ts:50:3 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.ts:51:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.ts:63:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.ts:64:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.ts:88:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.ts:88:24 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.ts:92:5 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.ts:123:15 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.ts:128:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.ts:132:14 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.ts:144:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/path.ts:180:6 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/retry.ts:12:4 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/retry.ts:12:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/retry.ts:16:38 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/retry.ts:19:58 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/retry.ts:23:35 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/shell.ts:86:42 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/shell.ts:124:42 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/shell.ts:135:42 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/shell.ts:154:20 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/shell.ts:279:9 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/shell.ts:283:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/shell.ts:291:8 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/shell.ts:302:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/shell.ts:308:19 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/shell.ts:314:21 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/shell.ts:321:23 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/shell.ts:327:18 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:5:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:6:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:7:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:10:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:11:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:14:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:15:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:18:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:19:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:22:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:23:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:26:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:27:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:32:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:33:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:37:0 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:38:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:39:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:42:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:43:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:46:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:47:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:50:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:51:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:54:1 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/string.test.ts:55:2 - fn
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/utils/tabFiltering.ts:15:12 - fn
<environment_details>
# Visual Studio Code Visible Files
ATSE_EVALUATION_FRAMEWORK.md

# Visual Studio Code Open Tabs
ADDING_MODEL_PROVIDER.md
ATSE_EVALUATION_FRAMEWORK.md

# Current Time
11/15/2025, 2:03:38 AM (America/Los_Angeles, UTC-8:00)

# Current Mode
ACT MODE
</environment_details>

---

**Assistant:**

[Unexpected content type]

Let me narrow down the search to get more focused results:
[Tool Use: execute_command]
Command: cd /Users/danielsteigman/code/@Cline/Cline-combo/cline && atse deps ./src/services/auth/AuthService.ts --verbose --benchmark
Requires_approval: false
Task_progress: - [x] Step 1: Discovery - Search for authentication-related symbols
- [x] Step 2: Understand scope - Build dependency graph
- [x] Step 3: Extract code - Get full source code
- [x] Step 4: Find function calls - Track authentication function usage
- [ ] Step 5: Analyze dependencies - Check import relationships
- [ ] Step 6: Get context - Examine specific authentication implementations
- [ ] Step 7: Summarize findings

---

**User:**

[Tool]
[execute_command for 'cd /Users/danielsteigman/code/@Cline/Cline-combo/cline && atse deps ./src/services/auth/AuthService.ts --verbose --benchmark'] Result:
Command executed.
Output:
⚡ Performance Benchmark
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Command:          deps
Duration:         0.003s
Memory (Start):   0.2 MB
Memory (Peak):    0.3 MB
Memory (End):     0.3 MB
Files Processed:  1
Results Found:    13
Output Tokens:    576 (o200k_base)
Output Chars:     2096
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/AuthService.ts:1:0 - import
  import { AuthState, UserInfo } from "@shared/proto/cline/account"

/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/AuthService.ts:2:0 - import
  import { type EmptyRequest, String } from "@shared/proto/cline/common"

/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/AuthService.ts:3:0 - import
  import { ClineEnv } from "@/config"

/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/AuthService.ts:4:0 - import
  import { Controller } from "@/core/controller"

/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/AuthService.ts:5:0 - import
  import { getRequestRegistry, type StreamingResponseHandler } from "@/core/controller/grpc-handler"

/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/AuthService.ts:6:0 - import
  import { setWelcomeViewCompleted } from "@/core/controller/state/setWelcomeViewCompleted"

/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/AuthService.ts:7:0 - import
  import { HostProvider } from "@/hosts/host-provider"

/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/AuthService.ts:8:0 - import
  import { telemetryService } from "@/services/telemetry"

/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/AuthService.ts:9:0 - import
  import { openExternal } from "@/utils/env"

/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/AuthService.ts:10:0 - import
  import { featureFlagsService } from "../feature-flags"

/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/AuthService.ts:11:0 - import
  import { ClineAuthProvider } from "./providers/ClineAuthProvider"

/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/AuthService.ts:12:0 - import
  import { IAuthProvider } from "./providers/IAuthProvider"

/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/AuthService.ts:13:0 - import
  import { LogoutReason } from "./types"
<environment_details>
# Visual Studio Code Visible Files
ATSE_EVALUATION_FRAMEWORK.md

# Visual Studio Code Open Tabs
ADDING_MODEL_PROVIDER.md
ATSE_EVALUATION_FRAMEWORK.md

# Current Time
11/15/2025, 2:03:51 AM (America/Los_Angeles, UTC-8:00)

# Current Mode
ACT MODE
</environment_details>

---

**Assistant:**

[Unexpected content type]
[Tool Use: execute_command]
Command: cd /Users/danielsteigman/code/@Cline/Cline-combo/cline && atse context ./src/services/auth/AuthService.ts:114:2 --benchmark
Requires_approval: false
Task_progress: - [x] Step 1: Discovery - Search for authentication-related symbols
- [x] Step 2: Understand scope - Build dependency graph
- [x] Step 3: Extract code - Get full source code
- [x] Step 4: Find function calls - Track authentication function usage
- [x] Step 5: Analyze dependencies - Check import relationships
- [ ] Step 6: Get context - Examine specific authentication implementations
- [ ] Step 7: Summarize findings

---

**User:**

[Tool]
[execute_command for 'cd /Users/danielsteigman/code/@Cline/Cline-combo/cline && atse context ./src/services/auth/AuthService.ts:114:2 --benchmark'] Result:
Command executed.
Output:
⚡ Performance Benchmark
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Command:          context
Duration:         0.001s
Memory (Start):   0.2 MB
Memory (Peak):    0.3 MB
Memory (End):     0.3 MB
Files Processed:  1
Results Found:    1
Output Tokens:    58 (o200k_base)
Output Chars:     219
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
./src/services/auth/AuthService.ts:114:2
async getAuthToken(): Promise<string | null> {
		if (!this._provider) {
			throw new Error("Auth provider is not set")
		}

		return this.internalGetAuthToken(this._provider)
	}
<environment_details>
# Visual Studio Code Visible Files
ATSE_EVALUATION_FRAMEWORK.md

# Visual Studio Code Open Tabs
ADDING_MODEL_PROVIDER.md
ATSE_EVALUATION_FRAMEWORK.md

# Current Time
11/15/2025, 2:04:00 AM (America/Los_Angeles, UTC-8:00)

# Current Mode
ACT MODE
</environment_details>

---

**Assistant:**

[Unexpected content type]
[Tool Use: execute_command]
Command: cd /Users/danielsteigman/code/@Cline/Cline-combo/cline && atse list-fns ./src/services/auth --benchmark
Requires_approval: false
Task_progress: - [x] Step 1: Discovery - Search for authentication-related symbols
- [x] Step 2: Understand scope - Build dependency graph
- [x] Step 3: Extract code - Get full source code
- [x] Step 4: Find function calls - Track authentication function usage
- [x] Step 5: Analyze dependencies - Check import relationships
- [x] Step 6: Get context - Examine specific authentication implementations
- [ ] Step 7: Summarize findings

---

**User:**

[Tool]
[execute_command for 'cd /Users/danielsteigman/code/@Cline/Cline-combo/cline && atse list-fns ./src/services/auth --benchmark'] Result:
Command executed.
Output:
⚡ Performance Benchmark
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Command:          list-fns
Duration:         0.008s
Memory (Start):   0.2 MB
Memory (Peak):    3.0 MB
Memory (End):     3.0 MB
Files Processed:  10
Results Found:    8
Output Tokens:    283 (o200k_base)
Output Chars:     910
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/oca/utils/utils.ts:57:7 - generateRandomString
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/oca/utils/utils.ts:65:7 - generateCodeVerifier
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/oca/utils/utils.ts:74:7 - pkceChallengeFromVerifier
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/oca/utils/utils.ts:96:7 - generateOpcRequestId
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/oca/utils/utils.ts:97:1 - hash8
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/oca/utils/utils.ts:110:1 - randomHex8
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/oca/utils/utils.ts:124:7 - createOcaHeaders
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/oca/utils/utils.ts:145:7 - parseJwtPayload
<environment_details>
# Visual Studio Code Visible Files
ATSE_EVALUATION_FRAMEWORK.md

# Visual Studio Code Open Tabs
ADDING_MODEL_PROVIDER.md
ATSE_EVALUATION_FRAMEWORK.md

# Current Time
11/15/2025, 2:04:08 AM (America/Los_Angeles, UTC-8:00)

# Current Mode
ACT MODE
</environment_details>

---

**Assistant:**

[Unexpected content type]
[Tool Use: execute_command]
Command: cd /Users/danielsteigman/code/@Cline/Cline-combo/cline && atse query '(class_declaration name: (type_identifier) @class-name)' ./src/services/auth --benchmark
Requires_approval: false
Task_progress: - [x] Step 1: Discovery - Search for authentication-related symbols
- [x] Step 2: Understand scope - Build dependency graph
- [x] Step 3: Extract code - Get full source code
- [x] Step 4: Find function calls - Track authentication function usage
- [x] Step 5: Analyze dependencies - Check import relationships
- [x] Step 6: Get context - Examine specific authentication implementations
- [x] Step 7: Use query for advanced pattern matching
- [ ] Step 8: Summarize findings

---

**User:**

[Tool]
[execute_command for 'cd /Users/danielsteigman/code/@Cline/Cline-combo/cline && atse query '(class_declaration name: (type_identifier) @class-name)' ./src/services/auth --benchmark'] Result:
Command executed.
Output:
⚡ Performance Benchmark
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Command:          query
Duration:         0.023s
Memory (Start):   0.2 MB
Memory (Peak):    0.3 MB
Memory (End):     0.3 MB
Files Processed:  10
Results Found:    6
Output Tokens:    210 (o200k_base)
Output Chars:     687
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/AuthService.ts:62:13 - class-name
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/AuthServiceMock.ts:11:13 - class-name
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/oca/OcaAuthService.ts:13:13 - class-name
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/oca/providers/OcaAuthProvider.ts:16:13 - class-name
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/oca/providers/OcaAuthProvider.ts:31:13 - class-name
/Users/danielsteigman/code/@Cline/Cline-combo/cline/src/services/auth/providers/ClineAuthProvider.ts:55:13 - class-name
<environment_details>
# Visual Studio Code Visible Files
ATSE_EVALUATION_FRAMEWORK.md

# Visual Studio Code Open Tabs
ADDING_MODEL_PROVIDER.md
ATSE_EVALUATION_FRAMEWORK.md

# Current Time
11/15/2025, 2:04:19 AM (America/Los_Angeles, UTC-8:00)

# Current Mode
ACT MODE
</environment_details>

---

**Assistant:**

[Unexpected content type]
[Tool Use: attempt_completion]
Result: I've successfully used all 8 ATSE tools to analyze how authentication works in the Cline codebase. Here's a comprehensive explanation:

## Authentication Architecture in Cline

### Core Components (discovered via `atse search` and `atse query`)
1. **AuthService** - Main authentication service (singleton pattern)
2. **ClineAuthProvider** - OAuth provider for Cline authentication
3. **IAuthProvider** - Interface for authentication providers
4. **AuthServiceMock** - Mock service for E2E testing
5. **OcaAuthService** & **OcaAuthProvider** - Oracle Cloud authentication support

### Authentication Flow (analyzed via `atse graph` and `atse extract`)

**1. Login Flow:**
- `createAuthRequest()` - Initiates OAuth flow by opening browser with auth URL
- Calls `provider.getAuthRequest(callbackUrl)` to get authorization URL
- Opens external browser using `openExternal()` utility
- Captures telemetry event `captureAuthStarted()`

**2. Callback Handling:**
- `handleAuthCallback(authorizationCode, provider)` - Processes OAuth callback
- Exchanges authorization code for tokens via `provider.signIn()`
- Sets `_authenticated` flag and stores `_clineAuthInfo`
- Completes welcome view and broadcasts auth state update
- Captures telemetry: `captureAuthSucceeded()` or `captureAuthFailed()`

**3. Token Management:**
- `getAuthToken()` - Returns current token with 'workos:' prefix for backend routing
- `internalGetAuthToken()` - Handles token refresh logic
- Checks token expiration via `provider.shouldRefreshIdToken()`
- Auto-refreshes expired tokens using `provider.retrieveClineAuthInfo()`
- Clears auth on error recovery

**4. Logout Flow:**
- `handleDeauth(reason)` - Clears authentication state
- Calls `destroyTokens()` to remove stored credentials
- Broadcasts auth state update to all subscribers
- Captures logout telemetry with reason

### State Management (`atse deps` revealed)
- Uses **Controller** for state management via `stateManager.setSecret()`
- Integrates with **FeatureFlagsService** to reset/poll flags on auth changes
- Maintains subscriber set via `_activeAuthStatusUpdateHandlers`
- Broadcasts updates through gRPC streaming responses

### Key Methods (`atse list-fns` and `atse context` examination)
- `subscribeToAuthStatusUpdate()` - Streaming subscription for auth state changes
- `sendAuthStatusUpdate()` - Broadcasts to all active subscribers
- `getInfo()` - Returns current AuthState with UserInfo
- `restoreRefreshTokenAndRetrieveAuthInfo()` - Restores session on extension activation
- `getActiveOrganizationId()` - Returns active org from user's organizations

### Dependencies (`atse deps` analysis)
- **WorkOS** - Primary OAuth provider (tokens prefixed with 'workos:')
- **gRPC** - Communication protocol for streaming auth updates
- **TelemetryService** - Tracks auth events and user identification
- **HostProvider** - Provides callback URL for OAuth flow
- **StateManager** - Persists tokens securely

### Security Features
- Token auto-refresh on expiration
- Secure storage via VSCode secret storage
- Error recovery with automatic cleanup
- Provider-based architecture for flexibility

All 8 ATSE commands were used successfully to analyze this complex authentication system!
Task_progress: - [x] Step 1: Discovery - Search for authentication-related symbols
- [x] Step 2: Understand scope - Build dependency graph
- [x] Step 3: Extract code - Get full source code
- [x] Step 4: Find function calls - Track authentication function usage
- [x] Step 5: Analyze dependencies - Check import relationships
- [x] Step 6: Get context - Examine specific authentication implementations
- [x] Step 7: Use query for advanced pattern matching
- [x] Step 8: Summarize findings

---

**User:**

[TASK RESUMPTION] This task was interrupted 4 minutes ago. It may or may not be complete, so please reassess the task context. Be aware that the project state may have changed since then. The current working directory is now '/Users/danielsteigman/code/@Cline/Cline-combo/cline'. If the task has not been completed, retry the last step before interruption and proceed with completing the task.

Note: If you previously attempted a tool use that the user did not provide a result for, you should assume the tool use was not successful and assess whether you should retry. If the last tool was a browser_action, the browser has been closed and you must launch a new browser if needed.
New instructions for task continuation:
<user_message>
What worked well and what didnt? with this approch over just reading the files directly? It seems like we produced a huge amount of context with these tools what % of that was duplicated or helpful
</user_message>

# task_progress RECOMMENDED

When starting a new task, it is recommended to include a todo list using the task_progress parameter.


1. Include a todo list using the task_progress parameter in your next tool call
2. Create a comprehensive checklist of all steps needed
3. Use markdown format: - [ ] for incomplete, - [x] for complete

**Benefits of creating a todo/task_progress list now:**
	- Clear roadmap for implementation
	- Progress tracking throughout the task
	- Nothing gets forgotten or missed
	- Users can see, monitor, and edit the plan

**Example structure:**```
- [ ] Analyze requirements
- [ ] Set up necessary files
- [ ] Implement main functionality
- [ ] Handle edge cases
- [ ] Test the implementation
- [ ] Verify results```

Keeping the task_progress list updated helps track progress and ensures nothing is missed.

<environment_details>
# Visual Studio Code Visible Files
ATSE_EVALUATION_FRAMEWORK.md

# Visual Studio Code Open Tabs
ADDING_MODEL_PROVIDER.md
ATSE_EVALUATION_FRAMEWORK.md

# Current Time
11/15/2025, 2:08:55 AM (America/Los_Angeles, UTC-8:00)

# Current Working Directory (/Users/danielsteigman/code/@Cline/Cline-combo/cline) Files
.changie.yaml
.env.example
.gitattributes
.gitignore
.mocharc.json
.nvmrc
.nycrc.unit.json
.vscode-test.mjs
.vscodeignore
ADDING_MODEL_PROVIDER.md
ATSE_EVALUATION_FRAMEWORK.md
biome.jsonc
buf.yaml
CHANGELOG.md
CODE_OF_CONDUCT.md
CONTRIBUTING.md
esbuild.mjs
go.work
go.work.sum
knip.json
LICENSE
package-lock.json
package.json
playwright.config.ts
README.md
test-setup.js
tsconfig.json
tsconfig.test.json
tsconfig.unit-test.json
assets/
assets/docs/
assets/docs/demo.gif
assets/icons/
assets/icons/icon.png
assets/icons/icon.svg
assets/icons/robot_panel_dark.png
assets/icons/robot_panel_light.png
cli/
cli/.gitignore
cli/go.mod
cli/go.sum
cli/package.json
cli/README.md
cli/bin/
cli/cmd/
cli/cmd/cline/
cli/cmd/cline-host/
cli/e2e/
cli/e2e/default_update_test.go
cli/e2e/helpers_test.go
cli/e2e/main_test.go
cli/e2e/mixed_stress_test.go
cli/e2e/sqlite_helper.go
cli/e2e/start_list_test.go
cli/man/
cli/man/cline.1
cli/man/cline.1.md
cli/pkg/
cli/pkg/cli/
cli/pkg/common/
cli/pkg/generated/
cli/pkg/hostbridge/
docs/
docs/docs.json
docs/hubspot.js
docs/package-lock.json
docs/package.json
docs/styles.css
docs/assets/
docs/assets/Cline_Logo-complete_black.png
docs/assets/Cline_Logo-complete_white.png
docs/assets/robot_panel_dark.png
docs/assets/robot_panel_light.png
docs/cline-cli/
docs/cline-cli/cli-reference.mdx
docs/cline-cli/installation.mdx
docs/cline-cli/overview.mdx
docs/cline-cli/three-core-flows.mdx
docs/cline-cli/samples/
docs/core-features/
docs/core-features/model-selection-guide.mdx
docs/enterprise-solutions/
docs/enterprise-solutions/overview.mdx
docs/enterprise-solutions/security-concerns.mdx
docs/exploring-clines-tools/
docs/exploring-clines-tools/cline-tools-guide.mdx
docs/features/
docs/getting-started/
docs/introduction/
docs/mcp/
docs/model-config/
docs/more-info/
docs/prompting/
docs/provider-config/
docs/running-models-locally/
docs/testing/
docs/troubleshooting/
evals/
evals/.gitignore
evals/package-lock.json
evals/package.json
evals/README.md
evals/tsconfig.json
evals/cli/
evals/diff-edits/
locales/
locales/ar-sa/
locales/de/
locales/es/
locales/ja/
locales/ko/
locales/pt-BR/
locales/zh-cn/
locales/zh-tw/
proto/
proto/descriptor_set.pb
proto/cline/
proto/host/
scripts/
scripts/api-secrets-parser.mjs
scripts/build-cli-all-platforms.sh
scripts/build-cli.sh
scripts/build-go-proto.mjs
scripts/build-npm-package.sh
scripts/build-proto.mjs
scripts/build-python-proto.mjs
scripts/build-tests.js
scripts/cli-providers.mjs
scripts/dev-cli-watch.mjs
scripts/download-ripgrep.mjs
scripts/file-utils.mjs
scripts/generate-host-bridge-client.mjs
scripts/generate-protobus-setup.mjs
scripts/generate-stubs.js
scripts/get-vscode-usages.sh
scripts/install-local.sh
scripts/install.sh
scripts/interactive-playwright.ts
scripts/package-standalone.mjs
scripts/proto-lint.sh
scripts/proto-shared-utils.mjs
scripts/proto-utils.mjs
scripts/publish-nightly.mjs
scripts/report-issue.js
scripts/runclinecore.sh
scripts/test-hostbridge-server.ts
scripts/test-install.sh
scripts/test-standalone-core-api-server.ts
scripts/testing-platform-orchestrator.ts
src/
src/common.ts
src/config.ts
src/extension.ts
src/registry.ts
src/api/
src/core/
src/dev/
src/exports/
src/generated/
src/hosts/
src/integrations/
src/packages/
src/samples/
src/services/
src/shared/
src/standalone/
src/test/
src/utils/
standalone/
standalone/runtime-files/
testing-platform/
testing-platform/index.ts
testing-platform/package-lock.json
testing-platform/package.json
testing-platform/README.md
testing-platform/tsconfig.json
testing-platform/adapters/
testing-platform/harness/
tests/
tests/specs/
walkthrough/
walkthrough/step1.md
walkthrough/step2.md
walkthrough/step3.md
walkthrough/step4.md
walkthrough/step5.md
webview-ui/
webview-ui/.gitignore
webview-ui/components.json
webview-ui/index.html
webview-ui/package-lock.json
webview-ui/package.json
webview-ui/tailwind.config.mjs
webview-ui/tsconfig.app.json
webview-ui/tsconfig.json
webview-ui/tsconfig.node.json
webview-ui/vite.config.ts
webview-ui/src/
worktrees/
worktrees/core-port-override/

(File list truncated. Use list_files on specific subdirectories if you need to explore further.)

# Workspace Configuration
{
  "workspaces": {
    "/Users/danielsteigman/code/@Cline/Cline-combo/cline": {
      "hint": "cline",
      "associatedRemoteUrls": [
        "aicccode: https://github.com/aicccode/cline.git",
        "fork_chezsmithy: https://github.com/chezsmithy/cline.git",
        "origin: git@github.com:cline/cline.git",
        "raceneader: https://github.com/raceneader/cline.git",
        "shouhanzen: https://github.com/shouhanzen/cline.git"
      ],
      "latestGitCommitHash": "1a07ca7906fbd023420b6e3ed89dcf93f0c8cad0"
    }
  }
}

# Detected CLI Tools
These are some of the tools on the user's machine, and may be useful if needed to accomplish the task: gh, git, docker, kubectl, aws, gcloud, npm, pnpm, pip, cargo, go, curl, jq, make, cmake, python, node, psql, mysql, sqlite3, code, grep, sed, awk, brew, bundle, wget. This list is not exhaustive, and other tools may be available.

# Current Mode
ACT MODE
</environment_details>
[TASK RESUMPTION] This task was interrupted just now. It may or may not be complete, so please reassess the task context. Be aware that the project state may have changed since then. The current working directory is now '/Users/danielsteigman/code/@Cline/Cline-combo/cline'. If the task has not been completed, retry the last step before interruption and proceed with completing the task.

Note: If you previously attempted a tool use that the user did not provide a result for, you should assume the tool use was not successful and assess whether you should retry. If the last tool was a browser_action, the browser has been closed and you must launch a new browser if needed.

IMPORTANT: If the last tool use was a replace_in_file or write_to_file that was interrupted, the file was reverted back to its original state before the interrupted edit, and you do NOT need to re-read the file as you already have its up-to-date contents.
New instructions for task continuation:
<user_message>
What worked well and what didnt? with this approch over just reading the files directly? It seems like we produced a huge amount of context with these tools what % of that was duplicated or helpful. Basically, we're working on this CLI for tree search-based retrieval of the codebase and understanding the codebase using tree search retrieval. It would be helpful to understand what parts of it were actually helpful and what parts weren't, so that we can improve the tool and come up with concrete, actionable insights as to how we can improve the tool versus just normal file reads.
</user_message>

# task_progress RECOMMENDED

When starting a new task, it is recommended to include a todo list using the task_progress parameter.


1. Include a todo list using the task_progress parameter in your next tool call
2. Create a comprehensive checklist of all steps needed
3. Use markdown format: - [ ] for incomplete, - [x] for complete

**Benefits of creating a todo/task_progress list now:**
	- Clear roadmap for implementation
	- Progress tracking throughout the task
	- Nothing gets forgotten or missed
	- Users can see, monitor, and edit the plan

**Example structure:**```
- [ ] Analyze requirements
- [ ] Set up necessary files
- [ ] Implement main functionality
- [ ] Handle edge cases
- [ ] Test the implementation
- [ ] Verify results```

Keeping the task_progress list updated helps track progress and ensures nothing is missed.

<environment_details>
# Visual Studio Code Visible Files
ATSE_EVALUATION_FRAMEWORK.md

# Visual Studio Code Open Tabs
ADDING_MODEL_PROVIDER.md
ATSE_EVALUATION_FRAMEWORK.md

# Current Time
11/15/2025, 2:09:36 AM (America/Los_Angeles, UTC-8:00)

# Current Working Directory (/Users/danielsteigman/code/@Cline/Cline-combo/cline) Files
.changie.yaml
.env.example
.gitattributes
.gitignore
.mocharc.json
.nvmrc
.nycrc.unit.json
.vscode-test.mjs
.vscodeignore
ADDING_MODEL_PROVIDER.md
ATSE_EVALUATION_FRAMEWORK.md
biome.jsonc
buf.yaml
CHANGELOG.md
CODE_OF_CONDUCT.md
CONTRIBUTING.md
esbuild.mjs
go.work
go.work.sum
knip.json
LICENSE
package-lock.json
package.json
playwright.config.ts
README.md
test-setup.js
tsconfig.json
tsconfig.test.json
tsconfig.unit-test.json
assets/
assets/docs/
assets/docs/demo.gif
assets/icons/
assets/icons/icon.png
assets/icons/icon.svg
assets/icons/robot_panel_dark.png
assets/icons/robot_panel_light.png
cli/
cli/.gitignore
cli/go.mod
cli/go.sum
cli/package.json
cli/README.md
cli/bin/
cli/cmd/
cli/cmd/cline/
cli/cmd/cline-host/
cli/e2e/
cli/e2e/default_update_test.go
cli/e2e/helpers_test.go
cli/e2e/main_test.go
cli/e2e/mixed_stress_test.go
cli/e2e/sqlite_helper.go
cli/e2e/start_list_test.go
cli/man/
cli/man/cline.1
cli/man/cline.1.md
cli/pkg/
cli/pkg/cli/
cli/pkg/common/
cli/pkg/generated/
cli/pkg/hostbridge/
docs/
docs/docs.json
docs/hubspot.js
docs/package-lock.json
docs/package.json
docs/styles.css
docs/assets/
docs/assets/Cline_Logo-complete_black.png
docs/assets/Cline_Logo-complete_white.png
docs/assets/robot_panel_dark.png
docs/assets/robot_panel_light.png
docs/cline-cli/
docs/cline-cli/cli-reference.mdx
docs/cline-cli/installation.mdx
docs/cline-cli/overview.mdx
docs/cline-cli/three-core-flows.mdx
docs/cline-cli/samples/
docs/core-features/
docs/core-features/model-selection-guide.mdx
docs/enterprise-solutions/
docs/enterprise-solutions/overview.mdx
docs/enterprise-solutions/security-concerns.mdx
docs/exploring-clines-tools/
docs/exploring-clines-tools/cline-tools-guide.mdx
docs/features/
docs/getting-started/
docs/introduction/
docs/mcp/
docs/model-config/
docs/more-info/
docs/prompting/
docs/provider-config/
docs/running-models-locally/
docs/testing/
docs/troubleshooting/
evals/
evals/.gitignore
evals/package-lock.json
evals/package.json
evals/README.md
evals/tsconfig.json
evals/cli/
evals/diff-edits/
locales/
locales/ar-sa/
locales/de/
locales/es/
locales/ja/
locales/ko/
locales/pt-BR/
locales/zh-cn/
locales/zh-tw/
proto/
proto/descriptor_set.pb
proto/cline/
proto/host/
scripts/
scripts/api-secrets-parser.mjs
scripts/build-cli-all-platforms.sh
scripts/build-cli.sh
scripts/build-go-proto.mjs
scripts/build-npm-package.sh
scripts/build-proto.mjs
scripts/build-python-proto.mjs
scripts/build-tests.js
scripts/cli-providers.mjs
scripts/dev-cli-watch.mjs
scripts/download-ripgrep.mjs
scripts/file-utils.mjs
scripts/generate-host-bridge-client.mjs
scripts/generate-protobus-setup.mjs
scripts/generate-stubs.js
scripts/get-vscode-usages.sh
scripts/install-local.sh
scripts/install.sh
scripts/interactive-playwright.ts
scripts/package-standalone.mjs
scripts/proto-lint.sh
scripts/proto-shared-utils.mjs
scripts/proto-utils.mjs
scripts/publish-nightly.mjs
scripts/report-issue.js
scripts/runclinecore.sh
scripts/test-hostbridge-server.ts
scripts/test-install.sh
scripts/test-standalone-core-api-server.ts
scripts/testing-platform-orchestrator.ts
src/
src/common.ts
src/config.ts
src/extension.ts
src/registry.ts
src/api/
src/core/
src/dev/
src/exports/
src/generated/
src/hosts/
src/integrations/
src/packages/
src/samples/
src/services/
src/shared/
src/standalone/
src/test/
src/utils/
standalone/
standalone/runtime-files/
testing-platform/
testing-platform/index.ts
testing-platform/package-lock.json
testing-platform/package.json
testing-platform/README.md
testing-platform/tsconfig.json
testing-platform/adapters/
testing-platform/harness/
tests/
tests/specs/
walkthrough/
walkthrough/step1.md
walkthrough/step2.md
walkthrough/step3.md
walkthrough/step4.md
walkthrough/step5.md
webview-ui/
webview-ui/.gitignore
webview-ui/components.json
webview-ui/index.html
webview-ui/package-lock.json
webview-ui/package.json
webview-ui/tailwind.config.mjs
webview-ui/tsconfig.app.json
webview-ui/tsconfig.json
webview-ui/tsconfig.node.json
webview-ui/vite.config.ts
webview-ui/src/
worktrees/
worktrees/core-port-override/

(File list truncated. Use list_files on specific subdirectories if you need to explore further.)

# Workspace Configuration
{
  "workspaces": {
    "/Users/danielsteigman/code/@Cline/Cline-combo/cline": {
      "hint": "cline",
      "associatedRemoteUrls": [
        "aicccode: https://github.com/aicccode/cline.git",
        "fork_chezsmithy: https://github.com/chezsmithy/cline.git",
        "origin: git@github.com:cline/cline.git",
        "raceneader: https://github.com/raceneader/cline.git",
        "shouhanzen: https://github.com/shouhanzen/cline.git"
      ],
      "latestGitCommitHash": "1a07ca7906fbd023420b6e3ed89dcf93f0c8cad0"
    }
  }
}

# Detected CLI Tools
These are some of the tools on the user's machine, and may be useful if needed to accomplish the task: gh, git, docker, kubectl, aws, gcloud, npm, pnpm, pip, cargo, go, curl, jq, make, cmake, python, node, psql, mysql, sqlite3, code, grep, sed, awk, brew, bundle, wget. This list is not exhaustive, and other tools may be available.

# Current Mode
ACT MODE
</environment_details>

---

**Assistant:**



[Response interrupted by user]

---

**User:**

[TASK RESUMPTION] This task was interrupted just now. It may or may not be complete, so please reassess the task context. Be aware that the project state may have changed since then. The current working directory is now '/Users/danielsteigman/code/@Cline/Cline-combo/cline'. If the task has not been completed, retry the last step before interruption and proceed with completing the task.

Note: If you previously attempted a tool use that the user did not provide a result for, you should assume the tool use was not successful and assess whether you should retry. If the last tool was a browser_action, the browser has been closed and you must launch a new browser if needed.

IMPORTANT: If the last tool use was a replace_in_file or write_to_file that was interrupted, the file was reverted back to its original state before the interrupted edit, and you do NOT need to re-read the file as you already have its up-to-date contents.
New instructions for task continuation:
<user_message>
What worked well and what didnt? with this approch over just reading the files directly? It seems like we produced a huge amount of context with these tools what % of that was duplicated or helpful. Basically, we're working on this CLI for tree search-based retrieval of the codebase and understanding the codebase using tree search retrieval. It would be helpful to understand what parts of it were actually helpful and what parts weren't, so that we can improve the tool and come up with concrete, actionable insights as to how we can improve the tool versus just normal file reads. So we can improve the tools we used
</user_message>

# task_progress RECOMMENDED

When starting a new task, it is recommended to include a todo list using the task_progress parameter.


1. Include a todo list using the task_progress parameter in your next tool call
2. Create a comprehensive checklist of all steps needed
3. Use markdown format: - [ ] for incomplete, - [x] for complete

**Benefits of creating a todo/task_progress list now:**
	- Clear roadmap for implementation
	- Progress tracking throughout the task
	- Nothing gets forgotten or missed
	- Users can see, monitor, and edit the plan

**Example structure:**```
- [ ] Analyze requirements
- [ ] Set up necessary files
- [ ] Implement main functionality
- [ ] Handle edge cases
- [ ] Test the implementation
- [ ] Verify results```

Keeping the task_progress list updated helps track progress and ensures nothing is missed.

<environment_details>
# Visual Studio Code Visible Files
ATSE_EVALUATION_FRAMEWORK.md

# Visual Studio Code Open Tabs
ADDING_MODEL_PROVIDER.md
ATSE_EVALUATION_FRAMEWORK.md

# Current Time
11/15/2025, 2:10:07 AM (America/Los_Angeles, UTC-8:00)

# Current Working Directory (/Users/danielsteigman/code/@Cline/Cline-combo/cline) Files
.changie.yaml
.env.example
.gitattributes
.gitignore
.mocharc.json
.nvmrc
.nycrc.unit.json
.vscode-test.mjs
.vscodeignore
ADDING_MODEL_PROVIDER.md
ATSE_EVALUATION_FRAMEWORK.md
biome.jsonc
buf.yaml
CHANGELOG.md
CODE_OF_CONDUCT.md
CONTRIBUTING.md
esbuild.mjs
go.work
go.work.sum
knip.json
LICENSE
package-lock.json
package.json
playwright.config.ts
README.md
test-setup.js
tsconfig.json
tsconfig.test.json
tsconfig.unit-test.json
assets/
assets/docs/
assets/docs/demo.gif
assets/icons/
assets/icons/icon.png
assets/icons/icon.svg
assets/icons/robot_panel_dark.png
assets/icons/robot_panel_light.png
cli/
cli/.gitignore
cli/go.mod
cli/go.sum
cli/package.json
cli/README.md
cli/bin/
cli/cmd/
cli/cmd/cline/
cli/cmd/cline-host/
cli/e2e/
cli/e2e/default_update_test.go
cli/e2e/helpers_test.go
cli/e2e/main_test.go
cli/e2e/mixed_stress_test.go
cli/e2e/sqlite_helper.go
cli/e2e/start_list_test.go
cli/man/
cli/man/cline.1
cli/man/cline.1.md
cli/pkg/
cli/pkg/cli/
cli/pkg/common/
cli/pkg/generated/
cli/pkg/hostbridge/
docs/
docs/docs.json
docs/hubspot.js
docs/package-lock.json
docs/package.json
docs/styles.css
docs/assets/
docs/assets/Cline_Logo-complete_black.png
docs/assets/Cline_Logo-complete_white.png
docs/assets/robot_panel_dark.png
docs/assets/robot_panel_light.png
docs/cline-cli/
docs/cline-cli/cli-reference.mdx
docs/cline-cli/installation.mdx
docs/cline-cli/overview.mdx
docs/cline-cli/three-core-flows.mdx
docs/cline-cli/samples/
docs/core-features/
docs/core-features/model-selection-guide.mdx
docs/enterprise-solutions/
docs/enterprise-solutions/overview.mdx
docs/enterprise-solutions/security-concerns.mdx
docs/exploring-clines-tools/
docs/exploring-clines-tools/cline-tools-guide.mdx
docs/features/
docs/getting-started/
docs/introduction/
docs/mcp/
docs/model-config/
docs/more-info/
docs/prompting/
docs/provider-config/
docs/running-models-locally/
docs/testing/
docs/troubleshooting/
evals/
evals/.gitignore
evals/package-lock.json
evals/package.json
evals/README.md
evals/tsconfig.json
evals/cli/
evals/diff-edits/
locales/
locales/ar-sa/
locales/de/
locales/es/
locales/ja/
locales/ko/
locales/pt-BR/
locales/zh-cn/
locales/zh-tw/
proto/
proto/descriptor_set.pb
proto/cline/
proto/host/
scripts/
scripts/api-secrets-parser.mjs
scripts/build-cli-all-platforms.sh
scripts/build-cli.sh
scripts/build-go-proto.mjs
scripts/build-npm-package.sh
scripts/build-proto.mjs
scripts/build-python-proto.mjs
scripts/build-tests.js
scripts/cli-providers.mjs
scripts/dev-cli-watch.mjs
scripts/download-ripgrep.mjs
scripts/file-utils.mjs
scripts/generate-host-bridge-client.mjs
scripts/generate-protobus-setup.mjs
scripts/generate-stubs.js
scripts/get-vscode-usages.sh
scripts/install-local.sh
scripts/install.sh
scripts/interactive-playwright.ts
scripts/package-standalone.mjs
scripts/proto-lint.sh
scripts/proto-shared-utils.mjs
scripts/proto-utils.mjs
scripts/publish-nightly.mjs
scripts/report-issue.js
scripts/runclinecore.sh
scripts/test-hostbridge-server.ts
scripts/test-install.sh
scripts/test-standalone-core-api-server.ts
scripts/testing-platform-orchestrator.ts
src/
src/common.ts
src/config.ts
src/extension.ts
src/registry.ts
src/api/
src/core/
src/dev/
src/exports/
src/generated/
src/hosts/
src/integrations/
src/packages/
src/samples/
src/services/
src/shared/
src/standalone/
src/test/
src/utils/
standalone/
standalone/runtime-files/
testing-platform/
testing-platform/index.ts
testing-platform/package-lock.json
testing-platform/package.json
testing-platform/README.md
testing-platform/tsconfig.json
testing-platform/adapters/
testing-platform/harness/
tests/
tests/specs/
walkthrough/
walkthrough/step1.md
walkthrough/step2.md
walkthrough/step3.md
walkthrough/step4.md
walkthrough/step5.md
webview-ui/
webview-ui/.gitignore
webview-ui/components.json
webview-ui/index.html
webview-ui/package-lock.json
webview-ui/package.json
webview-ui/tailwind.config.mjs
webview-ui/tsconfig.app.json
webview-ui/tsconfig.json
webview-ui/tsconfig.node.json
webview-ui/vite.config.ts
webview-ui/src/
worktrees/
worktrees/core-port-override/

(File list truncated. Use list_files on specific subdirectories if you need to explore further.)

# Workspace Configuration
{
  "workspaces": {
    "/Users/danielsteigman/code/@Cline/Cline-combo/cline": {
      "hint": "cline",
      "associatedRemoteUrls": [
        "aicccode: https://github.com/aicccode/cline.git",
        "fork_chezsmithy: https://github.com/chezsmithy/cline.git",
        "origin: git@github.com:cline/cline.git",
        "raceneader: https://github.com/raceneader/cline.git",
        "shouhanzen: https://github.com/shouhanzen/cline.git"
      ],
      "latestGitCommitHash": "1a07ca7906fbd023420b6e3ed89dcf93f0c8cad0"
    }
  }
}

# Detected CLI Tools
These are some of the tools on the user's machine, and may be useful if needed to accomplish the task: gh, git, docker, kubectl, aws, gcloud, npm, pnpm, pip, cargo, go, curl, jq, make, cmake, python, node, psql, mysql, sqlite3, code, grep, sed, awk, brew, bundle, wget. This list is not exhaustive, and other tools may be available.

# Current Mode
ACT MODE
</environment_details>

---

**Assistant:**

[Unexpected content type]
[Tool Use: attempt_completion]
Result: ## Critical Analysis: ATSE Tool Effectiveness vs Traditional File Reading

Based on the actual benchmark data from our authentication analysis session, here's a detailed breakdown:

### Performance Metrics Summary

**Total Context Generated:** ~368K tokens across 8 commands
**Total Execution Time:** ~102 seconds
**Files Processed:** 821 files (TypeScript codebase)

### What Worked Well ✅

**1. `atse search` (5.3s, 252 tokens)**
- **Value:** High - Quickly found 15 relevant auth symbols across the codebase
- **Efficiency:** Excellent - Low token cost, fast execution
- **Advantage over file reading:** Would have taken 5-10 manual file reads to discover these entry points
- **Recommendation:** ⭐ **KEEP AS-IS** - This is the ideal starting point

**2. `atse deps` (0.003s, 576 tokens)**
- **Value:** Very High - Shows all 13 imports for a file instantly
- **Efficiency:** Exceptional - Near-instant, focused output
- **Advantage:** Clear dependency understanding without reading entire files
- **Recommendation:** ⭐ **KEEP AS-IS** - Perfect for understanding module relationships

**3. `atse context` (0.001s, 58 tokens)**
- **Value:** High - Pinpoint code context for specific line/column
- **Efficiency:** Exceptional - Ultra-fast, minimal tokens
- **Advantage:** Extracts just the relevant function without full file read
- **Recommendation:** ⭐ **KEEP AS-IS** - Ideal for targeted inspection

**4. `atse list-fns` (0.008s, 283 tokens)**
- **Value:** High - Shows all function declarations in a directory
- **Efficiency:** Excellent - Fast with minimal overhead
- **Advantage:** Quick API surface area understanding
- **Recommendation:** ⭐ **KEEP AS-IS** - Great for discovery

**5. `atse query` (0.023s, 210 tokens)**
- **Value:** Medium-High - Advanced pattern matching for specific constructs
- **Efficiency:** Excellent - Fast, targeted results
- **Advantage:** Can find patterns impossible with grep/manual reading
- **Recommendation:** ⭐ **KEEP AS-IS** - Powerful for advanced use cases

### What Didn't Work ❌

**6. `atse find-fn "getAuthToken"` (2.8s, 363,267 tokens = 98.7% of total!)**
- **Value:** NEGATIVE - Returned 10,450 matches including test files, mocks, generated proto code
- **Efficiency:** TERRIBLE - Consumed 99% of tokens with mostly noise
- **Problem:** No filtering for meaningful vs test/generated code
- **Token Breakdown:**
  - Actual auth logic: ~50 results (~1%)
  - Test files: ~10,000 results (~95%)
  - Generated proto files: ~400 results (~4%)
- **Recommendation:** 🚨 **CRITICAL FIX NEEDED**
  - Add `--exclude` patterns by default: `*.test.ts`, `*.test.js`, `__tests__`, `proto/`, `generated/`
  - Add `--include` pattern support to filter by directory
  - Default limit should be 20-50, not unlimited
  - Consider categorizing results (production vs test vs generated)

**7. `atse graph "AuthService"` (74.9s, 332 tokens)**
- **Value:** Medium - Showed dependency tree but took 75 seconds
- **Efficiency:** POOR - 74% of total execution time for modest output
- **Problem:** Processed all 821 files to build graph even with depth=2
- **Recommendation:** ⚠️ **NEEDS OPTIMIZATION**
  - Should leverage symbol index from initial scan
  - Cache results for repeated queries
  - Add progress indicators (which --verbose provides, but slow regardless)
  - Consider "lazy loading" - only traverse paths as needed

**8. `atse extract "AuthService"` (19.6s, 2,981 tokens)**
- **Value:** Medium-High - Got source code but slower than reading 4 files directly
- **Efficiency:** POOR - 19 seconds to extract 4 components
- **Problem:** Rescans entire codebase to extract known files
- **Comparison to direct file read:**
  - ATSE: 19.6s for 4 files
  - Direct read: ~0.1s for 4 files (manual targeting required)
- **Recommendation:** ⚠️ **NEEDS OPTIMIZATION**
  - Should cache symbol locations from search/graph
  - Extract should be near-instant if symbols already discovered
  - Consider --shallow flag to skip dependency traversal

### Duplication Analysis

**Estimated Duplication Rate: ~95-98%**

Breaking down the 368K tokens:
- **find-fn noise:** 363K tokens (~99% of total) - mostly duplicated test code
- **Useful unique context:** ~5K tokens (~1%)
  - search: 252 tokens ✓
  - deps: 576 tokens ✓
  - extract: 2,981 tokens (partially duplicative with manual reads)
  - graph: 332 tokens ✓
  - context: 58 tokens ✓
  - list-fns: 283 tokens ✓
  - query: 210 tokens ✓

### Concrete Recommendations for Tool Improvement

**HIGH PRIORITY:**

1. **Smart Filtering (Addresses 98% of waste)**
   ```
   --exclude-defaults (enabled by default):
     - **/*.test.ts
     - **/__tests__/**
     - **/generated/**
     - **/proto/**/*.ts (generated protobuf code)
     - **/*.spec.ts
   
   --include-prod-only flag to auto-exclude test/generated
   ```

2. **Result Categorization**
   ```
   Group results by type:
   [Production Code] (15 results)
   [Test Code] (10,000 results) - collapsed by default
   [Generated Code] (400 results) - collapsed by default
   ```

3. **Symbol Index Caching**
   - Build symbol index once per session
   - Reuse for graph/extract/find-fn commands
   - Would reduce graph from 75s → ~2s
   - Would reduce extract from 20s → ~0.1s

**MEDIUM PRIORITY:**

4. **Smart Limits**
   - find-fn: default limit 20 (not unlimited)
   - search: default limit 10-15 (current seems good)
   - Add --show-all flag to override

5. **Token Usage Warnings**
   ```
   Warning: This query returned 10,450 results (363K tokens).
   Consider using --exclude "*.test.*" to filter test files.
   Showing first 20 results (use --limit to adjust)
   ```

6. **Performance Modes**
   ```
   --fast: Skip tests/generated, use cached index
   --full: Include everything (current behavior)
   --production-only: Only src/prod code
   ```

### ATSE vs Direct File Reading Comparison

**When ATSE is BETTER:**
- ✅ **Discovery** (search, list-fns, query) - Finding unknown entry points
- ✅ **Dependencies** (deps) - Understanding imports without reading files
- ✅ **Pinpoint context** (context) - Getting specific line context
- ✅ **Pattern matching** (query) - Finding specific AST patterns

**When Direct File Reading is BETTER:**
- ✅ **Deep understanding** - Reading full implementation details
- ✅ **Speed** - 0.1s to read a file vs 20s to extract
- ✅ **Token efficiency** - 1-2K tokens per file vs potential 363K waste

**Optimal Hybrid Approach:**
1. `atse search` → Find entry points (252 tokens)
2. `atse deps` → Understand imports (576 tokens)  
3. **Direct file read** → Read 3-4 key files (~3K tokens)
4. `atse context` → Targeted snippets as needed (~58 tokens each)

**Total with hybrid:** ~4K tokens vs 368K with pure ATSE

### Bottom Line

**Token Efficiency Gain Potential: 99% reduction possible**
- Current approach: 368K tokens
- Optimized approach: 4K tokens (with filtering + caching)
- Most value came from 5K tokens; 363K was noise

**Fix Priority:**
1. 🔥 **CRITICAL:** Auto-exclude test files from find-fn (99% waste reduction)
2. 🔥 **CRITICAL:** Add symbol index caching (98% speed improvement)
3. ⚠️ **HIGH:** Default sensible limits on all commands
4. ✅ **MEDIUM:** Result categorization and warnings

The tools have excellent potential but need production vs test code intelligence to be truly effective for AI agents.
Task_progress: - [x] Review benchmark data from all 8 ATSE commands
- [x] Calculate token usage and duplication rates
- [x] Identify what worked well (search, deps, context, list-fns, query)
- [x] Identify critical failures (find-fn: 98.7% waste, graph: slow)
- [x] Compare ATSE vs direct file reading performance
- [x] Provide concrete recommendations with priority levels
- [x] Suggest optimal hybrid approach
- [x] Document actionable improvements for tool development

---

**User:**

[Tool]
[attempt_completion] Result:
The user has provided feedback on the results. Consider their input to continue the task, and then attempt completion again.
<feedback>
Based on what we learned here what further strategies does this reveal about coding agents and tree sitter based tooling that we might want to add. seems like we need to improve the tooling more. Do you have any suggested new tools or change how the current tools work. for example graph tool could also work completely differently or represent the data differently
</feedback>
<environment_details>
# Visual Studio Code Visible Files
ATSE_EVALUATION_FRAMEWORK.md

# Visual Studio Code Open Tabs
ADDING_MODEL_PROVIDER.md
ATSE_EVALUATION_FRAMEWORK.md

# Current Time
11/15/2025, 2:18:55 AM (America/Los_Angeles, UTC-8:00)

# Current Mode
ACT MODE
</environment_details>

