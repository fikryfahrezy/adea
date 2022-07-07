package auth

import (
	"context"
	"encoding/hex"
	"errors"
	"time"

	"github.com/fikryfahrezy/adea/los-postgre/data"
	"github.com/fikryfahrezy/adea/los-postgre/model"
)

var (
	ErrDuplicateContraint = errors.New("some constraint are duplicate")
	ErrUserNotFound       = errors.New("user not found")
)

type Repository struct {
	db *data.JsonFile
}

func NewRepository(db *data.JsonFile) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) InsertUser(ctx context.Context, user model.User) (model.User, error) {
	t := time.Now()
	id := hex.EncodeToString([]byte(user.Username))

	user.Id = id
	user.CreatedDate = t

	r.db.Lock()
	defer r.db.Unlock()
	r.db.DbUser[id] = user

	return user, nil
}

func (r *Repository) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	id := hex.EncodeToString([]byte(username))

	r.db.Lock()
	defer r.db.Unlock()
	user, ok := r.db.DbUser[id]
	if !ok {
		return model.User{}, ErrUserNotFound
	}

	return user, nil
}
