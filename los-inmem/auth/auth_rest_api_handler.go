package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/fikryfahrezy/adea/los/resp"
	"github.com/fikryfahrezy/adea/los/session"
)

func (a *AuthApp) RegisterPost(session *session.Session) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var in RegisterIn
		err := json.NewDecoder(r.Body).Decode(&in)
		if err != nil {
			resp.NewResponse(http.StatusInternalServerError, "", err).HttpJSON(w, nil)
			return
		}

		out := a.Register(r.Context(), in)
		if out.Error == nil {
			session.Set(out.Res.Id, out.Res.Id, out.Res.IsOfficer, time.Now().Add(time.Hour).Unix())
		}

		out.HttpJSON(w, resp.NewHttpBody(out.Res))
	}
}

func (a *AuthApp) LoginPost(session *session.Session) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var in LoginIn
		err := json.NewDecoder(r.Body).Decode(&in)
		if err != nil {
			resp.NewResponse(http.StatusInternalServerError, "", err).HttpJSON(w, nil)
			return
		}

		out := a.Login(r.Context(), in)
		if out.Error == nil {
			session.Set(out.Res.Id, out.Res.Id, out.Res.IsOfficer, time.Now().Add(time.Hour).Unix())
		}

		out.HttpJSON(w, resp.NewHttpBody(out.Res))
	}
}
