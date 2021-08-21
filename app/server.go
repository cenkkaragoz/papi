package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Hello papi!..")
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("template/public"))))
	files := []string{
		"template/one.html",
	}

	t := template.Must(template.ParseFiles(files...))

	r.HandleFunc("/receipt", func(w http.ResponseWriter, r *http.Request) {
		err := t.ExecuteTemplate(w, "one.html", "")
		if err != nil {
			log.Println(err)
		}
	})

	r.HandleFunc("/receiptPdf", func(w http.ResponseWriter, r *http.Request) {

	})

	http.ListenAndServe(":8080", r)
}
