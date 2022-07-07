package loan

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/fikryfahrezy/adea/los/data"
	"github.com/fikryfahrezy/adea/los/model"
)

type Status struct {
	slug string
}

func (r Status) String() string {
	return r.slug
}

var (
	Unknown = Status{""}
	Wait    = Status{"wait"}
	Process = Status{"process"}
	Reject  = Status{"reject"}
	Approve = Status{"approve"}
)

func FromString(s string) (Status, error) {
	switch s {
	case Wait.slug:
		return Wait, nil
	case Process.slug:
		return Process, nil
	case Reject.slug:
		return Reject, nil
	case Approve.slug:
		return Approve, nil
	}

	return Unknown, errors.New("unknown status: " + s)
}

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrUserLoanNotFound = errors.New("user loan not found")
)

type Repository struct {
	db *data.JsonFile
}

func NewRepository(db *data.JsonFile) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetUser(ctx context.Context, userId string) (model.User, error) {
	r.db.Lock()
	defer r.db.Unlock()

	user, ok := r.db.DbUser[userId]
	if !ok {
		return model.User{}, ErrUserNotFound
	}

	return user, nil
}

func (r *Repository) GetUserLoans(ctx context.Context, userId string) (map[string]model.LoanApplication, error) {
	r.db.Lock()
	defer r.db.Unlock()

	userLoans := make(map[string]model.LoanApplication)
	for k, v := range r.db.DbLoan {
		if v.UserId == userId {
			userLoans[k] = v
		}
	}

	return userLoans, nil
}

func (r *Repository) GetUserLoan(ctx context.Context, loanId, userId string) (model.LoanApplication, error) {
	r.db.Lock()
	defer r.db.Unlock()

	for _, v := range r.db.DbLoan {
		if v.UserId == userId && v.Id == loanId {
			return v, nil
		}
	}

	return model.LoanApplication{}, ErrUserLoanNotFound
}

func (r *Repository) GetLoans(ctx context.Context) (map[string]model.LoanApplication, error) {
	r.db.Lock()
	defer r.db.Unlock()

	return r.db.DbLoan, nil
}

func (r *Repository) GetLoan(ctx context.Context, loanId string) (model.LoanApplication, error) {
	r.db.Lock()
	defer r.db.Unlock()

	for _, v := range r.db.DbLoan {
		if v.Id == loanId {
			return v, nil
		}
	}

	return model.LoanApplication{}, ErrUserLoanNotFound
}

func (r *Repository) InsertLoan(ctx context.Context, loan model.LoanApplication) (model.LoanApplication, error) {
	t := time.Now()
	tn := t.UnixNano()
	ra := rand.New(rand.NewSource(tn))
	id := hex.EncodeToString([]byte(fmt.Sprintf("%d-%d", tn, ra)))

	loan.Id = id
	loan.CreatedDate = t
	loan.UpdatedDate = t
	loan.Status = Wait.String()

	r.db.Lock()
	defer r.db.Unlock()
	r.db.DbLoan[id] = loan

	return loan, nil
}

func (r *Repository) RemoveLoan(ctx context.Context, loanId string) error {
	r.db.Lock()
	defer r.db.Unlock()

	if _, ok := r.db.DbLoan[loanId]; ok {
		delete(r.db.DbLoan, loanId)
	}

	return nil
}

func (r *Repository) UpdateLoan(ctx context.Context, loanId string, loan model.LoanApplication) error {
	r.db.Lock()
	defer r.db.Unlock()

	r.db.DbLoan[loanId] = loan

	return nil
}
