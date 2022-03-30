package infrastructure

import (
	"database/sql"

	_ "github.com/lib/pq"

	"toporet/hop/goclean/entity"
)

//UserRepository to manage users persistence
type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{db}
}

//Save user
func (r UserRepository) Save(user entity.User) error {
	//...
	return nil
}

//GetByEmail gets the user by email
func GetByEmail(email string) error {
	//...
	return nil
}
