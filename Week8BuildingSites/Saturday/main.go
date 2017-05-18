package main 

import ("net/http"
		"html/template")

var tpl *template.Template 

func init(){
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main () {
	http.HandleFunc("/", saturday)
	http.HandleFunc("/saturday.PNG", pic)
	http.HandleFunc("/main.css", css)
	http.ListenAndServe(":8080", nil)
}

func saturday (w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func pic (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "saturday.PNG")
}

func css (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "main.css")
}