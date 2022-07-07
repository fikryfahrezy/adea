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

	mux.HandleFunc("/loan/create", routeMWCompose(h.CreateLoanPost, postRoute, h.authRoute(false)))

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

func routeMWCompose(h http.HandlerFunc, mw ...RouteMiddleware) http.HandlerFunc {
	for i := len(mw) - 1; i >= 0; i-- {
		h = mw[i](h)
	}

	return h
}
