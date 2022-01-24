package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(rw http.ResponseWriter, r *http.Request) {
		u := usuario{"Dollynho", "dollynho@gmail.com"}
		templates.ExecuteTemplate(rw, "home.html", u)
	})

	fmt.Println("Servidor escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
