package auth_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/fikryfahrezy/adea/los-postgre/auth"
	"github.com/fikryfahrezy/adea/los-postgre/data"
	"github.com/fikryfahrezy/adea/los-postgre/model"
	"golang.org/x/crypto/bcrypt"
)

var (
	dbJson   = data.NewJson("")
	authRepo = auth.NewRepository(dbJson)
	authApp  = auth.NewApp(authRepo)
)

func clearDb() {
	dbJson.DbUser = make(map[string]model.User)
}

func TestLogin(t *testing.T) {
	clearDb()

	ctx := context.Background()
	hashed, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	authRepo.InsertUser(ctx, model.User{
		IsOfficer: true,
		Username:  "existusername",
		Password:  string(hashed),
	})

	testCases := []struct {
		expect int
		name   string
		input  auth.LoginIn
	}{
		{
			expect: http.StatusOK,
			name:   "Login successfully",
			input: auth.LoginIn{
				Username: "existusername",
				Password: "password",
			},
		},
		{
			expect: http.StatusBadRequest,
			name:   "Login fail, password not match",
			input: auth.LoginIn{
				Username: "existusername",
				Password: "passwordxxxxx",
			},
		},
		{
			expect: http.StatusNotFound,
			name:   "Login fail, user not found",
			input: auth.LoginIn{
				Username: "nonexistusername",
				Password: "password",
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Login fail, no input provided",
			input:  auth.LoginIn{},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Login fail, no username provided",
			input: auth.LoginIn{
				Username: "existusername",
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Login fail, no password provided",
			input: auth.LoginIn{
				Password: "password",
			},
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			out := authApp.Login(ctx, c.input)

			if out.StatusCode != c.expect {
				t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, c.expect, out.Error)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	clearDb()
	ctx := context.Background()

	hashed, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	authRepo.InsertUser(ctx, model.User{
		IsOfficer: true,
		Username:  "existusername",
		Password:  string(hashed),
	})

	testCases := []struct {
		expect int
		name   string
		input  auth.RegisterIn
	}{
		{
			expect: http.StatusCreated,
			name:   "Register as officer successfully",
			input: auth.RegisterIn{
				IsOfficer: true,
				Username:  "nonexsistofficerusername",
				Password:  "password",
			},
		},
		{
			expect: http.StatusCreated,
			name:   "Register as non officer successfully",
			input: auth.RegisterIn{
				IsOfficer: false,
				Username:  "nonexsistnonofficerusername",
				Password:  "password",
			},
		},
		{
			expect: http.StatusBadRequest,
			name:   "Register fail, username exist",
			input: auth.RegisterIn{
				IsOfficer: false,
				Username:  "existusername",
				Password:  "passwordxxxxx",
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Login fail, no input provided",
			input:  auth.RegisterIn{},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Login fail, no username provided",
			input: auth.RegisterIn{
				Username: "existusername",
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Login fail, no password provided",
			input: auth.RegisterIn{
				Password: "password",
			},
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			out := authApp.Register(ctx, c.input)

			if out.StatusCode != c.expect {
				t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, c.expect, out.Error)
			}
		})
	}
}
