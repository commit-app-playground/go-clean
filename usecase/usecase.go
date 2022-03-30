package usecase

import (
	"toporet/hop/goclean/entity"
)

//Hey use case there is out there some user saver that will save the //user
type UserSaver interface {
	Save(user entity.User) error
}

//UserRegisterUseCase will store his dependencies
type UserRegisterUseCase struct {
	userSaver UserSaver
}

func NewUserRegisterUseCase(userSaver UserSaver) *UserRegisterUseCase {
	return &UserRegisterUseCase{userSaver}
}

func (u UserRegisterUseCase) RegisterUser(user entity.User) error {
	err := user.Validate()
	if err == nil {
		err = u.userSaver.Save(user)
	}
	return err
}
