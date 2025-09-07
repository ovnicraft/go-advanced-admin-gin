# Repository Guidelines

## Project Structure & Module Organization
- This is a small Go library integrating the Go Advanced Admin Panel with Gin.
- Key files:
  - `integrator.go` (package `admingin`): exposes `Integrator` and HTTP helpers.
  - `go.mod`, `go.sum`: module and dependencies.
  - `README.md`, `LICENSE`.
- Tests are not yet present; add `*_test.go` files alongside sources.

## Build, Test, and Development Commands
- Build: `go build ./...` — verifies the module compiles.
- Test: `go test ./...` — runs unit tests; use `-v` for verbose and `-cover` for coverage.
- Vet: `go vet ./...` — catches common issues; fix warnings before PRs.
- Format: `gofmt -s -w .` and (optional) `goimports -w .`.
- Local replace (when developing against local admin repo): add in `go.mod`
  - `replace github.com/go-advanced-admin/admin => ../admin`

## Coding Style & Naming Conventions
- Follow idiomatic Go (Go 1.24+ as in `go.mod`).
- Use `gofmt` formatting; no custom style. Keep imports grouped/ordered.
- Naming:
  - Exported API uses PascalCase (e.g., `Integrator`, `NewIntegrator`).
  - Unexported helpers use lowerCamelCase. Package name remains `admingin`.
- Function signatures should accept/return Gin and admin types explicitly; avoid interface{} unless necessary.

## Testing Guidelines
- Use the standard `testing` package; table‑driven tests preferred.
- Place tests in `*_test.go` within the same package (`package admingin`).
- Example: `func TestHandleRoute(t *testing.T) { ... }`
- Run locally with `go test ./... -cover`; target ≥80% for new code.

## Commit & Pull Request Guidelines
- Use Conventional Commits: `feat:`, `fix:`, `docs:`, `chore:`, `refactor:`, `test:` (matches Git history).
- PRs must include:
  - Clear description, motivation, and scope; link related issues.
  - Tests for new behavior and updates to `README.md` when public API changes.
  - Passing build, `go vet`, and formatting checks.
  - Small, focused diffs; avoid unrelated refactors.

## Security & Maintenance Tips
- Do not commit secrets. Keep dependencies current: `go get -u ./...` and review changes.
- (Optional) Run `govulncheck ./...` if available.
