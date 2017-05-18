package	main

import ("net/http"
		"html/template"
		"fmt")

var tpl *template.Template

func init () {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/bowzer", bowzer)
	http.HandleFunc("/bowzer/pictures", bowzerpics)
	http.ListenAndServe(":8080", nil)

}

func index (w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
	}

func bowzer(w http.ResponseWriter, req *http.Request) {

	c := &http.Cookie{
		Name: "user-cookie",
		Value: "some value",
		Path: "/",}
		http.SetCookie(w, c)
		tpl.ExecuteTemplate(w, "bowzer.gohmtl", c)
}

func bowzerpics (w http.ResponseWriter, req *http.Request) {
	var c *http.Cookie
	c, err := req.Cookie("user-cookie")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%T/n", c)
		}
	tpl.ExecuteTemplate(w, "bowzerpics.gohtml", c)
	}
