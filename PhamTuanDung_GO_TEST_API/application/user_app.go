package application

import (
	"github.com/dungbk10t/test_api/domain/entity"
	"github.com/dungbk10t/test_api/domain/repository"
)

type userApp struct {
	us repository.UserRepository
}

//UserApp implements the UserAppInterface
var _ UserAppInterface = &userApp{}

type UserAppInterface interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	UpdateInfoUser(uint64, *entity.User) (*entity.User, error)
	GetUsers() ([]entity.User, error)
	GetUser(uint64) (*entity.User, error)
	DeleteUser(uint64) error
	GetUserByEmailAndPassword(*entity.User) (*entity.User, map[string]string)
}

func (u *userApp) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	return u.us.SaveUser(user)
}

func (u *userApp) UpdateInfoUser(userId uint64, user *entity.User) (*entity.User, error) {
	return u.us.UpdateInfoUser(userId, user)
}

func (u *userApp) GetUser(userId uint64) (*entity.User, error) {
	return u.us.GetUser(userId)
}

func (u *userApp) DeleteUser(userId uint64) error {
	return u.us.DeleteUser(userId)
}

func (u *userApp) GetUsers() ([]entity.User, error) {
	return u.us.GetUsers()
}

func (u *userApp) GetUserByEmailAndPassword(user *entity.User) (*entity.User, map[string]string) {
	return u.us.GetUserByEmailAndPassword(user)
}
