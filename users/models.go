package users

import (
	"errors"
	"github.com/rburawes/golang-demo/config"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// User object handles information about application's registered users.
type User struct {
	UserName  string
	Email     string
	Password  string
	Firstname string
	Lastname  string
}

// CheckUser looks for existing user using email or username
func CheckUser(un string) (User, error) {

	u, ok := FindUser(un)

	if ok {
		return u, errors.New("400. Bad Request. Username or email is taken")
	}

	return u, nil

}

// FindUser looks for registerd user by username.
func FindUser(un string) (User, bool) {

	u := User{}

	row := config.Database.QueryRow("SELECT u.username, u.email, u.password, u.firstname, u.lastname FROM users u WHERE u.username = $1", un)

	err := row.Scan(&u.UserName, &u.Email, &u.Password, &u.Firstname, &u.Lastname)

	if err != nil {
		return u, false
	}

	return u, true

}

// SaveUser create new user entry.
func SaveUser(r *http.Request) (User, error) {

	// Get form values and validate
	u, err := validateForm(r)

	if err != nil {
		return u, err
	}

	bs, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)

	if err != nil {
		return u, errors.New("500. Internal server error")
	}

	u.Password = string(bs)

	tx, err := config.Database.Begin()
	if err != nil {
		return u, err
	}

	stmt, err := tx.Prepare("INSERT INTO users (username, email, password, firstname, lastname) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return u, errors.New("500. Internal Server Error." + err.Error())
	}

	defer stmt.Close()
	// execute insert on 'books' table.
	if _, err := stmt.Exec(u.UserName, u.Email, u.Password, u.Firstname, u.Lastname); err != nil {
		tx.Rollback()
		return u, err
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		return u, errors.New("500. Internal Server Error." + err.Error())
	}

	return u, nil
}

func validateForm(r *http.Request) (User, error) {

	u := User{}
	p := r.FormValue("password")
	cp := r.FormValue("cpassword")
	f := r.FormValue("firstname")
	l := r.FormValue("lastname")
	e := r.FormValue("email")

	if p != cp {
		return u, errors.New("400. Bad Request. Password does not match")
	}

	if e == "" || p == "" || cp == "" {
		return u, errors.New("400. Bad Request. Fields cannot be empty")
	}

	_, err := CheckUser(e)

	if err != nil {
		return u, err
	}

	u.UserName = e
	u.Email = e
	u.Firstname = f
	u.Lastname = l
	u.Password = p

	return u, nil

}

// Validates the input password against the one in the database.
func (u *User) validatePassword(p string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	if err != nil {
		return false
	}

	return true

}
