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
	http.Handle("/assets/pics/", http.StripPrefix("/assets/pics/", http.FileServer(http.Dir("./pics"))))
	http.ListenAndServe(":8080", nil)
}

func index (w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func one (w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "one.gohtml", nil)
}

func onepic (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pics/1.PNG")
}

func two (w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "two.gohtml", nil)
}

func twopic (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pics/2.PNG")
}

func three (w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "three.gohtml", nil)
}

func threepic (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pics/3.PNG")
}

func four (w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "four.gohtml", nil)
}

func fourpic (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pics/4.PNG")
}

func five (w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "five.gohtml", nil)
}

func fivepic (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pics/5.PNG")
}