package db

import (
	"database/sql"
	"fmt"

	"github.com/sanchae/library-test/models"
)

func (db Database) GetAllBooks() (*models.BookList, error) {
	list := &models.BookList{}

	rows, err := db.Conn.Query(`
	SELECT * FROM books
	`)
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.BookTitle, &book.AvailableCopies)
		if err != nil {
			return list, err
		}
		list.Books = append(list.Books, book)
	}
	return list, nil
}

func (db Database) LendBook(bookID, userID int) error {
	tx, err := db.Conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var availableCopies int
	err = tx.QueryRow(`
	SELECT available_copies FROM books WHERE book_id = $1
	`, bookID).Scan(&availableCopies)

	if err != nil {
		return err
	}

	if availableCopies <= 0 {
		return fmt.Errorf("no available copies")
	}

	_, err = tx.Exec(`
	UPDATE books SET available_copies = available_copies - 1 WHERE book_id = $1
	`, bookID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
	INSERT INTO borrowed_books (user_id, book_id) VALUES ($1, $2)
	`, userID, bookID)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (db Database) ReturnBook(bookID, userID int) error {
	tx, err := db.Conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var borrowedID int
	err = tx.QueryRow(`
	SELECT borrowed_id FROM borrowed_books WHERE user_id = $1 AND book_id = $2
	`, userID, bookID).Scan(&borrowedID)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("book was not borrowed by this user")
		}
		return err
	}

	_, err = tx.Exec(`
	UPDATE books SET available_copies = available_copies + 1 WHERE book_id = $1
	`, bookID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
	DELETE FROM borrowed_books WHERE borrowed_id = $1
	`, borrowedID)
    if err != nil {
        return err
    }

    return tx.Commit()
}