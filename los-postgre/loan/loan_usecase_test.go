package loan_test

import (
	"context"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/fikryfahrezy/adea/los-postgre/auth"
	"github.com/fikryfahrezy/adea/los-postgre/data"
	"github.com/fikryfahrezy/adea/los-postgre/loan"
	"github.com/fikryfahrezy/adea/los-postgre/model"
)

var (
	uploadFunc = func(filename string, file io.Reader) (string, error) {
		return "", nil
	}
	dbJson   = data.NewJson("")
	authRepo = auth.NewRepository(dbJson)
	loanRepo = loan.NewRepository(dbJson)
	loanApp  = loan.NewApp(uploadFunc, loanRepo)
)

func clearDb() {
	dbJson.DbUser = make(map[string]model.User)
	dbJson.DbLoan = make(map[string]model.LoanApplication)
}

func TestGetUserLoans(t *testing.T) {
	clearDb()

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	loanRepo.InsertLoan(ctx, newLoan)

	out := loanApp.GetUserLoans(ctx, user.Id)

	if len(out.Res) != 1 {
		t.Fatalf("resulting: %d, expect: %d | err: %v", len(out.Res), 1, out.Error)
	}
}

func TestGetUserLoansButUserNotFound(t *testing.T) {
	clearDb()

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	loanRepo.InsertLoan(ctx, newLoan)

	out := loanApp.GetUserLoans(ctx, "some-random-user-id")

	if out.StatusCode != http.StatusNotFound {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusNotFound, out.Error)
	}
}

func TestGetUserLoanDetail(t *testing.T) {
	clearDb()

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	loanRepo.InsertLoan(ctx, newLoan)

	out := loanApp.GetUserLoanDetail(ctx, newLoan.Id, user.Id)

	if out.Res.LoanId != newLoan.Id {
		t.Fatalf("resulting: %v, expect: %v | err: %v", out.Res.LoanId, newLoan.Id, out.Error)
	}
}

func TestGetUserLoanDetailButLoanNotBelongToUser(t *testing.T) {
	clearDb()

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	user2 := model.User{
		Username:  "username2",
		Password:  "password",
		IsOfficer: false,
	}
	user2, _ = authRepo.InsertUser(ctx, user2)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	loanRepo.InsertLoan(ctx, newLoan)

	out := loanApp.GetUserLoanDetail(ctx, newLoan.Id, user2.Id)

	if out.StatusCode != http.StatusNotFound {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusNotFound, out.Error)
	}
}

func TestGetUserLoanDetailButUserNotFound(t *testing.T) {
	clearDb()

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	loanRepo.InsertLoan(ctx, newLoan)

	out := loanApp.GetUserLoanDetail(ctx, newLoan.Id, "some-random-user-id")

	if out.StatusCode != http.StatusNotFound {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusNotFound, out.Error)
	}
}

func TestGetUserLoanDetailButLoanNotFound(t *testing.T) {
	clearDb()

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	out := loanApp.GetUserLoanDetail(ctx, "some-random-loan-id", user.Id)

	if out.StatusCode != http.StatusNotFound {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusNotFound, out.Error)
	}
}

func TestCreateNewLoan(t *testing.T) {
	clearDb()

	ctx := context.Background()
	f, err := os.OpenFile("./loan_application.go", os.O_RDONLY, 0o444)
	if err != nil {
		t.Fatal(err)
	}

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	testCases := []struct {
		expect int
		name   string
		in     loan.CreateLoanIn
	}{
		{
			expect: http.StatusCreated,
			name:   "Create loan sucessfully",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, exp in year is 0",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    0,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, exp in year is < 0",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    -1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, active field number is 0",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            0,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, active field number is < 0",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            -1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, sow seed per cycle is 0",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             0,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, sow seed per cycle is < 0",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             -1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, needed fert per cycle in kg is 0",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 0,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, needed fert per cycle in kg is < 0",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: -1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, estimated yield is 0",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           0,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, estimated yield is < 0",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           -1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, estimated price harvest is 0",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 0,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, estimated price harvest is < 0",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: -1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, harvest cycle in months is 0",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         0,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, harvest cycle in months is < 0",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         -1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, loan idr is 0",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         0,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, loan idr is < 0",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         -1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, income idr is 0",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  0,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, income idr is < 0",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  -1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, outcome idr is 0",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 0,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, outcome idr is < 0",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: -1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, empty full name",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, empty birth date",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, empty full address",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, empty phone",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, phone not match min length",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, phone not match max length",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000000000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, phone not only numbers",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000xx",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Create loan fail, id card file required",
			in: loan.CreateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
				},
			},
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			out := loanApp.CreateLoan(ctx, user.Id, c.in)

			if out.StatusCode != c.expect {
				t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, c.expect, out.Error)
			}
		})
	}
}

func TestCreateNewLoanWithExist(t *testing.T) {
	clearDb()

	f, err := os.OpenFile("./loan_application.go", os.O_RDONLY, 0o444)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	newLoan, _ = loanRepo.InsertLoan(ctx, newLoan)

	newLoan.Status = loan.Process.String()
	loanRepo.UpdateLoan(ctx, newLoan.Id, newLoan)

	otherNewLoan := loan.CreateLoanIn{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		IdCard: loan.FileHeader{
			Filename: "test.img",
			File:     f,
		},
	}

	out := loanApp.CreateLoan(ctx, user.Id, otherNewLoan)
	if out.StatusCode != http.StatusBadRequest {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusBadRequest, out.Error)
	}
}

func TestCreateNewLoanWithExistProcess(t *testing.T) {
	clearDb()

	f, err := os.OpenFile("./loan_application.go", os.O_RDONLY, 0o444)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	newLoan, _ = loanRepo.InsertLoan(ctx, newLoan)

	newLoan.Status = loan.Process.String()
	loanRepo.UpdateLoan(ctx, newLoan.Id, newLoan)

	otherNewLoan := loan.CreateLoanIn{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		IdCard: loan.FileHeader{
			Filename: "test.img",
			File:     f,
		},
	}

	out := loanApp.CreateLoan(ctx, user.Id, otherNewLoan)
	if out.StatusCode != http.StatusBadRequest {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusBadRequest, out.Error)
	}
}

func TestCreateNewLoanWithExistReject(t *testing.T) {
	clearDb()

	f, err := os.OpenFile("./loan_application.go", os.O_RDONLY, 0o444)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	newLoan, _ = loanRepo.InsertLoan(ctx, newLoan)

	newLoan.Status = loan.Reject.String()
	loanRepo.UpdateLoan(ctx, newLoan.Id, newLoan)

	otherNewLoan := loan.CreateLoanIn{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		IdCard: loan.FileHeader{
			Filename: "test.img",
			File:     f,
		},
	}

	out := loanApp.CreateLoan(ctx, user.Id, otherNewLoan)
	if out.StatusCode != http.StatusCreated {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusCreated, out.Error)
	}
}

func TestCreateNewLoanWithExistApprove(t *testing.T) {
	clearDb()

	f, err := os.OpenFile("./loan_application.go", os.O_RDONLY, 0o444)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	newLoan, _ = loanRepo.InsertLoan(ctx, newLoan)

	newLoan.Status = loan.Approve.String()
	loanRepo.UpdateLoan(ctx, newLoan.Id, newLoan)

	otherNewLoan := loan.CreateLoanIn{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		IdCard: loan.FileHeader{
			Filename: "test.img",
			File:     f,
		},
	}

	out := loanApp.CreateLoan(ctx, user.Id, otherNewLoan)
	if out.StatusCode != http.StatusCreated {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusCreated, out.Error)
	}
}

func TestCreateNewLoanButUserNotExist(t *testing.T) {
	clearDb()

	f, err := os.OpenFile("./loan_application.go", os.O_RDONLY, 0o444)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	otherNewLoan := loan.CreateLoanIn{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		IdCard: loan.FileHeader{
			Filename: "test.img",
			File:     f,
		},
	}

	out := loanApp.CreateLoan(ctx, "some-random-user-id", otherNewLoan)
	if out.StatusCode != http.StatusNotFound {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusNotFound, out.Error)
	}
}

func TestUpdateLoan(t *testing.T) {
	clearDb()

	ctx := context.Background()
	f, err := os.OpenFile("./loan_application.go", os.O_RDONLY, 0o444)
	if err != nil {
		t.Fatal(err)
	}

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	newLoan, _ = loanRepo.InsertLoan(ctx, newLoan)

	testCases := []struct {
		expect int
		name   string
		in     loan.UpdateLoanIn
	}{
		{
			expect: http.StatusOK,
			name:   "Update loan sucessfully",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, exp in year is 0",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    0,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, exp in year is < 0",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    -1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, active field number is 0",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            0,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, active field number is < 0",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            -1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, sow seed per cycle is 0",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             0,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, sow seed per cycle is < 0",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             -1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, needed fert per cycle in kg is 0",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 0,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, needed fert per cycle in kg is < 0",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: -1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, estimated yield is 0",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           0,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, estimated yield is < 0",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           -1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, estimated price harvest is 0",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 0,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, estimated price harvest is < 0",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: -1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, harvest cycle in months is 0",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         0,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, harvest cycle in months is < 0",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         -1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, loan idr is 0",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         0,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, loan idr is < 0",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         -1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, income idr is 0",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  0,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, income idr is < 0",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  -1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, outcome idr is 0",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 0,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, outcome idr is < 0",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: -1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, empty full name",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, empty birth date",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, empty full address",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, empty phone",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, phone not match min length",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, phone not match max length",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000000000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, phone not only numbers",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000xx",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
					File:     f,
				},
			},
		},
		{
			expect: http.StatusUnprocessableEntity,
			name:   "Update loan fail, id card file required",
			in: loan.UpdateLoanIn{
				IsPrivateField:               true,
				ExpInYear:                    1,
				ActiveFieldNumber:            1,
				SowSeedsPerCycle:             1,
				NeededFertilizerPerCycleInKg: 1,
				EstimatedYieldInKg:           1,
				EstimatedPriceOfHarvestPerKg: 1,
				HarvestCycleInMonths:         1,
				LoanApplicationInIdr:         1,
				BusinessIncomePerMonthInIdr:  1,
				BusinessOutcomePerMonthInIdr: 1,
				FullName:                     "Full Name",
				BirthDate:                    "2006-01-02",
				FullAddress:                  "Full Address",
				Phone:                        "0000000000",
				OtherBussiness:               "-",
				IdCard: loan.FileHeader{
					Filename: "test.img",
				},
			},
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			out := loanApp.UpdateLoan(ctx, newLoan.Id, user.Id, c.in)

			if out.StatusCode != c.expect {
				t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, c.expect, out.Error)
			}
		})
	}
}

func TestUpdateLoanButUserNotExist(t *testing.T) {
	clearDb()

	f, err := os.OpenFile("./loan_application.go", os.O_RDONLY, 0o444)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		IdCardUrl:                    "http://random",
	}
	newLoan, _ = loanRepo.InsertLoan(ctx, newLoan)

	otherNewLoan := loan.UpdateLoanIn{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		IdCard: loan.FileHeader{
			Filename: "test.img",
			File:     f,
		},
	}

	out := loanApp.UpdateLoan(ctx, newLoan.Id, "some-random-user-id", otherNewLoan)
	if out.StatusCode != http.StatusNotFound {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusNotFound, out.Error)
	}
}

func TestUpdateLoanButLoanNotExist(t *testing.T) {
	clearDb()

	f, err := os.OpenFile("./loan_application.go", os.O_RDONLY, 0o444)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	otherNewLoan := loan.UpdateLoanIn{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		IdCard: loan.FileHeader{
			Filename: "test.img",
			File:     f,
		},
	}

	out := loanApp.UpdateLoan(ctx, "loan-not-found-id", user.Id, otherNewLoan)
	if out.StatusCode != http.StatusNotFound {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusNotFound, out.Error)
	}
}

func TestUpdateLoanButLoanNotBelongToUser(t *testing.T) {
	clearDb()

	f, err := os.OpenFile("./loan_application.go", os.O_RDONLY, 0o444)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	user2 := model.User{
		Username:  "username2",
		Password:  "password2",
		IsOfficer: false,
	}
	user2, _ = authRepo.InsertUser(ctx, user2)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	newLoan, _ = loanRepo.InsertLoan(ctx, newLoan)

	otherNewLoan := loan.UpdateLoanIn{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		IdCard: loan.FileHeader{
			Filename: "test.img",
			File:     f,
		},
	}

	out := loanApp.UpdateLoan(ctx, newLoan.Id, user2.Id, otherNewLoan)
	if out.StatusCode != http.StatusNotFound {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusNotFound, out.Error)
	}
}

func TestUpdateLoanWithStatusProcess(t *testing.T) {
	clearDb()

	f, err := os.OpenFile("./loan_application.go", os.O_RDONLY, 0o444)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	newLoan, _ = loanRepo.InsertLoan(ctx, newLoan)

	newLoan.Status = loan.Process.String()
	loanRepo.UpdateLoan(ctx, newLoan.Id, newLoan)

	otherNewLoan := loan.UpdateLoanIn{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		IdCard: loan.FileHeader{
			Filename: "test.img",
			File:     f,
		},
	}

	out := loanApp.UpdateLoan(ctx, newLoan.Id, user.Id, otherNewLoan)
	if out.StatusCode != http.StatusBadRequest {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusBadRequest, out.Error)
	}
}

func TestUpdateLoanWithStatusReject(t *testing.T) {
	clearDb()

	f, err := os.OpenFile("./loan_application.go", os.O_RDONLY, 0o444)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	newLoan, _ = loanRepo.InsertLoan(ctx, newLoan)

	newLoan.Status = loan.Reject.String()
	loanRepo.UpdateLoan(ctx, newLoan.Id, newLoan)

	otherNewLoan := loan.UpdateLoanIn{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		IdCard: loan.FileHeader{
			Filename: "test.img",
			File:     f,
		},
	}

	out := loanApp.UpdateLoan(ctx, newLoan.Id, user.Id, otherNewLoan)
	if out.StatusCode != http.StatusBadRequest {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusBadRequest, out.Error)
	}
}

func TestUpdateLoanWithStatusApprove(t *testing.T) {
	clearDb()

	f, err := os.OpenFile("./loan_application.go", os.O_RDONLY, 0o444)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	newLoan, _ = loanRepo.InsertLoan(ctx, newLoan)

	newLoan.Status = loan.Approve.String()
	loanRepo.UpdateLoan(ctx, newLoan.Id, newLoan)

	otherNewLoan := loan.UpdateLoanIn{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		IdCard: loan.FileHeader{
			Filename: "test.img",
			File:     f,
		},
	}

	out := loanApp.UpdateLoan(ctx, newLoan.Id, user.Id, otherNewLoan)
	if out.StatusCode != http.StatusBadRequest {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusBadRequest, out.Error)
	}
}

func TestDeleteUserLoan(t *testing.T) {
	clearDb()

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	loanRepo.InsertLoan(ctx, newLoan)

	out := loanApp.DeleteLoan(ctx, newLoan.Id, user.Id)

	if out.Res.Id != newLoan.Id {
		t.Fatalf("resulting: %v, expect: %v | err: %v", out.Res.Id, newLoan.Id, out.Error)
	}
}

func TestDeleteUserLoanButLoanNotBelongToUser(t *testing.T) {
	clearDb()

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	user2 := model.User{
		Username:  "username2",
		Password:  "password",
		IsOfficer: false,
	}
	user2, _ = authRepo.InsertUser(ctx, user2)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	loanRepo.InsertLoan(ctx, newLoan)

	out := loanApp.GetUserLoanDetail(ctx, newLoan.Id, user2.Id)

	if out.StatusCode != http.StatusNotFound {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusNotFound, out.Error)
	}
}

func TestDeleteUserLoanButUserNotFound(t *testing.T) {
	clearDb()

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	loanRepo.InsertLoan(ctx, newLoan)

	out := loanApp.GetUserLoanDetail(ctx, newLoan.Id, "some-random-user-id")

	if out.StatusCode != http.StatusNotFound {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusNotFound, out.Error)
	}
}

func TestDeleteUserLoanButLoanNotFound(t *testing.T) {
	clearDb()

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	out := loanApp.GetUserLoanDetail(ctx, "some-random-loan-id", user.Id)

	if out.StatusCode != http.StatusNotFound {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusNotFound, out.Error)
	}
}

func TestDeleteLoanWithStatusProcess(t *testing.T) {
	clearDb()
	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	newLoan, _ = loanRepo.InsertLoan(ctx, newLoan)

	newLoan.Status = loan.Process.String()
	loanRepo.UpdateLoan(ctx, newLoan.Id, newLoan)

	out := loanApp.DeleteLoan(ctx, newLoan.Id, user.Id)
	if out.StatusCode != http.StatusBadRequest {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusBadRequest, out.Error)
	}
}

func TestDeleteLoanWithStatusReject(t *testing.T) {
	clearDb()

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	newLoan, _ = loanRepo.InsertLoan(ctx, newLoan)

	newLoan.Status = loan.Reject.String()
	loanRepo.UpdateLoan(ctx, newLoan.Id, newLoan)

	out := loanApp.DeleteLoan(ctx, newLoan.Id, user.Id)
	if out.StatusCode != http.StatusBadRequest {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusBadRequest, out.Error)
	}
}

func TestDeleteLoanWithStatusApprove(t *testing.T) {
	clearDb()

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	newLoan, _ = loanRepo.InsertLoan(ctx, newLoan)

	newLoan.Status = loan.Approve.String()
	loanRepo.UpdateLoan(ctx, newLoan.Id, newLoan)

	out := loanApp.DeleteLoan(ctx, newLoan.Id, user.Id)
	if out.StatusCode != http.StatusBadRequest {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusBadRequest, out.Error)
	}
}

func TestGetLoans(t *testing.T) {
	clearDb()

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	loanRepo.InsertLoan(ctx, newLoan)

	out := loanApp.GetLoans(ctx)

	if len(out.Res) != 1 {
		t.Fatalf("resulting: %d, expect: %d | err: %v", len(out.Res), 1, out.Error)
	}
}

func TestGetLoanDetail(t *testing.T) {
	clearDb()

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	loanRepo.InsertLoan(ctx, newLoan)

	out := loanApp.GetLoanDetail(ctx, newLoan.Id)

	if out.Res.LoanId != newLoan.Id {
		t.Fatalf("resulting: %v, expect: %v | err: %v", out.Res.LoanId, newLoan.Id, out.Error)
	}
}

func TestGetLoanDetailButLoanNotFound(t *testing.T) {
	clearDb()

	ctx := context.Background()
	out := loanApp.GetLoanDetail(ctx, "some-random-loan-id")

	if out.StatusCode != http.StatusNotFound {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusNotFound, out.Error)
	}
}

func TestProceedLoan(t *testing.T) {
	clearDb()

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	loanRepo.InsertLoan(ctx, newLoan)

	out := loanApp.ProceedLoan(ctx, newLoan.Id)

	if out.Res.Id != newLoan.Id {
		t.Fatalf("resulting: %v, expect: %v | err: %v", out.Res.Id, newLoan.Id, out.Error)
	}
}

func TestProceedUserLoanButLoanNotFound(t *testing.T) {
	clearDb()

	ctx := context.Background()

	out := loanApp.ProceedLoan(ctx, "some-random-loan-id")

	if out.StatusCode != http.StatusNotFound {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusNotFound, out.Error)
	}
}

func TestProceedLoanWithStatusProcess(t *testing.T) {
	clearDb()
	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	newLoan, _ = loanRepo.InsertLoan(ctx, newLoan)

	newLoan.Status = loan.Process.String()
	loanRepo.UpdateLoan(ctx, newLoan.Id, newLoan)

	out := loanApp.ProceedLoan(ctx, newLoan.Id)
	if out.StatusCode != http.StatusBadRequest {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusBadRequest, out.Error)
	}
}

func TestProceedLoanWithStatusReject(t *testing.T) {
	clearDb()

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	newLoan, _ = loanRepo.InsertLoan(ctx, newLoan)

	newLoan.Status = loan.Reject.String()
	loanRepo.UpdateLoan(ctx, newLoan.Id, newLoan)

	out := loanApp.ProceedLoan(ctx, newLoan.Id)
	if out.StatusCode != http.StatusBadRequest {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusBadRequest, out.Error)
	}
}

func TestProceedLoanWithStatusApprove(t *testing.T) {
	clearDb()

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	newLoan, _ = loanRepo.InsertLoan(ctx, newLoan)

	newLoan.Status = loan.Approve.String()
	loanRepo.UpdateLoan(ctx, newLoan.Id, newLoan)

	out := loanApp.ProceedLoan(ctx, newLoan.Id)
	if out.StatusCode != http.StatusBadRequest {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusBadRequest, out.Error)
	}
}

func TestApproveLoan(t *testing.T) {
	clearDb()

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	loanRepo.InsertLoan(ctx, newLoan)

	out := loanApp.ApproveLoan(ctx, newLoan.Id, loan.ApproveLoanIn{
		IsApprove: true,
	})

	if out.Res.Id != newLoan.Id {
		t.Fatalf("resulting: %v, expect: %v | err: %v", out.Res.Id, newLoan.Id, out.Error)
	}
}

func TestApproveLoanButLoanNotFound(t *testing.T) {
	clearDb()

	ctx := context.Background()
	out := loanApp.ApproveLoan(ctx, "some-random-loan-id", loan.ApproveLoanIn{
		IsApprove: true,
	})

	if out.StatusCode != http.StatusNotFound {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusNotFound, out.Error)
	}
}

func TestApproveLoanWithStatusWait(t *testing.T) {
	clearDb()
	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	newLoan, _ = loanRepo.InsertLoan(ctx, newLoan)

	newLoan.Status = loan.Wait.String()
	loanRepo.UpdateLoan(ctx, newLoan.Id, newLoan)

	out := loanApp.ApproveLoan(ctx, newLoan.Id, loan.ApproveLoanIn{
		IsApprove: true,
	})
	if out.StatusCode != http.StatusBadRequest {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusBadRequest, out.Error)
	}
}

func TestApproveLoanWithStatusReject(t *testing.T) {
	clearDb()

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	newLoan, _ = loanRepo.InsertLoan(ctx, newLoan)

	newLoan.Status = loan.Reject.String()
	loanRepo.UpdateLoan(ctx, newLoan.Id, newLoan)

	out := loanApp.ApproveLoan(ctx, newLoan.Id, loan.ApproveLoanIn{
		IsApprove: true,
	})
	if out.StatusCode != http.StatusBadRequest {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusBadRequest, out.Error)
	}
}

func TestApproveLoanWithStatusApprove(t *testing.T) {
	clearDb()

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	newLoan, _ = loanRepo.InsertLoan(ctx, newLoan)

	newLoan.Status = loan.Approve.String()
	loanRepo.UpdateLoan(ctx, newLoan.Id, newLoan)

	out := loanApp.ApproveLoan(ctx, newLoan.Id, loan.ApproveLoanIn{
		IsApprove: true,
	})
	if out.StatusCode != http.StatusBadRequest {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusBadRequest, out.Error)
	}
}

func TestRejectLoan(t *testing.T) {
	clearDb()

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	loanRepo.InsertLoan(ctx, newLoan)

	out := loanApp.ApproveLoan(ctx, newLoan.Id, loan.ApproveLoanIn{
		IsApprove: false,
	})

	if out.Res.Id != newLoan.Id {
		t.Fatalf("resulting: %v, expect: %v | err: %v", out.Res.Id, newLoan.Id, out.Error)
	}
}

func TestRejectLoanButLoanNotFound(t *testing.T) {
	clearDb()

	ctx := context.Background()
	out := loanApp.ApproveLoan(ctx, "some-random-loan-id", loan.ApproveLoanIn{
		IsApprove: false,
	})

	if out.StatusCode != http.StatusNotFound {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusNotFound, out.Error)
	}
}

func TestRejectLoanWithStatusWait(t *testing.T) {
	clearDb()
	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	newLoan, _ = loanRepo.InsertLoan(ctx, newLoan)

	newLoan.Status = loan.Wait.String()
	loanRepo.UpdateLoan(ctx, newLoan.Id, newLoan)

	out := loanApp.ApproveLoan(ctx, newLoan.Id, loan.ApproveLoanIn{
		IsApprove: false,
	})
	if out.StatusCode != http.StatusBadRequest {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusBadRequest, out.Error)
	}
}

func TestRejectLoanWithStatusReject(t *testing.T) {
	clearDb()

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	newLoan, _ = loanRepo.InsertLoan(ctx, newLoan)

	newLoan.Status = loan.Reject.String()
	loanRepo.UpdateLoan(ctx, newLoan.Id, newLoan)

	out := loanApp.ApproveLoan(ctx, newLoan.Id, loan.ApproveLoanIn{
		IsApprove: false,
	})
	if out.StatusCode != http.StatusBadRequest {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusBadRequest, out.Error)
	}
}

func TestRejectLoanWithStatusApprove(t *testing.T) {
	clearDb()

	ctx := context.Background()

	user := model.User{
		Username:  "username",
		Password:  "password",
		IsOfficer: false,
	}
	user, _ = authRepo.InsertUser(ctx, user)

	newLoan := model.LoanApplication{
		IsPrivateField:               true,
		ExpInYear:                    1,
		ActiveFieldNumber:            1,
		SowSeedsPerCycle:             1,
		NeededFertilizerPerCycleInKg: 1,
		EstimatedYieldInKg:           1,
		EstimatedPriceOfHarvestPerKg: 1,
		HarvestCycleInMonths:         1,
		LoanApplicationInIdr:         1,
		BusinessIncomePerMonthInIdr:  1,
		BusinessOutcomePerMonthInIdr: 1,
		FullName:                     "Full Name",
		BirthDate:                    "2006-01-02",
		FullAddress:                  "Full Address",
		Phone:                        "0000000000",
		OtherBussiness:               "-",
		UserId:                       user.Id,
		IdCardUrl:                    "http://random",
	}
	newLoan, _ = loanRepo.InsertLoan(ctx, newLoan)

	newLoan.Status = loan.Approve.String()
	loanRepo.UpdateLoan(ctx, newLoan.Id, newLoan)

	out := loanApp.ApproveLoan(ctx, newLoan.Id, loan.ApproveLoanIn{
		IsApprove: false,
	})
	if out.StatusCode != http.StatusBadRequest {
		t.Fatalf("resulting: %d, expect: %d | err: %v", out.StatusCode, http.StatusBadRequest, out.Error)
	}
}
