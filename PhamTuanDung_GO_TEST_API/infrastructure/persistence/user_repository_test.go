package persistence

import (
	"github.com/dungbk10t/test_api/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveUser_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatal("want non error, got %#v", err)
	}
	var user = entity.User{}
	user.Email = "dungbk10t@gmail.com"
	user.Name = "trung02"
	user.Password = "123458"

	repo := NewUserRepository(conn)

	u, saveErr := repo.SaveUser(&user)
	assert.Nil(t, saveErr)
	assert.EqualValues(t, u.Email, "dung02@gmail.com")
	assert.EqualValues(t, u.Name, "trung02")
	// password has been hashed, so it should not the same old pass
	assert.EqualValues(t, u.Password, "123458")
}

func TestSaveUser_Failure(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	// sed user
	_, err = seedUser(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	var user = entity.User{}
	user.Email = "dung02@gmail.com"
	user.Name = "duc02"
	user.Password = "123458"

	repo := NewUserRepository(conn)
	u, saveErr := repo.SaveUser(&user)
	dbMsg := map[string]string{
		"email_taken": "email already taken",
	}
	assert.Nil(t, u)
	assert.EqualValues(t, dbMsg, saveErr)
}

func TestGetUser_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	// seed user
	user, err := seedUser(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	repo := NewUserRepository(conn)
	u, getErr := repo.GetUser(user.ID)

	assert.Nil(t, getErr)
	assert.EqualValues(t, u.Email, "dung02@gmail.com")
	assert.EqualValues(t, u.Name, "vic")
}

func TestGetUsers_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	// seed user
	_, err = seedUsers(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	repo := NewUserRepository(conn)
	users, getErr := repo.GetUsers()

	assert.Nil(t, getErr)
	assert.EqualValues(t, len(users), 2)
}

func TestGetUserByEmailAndPassword_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("wan non error, got %#v", err)
	}
	//seed user
	u, err := seedUser(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	var user = &entity.User{
		Email:    "dung02@gmail.com",
		Password: "123458",
	}
	repo := NewUserRepository(conn)
	u, getErr := repo.GetUserByEmailAndPassword(user)

	assert.Nil(t, getErr)
	assert.EqualValues(t, u.Email, user.Email)
	// the password from database should not be equal to a plane password, because it is hashed
}

func TestDeleteUser_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	food, err := seedUser(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	repo := NewUserRepository(conn)

	deleteErr := repo.DeleteUser(food.ID)

	assert.Nil(t, deleteErr)
}
