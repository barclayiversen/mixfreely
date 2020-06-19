package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.Handle("/static", http.StripPrefix("/static/", http.FileServer(http.Dir("."))))
	http.HandleFunc("/", dog)
	http.HandleFunc("/app/test", test)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func dog(w http.ResponseWriter, req *http.Request) {
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tpl.ExecuteTemplate(w, "index.html", nil)
}

func test(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "test.html", nil)
}
