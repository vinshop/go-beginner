package mock

import (
	"github.com/dungbk10t/test_api/domain/entity"
	"github.com/dungbk10t/test_api/infrastructure/auth"
	"net/http"
)

//UserAppInterface is a mock user app interface
type UserAppInterface struct {
	SaveUserFn                  func(*entity.User) (*entity.User, map[string]string)
	UpdateInfoUserFn            func(uint64, *entity.User) (*entity.User, error)
	UpdatePassWordUserFn        func(uint64, *entity.User) (*entity.User, error)
	GetUsersFn                  func() ([]entity.User, error)
	GetUserFn                   func(uint64) (*entity.User, error)
	DeleteUserFn                func(uint64) error
	GetUserByEmailAndPasswordFn func(*entity.User) (*entity.User, map[string]string)
}

//SaveUser calls the SaveUserFn
func (u *UserAppInterface) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	return u.SaveUserFn(user)
}

func (u *UserAppInterface) UpdateInfoUser(userId uint64, user *entity.User) (*entity.User, error) {
	return u.UpdateInfoUserFn(userId, user)
}

func (u *UserAppInterface) UpdatePassWordUser(userId uint64, user *entity.User) (*entity.User, error) {
	return u.UpdatePassWordUserFn(userId, user)
}

//GetUsersFn calls the GetUsers
func (u *UserAppInterface) GetUsers() ([]entity.User, error) {
	return u.GetUsersFn()
}

//GetUserFn calls the GetUser
func (u *UserAppInterface) GetUser(userId uint64) (*entity.User, error) {
	return u.GetUserFn(userId)
}

func (u *UserAppInterface) DeleteUser(userId uint64) error {
	return u.DeleteUserFn(userId)
}

//GetUserByEmailAndPasswordFn calls the GetUserByEmailAndPassword
func (u *UserAppInterface) GetUserByEmailAndPassword(user *entity.User) (*entity.User, map[string]string) {
	return u.GetUserByEmailAndPasswordFn(user)
}

//AuthInterface is a mock auth interface
type AuthInterface struct {
	CreateAuthFn    func(uint64, *auth.TokenDetails) error
	FetchAuthFn     func(string) (uint64, error)
	DeleteRefreshFn func(string) error
	DeleteTokensFn  func(*auth.AccessDetails) error
}

func (f *AuthInterface) DeleteRefresh(refreshUuid string) error {
	return f.DeleteRefreshFn(refreshUuid)
}
func (f *AuthInterface) DeleteTokens(authD *auth.AccessDetails) error {
	return f.DeleteTokensFn(authD)
}
func (f *AuthInterface) FetchAuth(uuid string) (uint64, error) {
	return f.FetchAuthFn(uuid)
}
func (f *AuthInterface) CreateAuth(userId uint64, authD *auth.TokenDetails) error {
	return f.CreateAuthFn(userId, authD)
}

//TokenInterface is a mock token interface
type TokenInterface struct {
	CreateTokenFn          func(userId uint64) (*auth.TokenDetails, error)
	ExtractTokenMetadataFn func(*http.Request) (*auth.AccessDetails, error)
}

func (f *TokenInterface) CreateToken(userid uint64) (*auth.TokenDetails, error) {
	return f.CreateTokenFn(userid)
}
func (f *TokenInterface) ExtractTokenMetadata(r *http.Request) (*auth.AccessDetails, error) {
	return f.ExtractTokenMetadataFn(r)
}

//type UploadFileInterface struct {
//	UploadFileFn func(file *multipart.FileHeader) (string, error)
//}
//
//func (up *UploadFileInterface) UploadFile(file *multipart.FileHeader) (string, error) {
//	return up.UploadFileFn(file)
//}
