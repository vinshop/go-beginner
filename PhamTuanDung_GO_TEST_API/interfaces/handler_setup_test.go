package interfaces

import "github.com/dungbk10t/test_api/utils/mock"

var (
	userApp   mock.UserAppInterface
	fakeAuth  mock.AuthInterface
	fakeToken mock.TokenInterface

	s  = NewUsers(&userApp, &fakeAuth, &fakeToken)        //We use all mocked data here
	au = NewAuthenticate(&userApp, &fakeAuth, &fakeToken) //We use all mocked data here

)
