package main

import (
	"log"
	"net/http"
)

func home(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Opa gente boa!!!"))
}

func usuarios(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Carregar pagina de usuarios!!!"))
}

func main() {
	// URI - identificador do recurso
	// METODO - GET, POST, DELETE, UPDATE

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
