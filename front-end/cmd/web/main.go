package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	loadConfig()
	port := os.Getenv("FRONTEND_PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.gohtml")
	})

	log.Printf("Server listening on port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Panic(err)
	}

}

func loadConfig() {
	if err := godotenv.Load(".env"); err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, t string) {
	execPath, err := os.Executable()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	baseDir := filepath.Dir(execPath)
	templateDir := filepath.Join(baseDir, "templates")

	templateSlices := []string{
		filepath.Join(templateDir, t),
		filepath.Join(templateDir, "base.layout.gohtml"),
		filepath.Join(templateDir, "footer.partial.gohtml"),
		filepath.Join(templateDir, "header.partial.gohtml"),
	}

	tmpl, err := template.ParseFiles(templateSlices...)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
