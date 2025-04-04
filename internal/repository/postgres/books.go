package postgres

import (
	"github.com/indienSs/go-std/internal/models"
	"github.com/indienSs/go-std/internal/types"
)

func (pg *Postgres) GetBooks() ([]models.Book, error) {
	rows, err := pg.db.Query("SELECT id, title, author, publicationDate FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublicationDate); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	
	return books, nil
}

func (pg *Postgres) GetBook(id int) (models.Book, error) {
	var book models.Book
	err := pg.db.QueryRow("SELECT id, title, author, publicationDate FROM books WHERE id = $1", id).
		Scan(&book.ID, &book.Title, &book.Author, &book.PublicationDate)
	
	if err != nil {
		return book, err
	}

	return book, nil
}

func (pg *Postgres) CreateBook(book *models.Book) error {
	err := pg.db.QueryRow(
		"INSERT INTO books (title, author, publicationDate) VALUES ($1, $2, $3) RETURNING id",
		book.Title, book.Author, book.PublicationDate,
	).Scan(&book.ID)
	
	if err != nil {
		return err
	}

	return nil
}

func (pg *Postgres) UpdateBook(id int, book models.Book) error {
	_, err := pg.db.Exec(
		"UPDATE books SET title = $1, author = $2, publicationDate = $3 WHERE id = $4",
		book.Title, book.Author, book.PublicationDate, id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (pg *Postgres) DeleteBook(id int) error {
	result, err := pg.db.Exec("DELETE FROM books WHERE id = $1", id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return types.ErrNotFound
	}

	return nil
}