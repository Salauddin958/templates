package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", serve)
	http.ListenAndServe(":8080", nil)
}

func serve(w http.ResponseWriter, r *http.Request) {
	sages := []string{"Tom Cruise", "Jack Nicholson", "Demi Moore", "Kevin Bacon", "Wolfgang Bodison"}
	tpl, err := template.ParseFiles("index_array.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.Execute(w, sages)
}
