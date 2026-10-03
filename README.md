# Emogo

Emogo creates a small Go web project with a server-rendered [Templ](https://templ.guide/) UI, [HTMX](https://htmx.org/), [Alpine.js](https://alpinejs.dev/), [Tailwind CSS](https://tailwindcss.com/), and [shadcn-templ](https://shadcn-templ.com/) components. It also includes PostgreSQL connection and migration scaffolding, plus a `sqlc.yaml` configuration file.

## Prerequisites

Install the following before creating a project:

- [Go](https://go.dev/dl/)
- [Bun](https://bun.sh/docs/installation)
- [Task](https://taskfile.dev/installation/)
- [Air](https://github.com/air-verse/air): `go install github.com/air-verse/air@latest`
- [Templ](https://templ.guide/quick-start/installation): `go install github.com/a-h/templ/cmd/templ@latest`
- [Goose](https://github.com/pressly/goose): `go install github.com/pressly/goose/v3/cmd/goose@latest`
- [shadcn-templ](https://shadcn-templ.com/docs/cli): `go install github.com/axadrn/shadcn-templ/v2/cmd/shadcn-templ@latest`

Ensure Go's binary directory is on your `PATH` so `air`, `templ`, `goose`, and `shadcn-templ` are available in your shell.

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

Emogo prompts for a project name, whether to install all shadcn-templ components, and whether to initialize Git. If you decline all components, it installs the button and icon components used by the starter page. Run it in the directory that should contain the new project; it writes files into the current directory.

## What Emogo creates

For the supplied project name, Emogo:

1. Initializes the Go module and installs `pgx`, `godotenv`, Templ, and the shadcn-templ Tailwind merge dependency.
2. Initializes Bun and installs Tailwind CSS, HTMX, Alpine.js, and Alpine.js types.
3. Creates the Go server, Templ templates, TypeScript entry point, Tailwind source CSS, database connection helper, sqlc configuration, initial Goose migration, Air configuration, Taskfile, and project README.
4. Writes a self-hosted Geist Latin variable font and its OFL license. The font covers weights 100–900 and is already connected to Tailwind’s `font-sans` utility.
5. Initializes shadcn-templ and installs the selected components. The starter page includes a theme switcher, an HTMX button, and an Alpine.js counter.
6. Initializes a Git repository only when you answer `y` to the Git prompt.
7. Starts `task dev`.

The generated server serves `/static/` from `web/static`, renders the home page at `/`, and includes an HTMX example endpoint at `/htmx-test`.

When development starts, open [http://localhost:8090](http://localhost:8090). Air runs the Go server on port `8080` and provides the browser-facing proxy on port `8090`.

## Generated project

Each generated project contains its own `README.md` with setup guidance and a reference for all `Taskfile.yaml` commands. The initial application includes:

- Templ templates under `web/templates`
- Tailwind source and generated browser assets under `web/static`
- A self-hosted Geist Latin variable font, with weights from 100 through 900
- An HTMX starter route at `/htmx-test`
- A theme switcher and Alpine.js counter on the starter page
- shadcn-templ components under `components`
- PostgreSQL connection code in `internal/db`, migrations in `sql/schema`, and `sqlc.yaml`

The theme follows the system preference until the switcher is clicked. The selected light or dark theme is saved in browser local storage for future visits.

Set `DATABASE_URL` in the generated `.env` before using the database migration tasks.
