package books

import (
	"errors"
	"fmt"
	"github.com/rburawes/golang-demo/config"
	"net/http"
	"strconv"
)

// Book holds information about the book entry.
type Book struct {
	Isbn      string
	Title     string
	Author    string
	Price     float32
	AuthorID  int32
	TheAuthor string
}

// AllBooks retriev all the books from the database.
func AllBooks() ([]Book, error) {

	rows, err := config.Database.Query("SELECT b.isbn, b.title,  concat(a.firstname, ' ', a.lastname) as author, b.price, a.id FROM books b INNER JOIN book_authors ba on b.isbn = ba.book_isbn INNER JOIN authors a ON ba.author_id = a.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price, &bk.AuthorID) // order matters
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil

}

// GetBook retrieve specific book record from the database.
func GetBook(r *http.Request) (Book, error) {

	bk := Book{}
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return bk, errors.New("400. Bad Request")
	}

	row := config.Database.QueryRow("SELECT b.isbn, b.title, concat(a.firstname, ' ', a.lastname) as author, b.price, a.id, a.about FROM books b INNER JOIN book_authors ba on b.isbn = ba.book_isbn INNER JOIN authors a ON ba.author_id = a.id WHERE b.isbn = $1", isbn)

	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price, &bk.AuthorID, &bk.TheAuthor)
	if err != nil {
		return bk, err
	}

	return bk, nil

}

// SaveBook create new book entry.
func SaveBook(r *http.Request) (Book, error) {

	// Get form values and validate
	bk, err := validateForm(r)

	if err != nil {
		return bk, err
	}

	tx, err := config.Database.Begin()
	if err != nil {
		return bk, err
	}

	stmt, err := tx.Prepare("INSERT INTO books (isbn, title, price) VALUES ($1, $2, $3)")
	if err != nil {
		return bk, errors.New("500. Internal Server Error." + err.Error())
	}

	defer stmt.Close()
	// execute insert on 'books' table.
	if _, err := stmt.Exec(bk.Isbn, bk.Title, bk.Price); err != nil {
		tx.Rollback()
		return bk, err
	}

	stmt, err = tx.Prepare("INSERT INTO book_authors (book_isbn, author_id) VALUES ($1, $2)")
	if err != nil {
		return bk, err
	}

	defer stmt.Close()

	// execute insert on 'book_authors' table.
	if _, err := stmt.Exec(bk.Isbn, bk.AuthorID); err != nil {
		tx.Rollback()
		return bk, err
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		return bk, errors.New("500. Internal Server Error." + err.Error())
	}

	return GetBook(r)

}

// UpdateBook modifies existing book details.
func UpdateBook(r *http.Request) (Book, error) {

	// Get form values and validate
	bk, err := validateForm(r)

	if err != nil {
		return bk, err
	}

	tx, err := config.Database.Begin()
	if err != nil {
		return bk, err
	}

	stmt, err := tx.Prepare("UPDATE books SET isbn = $1, title=$2, price=$3 WHERE isbn=$1;")
	if err != nil {
		return bk, err
	}

	defer stmt.Close()
	// execute update on 'books' table.
	if _, err := stmt.Exec(bk.Isbn, bk.Title, bk.Price); err != nil {
		tx.Rollback()
		return bk, err
	}

	stmt, err = tx.Prepare("UPDATE book_authors SET author_id = $1 WHERE book_isbn=$2")

	if err != nil {
		return bk, err
	}

	defer stmt.Close()

	// execute update on 'book_authors' table.
	if _, err := stmt.Exec(bk.AuthorID, bk.Isbn); err != nil {
		tx.Rollback()
		return bk, err
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		return bk, err
	}

	return GetBook(r)

}

// DeleteBook removes book entry from the database.
func DeleteBook(r *http.Request) error {

	isbn := r.FormValue("isbn")
	if isbn == "" {
		return errors.New("400. Bad Request")
	}

	tx, err := config.Database.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("DELETE FROM books WHERE isbn=$1;")
	if err != nil {
		return err
	}

	defer stmt.Close()

	// execute delete on 'books' table.
	if _, err := stmt.Exec(isbn); err != nil {
		tx.Rollback()
		return errors.New("500. Internal Server Error. Unable to delete the book")
	}

	stmt, err = tx.Prepare("DELETE FROM book_authors WHERE book_isbn=$1;")
	if err != nil {
		return errors.New("500. Internal Server Error. Unable to delete the book")
	}

	defer stmt.Close()

	// execute delete on 'books' table.
	if _, err := stmt.Exec(isbn); err != nil {
		tx.Rollback()
		return errors.New("500. Internal Server Error. Unable to delete the book")
	}

	tx.Commit()
	if err != nil {
		return errors.New("500. Internal Server Error. Unable to delete the book")
	}

	return nil

}

// FormatBookPrice formats the price of the book.
func (book *Book) FormatBookPrice() string {

	return fmt.Sprintf("$%.2f", book.Price)

}

func validateForm(r *http.Request) (Book, error) {

	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	p := r.FormValue("price")
	a := r.FormValue("author")

	if bk.Isbn == "" || bk.Title == "" {
		return bk, errors.New("400. Bad Request. Fields cannot be empty")
	}

	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. Not Acceptable. Price must be a number")
	}

	bk.Price = float32(f64)

	int64, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		return bk, errors.New("406. Not Acceptable. Author ID can't be processed")
	}

	if int32(int64) <= 0 {
		return bk, errors.New("400. Bad Request. Invalid author id")
	}

	bk.AuthorID = int32(int64)
	return bk, nil

}
