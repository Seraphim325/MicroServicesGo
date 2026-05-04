package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	port := "80"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.gohtml")
	})

	log.Printf("Server listening on port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
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
