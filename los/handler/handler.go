package handler

import (
	"fmt"
	"net/http"

	"github.com/fikryfahrezy/adea/los/auth"
	"github.com/fikryfahrezy/adea/los/loan"
	"github.com/fikryfahrezy/adea/los/session"
	"github.com/fikryfahrezy/adea/los/setting"
)

type Handler struct {
	*session.Session
	*setting.SettingApp
	*auth.AuthApp
	*loan.LoanApp
}

func NewHandler(
	session *session.Session,
	settingApp *setting.SettingApp,
	authApp *auth.AuthApp,
	loanApp *loan.LoanApp,
) *Handler {
	return &Handler{
		Session:    session,
		SettingApp: settingApp,
		AuthApp:    authApp,
		LoanApp:    loanApp,
	}
}

func (h *Handler) ServeRestAPI() {
	mux := http.NewServeMux()

	mux.Handle("/tmp/", http.StripPrefix("/tmp/", http.FileServer(http.Dir("./tmp"))))

	mux.HandleFunc("/setting/generatejsondb", routeMWCompose(h.GanerateJsonDB, getRoute, h.authRoute(true)))
	mux.HandleFunc("/setting/loadjsondb", routeMWCompose(h.LoadJsonDB, postRoute, h.authRoute(true)))
	mux.HandleFunc("/setting/ziptmp", routeMWCompose(h.ZipTmp, getRoute, h.authRoute(true)))
	mux.HandleFunc("/setting/unziptmp", routeMWCompose(h.LoadZipTmp, postRoute, h.authRoute(true)))

	mux.HandleFunc("/auth/login", routeMWCompose(h.LoginPost(h.Session), postRoute))
	mux.HandleFunc("/auth/register", routeMWCompose(h.RegisterPost(h.Session), postRoute))

	mux.HandleFunc("/loan/getall", routeMWCompose(h.UserLoansGet, getRoute, h.authRoute(false)))
	mux.HandleFunc("/loan/get", routeMWCompose(h.UserLoanDetailGet, getRoute, h.authRoute(false)))
	mux.HandleFunc("/loan/create", routeMWCompose(h.CreateLoanPost, postRoute, h.authRoute(false)))
	mux.HandleFunc("/loan/update", routeMWCompose(h.UpdateLoanPut, putRoute, h.authRoute(false)))
	mux.HandleFunc("/loan/delete", routeMWCompose(h.UserLoanDelete, deleteRoute, h.authRoute(false)))

	mux.HandleFunc("/loan/getall/admin", routeMWCompose(h.LoansGet, getRoute, h.authRoute(true)))
	mux.HandleFunc("/loan/get/admin", routeMWCompose(h.LoanDetailGet, getRoute, h.authRoute(true)))
	mux.HandleFunc("/loan/proceedloan", routeMWCompose(h.ProceedLoanPatch, patchRoute, h.authRoute(true)))
	mux.HandleFunc("/loan/approveloan", routeMWCompose(h.ApproveLoanPatch, patchRoute, h.authRoute(true)))

	fmt.Println("You are ready to rock and roll!")
	http.ListenAndServe(":4000", mux)
}

func (h *Handler) authRoute(isPrivate bool) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("authorization")
			if auth == "" {
				http.Error(w, "forbidden", http.StatusForbidden)
				return
			}

			var ok bool
			var sess session.SessionObj
			if isPrivate {
				if sess, ok = h.Session.IsKeyPrivate(auth); !ok {
					http.Error(w, "forbidden private route", http.StatusForbidden)
					return
				}
			}

			if sess.Key == "" {
				if sess, ok = h.Session.Get(auth); !ok {
					http.Error(w, "forbidden", http.StatusForbidden)
					return
				}
			}

			if h.Session.IsExpired(sess) {
				http.Error(w, "forbidden session expired", http.StatusForbidden)
				return
			}

			next(w, r)
		}
	}
}

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

func putRoute(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.NotFound(w, r)
			return
		}
		next(w, r)
	}
}

func deleteRoute(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.NotFound(w, r)
			return
		}
		next(w, r)
	}
}

func patchRoute(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			http.NotFound(w, r)
			return
		}
		next(w, r)
	}
}

type RouteMiddleware func(next http.HandlerFunc) http.HandlerFunc

func routeMWCompose(h http.HandlerFunc, mw ...RouteMiddleware) http.HandlerFunc {
	for i := len(mw) - 1; i >= 0; i-- {
		h = mw[i](h)
	}

	return h
}
