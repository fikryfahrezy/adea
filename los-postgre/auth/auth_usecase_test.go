package auth_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/fikryfahrezy/adea/los-postgre/auth"
	"github.com/fikryfahrezy/adea/los-postgre/model"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest"
	"golang.org/x/crypto/bcrypt"
)

var (
	dbPg     *pgx.Conn
	authRepo *auth.Repository
	authApp  *auth.AuthApp
)

func loadTables(conn *pgx.Conn) error {
	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	f, err := os.ReadFile("../docs/db.sql")
	if err != nil {
		return err
	}

	_, err = tx.Exec(context.Background(),
		string(f),
	)
	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func clearDb() error {
	tx, err := dbPg.Begin(context.Background())
	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	// This should be in order of which table truncate first before the other
	queries := []string{
		`TRUNCATE users CASCADE`,
	}

	for _, v := range queries {
		_, err = tx.Exec(context.Background(),
			v,
		)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func TestMain(m *testing.M) {
	var err error
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{Repository: "cockroachdb/cockroach", Tag: "v21.2.13", Cmd: []string{"start-single-node", "--insecure"}})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	databaseUrl := fmt.Sprintf("postgresql://root@localhost:%s/defaultdb?sslmode=disable", resource.GetPort("26257/tcp"))

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 120 * time.Second
	if err = pool.Retry(func() error {
		dbConfig, err := pgx.ParseConfig(databaseUrl)
		if err != nil {
			return err
		}

		dbPg, err = pgx.ConnectConfig(context.Background(), dbConfig)
		if err != nil {
			return err
		}

		return dbPg.Ping(context.Background())
	}); err != nil {
		log.Fatalf("Could not connect to cockroach container: %s", err)
	}

	authRepo = auth.NewRepository(dbPg)
	authApp = auth.NewApp(authRepo)

	loadTables(dbPg)

	code := m.Run()

	// When you're done, kill and remove the container
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestLogin(t *testing.T) {
	err := clearDb()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	hashed, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	authRepo.InsertUser(ctx, model.User{
		IsOfficer: true,
		Username:  "existusername",
		Password:  string(hashed),
	})

	testCases := []struct {
		expect int
		name   string
		input  auth.LoginIn
	}{
		{
			expect: http.StatusOK,
			name:   "Login successfully",
			input: auth.LoginIn{
				Username: "existusername",
				Password: "password",
			},
		},
		{
			expect: http.StatusBadRequest,
			name:   "Login fail, password not match",
			input: auth.LoginIn{
				Username: "existusername",
				Password: "passwordxxxxx",
			},
		},
		{
			expect: http.StatusNotFound,
			name:   "Login fail, user not found",
			input: auth.LoginIn{
				Username: "nonexistusername",
				Password: "password",
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Login fail, no input provided",
			input:  auth.LoginIn{},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Login fail, no username provided",
			input: auth.LoginIn{
				Username: "existusername",
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Login fail, no password provided",
			input: auth.LoginIn{
				Password: "password",
			},
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			out := authApp.Login(ctx, c.input)

			if out.StatusCode != c.expect {
				t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, c.expect, out.Error)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	err := clearDb()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	hashed, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	authRepo.InsertUser(ctx, model.User{
		IsOfficer: true,
		Username:  "existusername",
		Password:  string(hashed),
	})

	testCases := []struct {
		expect int
		name   string
		input  auth.RegisterIn
	}{
		{
			expect: http.StatusCreated,
			name:   "Register as officer successfully",
			input: auth.RegisterIn{
				IsOfficer: true,
				Username:  "nonexsistofficerusername",
				Password:  "password",
			},
		},
		{
			expect: http.StatusCreated,
			name:   "Register as non officer successfully",
			input: auth.RegisterIn{
				IsOfficer: false,
				Username:  "nonexsistnonofficerusername",
				Password:  "password",
			},
		},
		{
			expect: http.StatusBadRequest,
			name:   "Register fail, username exist",
			input: auth.RegisterIn{
				IsOfficer: false,
				Username:  "existusername",
				Password:  "passwordxxxxx",
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Login fail, no input provided",
			input:  auth.RegisterIn{},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Login fail, no username provided",
			input: auth.RegisterIn{
				Username: "existusername",
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Login fail, no password provided",
			input: auth.RegisterIn{
				Password: "password",
			},
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			out := authApp.Register(ctx, c.input)

			if out.StatusCode != c.expect {
				t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, c.expect, out.Error)
			}
		})
	}
}
