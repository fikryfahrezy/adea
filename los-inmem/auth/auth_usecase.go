package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/fikryfahrezy/adea/los/model"
	"github.com/fikryfahrezy/adea/los/resp"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrAuthPwNotMatch = errors.New("authentication password not match")
	ErrUsernameExist  = errors.New("username already exist")
)

type (
	RegisterIn struct {
		IsOfficer bool   `json:"is_officer"`
		Username  string `json:"username"`
		Password  string `json:"password"`
	}
	RegisterRes struct {
		IsOfficer bool   `json:"is_officer"`
		Id        string `json:"id"`
	}
	RegisterOut struct {
		resp.Response
		Res RegisterRes
	}
)

func (a *AuthApp) Register(ctx context.Context, in RegisterIn) (out RegisterOut) {
	out.Response = resp.NewResponse(http.StatusCreated, "", nil)

	if err := validateRegister(in); err != nil {
		out.Response = resp.NewResponse(http.StatusUnprocessableEntity, "", err)
		return
	}

	_, err := a.repository.GetUserByUsername(ctx, in.Username)
	if err == nil {
		out.Response = resp.NewResponse(http.StatusBadRequest, "", ErrUsernameExist)
		return
	}
	if !errors.Is(err, ErrUserNotFound) {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	newUser := model.User{
		Username:  in.Username,
		IsOfficer: in.IsOfficer,
		Password:  string(hashed),
	}

	if newUser, err = a.repository.InsertUser(ctx, newUser); err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	out.Res = RegisterRes{
		IsOfficer: newUser.IsOfficer,
		Id:        newUser.Id,
	}

	return
}

type (
	LoginIn struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	LoginRes struct {
		IsOfficer bool   `json:"is_officer"`
		Id        string `json:"id"`
	}
	LoginOut struct {
		resp.Response
		Res LoginRes
	}
)

func (a *AuthApp) Login(ctx context.Context, in LoginIn) (out LoginOut) {
	out.Response = resp.NewResponse(http.StatusOK, "", nil)

	if err := validateLogin(in); err != nil {
		out.Response = resp.NewResponse(http.StatusUnprocessableEntity, "", err)
		return
	}

	user, err := a.repository.GetUserByUsername(ctx, in.Username)
	if errors.Is(err, ErrUserNotFound) {
		out.Response = resp.NewResponse(http.StatusNotFound, "", err)
		return
	}
	if err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		out.Response = resp.NewResponse(http.StatusBadRequest, "", ErrAuthPwNotMatch)
		return
	}
	if err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	out.Res = LoginRes{
		IsOfficer: user.IsOfficer,
		Id:        user.Id,
	}

	return
}
