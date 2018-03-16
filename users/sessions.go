package users

import (
	"github.com/satori/go.uuid"
	"net/http"
)

// CurrentUsers are the currently logged in users.
var CurrentUsers = map[string]User{}

// StoredSessions has session id the user ID
var StoredSessions = map[string]string{} // session ID, user ID

// GetUser retrieves the currently logged in user.
func GetUser(w http.ResponseWriter, req *http.Request) User {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u User
	if un, ok := StoredSessions[c.Value]; ok {
		u = CurrentUsers[un]
	}
	return u
}

// IsLoggedIn verifies if the user is already in the session and logged in.
func IsLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := StoredSessions[c.Value]
	_, ok := CurrentUsers[un]
	return ok
}
