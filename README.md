# Emogo

Emogo creates a small Go web project with a server-rendered [Templ](https://templ.guide/) UI, [HTMX](https://htmx.org/), [Tailwind CSS](https://tailwindcss.com/), and a Bun-based frontend build. It also includes PostgreSQL connection and migration scaffolding, plus a `sqlc.yaml` configuration file.

## Prerequisites

Install the following before creating a project:

- [Go](https://go.dev/dl/)
- [Bun](https://bun.sh/docs/installation)
- [Task](https://taskfile.dev/installation/)
- [Air](https://github.com/air-verse/air): `go install github.com/air-verse/air@latest`
- [Templ](https://templ.guide/quick-start/installation): `go install github.com/a-h/templ/cmd/templ@latest`
- [Goose](https://github.com/pressly/goose): `go install github.com/pressly/goose/v3/cmd/goose@latest`

Ensure Go's binary directory is on your `PATH` so `air`, `templ`, and `goose` are available in your shell.

## Build and use the CLI

Build Emogo from this repository:

```sh
go build -o emogo .
```

Create a directory for the new application, enter it, and run the binary from there:

```sh
mkdir my-app
cd my-app
/path/to/emogo
```

Emogo prompts for a project name, initializes the Go module and Bun project, installs the Go and Bun dependencies, writes the starter application files, and starts the development task. Run it in the directory that should contain the new project; it writes files into the current directory.

When development starts, open [http://localhost:8090](http://localhost:8090). Air runs the Go server on port `8080` and provides the browser-facing proxy on port `8090`.

## Generated project

Each generated project contains its own `README.md` with setup guidance and a reference for all `Taskfile.yaml` commands. The initial application includes:

- Templ templates under `web/templates`
- Tailwind source and generated browser assets under `web/static`
- A self-hosted Geist Latin variable font, with weights from 100 through 900
- An HTMX starter route at `/htmx-test`
- PostgreSQL connection code in `internal/db`, migrations in `sql/schema`, and `sqlc.yaml`

Set `DATABASE_URL` in the generated `.env` before using the database migration tasks.
