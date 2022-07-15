package repository

import "github.com/dungbk10t/test_api/domain/entity"

type UserRepository interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	UpdateInfoUser(uint64, *entity.User) (*entity.User, error)
	GetUser(uint64) (*entity.User, error)
	DeleteUser(uint64) error
	GetUsers() ([]entity.User, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, map[string]string)
}
