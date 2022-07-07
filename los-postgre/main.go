package main

import (
	"github.com/fikryfahrezy/adea/los/auth"
	"github.com/fikryfahrezy/adea/los/data"
	"github.com/fikryfahrezy/adea/los/file"
	"github.com/fikryfahrezy/adea/los/handler"
	"github.com/fikryfahrezy/adea/los/loan"
	"github.com/fikryfahrezy/adea/los/session"
	"github.com/fikryfahrezy/adea/los/setting"
)

func main() {
	dbJson := data.NewJson("")
	file := file.New()
	session := session.New()

	authRepo := auth.NewRepository(dbJson)
	loanRepo := loan.NewRepository(dbJson)

	setting := setting.NewSetting(dbJson)
	authApp := auth.NewApp(authRepo)
	loanApp := loan.NewApp(file.Save, loanRepo)

	handler := handler.NewHandler(session, setting, authApp, loanApp)

	handler.ServeRestAPI()
}
