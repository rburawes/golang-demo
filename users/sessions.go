package users

import (
	"github.com/satori/go.uuid"
	"net/http"
	"time"
)

type session struct {
	un           string
	lastActivity time.Time
}

// StoredSessions has session id the user ID
var storedSessions = make(map[string]session) // session ID, user ID

// Amount of time for the application to clean all the sessions after log out
var storedSessionClean time.Time

func init() {
	storedSessionClean = time.Now()
}

// CreateSession creates a session for a logged user.
func CreateSession(w http.ResponseWriter, u User) {

	// create session
	sID, _ := uuid.NewV4()
	c := &http.Cookie{
		Name:  "session",
		Value: sID.String(),
	}

	refreshCookie(w, c)
	storedSessions[c.Value] = session{u.UserName, time.Now()}

}

// RemoveSession deletes expires user's session.
func RemoveSession(w http.ResponseWriter, r *http.Request) {

	c, _ := r.Cookie("session")
	// delete the session
	delete(storedSessions, c.Value)
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

	s, ok := storedSessions[c.Value]
	if ok {
		s.lastActivity = time.Now()
		storedSessions[c.Value] = s
	}

	_, ok = FindUser(s.un)
	return ok

}

// remove all the sessions after logout, all that
// has been inactive for more than 30 seconds
func cleanSessions() {
	for k, v := range storedSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(storedSessions, k)
		}
	}
	storedSessionClean = time.Now()
}

// refreshes cookie value
func refreshCookie(w http.ResponseWriter, c *http.Cookie) {
	c.Path = "/"
	http.SetCookie(w, c)
}
