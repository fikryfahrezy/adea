package main

import (
	"context"
	"log"
	"os"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"
	"github.com/fikryfahrezy/adea/los-postgre/auth"
	"github.com/fikryfahrezy/adea/los-postgre/file"
	"github.com/fikryfahrezy/adea/los-postgre/handler"
	"github.com/fikryfahrezy/adea/los-postgre/loan"
	"github.com/fikryfahrezy/adea/los-postgre/session"
	"github.com/fikryfahrezy/adea/los-postgre/setting"
	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return conn.Ping(context.Background())
	})
	if err != nil {
		log.Fatal(err)
	}

	file := file.New()
	session := session.New()

	authRepo := auth.NewRepository(conn)
	loanRepo := loan.NewRepository(conn)

	setting := setting.NewSetting(file)
	authApp := auth.NewApp(authRepo)
	loanApp := loan.NewApp(file.Save, loanRepo)

	handler := handler.NewHandler(session, setting, authApp, loanApp)

	handler.ServeRestAPI()
}
