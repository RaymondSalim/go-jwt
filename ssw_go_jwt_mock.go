// Code generated by mockery v2.26.1. DO NOT EDIT.

package ssw_go_jwt

import (
	jwt "github.com/golang-jwt/jwt/v5"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// MockSSWGoJWT is an autogenerated mock type for the SSWGoJWT type
type MockSSWGoJWT struct {
	mock.Mock
}

// GenerateTokens provides a mock function with given fields: claims
func (_m *MockSSWGoJWT) GenerateTokens(claims map[string]interface{}) (Tokens, error) {
	ret := _m.Called(claims)

	var r0 Tokens
	var r1 error
	if rf, ok := ret.Get(0).(func(map[string]interface{}) (Tokens, error)); ok {
		return rf(claims)
	}
	if rf, ok := ret.Get(0).(func(map[string]interface{}) Tokens); ok {
		r0 = rf(claims)
	} else {
		r0 = ret.Get(0).(Tokens)
	}

	if rf, ok := ret.Get(1).(func(map[string]interface{}) error); ok {
		r1 = rf(claims)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Init provides a mock function with given fields:
func (_m *MockSSWGoJWT) Init() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RenewToken provides a mock function with given fields: refreshToken, newAccessTokenClaims
func (_m *MockSSWGoJWT) RenewToken(refreshToken string, newAccessTokenClaims map[string]interface{}) (Tokens, error) {
	ret := _m.Called(refreshToken, newAccessTokenClaims)

	var r0 Tokens
	var r1 error
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}) (Tokens, error)); ok {
		return rf(refreshToken, newAccessTokenClaims)
	}
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}) Tokens); ok {
		r0 = rf(refreshToken, newAccessTokenClaims)
	} else {
		r0 = ret.Get(0).(Tokens)
	}

	if rf, ok := ret.Get(1).(func(string, map[string]interface{}) error); ok {
		r1 = rf(refreshToken, newAccessTokenClaims)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateAccessTokenWithClaims provides a mock function with given fields: signedToken, target
func (_m *MockSSWGoJWT) ValidateAccessTokenWithClaims(signedToken string, target *jwt.MapClaims) error {
	ret := _m.Called(signedToken, target)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *jwt.MapClaims) error); ok {
		r0 = rf(signedToken, target)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateToken provides a mock function with given fields: signedToken, tokenType
func (_m *MockSSWGoJWT) ValidateToken(signedToken string, tokenType TokenType) error {
	ret := _m.Called(signedToken, tokenType)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, TokenType) error); ok {
		r0 = rf(signedToken, tokenType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// generateToken provides a mock function with given fields: claims, expiresAt, signingKeyOrSecret
func (_m *MockSSWGoJWT) generateToken(claims map[string]interface{}, expiresAt time.Time, signingKeyOrSecret interface{}) (string, error) {
	ret := _m.Called(claims, expiresAt, signingKeyOrSecret)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(map[string]interface{}, time.Time, interface{}) (string, error)); ok {
		return rf(claims, expiresAt, signingKeyOrSecret)
	}
	if rf, ok := ret.Get(0).(func(map[string]interface{}, time.Time, interface{}) string); ok {
		r0 = rf(claims, expiresAt, signingKeyOrSecret)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(map[string]interface{}, time.Time, interface{}) error); ok {
		r1 = rf(claims, expiresAt, signingKeyOrSecret)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// getKeyFunc provides a mock function with given fields: tokenType
func (_m *MockSSWGoJWT) getKeyFunc(tokenType TokenType) func(*jwt.Token) (interface{}, error) {
	ret := _m.Called(tokenType)

	var r0 func(*jwt.Token) (interface{}, error)
	if rf, ok := ret.Get(0).(func(TokenType) func(*jwt.Token) (interface{}, error)); ok {
		r0 = rf(tokenType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func(*jwt.Token) (interface{}, error))
		}
	}

	return r0
}

// getSigningKeyOrSecret provides a mock function with given fields: tokenType
func (_m *MockSSWGoJWT) getSigningKeyOrSecret(tokenType TokenType) interface{} {
	ret := _m.Called(tokenType)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(TokenType) interface{}); ok {
		r0 = rf(tokenType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// validateConfig provides a mock function with given fields:
func (_m *MockSSWGoJWT) validateConfig() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockSSWGoJWT interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockSSWGoJWT creates a new instance of MockSSWGoJWT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockSSWGoJWT(t mockConstructorTestingTNewMockSSWGoJWT) *MockSSWGoJWT {
	mock := &MockSSWGoJWT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}