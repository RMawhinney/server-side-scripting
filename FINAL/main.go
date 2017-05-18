package main 

import (
	"net/http"
	"html/template"
)

var tpl *template.Template 

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)	
	http.HandleFunc("/main.css", css)
	http.HandleFunc("/favicon.ico", favicon)
	http.ListenAndServe(":8080", nil)
}

func index (w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func css (w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/main.css")
}

func favicon (w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/favicon.ico")
}