package main

import (
	"log"
	"net/http"
	"text/template"
	"time"
)

func main() {
	http.HandleFunc("/", serveFunc)
	http.ListenAndServe(":8090", nil)
}

func serveFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var fm = template.FuncMap{
		"fdateMDY": monthDayYear,
	}
	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("index_func.gohtml"))
	if err := tpl.ExecuteTemplate(w, "index_func.gohtml", time.Now()); err != nil {
		log.Fatal(err)
	}
}

func monthDayYear(t time.Time) string {
	return t.Format(time.UnixDate)
}
