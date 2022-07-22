package persistence

import (
	"github.com/dungbk10t/test_api/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveUser_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	var user = entity.User{}
	user.Email = "dung101@gmail.com"
	user.Name = "dung101"
	user.Password = "123456"

	repo := NewUserRepository(conn)

	u, saveErr := repo.SaveUser(&user)
	assert.Nil(t, saveErr)
	assert.EqualValues(t, u.Email, "dung101@gmail.com")
	assert.EqualValues(t, u.Name, "dung101")
	// password has been hashed, so it should not the same old pass (use NotEqual)
	assert.NotEqual(t, u.Password, "123456")
}

//Test truong hop Tao user nhung da trung email voi 1 use co san trong DB. (O day la user trong co san trong seedUser)
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
	user.Email = "dung100@gmail.com"
	user.Name = "dung102"
	user.Password = "123456"

	repo := NewUserRepository(conn)
	u, saveErr := repo.SaveUser(&user)
	dbMsg := map[string]string{
		"email_taken": "email already taken",
	}
	assert.Nil(t, u)
	assert.EqualValues(t, dbMsg, saveErr)
}

//Test get use sao cho tra ra test case dung trung voi email va name cua use co san trong DB.
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
	assert.EqualValues(t, u.Email, "dung100@gmail.com")
	assert.EqualValues(t, u.Name, "dung100")
}

// Kiem tra TH get user by id TH fail khi tra ve ID khac voi ID user can get.
func TestGetUser_Failure(t *testing.T) {
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
	assert.NotEqual(t, u.ID, "100")
}

// Tra ra so luong 3 object user dung voi so luong khoi tao co trong db seedUsers
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
	assert.EqualValues(t, len(users), 3)
}

func TestGetUserByEmailAndPassword_Failure(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("wan non error, got %#v", err)
	}
	//seed user
	u, err := seedUser(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	// TH1 : Dia chi email khong co trong DB
	var user_case_1 = &entity.User{
		Email:    "dung200@gmail.com",
		Password: "123456",
	}
	// TH2 : Dia chi Email dung, nhung mat khau khong hop le
	var user_case_2 = &entity.User{
		Email:    "dung100@gmail.com",
		Password: "123457",
	}
	repo := NewUserRepository(conn)
	u, saveErr1 := repo.GetUserByEmailAndPassword(user_case_1)
	u, saveErr2 := repo.GetUserByEmailAndPassword(user_case_2)
	dbMsg1 := map[string]string{
		"no_user": "user not found",
		//"incorrect_password": "incorrect password",
	}
	dbMsg2 := map[string]string{
		//"no_user": "user not found",
		"incorrect_password": "incorrect password",
	}
	assert.Nil(t, u)
	assert.EqualValues(t, dbMsg1, saveErr1)
	assert.EqualValues(t, dbMsg2, saveErr2)
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
		Email:    "dung100@gmail.com",
		Password: "123456",
	}
	repo := NewUserRepository(conn)
	u, getErr := repo.GetUserByEmailAndPassword(user)

	assert.Nil(t, getErr)
	assert.EqualValues(t, u.Email, user.Email)
	// the password from database should not be equal to a plane password, because it is hashed
	assert.NotEqual(t, u.Password, user.Password)
}

func TestDeleteUser_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	user, err := seedUser(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	repo := NewUserRepository(conn)
	deleteErr := repo.DeleteUser(user.ID)
	assert.Nil(t, deleteErr)
}
