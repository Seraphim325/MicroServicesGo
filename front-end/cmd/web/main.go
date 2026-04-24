package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.gohtml")
	})

	log.Println("Server listening on port 80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Panic(err)
	}

}

func render(w http.ResponseWriter, t string) {

	templateSlices := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.gohtml",
		"./templates/footer.partial.gohtml",
		"./templates/header.partial.gohtml",
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
