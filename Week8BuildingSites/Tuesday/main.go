package main 

import ("net/http"
		"html/template")

var tpl *template.Template 

func init(){
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main () {
	http.HandleFunc("/", thursday)
	http.HandleFunc("/tuesday.jpg", pic)
	http.HandleFunc("/main.css", css)
	http.ListenAndServe(":8080", nil)
}

func thursday (w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func pic (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "tuesday.jpg")
}

func css (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "main.css")
}