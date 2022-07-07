package main

import (
	"github.com/fikryfahrezy/adea/los-postgre/auth"
	"github.com/fikryfahrezy/adea/los-postgre/data"
	"github.com/fikryfahrezy/adea/los-postgre/file"
	"github.com/fikryfahrezy/adea/los-postgre/handler"
	"github.com/fikryfahrezy/adea/los-postgre/loan"
	"github.com/fikryfahrezy/adea/los-postgre/session"
	"github.com/fikryfahrezy/adea/los-postgre/setting"
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
