package gateway

import (
	"database/sql"

	"toporet/hop/goclean/entity"
)

// type IGetAllBooks func() ([]entity.Book, error)

type TaskStore struct {
	db *sql.DB
}

func NewTaskStore(db *sql.DB) TaskStore {
	return TaskStore{db}
}

func (s TaskStore) SaveNewTask(t entity.Task) (*entity.TaskId, error) {
	//
	// TODO: implement saving for realz
	//
	return entity.NewTaskId("foo-bar-42")
}

func (s TaskStore) SaveTask(t entity.Task) error {
	//
	// TODO: implement saving for realz
	//
	return nil
}

// func (s TaskStore) RetrieveAll() ([]entity.Book, error) {
// 	rows, err := s.db.Query("SELECT * FROM books")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var books []entity.Book

// 	for rows.Next() {
// 		var bk entity.Book

// 		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
// 		if err != nil {
// 			return nil, err
// 		}

// 		books = append(books, bk)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return books, nil
// }
