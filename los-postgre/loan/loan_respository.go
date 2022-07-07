package loan

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"
	"github.com/fikryfahrezy/adea/los-postgre/model"
	"github.com/jackc/pgx/v4"
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
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetUser(ctx context.Context, userId string) (model.User, error) {
	var user model.User
	err := crdbpgx.ExecuteTx(context.Background(), r.db, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return tx.QueryRow(ctx,
			`SELECT id, username, password, is_officer, created_date
		FROM users WHERE id = $1`,
			userId,
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

func (r *Repository) GetUserLoans(ctx context.Context, userId string) ([]model.LoanApplication, error) {
	var userLoans []model.LoanApplication
	err := crdbpgx.ExecuteTx(context.Background(), r.db, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := tx.Query(ctx,
			`SELECT 
				id,
				user_id,
				officer_id,
				full_name,
				birth_date,
				full_address,
				phone,
				id_card_url,
				other_business,
				status,
				is_private_field,
				exp_in_year,
				active_field_number,
				sow_seeds_per_cycle,
				needed_fertilizier_per_cycle_in_kg,
				estimated_yield_in_kg,
				estimated_price_of_harvest_per_kg,
				harvest_cycle_in_months,
				loan_application_in_idr,
				business_income_per_month_in_idr,
				business_outcome_per_month_in_idr,
				created_date,
				updated_date
			FROM loan_applications
			WHERE user_id = $1`,
			userId,
		)
		if err != nil {
			return err
		}

		for rows.Next() {
			var userLoan model.LoanApplication
			if err := rows.Scan(
				&userLoan.Id,
				&userLoan.UserId,
				&userLoan.OfficerId,
				&userLoan.FullName,
				&userLoan.BirthDate,
				&userLoan.FullAddress,
				&userLoan.Phone,
				&userLoan.IdCardUrl,
				&userLoan.OtherBusiness,
				&userLoan.Status,
				&userLoan.IsPrivateField,
				&userLoan.ExpInYear,
				&userLoan.ActiveFieldNumber,
				&userLoan.SowSeedsPerCycle,
				&userLoan.NeededFertilizerPerCycleInKg,
				&userLoan.EstimatedYieldInKg,
				&userLoan.EstimatedPriceOfHarvestPerKg,
				&userLoan.HarvestCycleInMonths,
				&userLoan.LoanApplicationInIdr,
				&userLoan.BusinessIncomePerMonthInIdr,
				&userLoan.BusinessOutcomePerMonthInIdr,
				&userLoan.CreatedDate,
				&userLoan.UpdatedDate,
			); err != nil {
				return err
			}
			userLoans = append(userLoans, userLoan)
		}

		return nil
	})
	if err != nil {
		return []model.LoanApplication{}, err
	}

	return userLoans, nil
}

func (r *Repository) GetUserLoan(ctx context.Context, loanId, userId string) (model.LoanApplication, error) {
	var userLoan model.LoanApplication
	err := crdbpgx.ExecuteTx(context.Background(), r.db, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return tx.QueryRow(ctx,
			`SELECT 
				id,
				user_id,
				officer_id,
				full_name,
				birth_date,
				full_address,
				phone,
				id_card_url,
				other_business,
				status,
				is_private_field,
				exp_in_year,
				active_field_number,
				sow_seeds_per_cycle,
				needed_fertilizier_per_cycle_in_kg,
				estimated_yield_in_kg,
				estimated_price_of_harvest_per_kg,
				harvest_cycle_in_months,
				loan_application_in_idr,
				business_income_per_month_in_idr,
				business_outcome_per_month_in_idr,
				created_date,
				updated_date
			FROM loan_applications
			WHERE id = $1 AND user_id = $2`,
			loanId,
			userId,
		).Scan(
			&userLoan.Id,
			&userLoan.UserId,
			&userLoan.OfficerId,
			&userLoan.FullName,
			&userLoan.BirthDate,
			&userLoan.FullAddress,
			&userLoan.Phone,
			&userLoan.IdCardUrl,
			&userLoan.OtherBusiness,
			&userLoan.Status,
			&userLoan.IsPrivateField,
			&userLoan.ExpInYear,
			&userLoan.ActiveFieldNumber,
			&userLoan.SowSeedsPerCycle,
			&userLoan.NeededFertilizerPerCycleInKg,
			&userLoan.EstimatedYieldInKg,
			&userLoan.EstimatedPriceOfHarvestPerKg,
			&userLoan.HarvestCycleInMonths,
			&userLoan.LoanApplicationInIdr,
			&userLoan.BusinessIncomePerMonthInIdr,
			&userLoan.BusinessOutcomePerMonthInIdr,
			&userLoan.CreatedDate,
			&userLoan.UpdatedDate,
		)
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return model.LoanApplication{}, ErrUserLoanNotFound
	}
	if err != nil {
		return model.LoanApplication{}, err
	}

	return userLoan, nil
}

func (r *Repository) GetLoans(ctx context.Context) ([]model.LoanApplication, error) {
	var loans []model.LoanApplication
	err := crdbpgx.ExecuteTx(context.Background(), r.db, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := tx.Query(ctx,
			`SELECT 
				id,
				user_id,
				officer_id,
				full_name,
				birth_date,
				full_address,
				phone,
				id_card_url,
				other_business,
				status,
				is_private_field,
				exp_in_year,
				active_field_number,
				sow_seeds_per_cycle,
				needed_fertilizier_per_cycle_in_kg,
				estimated_yield_in_kg,
				estimated_price_of_harvest_per_kg,
				harvest_cycle_in_months,
				loan_application_in_idr,
				business_income_per_month_in_idr,
				business_outcome_per_month_in_idr,
				created_date,
				updated_date
			FROM loan_applications`,
		)
		if err != nil {
			return err
		}

		for rows.Next() {
			var userLoan model.LoanApplication
			if err := rows.Scan(
				&userLoan.Id,
				&userLoan.UserId,
				&userLoan.OfficerId,
				&userLoan.FullName,
				&userLoan.BirthDate,
				&userLoan.FullAddress,
				&userLoan.Phone,
				&userLoan.IdCardUrl,
				&userLoan.OtherBusiness,
				&userLoan.Status,
				&userLoan.IsPrivateField,
				&userLoan.ExpInYear,
				&userLoan.ActiveFieldNumber,
				&userLoan.SowSeedsPerCycle,
				&userLoan.NeededFertilizerPerCycleInKg,
				&userLoan.EstimatedYieldInKg,
				&userLoan.EstimatedPriceOfHarvestPerKg,
				&userLoan.HarvestCycleInMonths,
				&userLoan.LoanApplicationInIdr,
				&userLoan.BusinessIncomePerMonthInIdr,
				&userLoan.BusinessOutcomePerMonthInIdr,
				&userLoan.CreatedDate,
				&userLoan.UpdatedDate,
			); err != nil {
				return err
			}
			loans = append(loans, userLoan)
		}

		return nil
	})
	if err != nil {
		return []model.LoanApplication{}, err
	}

	return loans, nil
}

func (r *Repository) GetLoan(ctx context.Context, loanId string) (model.LoanApplication, error) {
	var userLoan model.LoanApplication
	err := crdbpgx.ExecuteTx(context.Background(), r.db, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return tx.QueryRow(ctx,
			`SELECT 
				id,
				user_id,
				officer_id,
				full_name,
				birth_date,
				full_address,
				phone,
				id_card_url,
				other_business,
				status,
				is_private_field,
				exp_in_year,
				active_field_number,
				sow_seeds_per_cycle,
				needed_fertilizier_per_cycle_in_kg,
				estimated_yield_in_kg,
				estimated_price_of_harvest_per_kg,
				harvest_cycle_in_months,
				loan_application_in_idr,
				business_income_per_month_in_idr,
				business_outcome_per_month_in_idr,
				created_date,
				updated_date
			FROM loan_applications
			WHERE id = $1`,
			loanId,
		).Scan(
			&userLoan.Id,
			&userLoan.UserId,
			&userLoan.OfficerId,
			&userLoan.FullName,
			&userLoan.BirthDate,
			&userLoan.FullAddress,
			&userLoan.Phone,
			&userLoan.IdCardUrl,
			&userLoan.OtherBusiness,
			&userLoan.Status,
			&userLoan.IsPrivateField,
			&userLoan.ExpInYear,
			&userLoan.ActiveFieldNumber,
			&userLoan.SowSeedsPerCycle,
			&userLoan.NeededFertilizerPerCycleInKg,
			&userLoan.EstimatedYieldInKg,
			&userLoan.EstimatedPriceOfHarvestPerKg,
			&userLoan.HarvestCycleInMonths,
			&userLoan.LoanApplicationInIdr,
			&userLoan.BusinessIncomePerMonthInIdr,
			&userLoan.BusinessOutcomePerMonthInIdr,
			&userLoan.CreatedDate,
			&userLoan.UpdatedDate,
		)
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return model.LoanApplication{}, ErrUserLoanNotFound
	}
	if err != nil {
		return model.LoanApplication{}, err
	}

	return userLoan, nil
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

	err := crdbpgx.ExecuteTx(context.Background(), r.db, pgx.TxOptions{}, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx,
			`INSERT INTO loan_applications (
				id,
				user_id,
				full_name,
				birth_date,
				full_address,
				phone,
				id_card_url,
				other_business,
				status,
				is_private_field,
				exp_in_year,
				active_field_number,
				sow_seeds_per_cycle,
				needed_fertilizier_per_cycle_in_kg,
				estimated_yield_in_kg,
				estimated_price_of_harvest_per_kg,
				harvest_cycle_in_months,
				loan_application_in_idr,
				business_income_per_month_in_idr,
				business_outcome_per_month_in_idr,
				created_date,
				updated_date
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)`,
			loan.Id,
			loan.UserId,
			loan.FullName,
			loan.BirthDate,
			loan.FullAddress,
			loan.Phone,
			loan.IdCardUrl,
			loan.OtherBusiness,
			loan.Status,
			loan.IsPrivateField,
			loan.ExpInYear,
			loan.ActiveFieldNumber,
			loan.SowSeedsPerCycle,
			loan.NeededFertilizerPerCycleInKg,
			loan.EstimatedYieldInKg,
			loan.EstimatedPriceOfHarvestPerKg,
			loan.HarvestCycleInMonths,
			loan.LoanApplicationInIdr,
			loan.BusinessIncomePerMonthInIdr,
			loan.BusinessOutcomePerMonthInIdr,
			loan.CreatedDate,
			loan.UpdatedDate,
		); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return model.LoanApplication{}, err
	}

	return loan, nil
}

func (r *Repository) RemoveLoan(ctx context.Context, loanId string) error {
	err := crdbpgx.ExecuteTx(context.Background(), r.db, pgx.TxOptions{}, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx,
			`DELETE FROM loan_applications WHERE id = $1`,
			loanId,
		); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateLoan(ctx context.Context, loanId string, loan model.LoanApplication) error {
	t := time.Now()
	err := crdbpgx.ExecuteTx(context.Background(), r.db, pgx.TxOptions{}, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx,
			`UPDATE loan_applications SET (
				officer_id,
				full_name,
				birth_date,
				full_address,
				phone,
				id_card_url,
				other_business,
				status,
				is_private_field,
				exp_in_year,
				active_field_number,
				sow_seeds_per_cycle,
				needed_fertilizier_per_cycle_in_kg,
				estimated_yield_in_kg,
				estimated_price_of_harvest_per_kg,
				harvest_cycle_in_months,
				loan_application_in_idr,
				business_income_per_month_in_idr,
				business_outcome_per_month_in_idr,
				updated_date
			) = ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)`,
			loan.OfficerId,
			loan.FullName,
			loan.BirthDate,
			loan.FullAddress,
			loan.Phone,
			loan.IdCardUrl,
			loan.OtherBusiness,
			loan.Status,
			loan.IsPrivateField,
			loan.ExpInYear,
			loan.ActiveFieldNumber,
			loan.SowSeedsPerCycle,
			loan.NeededFertilizerPerCycleInKg,
			loan.EstimatedYieldInKg,
			loan.EstimatedPriceOfHarvestPerKg,
			loan.HarvestCycleInMonths,
			loan.LoanApplicationInIdr,
			loan.BusinessIncomePerMonthInIdr,
			loan.BusinessOutcomePerMonthInIdr,
			t,
		); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
