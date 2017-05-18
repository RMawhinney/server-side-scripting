package main

import ("html/template"
	"net/http"
)

var tpl *template.Template

func init () {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main () {
	files := http.FileServer(http.Dir("public"))
	http.Handle("/pics/", files)
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo (w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate (w, "index.gohtml", nil)
}