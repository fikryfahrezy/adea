package loan

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/fikryfahrezy/adea/los-postgre/model"
	"github.com/fikryfahrezy/adea/los-postgre/resp"
)

var (
	ErrProcessLoanExist  = errors.New("already have processed loan")
	ErrModifyProcessLoan = errors.New("cannot modify processed loan")
	ErrUserForbidden     = errors.New("officer only")
)

type File interface {
	io.Reader
	io.Closer
}

type FileHeader struct {
	Filename string
	File     File
}

type (
	GetUserLoanRes struct {
		LoanId          string `json:"loan_id"`
		UserId          string `json:"user_id"`
		FullName        string `json:"full_name"`
		LoanStatus      string `json:"loan_status"`
		LoanCreatedDate string `json:"loan_created_date"`
	}
	GetUserLoanOut struct {
		resp.Response
		Res []GetUserLoanRes
	}
)

func (a *LoanApp) GetUserLoans(ctx context.Context, userId string) (out GetUserLoanOut) {
	out.Response = resp.NewResponse(http.StatusOK, "", nil)

	_, err := a.repository.GetUser(ctx, userId)
	if errors.Is(err, ErrUserNotFound) {
		out.Response = resp.NewResponse(http.StatusNotFound, "", err)
		return
	}
	if err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	userLoans, err := a.repository.GetUserLoans(ctx, userId)
	if err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	res := make([]GetUserLoanRes, 0, 0)
	for _, loan := range userLoans {
		res = append(res, GetUserLoanRes{
			LoanId:          loan.Id,
			UserId:          loan.UserId,
			FullName:        loan.FullName,
			LoanStatus:      loan.Status,
			LoanCreatedDate: loan.CreatedDate.Format("2006-01-02"),
		})
	}

	out.Res = res

	return
}

type (
	GetUserLoanDetailRes struct {
		IsPrivateField               bool   `json:"is_private_field"`
		ExpInYear                    int64  `json:"exp_in_year"`
		ActiveFieldNumber            int64  `json:"active_field_number"`
		SowSeedsPerCycle             int64  `json:"sow_seeds_per_cycle"`
		NeededFertilizerPerCycleInKg int64  `json:"needed_fertilizer_per_cycle_in_kg"`
		EstimatedYieldInKg           int64  `json:"estimated_yield_in_kg"`
		EstimatedPriceOfHarvestPerKg int64  `json:"estimated_price_of_harvest_per_kg"`
		HarvestCycleInMonths         int64  `json:"harvest_cycle_in_months"`
		LoanApplicationInIdr         int64  `json:"loan_application_in_idr"`
		BusinessIncomePerMonthInIdr  int64  `json:"business_income_per_month_in_idr"`
		BusinessOutcomePerMonthInIdr int64  `json:"business_outcome_per_month_in_idr"`
		LoanId                       string `json:"loan_id"`
		UserId                       string `json:"user_id"`
		FullName                     string `json:"full_name"`
		BirthDate                    string `json:"birth_date"`
		FullAddress                  string `json:"full_address"`
		Phone                        string `json:"phone"`
		OtherBusiness                string `json:"other_business"`
		IdCardUrl                    string `json:"id_card_url"`
		Status                       string `json:"status"`
	}
	GetUserLoanDetailOut struct {
		resp.Response
		Res GetUserLoanDetailRes
	}
)

func (a *LoanApp) GetUserLoanDetail(ctx context.Context, loanId, userId string) (out GetUserLoanDetailOut) {
	out.Response = resp.NewResponse(http.StatusOK, "", nil)

	_, err := a.repository.GetUser(ctx, userId)
	if errors.Is(err, ErrUserNotFound) {
		out.Response = resp.NewResponse(http.StatusNotFound, "", err)
		return
	}
	if err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	userLoan, err := a.repository.GetUserLoan(ctx, loanId, userId)
	if errors.Is(err, ErrUserLoanNotFound) {
		out.Response = resp.NewResponse(http.StatusNotFound, "", err)
		return
	}
	if err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	out.Res = GetUserLoanDetailRes{
		IsPrivateField:               userLoan.IsPrivateField,
		ExpInYear:                    userLoan.ExpInYear,
		ActiveFieldNumber:            userLoan.ActiveFieldNumber,
		SowSeedsPerCycle:             userLoan.SowSeedsPerCycle,
		NeededFertilizerPerCycleInKg: userLoan.NeededFertilizerPerCycleInKg,
		EstimatedYieldInKg:           userLoan.EstimatedPriceOfHarvestPerKg,
		EstimatedPriceOfHarvestPerKg: userLoan.EstimatedPriceOfHarvestPerKg,
		HarvestCycleInMonths:         userLoan.HarvestCycleInMonths,
		LoanApplicationInIdr:         userLoan.LoanApplicationInIdr,
		BusinessIncomePerMonthInIdr:  userLoan.BusinessIncomePerMonthInIdr,
		BusinessOutcomePerMonthInIdr: userLoan.BusinessOutcomePerMonthInIdr,
		LoanId:                       userLoan.Id,
		UserId:                       userLoan.UserId,
		FullName:                     userLoan.FullName,
		BirthDate:                    userLoan.BirthDate,
		FullAddress:                  userLoan.FullAddress,
		Phone:                        userLoan.Phone,
		OtherBusiness:                userLoan.OtherBusiness,
		IdCardUrl:                    userLoan.IdCardUrl,
		Status:                       userLoan.Status,
	}

	return
}

type (
	CreateLoanIn struct {
		IsPrivateField               bool
		ExpInYear                    int64
		ActiveFieldNumber            int64
		SowSeedsPerCycle             int64
		NeededFertilizerPerCycleInKg int64
		EstimatedYieldInKg           int64
		EstimatedPriceOfHarvestPerKg int64
		HarvestCycleInMonths         int64
		LoanApplicationInIdr         int64
		BusinessIncomePerMonthInIdr  int64
		BusinessOutcomePerMonthInIdr int64
		FullName                     string
		BirthDate                    string
		FullAddress                  string
		Phone                        string
		OtherBusiness                string
		IdCard                       FileHeader
	}
	CreateLoanRes struct {
		Id string `json:"id"`
	}
	CreateLoanOut struct {
		resp.Response
		Res CreateLoanRes
	}
)

func (a *LoanApp) CreateLoan(ctx context.Context, userId string, in CreateLoanIn) (out CreateLoanOut) {
	out.Response = resp.NewResponse(http.StatusCreated, "", nil)

	if err := validateCreateLoan(in); err != nil {
		out.Response = resp.NewResponse(http.StatusUnprocessableEntity, "", err)
		return
	}

	_, err := a.repository.GetUser(ctx, userId)
	if errors.Is(err, ErrUserNotFound) {
		out.Response = resp.NewResponse(http.StatusNotFound, "", err)
		return
	}
	if err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	userLoans, err := a.repository.GetUserLoans(ctx, userId)
	if err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	for _, loan := range userLoans {
		if loan.Status == Wait.String() || loan.Status == Process.String() {
			out.Response = resp.NewResponse(http.StatusBadRequest, "", ErrProcessLoanExist)
			return
		}
	}

	var fileUrl string
	if in.IdCard.File != nil {
		var err error
		fileUrl, err = a.saveFile(in.IdCard.Filename, in.IdCard.File)
		if err != nil {
			out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
			return
		}
	}

	newLoan := model.LoanApplication{
		IsPrivateField:               in.IsPrivateField,
		ExpInYear:                    in.ExpInYear,
		ActiveFieldNumber:            in.ActiveFieldNumber,
		SowSeedsPerCycle:             in.SowSeedsPerCycle,
		NeededFertilizerPerCycleInKg: in.NeededFertilizerPerCycleInKg,
		EstimatedYieldInKg:           in.EstimatedPriceOfHarvestPerKg,
		EstimatedPriceOfHarvestPerKg: in.EstimatedPriceOfHarvestPerKg,
		HarvestCycleInMonths:         in.HarvestCycleInMonths,
		LoanApplicationInIdr:         in.LoanApplicationInIdr,
		BusinessIncomePerMonthInIdr:  in.BusinessIncomePerMonthInIdr,
		BusinessOutcomePerMonthInIdr: in.BusinessOutcomePerMonthInIdr,
		UserId:                       userId,
		FullName:                     in.FullName,
		BirthDate:                    in.BirthDate,
		FullAddress:                  in.FullAddress,
		Phone:                        in.Phone,
		IdCardUrl:                    fileUrl,
		OtherBusiness:                in.OtherBusiness,
	}

	if newLoan, err = a.repository.InsertLoan(ctx, newLoan); err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	out.Res = CreateLoanRes{
		Id: newLoan.Id,
	}

	return
}

type (
	UpdateLoanIn struct {
		IsPrivateField               bool
		ExpInYear                    int64
		ActiveFieldNumber            int64
		SowSeedsPerCycle             int64
		NeededFertilizerPerCycleInKg int64
		EstimatedYieldInKg           int64
		EstimatedPriceOfHarvestPerKg int64
		HarvestCycleInMonths         int64
		LoanApplicationInIdr         int64
		BusinessIncomePerMonthInIdr  int64
		BusinessOutcomePerMonthInIdr int64
		FullName                     string
		BirthDate                    string
		FullAddress                  string
		Phone                        string
		OtherBusiness                string
		IdCard                       FileHeader
	}
	UpdateLoanRes struct {
		Id string `json:"id"`
	}
	UpdateLoanOut struct {
		resp.Response
		Res UpdateLoanRes
	}
)

func (a *LoanApp) UpdateLoan(ctx context.Context, loanId string, userId string, in UpdateLoanIn) (out UpdateLoanOut) {
	out.Response = resp.NewResponse(http.StatusOK, "", nil)

	if err := validateUpdateLoan(in); err != nil {
		out.Response = resp.NewResponse(http.StatusUnprocessableEntity, "", err)
		return
	}

	_, err := a.repository.GetUser(ctx, userId)
	if errors.Is(err, ErrUserNotFound) {
		out.Response = resp.NewResponse(http.StatusNotFound, "", err)
		return
	}
	if err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	userLoan, err := a.repository.GetUserLoan(ctx, loanId, userId)
	if errors.Is(err, ErrUserLoanNotFound) {
		out.Response = resp.NewResponse(http.StatusNotFound, "", err)
		return
	}
	if err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	if userLoan.Status != Wait.String() {
		out.Response = resp.NewResponse(http.StatusBadRequest, "", ErrModifyProcessLoan)
		return
	}

	var fileUrl string
	if in.IdCard.File != nil {
		var err error
		fileUrl, err = a.saveFile(in.IdCard.Filename, in.IdCard.File)
		if err != nil {
			out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
			return
		}
	}

	userLoan.IsPrivateField = in.IsPrivateField
	userLoan.ExpInYear = in.ExpInYear
	userLoan.ActiveFieldNumber = in.ActiveFieldNumber
	userLoan.SowSeedsPerCycle = in.SowSeedsPerCycle
	userLoan.NeededFertilizerPerCycleInKg = in.NeededFertilizerPerCycleInKg
	userLoan.EstimatedYieldInKg = in.EstimatedPriceOfHarvestPerKg
	userLoan.EstimatedPriceOfHarvestPerKg = in.EstimatedPriceOfHarvestPerKg
	userLoan.HarvestCycleInMonths = in.HarvestCycleInMonths
	userLoan.LoanApplicationInIdr = in.LoanApplicationInIdr
	userLoan.BusinessIncomePerMonthInIdr = in.BusinessIncomePerMonthInIdr
	userLoan.BusinessOutcomePerMonthInIdr = in.BusinessOutcomePerMonthInIdr
	userLoan.FullName = in.FullName
	userLoan.BirthDate = in.BirthDate
	userLoan.FullAddress = in.FullAddress
	userLoan.Phone = in.Phone
	userLoan.IdCardUrl = fileUrl
	userLoan.OtherBusiness = in.OtherBusiness

	if err = a.repository.UpdateLoan(ctx, loanId, userLoan); err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	out.Res = UpdateLoanRes{
		Id: loanId,
	}

	return
}

type (
	DeleteLoanRes struct {
		Id string `json:"id"`
	}
	DeleteLoanOut struct {
		resp.Response
		Res DeleteLoanRes
	}
)

func (a *LoanApp) DeleteLoan(ctx context.Context, loanId string, userId string) (out DeleteLoanOut) {
	out.Response = resp.NewResponse(http.StatusOK, "", nil)
	_, err := a.repository.GetUser(ctx, userId)
	if errors.Is(err, ErrUserNotFound) {
		out.Response = resp.NewResponse(http.StatusNotFound, "", err)
		return
	}
	if err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	userLoan, err := a.repository.GetUserLoan(ctx, loanId, userId)
	if errors.Is(err, ErrUserLoanNotFound) {
		out.Response = resp.NewResponse(http.StatusNotFound, "", err)
		return
	}
	if err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	if userLoan.Status != Wait.String() {
		out.Response = resp.NewResponse(http.StatusBadRequest, "", ErrModifyProcessLoan)
		return
	}

	if err = a.repository.RemoveLoan(ctx, loanId); err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	out.Res = DeleteLoanRes{
		Id: loanId,
	}

	return
}

type (
	GetLoanRes struct {
		LoanId          string `json:"loan_id"`
		UserId          string `json:"user_id"`
		FullName        string `json:"full_name"`
		LoanStatus      string `json:"loan_status"`
		LoanCreatedDate string `json:"loan_created_date"`
	}
	GetLoanOut struct {
		resp.Response
		Res []GetUserLoanRes
	}
)

func (a *LoanApp) GetLoans(ctx context.Context) (out GetUserLoanOut) {
	out.Response = resp.NewResponse(http.StatusOK, "", nil)

	userLoans, err := a.repository.GetLoans(ctx)
	if err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	res := make([]GetUserLoanRes, 0, 0)
	for _, loan := range userLoans {
		res = append(res, GetUserLoanRes{
			LoanId:          loan.Id,
			UserId:          loan.UserId,
			FullName:        loan.FullName,
			LoanStatus:      loan.Status,
			LoanCreatedDate: loan.CreatedDate.Format("2006-01-02"),
		})
	}

	out.Res = res

	return
}

type (
	GetLoanDetailRes struct {
		IsPrivateField               bool   `json:"is_private_field"`
		ExpInYear                    int64  `json:"exp_in_year"`
		ActiveFieldNumber            int64  `json:"active_field_number"`
		SowSeedsPerCycle             int64  `json:"sow_seeds_per_cycle"`
		NeededFertilizerPerCycleInKg int64  `json:"needed_fertilizer_per_cycle_in_kg"`
		EstimatedYieldInKg           int64  `json:"estimated_yield_in_kg"`
		EstimatedPriceOfHarvestPerKg int64  `json:"estimated_price_of_harvest_per_kg"`
		HarvestCycleInMonths         int64  `json:"harvest_cycle_in_months"`
		LoanApplicationInIdr         int64  `json:"loan_application_in_idr"`
		BusinessIncomePerMonthInIdr  int64  `json:"business_income_per_month_in_idr"`
		BusinessOutcomePerMonthInIdr int64  `json:"business_outcome_per_month_in_idr"`
		LoanId                       string `json:"loan_id"`
		UserId                       string `json:"user_id"`
		FullName                     string `json:"full_name"`
		BirthDate                    string `json:"birth_date"`
		FullAddress                  string `json:"full_address"`
		Phone                        string `json:"phone"`
		OtherBusiness                string `json:"other_business"`
		IdCardUrl                    string `json:"id_card_url"`
		Status                       string `json:"status"`
	}
	GetLoanDetailOut struct {
		resp.Response
		Res GetLoanDetailRes
	}
)

func (a *LoanApp) GetLoanDetail(ctx context.Context, loanId string) (out GetLoanDetailOut) {
	out.Response = resp.NewResponse(http.StatusOK, "", nil)

	userLoan, err := a.repository.GetLoan(ctx, loanId)
	if errors.Is(err, ErrUserLoanNotFound) {
		out.Response = resp.NewResponse(http.StatusNotFound, "", err)
		return
	}
	if err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	out.Res = GetLoanDetailRes{
		IsPrivateField:               userLoan.IsPrivateField,
		ExpInYear:                    userLoan.ExpInYear,
		ActiveFieldNumber:            userLoan.ActiveFieldNumber,
		SowSeedsPerCycle:             userLoan.SowSeedsPerCycle,
		NeededFertilizerPerCycleInKg: userLoan.NeededFertilizerPerCycleInKg,
		EstimatedYieldInKg:           userLoan.EstimatedPriceOfHarvestPerKg,
		EstimatedPriceOfHarvestPerKg: userLoan.EstimatedPriceOfHarvestPerKg,
		HarvestCycleInMonths:         userLoan.HarvestCycleInMonths,
		LoanApplicationInIdr:         userLoan.LoanApplicationInIdr,
		BusinessIncomePerMonthInIdr:  userLoan.BusinessIncomePerMonthInIdr,
		BusinessOutcomePerMonthInIdr: userLoan.BusinessOutcomePerMonthInIdr,
		LoanId:                       userLoan.Id,
		UserId:                       userLoan.UserId,
		FullName:                     userLoan.FullName,
		BirthDate:                    userLoan.BirthDate,
		FullAddress:                  userLoan.FullAddress,
		Phone:                        userLoan.Phone,
		OtherBusiness:                userLoan.OtherBusiness,
		IdCardUrl:                    userLoan.IdCardUrl,
		Status:                       userLoan.Status,
	}

	return
}

type (
	ProceedLoanRes struct {
		Id string `json:"id"`
	}
	ProceedLoanOut struct {
		resp.Response
		Res ProceedLoanRes
	}
)

func (a *LoanApp) ProceedLoan(ctx context.Context, loanId, userId string) (out ProceedLoanOut) {
	out.Response = resp.NewResponse(http.StatusOK, "", nil)

	user, err := a.repository.GetUser(ctx, userId)
	if errors.Is(err, ErrUserNotFound) {
		out.Response = resp.NewResponse(http.StatusNotFound, "", err)
		return
	}
	if err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	if !user.IsOfficer {
		out.Response = resp.NewResponse(http.StatusForbidden, "", ErrUserForbidden)
		return
	}

	userLoan, err := a.repository.GetLoan(ctx, loanId)
	if errors.Is(err, ErrUserLoanNotFound) {
		out.Response = resp.NewResponse(http.StatusNotFound, "", err)
		return
	}
	if err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	if userLoan.Status != Wait.String() {
		out.Response = resp.NewResponse(http.StatusBadRequest, "", ErrModifyProcessLoan)
		return
	}

	userLoan.Status = Process.String()
	userLoan.OfficerId.Scan(userId)
	if err = a.repository.UpdateLoan(ctx, loanId, userLoan); err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	out.Res = ProceedLoanRes{
		Id: loanId,
	}

	return
}

type (
	ApproveLoanIn struct {
		IsApprove bool `json:"is_approve"`
	}
	ApproveLoanRes struct {
		Id string `json:"id"`
	}
	ApproveLoanOut struct {
		resp.Response
		Res ApproveLoanRes
	}
)

func (a *LoanApp) ApproveLoan(ctx context.Context, loanId, userId string, in ApproveLoanIn) (out ApproveLoanOut) {
	out.Response = resp.NewResponse(http.StatusOK, "", nil)

	user, err := a.repository.GetUser(ctx, userId)
	if errors.Is(err, ErrUserNotFound) {
		out.Response = resp.NewResponse(http.StatusNotFound, "", err)
		return
	}
	if err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	if !user.IsOfficer {
		out.Response = resp.NewResponse(http.StatusForbidden, "", ErrUserForbidden)
		return
	}

	userLoan, err := a.repository.GetLoan(ctx, loanId)
	if errors.Is(err, ErrUserLoanNotFound) {
		out.Response = resp.NewResponse(http.StatusNotFound, "", err)
		return
	}
	if err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	if userLoan.Status != Process.String() {
		out.Response = resp.NewResponse(http.StatusBadRequest, "", ErrModifyProcessLoan)
		return
	}

	userLoan.OfficerId.Scan(userId)

	userLoan.Status = Reject.String()
	if in.IsApprove {
		userLoan.Status = Approve.String()
	}

	if err = a.repository.UpdateLoan(ctx, loanId, userLoan); err != nil {
		out.Response = resp.NewResponse(http.StatusInternalServerError, "", err)
		return
	}

	out.Res = ApproveLoanRes{
		Id: loanId,
	}

	return
}
