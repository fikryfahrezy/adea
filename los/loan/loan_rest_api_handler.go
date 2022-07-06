package loan

import (
	"net/http"

	"github.com/fikryfahrezy/adea/los/resp"
)

func (a *LoanApp) CreateLoanPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(1024); err != nil {
		resp.NewResponse(http.StatusInternalServerError, "", err).HttpJSON(w, nil)
		return
	}

	_, header, err := r.FormFile("file")
	if err != nil {
		resp.NewResponse(http.StatusInternalServerError, "", err).HttpJSON(w, nil)
		return
	}

	in := CreateLoanIn{
		File: header,
	}

	out := a.CreateLoan(r.Context(), in)
	out.HttpJSON(w, resp.NewHttpBody(out.Res))
}
