package model_test

import (
	"test_task/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

var TestUser = &model.User{
	Email:    "user@example.com",
	Passowrd: "password",
}

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return TestUser
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *model.User {
				u := TestUser
				u.Email = ""

				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *model.User {
				u := TestUser
				u.Passowrd = ""

				return u
			},
			isValid: false,
		},
		{
			name: "with encrypted password",
			u: func() *model.User {
				u := TestUser
				u.Passowrd = ""
				u.EncryptedPassword = "encrypted_password"

				return u
			},
			isValid: true,
		},
		{
			name: "invalid email",
			u: func() *model.User {
				u := TestUser
				u.Email = "invalid email"

				return u
			},
			isValid: false,
		},
		{
			name: "invalid password",
			u: func() *model.User {
				u := TestUser
				u.Passowrd = "short"

				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

func TestUser_BeforeUserCreation(t *testing.T) {
	u := TestUser

	assert.NoError(t, u.BeforeUserCreation())
	assert.NotEmpty(t, u.EncryptedPassword)
}
