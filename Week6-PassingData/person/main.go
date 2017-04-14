package main 

import ("net/http"
		"html/template")

var tpl *template.Template 

type person struct {
	First string
	Age int 
	Drunk bool
}

func init(){
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}


func index (w http.ResponseWriter, r *http.Request) {
	j := person {"Jason", 24, true}
	s := person {"Sid", 23, true}
	xp := []person {}
	xp = append(xp, j)
	xp = append(xp, s)
	tpl.ExecuteTemplate(w, "index.gohtml", xp)	
}