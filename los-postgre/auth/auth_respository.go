package auth

import (
	"context"
	"encoding/hex"
	"errors"
	"time"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"
	"github.com/fikryfahrezy/adea/los-postgre/model"
	"github.com/jackc/pgx/v4"
)

var (
	ErrDuplicateContraint = errors.New("some constraint are duplicate")
	ErrUserNotFound       = errors.New("user not found")
)

type Repository struct {
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) InsertUser(ctx context.Context, user model.User) (model.User, error) {
	t := time.Now()
	id := hex.EncodeToString([]byte(user.Username))

	user.Id = id
	user.CreatedDate = t

	err := crdbpgx.ExecuteTx(context.Background(), r.db, pgx.TxOptions{}, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx,
			`INSERT INTO users (id, username, password, is_officer, created_date)
			VALUES ($1, $2, $3, $4, $5)`,
			user.Id, user.Username, user.Password, user.IsOfficer, user.CreatedDate,
		); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *Repository) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	id := hex.EncodeToString([]byte(username))

	var user model.User
	err := crdbpgx.ExecuteTx(context.Background(), r.db, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return tx.QueryRow(ctx,
			`SELECT id, username, password, is_officer, created_date
			FROM users WHERE id = $1`,
			id,
		).Scan(&user.Id, &user.Username, &user.Password, &user.IsOfficer, &user.CreatedDate)
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return model.User{}, ErrUserNotFound
	}
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
