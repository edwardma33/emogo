package boilerplate

var READMEFileData string = `# %s

This project was created with Emogo. It uses Go, Templ, HTMX, Alpine.js, Tailwind CSS, shadcn-templ components, Bun, and scaffolding for the database selected during project creation.

Emogo installed the Go and Bun dependencies, generated the starter files, and optionally initialized Git during project creation.

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

The generated Taskfile expects these commands to be installed and available on your PATH:

- [Go](https://go.dev/dl/)
- [Bun](https://bun.sh/docs/installation)
- [Task](https://taskfile.dev/installation/)
- [Air](https://github.com/air-verse/air): <code>go install github.com/air-verse/air@latest</code>
- [Templ](https://templ.guide/quick-start/installation): <code>go install github.com/a-h/templ/cmd/templ@latest</code>
- [Goose](https://github.com/pressly/goose): <code>go install github.com/pressly/goose/v3/cmd/goose@latest</code>

If you selected SQLite, building the Go app with <code>go-sqlite3</code> requires CGo and a C compiler.

## Get started

The generated <code>.env</code> sets <code>DATABASE_URL</code> for the selected database. If you selected PostgreSQL, replace the placeholder connection string with your own. If you selected SQLite, the default <code>data.db</code> is a database file in the project directory; change the path if needed. The connection helper and migration tasks use this value.

Start the full development workflow:

~~~sh
task dev
~~~

Open [http://localhost:8090](http://localhost:8090). Air runs the Go application on port 8080 and proxies browser requests through port 8090.

The server renders the home page at <code>/</code>, serves assets from <code>/static/</code>, and includes a starter HTMX endpoint at <code>/htmx-test</code>.

The theme follows the system preference until the switcher is clicked. The selected light or dark theme is saved in browser local storage for future visits.

## Task commands

| Command | What it does |
| --- | --- |
| <code>task dev</code> | Starts the Go server, Templ watch, Tailwind watch, and Bun watch together. |
| <code>task air</code> | Runs Air, which rebuilds and reloads the Go server as Go files change. |
| <code>task templ:generate</code> | Generates Go files from <code>.templ</code> files once. Run it after changing a template when not using the watcher. |
| <code>task templ:watch</code> | Watches <code>.templ</code> files and regenerates their Go files when they change. |
| <code>task tailwind</code> | Builds <code>web/static/input.css</code> into <code>web/static/output.css</code> once. |
| <code>task tailwind:watch</code> | Watches Tailwind source files and continually rebuilds <code>web/static/output.css</code>. |
| <code>task bun:build</code> | Bundles and minifies <code>web/static/app.ts</code> into browser-ready <code>web/static/app.js</code> once. |
| <code>task bun:watch</code> | Watches <code>web/static/app.ts</code> and continually rebuilds <code>web/static/app.js</code>. |
| <code>task goose:up</code> | Applies pending migrations from <code>sql/schema</code> to the selected database using <code>DATABASE_URL</code>. |
| <code>task goose:down</code> | Rolls back the most recently applied migration using <code>DATABASE_URL</code>. |

## Project layout

- <code>web/templates</code> contains the starter page with an HTMX button that calls <code>/htmx-test</code> and an Alpine.js counter.
- <code>components</code> contains shadcn-templ components and the theme switcher.
- <code>web/static</code> contains the Tailwind input/output CSS, TypeScript/JavaScript browser bundle, and a self-hosted Geist Latin variable font with weights from 100 through 900. Its OFL license is in <code>LICENSES/Geist-OFL.txt</code>.
- <code>internal/db</code> contains connection code for the selected database. <code>sql/schema</code> contains Goose migrations, and <code>sqlc.yaml</code> configures sqlc for future query generation.
`
