package main

import (
	"github.com/satori/go.uuid"
	"html/template"
	"net/http"
	"fmt"
)

type user struct {
	Email    string
	Password string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/login", login)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	tpl.ExecuteTemplate(w, "index.gohtml", u)

}

func login(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)

	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/bar", http.StatusSeeOther)
		fmt.Println(u.First, ` `, u.Last, ` logged in`)
		return
	}

	// if r.Method == http.MethodPost {
	// 			e := r.FormValue("email")
	// 			p := r.FormValue("password")

	//does username exist within dbUsers?

	//if username exists, does password?

	//yes

	// }
	fmt.Println(u.First, u.Last, `logged in`)
	tpl.ExecuteTemplate(w, "login.gohtml", u)
}

func signup(w http.ResponseWriter, r *http.Request) {
	// if !alreadyLoggedIn(r) {
	// 	http.Redirect(w, r, "/", http.StatusSeeOther)
	// 	return
	// }

	if r.Method == http.MethodPost {
		e := r.FormValue("email")
		p := r.FormValue("password")
		f := r.FormValue("first")
		l := r.FormValue("last")

		//is username taken, eg, user already signed up
		if _, ok := dbUsers[e]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		//create session
		id := uuid.NewV4()
		c := &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = e

		//Store user
		u := user{e, p, f, l}
		dbSessions[c.Value] = e
		dbUsers[e] = u

		//all done processing form submission
		fmt.Println(`New user signed up: {`, u.Email, `:`, u.Password, `;`, u.First, u.Last, `}`)
		fmt.Println(`New sessions started: {`, c.Value, `{`, u.Email,`}`)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func bar(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	u := getUser(w, r)
	fmt.Println(u.First, u.Last, `is at the bar`)
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

// func protected (f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
// 		return func(http.ResponseWriter, *http.Request) {}
// 		if !alreadyLoggedIn(r) {
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 		return
// 	}
// 	return f
// }
