package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

type person struct {
	Name string
	Rank string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/createuser", createuser)
	http.HandleFunc("/lookie", lookie)
	http.HandleFunc("/overhere", overhere)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("at Index")
	// tt := getGuy(w, r)
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func createuser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("at Create User")
	p, err := getPerson(w, r)
	if err != nil {
		fmt.Println("err getPerson", err)
	}

	if r.Method == http.MethodPost {
		nm := r.FormValue("name")
		rk := r.FormValue("rank")

		p = person{
			Name: nm,
			Rank: rk,
		}

		//encode into json
		pj, erc := json.Marshal(p)
		if erc != nil {
			fmt.Println("json encode error")
			return
		}
		// fmt.Println(`encoding into json`, p)

		//encode into b64
		pj64 := base64.StdEncoding.EncodeToString([]byte(pj))

		// fmt.Println(`encoding into b64`, userjs)

		//create Cookie with second encode as value
		c := &http.Cookie{
			Name:     "Cookie",
			Value:    pj64,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, c)

		fmt.Println(`Cookie set, person:`, p)
		fmt.Println(pj64)

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "createuser.html", p)
}

func getPerson(w http.ResponseWriter, r *http.Request) (person, error) {
	p := person{
		Name: "Human",
		Rank: "Man",
	}

	c, err := r.Cookie("Cookie")
	if err != nil {
		// fmt.Println(`No cookie; baking now`, p)

		pj, err := json.Marshal(p)
		if err != nil {
			// fmt.Println("json encode error")
			return p, err
		}
		// fmt.Println(`encoding into json`, p)

		pj64 := base64.StdEncoding.EncodeToString([]byte(pj))

		// fmt.Println(`encoding into b64`, pj)

		c = &http.Cookie{
			Name:     "Cookie",
			Value:    pj64,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, c)
		// fmt.Println("cookie baked", p)
	}

	pj, err := base64.StdEncoding.DecodeString(c.Value)
	if err != nil {
		// fmt.Println(`b64 decode failed`)
		return p, err
	}
	// fmt.Println(`decoding from b64`, c.Value)

	err = json.Unmarshal([]byte(pj), &p)
	if err != nil {
		// fmt.Println(`json decode failed: `, err)
		return p, err
	}
	//fmt.Println(`decoding from json`, pj)
	// fmt.Println(`found`, p)
	return p, nil

}

// func andYouAre (w http.ResponseWriter, r *http.Request) person {
// 		c, err := r.Cookie("Cookie")
// 		if err != nil {
// 			fmt.Println("no cookie here")
// 		}

// 	de1, erd := base64.StdEncoding.DecodeString(c.Value)
// 		if erd != nil {
// 			fmt.Println(`b64 decode failed`)
// 		}
// 	fmt.Println(`decoding from b64`, de1)

// 	var guy person
// 	recog := json.Unmarshal([]byte(de1), &guy)
// 	if recog != nil {
// 			fmt.Println(`json decode failed: `, recog)
// 		}

// 	return guy
// }

func lookie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("at Hey A Thing")
	p, err := getPerson(w, r)
	if err != nil {
		fmt.Println("err getPerson", err)
	}
	tpl.ExecuteTemplate(w, "lookie.html", p)
}

func overhere(w http.ResponseWriter, r *http.Request) {
	fmt.Println("at Also Here")
	p, err := getPerson(w, r)
	if err != nil {
		fmt.Println("err getPerson", err)
	}
	tpl.ExecuteTemplate(w, "overhere.html", p)
}
