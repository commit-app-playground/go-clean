package gateway

import (
	"database/sql"
	"toporet/hop/goclean/entity"
)

// type IGetAllBooks interface {
// 	 AllBooks() ([]entity.Book, error)
// }

type IGetAllBooks func() ([]entity.Book, error)

type BookRepository struct {
	db *sql.DB
	// usecase.BookRetriever
}

func NewBookRepository(db *sql.DB) BookRepository {
	return BookRepository{
		db: db,
		// IGetAllBooks2: AllBooks1,
	}
}

// Update the AllBooks function so it accepts the connection pool as a
// parameter.
// func (br BookRepository) AllBooks() ([]entity.Book, error) {
func (br BookRepository) RetrieveAllBooks() ([]entity.Book, error) {
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
