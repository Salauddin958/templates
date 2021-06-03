package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveStatic)
	http.ListenAndServe(":8000", nil)
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	type User struct {
		Name   string
		Coupon string
		Amount int
	}
	user := User{
		Name:   "Rick",
		Coupon: "IAMAWESOMEGOPHER",
		Amount: 5000,
	}
	err = tpl.Execute(w, user)
	if err != nil {
		fmt.Println(err)
	}
}
