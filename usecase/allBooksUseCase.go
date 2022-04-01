package usecase

import (
	"toporet/hop/goclean/entity"
)

type BookRetriever interface {
	RetrieveAllBooks( /* user entity.User */ ) ([]entity.Book, error)
}

//UserRegisterUseCase will store his dependencies
type AllBoksUseCase struct {
	bookRetrieverGateway BookRetriever
}

func NewAllBooksUseCase(br BookRetriever) *AllBoksUseCase {
	return &AllBoksUseCase{br}
}

func (u AllBoksUseCase) GetAllBooks( /* pass an input model later | user entity.User */ ) ([]entity.Book, error) {
	// err := user.Validate()
	// if err == nil {
	// 	err = u.userSaver.Save(user)
	// }
	// return err

	// potentially more IO interactions and business logic here

	return u.bookRetrieverGateway.RetrieveAllBooks()
}
