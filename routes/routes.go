package routes

import (
	"github.com/rburawes/golang-demo/authors"
	"github.com/rburawes/golang-demo/books"
	"github.com/rburawes/golang-demo/users"
	"net/http"
)

// LoadRoutes handles routes to pages of the application.
func LoadRoutes() {

	// Index or main page.
	http.HandleFunc("/", index)
	// Book related routes
	http.HandleFunc("/books", books.Index)
	http.HandleFunc("/books/show", books.Show)
	http.HandleFunc("/books/create", books.Create)
	http.HandleFunc("/books/create/process", books.CreateProcess)
	http.HandleFunc("/books/update", books.Update)
	http.HandleFunc("/books/update/process", books.UpdateProcess)
	http.HandleFunc("/books/delete/process", books.DeleteProcess)
	// Author related route(s)
	http.HandleFunc("/authors", authors.GetAuthors)
	// User related route(s)
	http.HandleFunc("/signup", users.Signup)
	http.HandleFunc("/login", users.Login)
	http.HandleFunc("/logout", users.Logout)
	// CSS, JS and images
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./resource"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	// Listens and serve requests.
	http.ListenAndServe(":8080", nil)

}

// Redirect to list of books.
func index(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/books", http.StatusSeeOther)

}
