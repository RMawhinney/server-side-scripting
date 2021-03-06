package main

import("html/template" 
		"net/http")

var tpl *template.Template

func init() {
	tpl = template.Must (template.ParseGlob ("templates/*.gohtml"))
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe (":8080", nil)
}