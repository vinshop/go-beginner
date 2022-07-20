package interfaces

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dungbk10t/test_api/domain/entity"
	"github.com/dungbk10t/test_api/infrastructure/auth"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

//Don't use mock, use table test to cover all validation errors
func Test_Login_Invalid_Data(t *testing.T) {
	samples := []struct {
		inputJSON  string
		statusCode int
	}{
		{
			//empty email
			inputJSON:  `{"email": "","password": "123456"}`,
			statusCode: 422,
		},
		{
			//empty password
			inputJSON:  `{"email": "dung01@gmail.com","password": ""}`,
			statusCode: 422,
		},
		{
			//invalid email
			inputJSON:  `{"email": "dung01com","password": ""}`,
			statusCode: 422,
		},
	}

	for _, v := range samples {

		r := gin.Default()
		r.POST("/login", au.Login)
		req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v\n", err)
		}
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		validationErr := make(map[string]string)

		err = json.Unmarshal(rr.Body.Bytes(), &validationErr)
		if err != nil {
			t.Errorf("error unmarshalling error %s\n", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)

		if validationErr["email_required"] != "" {
			assert.Equal(t, validationErr["email_required"], "email is required")
		}
		if validationErr["invalid_email"] != "" {
			assert.Equal(t, validationErr["invalid_email"], "please provide a valid email")
		}
		if validationErr["password_required"] != "" {
			assert.Equal(t, validationErr["password_required"], "password is required")
		}
	}
}

func Test_Login_Success(t *testing.T) {

	userApp.GetUserByEmailAndPasswordFn = func(*entity.User) (*entity.User, map[string]string) {
		return &entity.User{
			ID:   1,
			Name: "dung01",
		}, nil
	}
	fakeToken.CreateTokenFn = func(userid uint64) (*auth.TokenDetails, error) {
		return &auth.TokenDetails{
			AccessToken:  "this-is-the-access-token",
			RefreshToken: "this-is-the-refresh-token",
			TokenUuid:    "dfsdf-342-34-23-4234-234",
			RefreshUuid:  "sfd-3234-sdfew-34234-df3",
			AtExpires:    12345,
			RtExpires:    1234555,
		}, nil
	}
	fakeAuth.CreateAuthFn = func(uint64, *auth.TokenDetails) error {
		return nil
	}

	inputJSON := `{"email": "dung01@gmail.com.com","password": "123456"}`
	r := gin.Default()
	r.POST("/login", au.Login)
	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(inputJSON))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	fmt.Println("The response: ", string(rr.Body.Bytes()))

	response := make(map[string]interface{})

	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("error unmarshalling error %s\n", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.EqualValues(t, response["access_token"], "this-is-the-access-token")
	assert.EqualValues(t, response["refresh_token"], "this-is-the-refresh-token")
	assert.EqualValues(t, response["name"], "dung01")
}

func TestLogout_Success(t *testing.T) {
	//Mock extracting metadata
	fakeToken.ExtractTokenMetadataFn = func(r *http.Request) (*auth.AccessDetails, error) {
		return &auth.AccessDetails{
			TokenUuid: "0237817a-1546-4ca3-96a4-17621c237f6b",
			UserId:    1,
		}, nil
	}
	//Mock the methods that Logout depend on
	fakeAuth.DeleteTokensFn = func(*auth.AccessDetails) error {
		return nil
	}
	//This can be anything, since we have already mocked the method that checks if the token is valid or not and have told it what to return for us.
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6ImEzNmIwZjJmLWM5ODctNDk0My05MjBiLWVjNDc2ZGNjMTAzYyIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTY1ODI5MjA3MSwidXNlcl9pZCI6MTA4fQ.bEfJDdBx1rq3KSNshvCM8cK1utMCtLgx5jQLl92r9NU"

	tokenString := fmt.Sprintf("Bearer %v", token)

	req, err := http.NewRequest(http.MethodPost, "/logout", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r := gin.Default()
	r.POST("/logout", au.Logout)
	req.Header.Set("Authorization", tokenString)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	response := ""
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("cannot unmarshal response: %v\n", err)
	}
	assert.EqualValues(t, rr.Code, http.StatusOK)
	assert.EqualValues(t, response, "Successfully logged out")
}

func TestRefresh_Success(t *testing.T) {

	fakeAuth.DeleteRefreshFn = func(string) error {
		return nil
	}
	fakeToken.CreateTokenFn = func(userid uint64) (*auth.TokenDetails, error) {
		return &auth.TokenDetails{
			AccessToken:  "this-is-the-NEW-access-token",
			RefreshToken: "this-is-the-NEW-refresh-token",
			TokenUuid:    "dfsdf-342-34-23-4234-234",
			RefreshUuid:  "sfd-3234-sdfew-34234-df3",
			AtExpires:    12345,
			RtExpires:    1234555,
		}, nil
	}
	fakeAuth.CreateAuthFn = func(uint64, *auth.TokenDetails) error {
		return nil
	}

	r := gin.Default()
	r.POST("/refresh", au.Refresh)

	//Note that since we will be cheking this token, A secret is needed. THis secret was used to create the token,
	//lets set it, so that this test can retrieve it. Setting it this way we save us from importing the .env, which we dont really need.
	os.Setenv("REFRESH_SECRET", "786dfdbjhsb")

	inputJSON := `{
		"refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg4OTU5NzEsInJlZnJlc2hfdXVpZCI6ImEzNmIwZjJmLWM5ODctNDk0My05MjBiLWVjNDc2ZGNjMTAzYysrMTA4IiwidXNlcl9pZCI6MTA4fQ.wB1tLQjlV8w8YNZvdSkk-MkI5DbA4UxaQdtAtWke5lc"
		}`
	req, err := http.NewRequest(http.MethodPost, "/refresh", bytes.NewBufferString(inputJSON))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	fmt.Println("the response: ", string(rr.Body.Bytes()))

	tokens := make(map[string]string)
	err = json.Unmarshal(rr.Body.Bytes(), &tokens)
	if err != nil {
		t.Errorf("cannot unmarshal response: %v\n", err)
	}
	assert.Equal(t, 201, rr.Code)
	assert.EqualValues(t, "this-is-the-NEW-access-token", tokens["access_token"])
	assert.EqualValues(t, "this-is-the-NEW-refresh-token", tokens["refresh_token"])
}
