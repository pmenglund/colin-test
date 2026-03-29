# Repository Guidelines

## Project Structure & Module Organization

This repository is currently a minimal Go module. The entrypoint is [`main.go`](/Users/pme/src/pmenglund/colin-test/main.go), module metadata lives in [`go.mod`](/Users/pme/src/pmenglund/colin-test/go.mod), and the short project description is in [`README.md`](/Users/pme/src/pmenglund/colin-test/README.md).

Keep small applications simple at the root. If the codebase grows, move reusable packages under `internal/` and keep `main.go` focused on wiring and startup.

## Build, Test, and Development Commands

- `go run .` runs the application locally.
- `go build ./...` compiles all packages in the module and should stay clean before merging.
- `go test ./...` runs the full test suite. It currently reports `no test files`, which is expected.
- `gofmt -w .` formats all Go files in place.

Run commands from the repository root: `/Users/pme/src/pmenglund/colin-test`.

## Coding Style & Naming Conventions

Use standard Go formatting and let `gofmt` decide spacing and indentation. Prefer short, descriptive package and file names, all lowercase unless Go export rules require `CamelCase`.

Follow normal Go naming:

- exported identifiers: `CamelCase`
- unexported identifiers: `camelCase`
- packages: short lowercase names without underscores

Keep `main.go` thin. Business logic should move into functions or packages once it exceeds trivial setup.

## Testing Guidelines

Use Go's built-in `testing` package. Place tests in `*_test.go` files beside the code they cover, and name test functions `TestXxx`.

Prefer table-driven tests for logic with multiple cases. Before opening a PR, run `go test ./...`. If coverage becomes meaningful, use `go test -cover ./...` to check it.

## Commit & Pull Request Guidelines

The current history uses a short imperative subject (`Initial commit`). Continue with concise imperative commit messages such as `Add health check handler` or `Extract config loader`.

PRs should include:

- a brief summary of the change
- test evidence, for example `go test ./...`
- linked issues if applicable

## Agent-Specific Notes

This repository may change between turns. Before editing, check the working tree with `git status --short` so you do not overwrite user changes.
