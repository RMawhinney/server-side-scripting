package main

import ("net/http"
		"io"
		"html/template"
)

var tpl *template.Template

func init () {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/dog.jpg", dogpic)
	http.ListenAndServe(":8080", nil)
}

func foo (w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog (w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func dogpic (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dog.jpg")
}