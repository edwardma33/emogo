# Emogo

Emogo creates a small Go web project with a server-rendered [Templ](https://templ.guide/) UI, [HTMX](https://htmx.org/), [Alpine.js](https://alpinejs.dev/), [Tailwind CSS](https://tailwindcss.com/), and [shadcn-templ](https://shadcn-templ.com/) components. Choose PostgreSQL or SQLite for the generated database connection, migration tasks, and `sqlc.yaml` configuration.

## Technology docs

| Technology | Used for | Documentation |
| --- | --- | --- |
| Go | HTTP server and application code | [Go docs](https://go.dev/doc/) |
| Templ | Server-rendered HTML components | [Templ docs](https://templ.guide/) |
| shadcn-templ | Starter UI components and theme styles | [shadcn-templ docs](https://shadcn-templ.com/docs/introduction) |
| HTMX | HTML-driven server requests | [HTMX docs](https://htmx.org/docs/) |
| Alpine.js | Client-side interactions | [Alpine.js docs](https://alpinejs.dev/start-here) |
| Tailwind CSS | Utility-based styling | [Tailwind CSS docs](https://tailwindcss.com/docs/) |
| Bun | Frontend dependencies and TypeScript bundling | [Bun docs](https://bun.com/docs) |
| Task | Development and migration commands | [Task docs](https://taskfile.dev/docs/) |
| Air | Go server reloads during development | [Air docs](https://github.com/air-verse/air#readme) |
| PostgreSQL and pgx | PostgreSQL option and Go driver | [PostgreSQL docs](https://www.postgresql.org/docs/), [pgx docs](https://pkg.go.dev/github.com/jackc/pgx/v5) |
| SQLite and go-sqlite3 | SQLite option and Go driver | [SQLite docs](https://www.sqlite.org/docs.html), [go-sqlite3 docs](https://github.com/mattn/go-sqlite3) |
| Goose | Database migrations | [Goose docs](https://pressly.github.io/goose/) |
| sqlc | SQL query code generation configuration | [sqlc docs](https://docs.sqlc.dev/en/latest/) |

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

SQLite projects use `go-sqlite3`, which requires CGo and a C compiler when building the generated app.

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

Emogo prompts for a project name, a database choice, whether to install all shadcn-templ components, and whether to initialize Git. If you decline all components, it installs the button and icon components used by the starter page. Run it in the directory that should contain the new project; it writes files into the current directory.

## What Emogo creates

For the supplied project name, Emogo:

1. Initializes the Go module and installs the selected database driver (`pgx` or `go-sqlite3`), `godotenv`, Templ, and the shadcn-templ Tailwind merge dependency.
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
- Connection code for the selected database in `internal/db`, migrations in `sql/schema`, and `sqlc.yaml`

The theme follows the system preference until the switcher is clicked. The selected light or dark theme is saved in browser local storage for future visits.

The generated `.env` contains `DATABASE_URL`. Replace the PostgreSQL placeholder with your connection string, or use the SQLite default `data.db` file in the project directory. The database migration tasks use this value.
