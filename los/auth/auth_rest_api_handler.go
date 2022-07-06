package auth

import (
	"encoding/json"
	"net/http"

	"github.com/fikryfahrezy/adea/los/resp"
)

func (a *AuthApp) RegisterPost(w http.ResponseWriter, r *http.Request) {
	var in RegisterIn
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		resp.NewResponse(http.StatusInternalServerError, "", err).HttpJSON(w, nil)
		return
	}

	out := a.Register(r.Context(), in)
	out.HttpJSON(w, resp.NewHttpBody(out.Res))
}

func (a *AuthApp) LoginPost(w http.ResponseWriter, r *http.Request) {
	var in RegisterIn
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		resp.NewResponse(http.StatusInternalServerError, "", err).HttpJSON(w, nil)
		return
	}

	out := a.Register(r.Context(), in)
	out.HttpJSON(w, resp.NewHttpBody(out.Res))
}
