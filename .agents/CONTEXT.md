# Emogo codebase context

## Purpose

Emogo is a small interactive Go CLI that scaffolds a Go web project. It is intentionally direct and minimal: the generator runs shell commands, then writes starter files into the current working directory.

## Architecture

- `main.go` owns the generation flow: prompt for the project name and Git choice, initialize Go and Bun, create directories, write text/binary files, optionally run `git init`, then run `task dev`.
- `boilerplate/` contains Go string templates for every generated text file. Keep generated-file content in this package rather than adding a separate templating framework.
- `boilerplate/assets/` holds vendored binary/text assets. `boilerplate/font.go` embeds them into the CLI so they are written into each generated project without a runtime download.
- `utils/` contains the simple stdin prompt helper and shell-command runner used by `main.go`.

## Generated project contract

- Stack: Go standard-library HTTP server, Templ, HTMX, Tailwind CSS, Bun, pgx, godotenv, Goose migration layout, and sqlc configuration.
- Runtime: the server listens on port 8080; Air proxies browser traffic at `http://localhost:8090`.
- Static files are served from `web/static` at `/static/`.
- `Taskfile.yaml` owns the development workflow: Air, Templ generation, Tailwind compilation, Bun bundling, and Goose migrations.
- Geist is self-hosted. The embedded Google Fonts Latin variable WOFF2 is emitted to `web/static/files/geist-latin-wght-normal.woff2`; `input.css` declares `@font-face` with that absolute static URL and defines Tailwind’s `--font-sans`. It covers normal-style weights 100–900. `LICENSES/Geist-OFL.txt` is emitted with the project.

## Maintenance conventions

- Preserve the existing small, imperative style. Do not introduce abstractions, frameworks, or broad refactors solely for code quality.
- When adding a generated text file, add its template under `boilerplate/`, create any parent directory in `dirs`, and register it in the `files` map in `main.go`.
- When adding a generated binary asset, vendor it under `boilerplate/assets/`, expose it with `go:embed`, add its target parent directory to `dirs`, and write it through `CreateBinaryFiles`.
- Keep the root README and generated-project README template synchronized whenever the generated workflow, prerequisites, stack, or Taskfile changes.
- Validate changes with `gofmt`, `go test ./...`, `git diff --check`, and, when CSS/assets change, a temporary Tailwind build that verifies generated output paths.
