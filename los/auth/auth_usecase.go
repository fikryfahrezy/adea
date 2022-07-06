package auth

import (
	"context"
	"net/http"

	"github.com/fikryfahrezy/adea/los/resp"
	"golang.org/x/crypto/bcrypt"
)

type (
	RegisterIn  struct{}
	RegisterRes struct{}
	RegisterOut struct {
		resp.Response
		Res RegisterRes
	}
)

func (a *AuthApp) Register(ctx context.Context, in RegisterIn) (out RegisterOut) {
	if err := validateRegister(in); err != nil {
		out.Response = resp.NewResponse(http.StatusUnprocessableEntity, "", err)
		return
	}
	out.Response = resp.NewResponse(http.StatusCreated, "", nil)

	bcrypt.GenerateFromPassword([]byte("hifjdsfds"), bcrypt.DefaultCost)

	return
}

type (
	LoginIn  struct{}
	LoginRes struct{}
	LoginOut struct {
		resp.Response
		Res LoginRes
	}
)

func (a *AuthApp) Login(ctx context.Context, in LoginIn) (out LoginOut) {
	if err := validateLogin(in); err != nil {
		out.Response = resp.NewResponse(http.StatusUnprocessableEntity, "", err)
		return
	}
	out.Response = resp.NewResponse(http.StatusCreated, "", nil)

	bcrypt.CompareHashAndPassword([]byte("hifjdsfds"), []byte("hifjdsfds"))

	return
}
