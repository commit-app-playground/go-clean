package gateway

import (
	"database/sql"
	"toporet/hop/goclean/entity"
)

type IGetAllBooks func() ([]entity.Book, error)

type BookStore struct {
	db *sql.DB
}

func NewBookStore(db *sql.DB) BookStore {
	return BookStore{db}
}

func (br BookStore) RetrieveAll() ([]entity.Book, error) {
	rows, err := br.db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bks []entity.Book

	for rows.Next() {
		var bk entity.Book

		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
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
