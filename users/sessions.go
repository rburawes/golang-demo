package users

import (
	"github.com/satori/go.uuid"
	"net/http"
)

// StoredSessions has session id the user ID
var StoredSessions = map[string]string{} // session ID, user ID

// CreateSession creates a session for a logged user.
func CreateSession(w http.ResponseWriter, u User) {

	// create session
	sID, _ := uuid.NewV4()
	c := &http.Cookie{
		Name:  "session",
		Value: sID.String(),
	}

	refreshCookie(w, c)
	StoredSessions[c.Value] = u.UserName

}

// RemoveSession deletes expires user's session.
func RemoveSession(w http.ResponseWriter, r *http.Request) {

	c, _ := r.Cookie("session")
	// delete the session
	delete(StoredSessions, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	refreshCookie(w, c)
}

// IsLoggedIn verifies if the user is already in the session and logged in.
func IsLoggedIn(r *http.Request) bool {

	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	un := StoredSessions[c.Value]
	_, ok := FindUser(un)
	return ok

}

// refreshes cookie value
func refreshCookie(w http.ResponseWriter, c *http.Cookie)  {
	c.Path = "/"
	http.SetCookie(w, c)
}
