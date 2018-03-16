package users

import (
	"github.com/rburawes/golang-demo/config"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		StoredSessions[c.Value] = u.UserName

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
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the password provided match the password in the db
		err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		CurrentUsers[un] = u
		http.SetCookie(w, c)
		StoredSessions[c.Value] = un
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	config.TPL.ExecuteTemplate(w, "login.gohtml", u)
}
