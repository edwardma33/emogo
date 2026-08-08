package main

import (
	"emogo/boilerplate"
	"emogo/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

var dirs = []string{
	"web/templates", "web/static", "internal/db", "sql/schema", "sql/queries",
}

func main() {
	projectName := GetProjectName()

	var files = map[string]string{
		"main.go":                    fmt.Sprintf(boilerplate.MainGoFileData, projectName),
		"README.md":                  fmt.Sprintf(boilerplate.READMEFileData, projectName),
		"sqlc.yaml":                  boilerplate.SQLCYAMLFileData,
		".air.toml":                  boilerplate.AirTOMLFileData,
		".env":                       "DATABASE_URL=\"\"",
		"Taskfile.yaml":              boilerplate.TaskfileYAMLFileData,
		"internal/db/conn.go":        boilerplate.DbConnGoFileData,
		"sql/schema/001_init.sql":    "-- +goose Up\n\n-- +goose Down",
		"web/templates/layout.templ": boilerplate.LayoutTemplFileData,
		"web/templates/index.templ":  boilerplate.IndexTemplFileData,
		"web/static/input.css":       "@import \"tailwindcss\";",
		"web/static/app.ts":          boilerplate.AppTsFileData,
	}

	// go mod init and install deps
	utils.RunCommand(fmt.Sprintf("go mod init %s", projectName))
	var deps = []string{
		"github.com/jackc/pgx/v5",
		"github.com/joho/godotenv",
		"github.com/a-h/templ",
	}

	for _, dep := range deps {
		utils.RunCommand(fmt.Sprintf("go get %s", dep))
	}

	// init bun and install tw
	utils.RunCommand("bun init -y --minimal")
	utils.RunCommand("bun add tailwindcss htmx.org")

	err := CreateDirAll(dirs)
	if err != nil {
		log.Fatal(err)
	}

	err = CreateFiles(files)
	if err != nil {
		log.Fatal(err)
	}

	utils.RunCommand("task dev")
}

func GetProjectName() string {
	pn := utils.GetUserInput("Project Name: ")
	if pn == "" {
		return GetProjectName()
	}

	return strings.ReplaceAll(pn, " ", "-")
}

func CreateDirAll(paths []string) error {
	var err error
	for _, p := range paths {
		err = os.MkdirAll(p, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateFiles(files map[string]string) error {
	var err error
	for p, data := range files {
		err = os.WriteFile(p, []byte(data), 0666)
		if err != nil {
			return err
		}
	}
	return nil
}
