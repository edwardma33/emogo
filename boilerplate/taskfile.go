package boilerplate

var TaskfileYAMLFileData string = `
version: '3'

dotenv: ['.env']

tasks:
  goose:up:
    cmds:
      - goose -dir sql/schema postgres "$DATABASE_URL" up

  goose:down:
    cmds:
      - goose -dir sql/schema postgres "$DATABASE_URL" down

  templ:generate:
    cmds:
      - templ generate

  air:
    cmds:
      - air

  tailwind:
    cmds:
      - tailwindcss -i ./web/static/input.css -o ./web/static/output.css

  templ:watch:
    cmds:
      - templ generate --watch

  tailwind:watch:
    cmds:
      - tailwindcss -i ./web/static/input.css -o ./web/static/output.css --watch

  bun:build:
    cmds:
      - bun build ./web/static/app.ts --outfile ./web/static/app.js --target browser --minify

  bun:watch:
    cmds:
      - bun build ./web/static/app.ts --outfile ./web/static/app.js --target browser --minify --watch

  dev:
    deps:
      - air
      - templ:watch
      - tailwind:watch
      - bun:watch
`
