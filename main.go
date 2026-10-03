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
	"web/templates", "web/static/files", "internal/db", "sql/schema", "sql/queries", "LICENSES", "components/themeswitcher",
}

func main() {
	projectName := GetProjectName()

	var files = map[string]string{
		"main.go":                    fmt.Sprintf(boilerplate.MainGoFileData, projectName),
		"README.md":                  fmt.Sprintf(boilerplate.READMEFileData, projectName),
		"sqlc.yaml":                  boilerplate.SQLCYAMLFileData,
		".air.toml":                  boilerplate.AirTOMLFileData,
		".env":                       "DATABASE_URL=\"postgresql://[user[:password]@][host][:port][/dbname]?sslmode=disable\"",
		"Taskfile.yaml":              boilerplate.TaskfileYAMLFileData,
		"internal/db/conn.go":        boilerplate.DbConnGoFileData,
		"sql/schema/001_init.sql":    "-- +goose Up\n\n-- +goose Down",
		"web/templates/layout.templ": boilerplate.LayoutTemplFileData,
		"web/templates/index.templ":  fmt.Sprintf(boilerplate.IndexTemplFileData, projectName, projectName),
		"web/static/input.css":       boilerplate.InputCSSFileData,
		"web/static/app.ts":          boilerplate.AppTsFileData,
		"components/themeswitcher/themeswitcher.templ": fmt.Sprintf(boilerplate.ThemeswitcherData, projectName, projectName),
	}

	// go mod init and install deps
	utils.RunCommand(fmt.Sprintf("go mod init %s", projectName))
	var deps = []string{
		"github.com/jackc/pgx/v5",
		"github.com/joho/godotenv",
		"github.com/a-h/templ",
		"github.com/Oudwins/tailwind-merge-go",
	}

	for _, dep := range deps {
		utils.RunCommand(fmt.Sprintf("go get %s", dep))
	}

	// init bun and install tw
	utils.RunCommand("bun init -y --minimal")
	utils.RunCommand("bun add tailwindcss htmx.org alpinejs")
	utils.RunCommand("bun add -d @types/alpinejs")

	err := CreateDirAll(dirs)
	if err != nil {
		log.Fatal(err)
	}

	err = CreateFiles(files)
	if err != nil {
		log.Fatal(err)
	}

	err = CreateBinaryFiles(map[string][]byte{
		"web/static/files/geist-latin-wght-normal.woff2": boilerplate.GeistLatinVariableFont,
		"LICENSES/Geist-OFL.txt":                         boilerplate.GeistOFL,
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := utils.RunCommand("shadcn-templ init"); err != nil {
		log.Fatal(err)
	}
	if strings.ToLower(utils.GetUserInput("Add all components? [Y, n]")) != "n" {
		if err := utils.RunCommand("shadcn-templ add --all"); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := utils.RunCommand("shadcn-templ add button"); err != nil {
			log.Fatal(err)
		}
		if err := utils.RunCommand("shadcn-templ add icon"); err != nil {
			log.Fatal(err)
		}
	}
	if err := fixGeneratedDarkThemeCSS("web/static/input.css"); err != nil {
		log.Fatal(err)
	}

	if IsGitEnabled() {
		utils.RunCommand("git init .")
	}

	utils.RunCommand("task dev")
}

func fixGeneratedDarkThemeCSS(path string) error {
	css, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	const darkSelector = "\n.dark {"
	if strings.Count(string(css), darkSelector) != 1 {
		return fmt.Errorf("expected exactly one dark theme block in %s", path)
	}

	updated := strings.Replace(string(css), darkSelector, "\n:root.dark {", 1)
	return os.WriteFile(path, []byte(updated), 0666)
}

func GetProjectName() string {
	pn := utils.GetUserInput("Project Name: ")
	if pn == "" {
		return GetProjectName()
	}

	return strings.ReplaceAll(pn, " ", "-")
}

func IsGitEnabled() bool {
	v := utils.GetUserInput("Enable git? (y/N): ")
	return strings.ToLower(v) == "y"
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

func CreateBinaryFiles(files map[string][]byte) error {
	var err error
	for p, data := range files {
		err = os.WriteFile(p, data, 0666)
		if err != nil {
			return err
		}
	}
	return nil
}
