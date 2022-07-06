package loan

import (
	"context"
	"mime/multipart"
	"net/http"

	"github.com/fikryfahrezy/adea/los/resp"
	"golang.org/x/crypto/bcrypt"
)

type (
	CreateLoanIn struct {
		File *multipart.FileHeader
	}
	CreateLoanRes struct{}
	CreateLoanOut struct {
		resp.Response
		Res CreateLoanRes
	}
)

func (a *LoanApp) CreateLoan(ctx context.Context, in CreateLoanIn) (out CreateLoanOut) {
	if err := validateRegister(in); err != nil {
		out.Response = resp.NewResponse(http.StatusUnprocessableEntity, "", err)
		return
	}
	out.Response = resp.NewResponse(http.StatusCreated, "", nil)

	var fileUrl string
	if in.File != nil {
		file, err := in.File.Open()
		if err != nil {
			out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
			return
		}

		fileUrl, err = a.saveFile(in.File.Filename, file)
		if err != nil {
			out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
			return
		}
	}

	_ = fileUrl

	bcrypt.GenerateFromPassword([]byte("hifjdsfds"), bcrypt.DefaultCost)

	return
}
