package main 

import ("html/template"
		"net/http")

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("assets/*.gohtml"))
}

func main(){
	http.HandleFunc("/", index)
	http.HandleFunc("/one", one)
	http.HandleFunc("/two", two)
	http.HandleFunc("/three", three)
	http.HandleFunc("/four", four)
	http.HandleFunc("/five", five)
	http.HandleFunc("/1.PNG", onepic)
	http.HandleFunc("/2.PNG", twopic)
	http.HandleFunc("/3.PNG", threepic)
	http.HandleFunc("/4.PNG", fourpic)
	http.HandleFunc("/5.PNG", fivepic)
	http.ListenAndServe(":8080", nil)
}

func index (w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func one (w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "one.gohtml", nil)
}

func onepic (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/pics/1.PNG")
}

func two (w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "two.gohtml", nil)
}

func twopic (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/pics/2.PNG")
}

func three (w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "three.gohtml", nil)
}

func threepic (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/pics/3.PNG")
}

func four (w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "four.gohtml", nil)
}

func fourpic (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/pics/4.PNG")
}

func five (w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "five.gohtml", nil)
}

func fivepic (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/pics/5.PNG")
}