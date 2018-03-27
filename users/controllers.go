package users

import (
	"github.com/rburawes/golang-demo/config"
	"net/http"
	"time"
)

// Signup allows the user to create an account.
func Signup(w http.ResponseWriter, r *http.Request) {

	if IsLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var u User
	// process form submission
	if r.Method == http.MethodPost {

		u, err := SaveUser(r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// create session
		CreateSession(w, u)

		// redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	config.TPL.ExecuteTemplate(w, "signup.gohtml", u)

}

// Login allows registered user to access the application.
func Login(w http.ResponseWriter, r *http.Request) {

	if IsLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var u User
	// process form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")
		e := r.FormValue("email")

		if un == "" {
			un = e
		}

		// check if the user exists
		u, ok := FindUser(un)
		if !ok {
			http.Error(w, "username and/or password do not match", http.StatusForbidden)
			return
		}

		if !u.validatePassword(p) {
			http.Error(w, "username and/or password do not match", http.StatusForbidden)
			return
		}

		// create session
		CreateSession(w, u)

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	config.TPL.ExecuteTemplate(w, "login.gohtml", u)

}

// Logout method to call when the user signed out of the application.
func Logout(w http.ResponseWriter, r *http.Request) {

	if !IsLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	RemoveSession(w, r)

	// not the best place and not to be used in production
	if time.Now().Sub(storedSessionClean) > (time.Second * 30) {
		go cleanSessions()
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
