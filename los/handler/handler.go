package handler

import (
	"fmt"
	"net/http"

	"github.com/fikryfahrezy/adea/los/auth"
	"github.com/fikryfahrezy/adea/los/loan"
	"github.com/fikryfahrezy/adea/los/setting"
)

type Handler struct {
	*setting.SettingApp
	*auth.AuthApp
	*loan.LoanApp
}

func NewHandler(
	settingApp *setting.SettingApp,
	authApp *auth.AuthApp,
	loanApp *loan.LoanApp,
) *Handler {
	return &Handler{
		SettingApp: settingApp,
		AuthApp:    authApp,
		LoanApp:    loanApp,
	}
}

func (h *Handler) ServeRestAPI() {
	mux := http.NewServeMux()

	mux.Handle("/tmp/", http.StripPrefix("/tmp/", http.FileServer(http.Dir("./tmp"))))

	mux.HandleFunc("/setting/generatejsondb", routeMWCompose(h.GanerateJsonDB, getRoute))

	mux.HandleFunc("/setting/loadjsondb", routeMWCompose(h.LoadJsonDB, postRoute))
	mux.HandleFunc("/setting/ziptmp", routeMWCompose(h.ZipTmp, getRoute))
	mux.HandleFunc("/setting/unziptmp", routeMWCompose(h.LoadZipTmp, postRoute))

	mux.HandleFunc("/auth/login", routeMWCompose(h.LoginPost, postRoute))
	mux.HandleFunc("/auth/register", routeMWCompose(h.RegisterPost, postRoute))

	mux.HandleFunc("/loan/create", routeMWCompose(h.CreateLoanPost, postRoute))

	fmt.Println("You are ready to rock and roll!")
	http.ListenAndServe(":4000", mux)
}

type RouteMiddleware func(next http.HandlerFunc) http.HandlerFunc

func getRoute(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.NotFound(w, r)
			return
		}
		next(w, r)
	}
}

func postRoute(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.NotFound(w, r)
			return
		}
		next(w, r)
	}
}

func routeMWCompose(h http.HandlerFunc, mw ...RouteMiddleware) http.HandlerFunc {
	for i := len(mw) - 1; i >= 0; i-- {
		h = mw[i](h)
	}

	return h
}
