package main

import (
	"github.com/fikryfahrezy/adea/los-inmen/auth"
	"github.com/fikryfahrezy/adea/los-inmen/data"
	"github.com/fikryfahrezy/adea/los-inmen/file"
	"github.com/fikryfahrezy/adea/los-inmen/handler"
	"github.com/fikryfahrezy/adea/los-inmen/loan"
	"github.com/fikryfahrezy/adea/los-inmen/session"
	"github.com/fikryfahrezy/adea/los-inmen/setting"
)

func main() {
	dbJson := data.NewJson("")
	file := file.New()
	session := session.New()

	authRepo := auth.NewRepository(dbJson)
	loanRepo := loan.NewRepository(dbJson)

	setting := setting.NewSetting(file, dbJson)
	authApp := auth.NewApp(authRepo)
	loanApp := loan.NewApp(file.Save, loanRepo)

	handler := handler.NewHandler(session, setting, authApp, loanApp)

	handler.ServeRestAPI()
}
