package application

import (
	"testing"

	"github.com/dungbk10t/test_api/domain/entity"
	"github.com/stretchr/testify/assert"
)

var (
	saveUserRepo                func(user *entity.User) (*entity.User, map[string]string)
	updateInfoUserRepo          func(userId uint64, user *entity.User) (*entity.User, error)
	updatePassWordUserRepo      func(userId uint64, user *entity.User) (*entity.User, error)
	getUserRepo                 func(userId uint64) (*entity.User, error)
	getUsersRepo                func() ([]entity.User, error)
	getUserEmailAndPasswordRepo func(user *entity.User) (*entity.User, map[string]string)
	deleteUserRepo              func(userId uint64) error
)

type fakeUserRepo struct{}

func (u *fakeUserRepo) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	return saveUserRepo(user)
}

func (u *fakeUserRepo) UpdateInfoUser(userId uint64, user *entity.User) (*entity.User, error) {
	return updateInfoUserRepo(userId, user)
}

func (u *fakeUserRepo) UpdatePassWordUser(userId uint64, user *entity.User) (*entity.User, error) {
	return updatePassWordUserRepo(userId, user)
}

func (u *fakeUserRepo) GetUser(userId uint64) (*entity.User, error) {
	return getUserRepo(userId)
}
func (u *fakeUserRepo) GetUsers() ([]entity.User, error) {
	return getUsersRepo()
}
func (u *fakeUserRepo) DeleteUser(userId uint64) error {
	return deleteUserRepo(userId)
}

func (u *fakeUserRepo) GetUserByEmailAndPassword(user *entity.User) (*entity.User, map[string]string) {
	return getUserEmailAndPasswordRepo(user)
}

var userAppFake UserAppInterface = &fakeUserRepo{}

func TestSaveUser_Success(t *testing.T) {
	saveUserRepo = func(user *entity.User) (*entity.User, map[string]string) {
		return &entity.User{
			ID:       1,
			Name:     "dung01",
			Email:    "dung01@gmail.com",
			Password: "123456",
		}, nil
	}
	user := &entity.User{
		ID:       1,
		Name:     "dung01",
		Email:    "dung01@gmail.com",
		Password: "123456",
	}
	u, err := userAppFake.SaveUser(user)
	assert.Nil(t, err)
	assert.EqualValues(t, u.Name, "dung01")
	assert.EqualValues(t, u.Email, "dung01@gmail.com")
}

func TestUpdateInfoUser_Success(t *testing.T) {
	updateInfoUserRepo = func(userId uint64, user *entity.User) (*entity.User, error) {
		return &entity.User{
			ID:       1,
			Name:     "dung01",
			Email:    "dung01@gmail.com",
			Password: "123456",
		}, nil
	}
	user := &entity.User{
		ID:       1,
		Name:     "dung01",
		Email:    "dung01@gmail.com",
		Password: "123456",
	}
	userId := uint64(1)
	u, err := userAppFake.UpdateInfoUser(userId, user)
	assert.Nil(t, err)
	assert.EqualValues(t, u.Name, "dung01")
	assert.EqualValues(t, u.Email, "dung01@gmail.com")
}

func TestUpdatePassWordUser_Success(t *testing.T) {
	updatePassWordUserRepo = func(userId uint64, user *entity.User) (*entity.User, error) {
		return &entity.User{
			ID:       1,
			Name:     "dung01",
			Email:    "dung01@gmail.com",
			Password: "123456",
		}, nil
	}
	user := &entity.User{
		ID:       1,
		Name:     "dung01",
		Email:    "dung01@gmail.com",
		Password: "123456",
	}
	userId := uint64(1)
	u, err := userAppFake.UpdatePassWordUser(userId, user)
	assert.Nil(t, err)
	assert.EqualValues(t, u.Password, "123456")
	assert.EqualValues(t, u.Email, "dung01@gmail.com")
}

func TestGetUser_Success(t *testing.T) {
	getUserRepo = func(userId uint64) (*entity.User, error) {
		return &entity.User{
			ID:       1,
			Name:     "dung01",
			Email:    "dung01@gmail.com",
			Password: "123456",
		}, nil
	}
	userId := uint64(1)
	u, err := userAppFake.GetUser(userId)
	assert.Nil(t, err)
	assert.EqualValues(t, u.Name, "dung01")
	assert.EqualValues(t, u.Email, "dung01@gmail.com")
}

func TestDeleteUser_Success(t *testing.T) {
	deleteUserRepo = func(userId uint64) error {
		return nil
	}
	userId := uint64(1)
	err := userAppFake.DeleteUser(userId)
	assert.Nil(t, err)
}
func TestGetUsers_Success(t *testing.T) {
	getUsersRepo = func() ([]entity.User, error) {
		return []entity.User{
			{
				ID:       1,
				Name:     "dung01",
				Email:    "dung01@gmail.com",
				Password: "123456",
			},
			{
				ID:       2,
				Name:     "dung02",
				Email:    "dung02@gmail.com",
				Password: "123456",
			},
		}, nil
	}
	users, err := userAppFake.GetUsers()
	assert.Nil(t, err)
	assert.EqualValues(t, len(users), 2)
}

func TestGetUserByEmailAndPassword_Success(t *testing.T) {
	//Mock the response coming from the infrastructure
	getUserEmailAndPasswordRepo = func(user *entity.User) (*entity.User, map[string]string) {
		return &entity.User{
			ID:       1,
			Name:     "dung01",
			Email:    "dung01@gmail.com",
			Password: "123456",
		}, nil
	}
	user := &entity.User{
		ID:       1,
		Name:     "dung01",
		Email:    "dung01@gmail.com",
		Password: "123456",
	}
	u, err := userAppFake.GetUserByEmailAndPassword(user)
	assert.Nil(t, err)
	assert.EqualValues(t, u.Name, "dung01")
	assert.EqualValues(t, u.Email, "dung01@gmail.com")
}
