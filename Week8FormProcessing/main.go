package main

import ("net/http"
	"html/template"
)

var tpl *template.Template

type user struct {
	Name string
	Guild string
	Rank string
}


func init () {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index (w http.ResponseWriter, req *http.Request) {

	// if req.Method == http.MethodPost {
	// 	n := req.FormValue("name")
	// 	g := req.FormValue("guild")
	// 	r := req.FormValue("rank")

	// // 	// uz := user{n, g, r}
	// 	return
	// 	}

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
