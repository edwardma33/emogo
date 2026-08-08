package boilerplate

var MainGoFileData = `
package main

import (
	"context"
	"log"
	"net/http"
	"%s/web/templates"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	mux := http.NewServeMux()

	s := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.Layout(templates.Home()).Render(context.Background(), w)
	})

	mux.HandleFunc("/htmx-test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("It's working!"))
	})

	log.Fatal(s.ListenAndServe())
}
`
