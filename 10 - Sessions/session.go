package main

import (
	"github.com/satori/go.uuid"
	"net/http"
	"fmt"
)

func getUser(w http.ResponseWriter, r *http.Request) user {
	c, err := r.Cookie("session")
	
	if err != nil {
		fmt.Println(`No cookie; baking now`)
		id := uuid.NewV4()
		c = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, c)
		fmt.Println(`Cookie baked`)
	}

	var u user
	if uid, ok := dbSessions[c.Value]; ok {
		u = dbUsers[uid]
		return u
	}
	
	// create user
	if r.Method == http.MethodPost {
		e := r.FormValue("email")
		p := r.FormValue("password")
		f := r.FormValue("first")
		l := r.FormValue("last")
		u = user{e, p, f, l}
		dbSessions[c.Value] = e
		dbUsers[e] = u
	}
	return u
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}

	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}
