package main

import (
	"github.com/fikryfahrezy/adea/los/auth"
	"github.com/fikryfahrezy/adea/los/data"
	"github.com/fikryfahrezy/adea/los/file"
	"github.com/fikryfahrezy/adea/los/handler"
	"github.com/fikryfahrezy/adea/los/loan"
	"github.com/fikryfahrezy/adea/los/setting"
)

func main() {
	dbJson := data.NewJson("")
	file := file.New()

	authRepo := auth.NewRepository(dbJson)
	loanRepo := loan.NewRepository(dbJson)

	setting := setting.NewSetting(dbJson)
	auth := auth.NewApp(authRepo)
	loan := loan.NewApp(file.Save, loanRepo)

	handler := handler.NewHandler(setting, auth, loan)

	handler.ServeRestAPI()
}
