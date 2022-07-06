package auth

import (
	"context"

	"github.com/fikryfahrezy/adea/los/data"
	"github.com/fikryfahrezy/adea/los/model"
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
	r.db.Lock()
	defer r.db.Unlock()

	r.db.DbUser["test"] = user
	return model.User{}, nil
}

func (r *Repository) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	return model.User{}, nil
}
