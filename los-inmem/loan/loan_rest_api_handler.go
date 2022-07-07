package loan

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fikryfahrezy/adea/los-inmen/resp"
)

func (a *LoanApp) UserLoansGet(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("authorization")
	out := a.GetUserLoans(r.Context(), userId)
	out.HttpJSON(w, resp.NewHttpBody(out.Res))
}

func (a *LoanApp) UserLoanDetailGet(w http.ResponseWriter, r *http.Request) {
	loanId := r.URL.Query().Get("id")
	if loanId == "" {
		http.NotFound(w, r)
		return
	}

	userId := r.Header.Get("authorization")
	out := a.GetUserLoanDetail(r.Context(), loanId, userId)
	out.HttpJSON(w, resp.NewHttpBody(out.Res))
}

func (a *LoanApp) CreateLoanPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(1024); err != nil {
		resp.NewResponse(http.StatusInternalServerError, "", err).HttpJSON(w, nil)
		return
	}

	in := CreateLoanIn{
		FullName:       r.FormValue("full_name"),
		BirthDate:      r.FormValue("birth_date"),
		FullAddress:    r.FormValue("full_address"),
		Phone:          r.FormValue("phone"),
		OtherBussiness: r.FormValue("other_business"),
	}

	in.IsPrivateField, _ = strconv.ParseBool(r.FormValue("is_private_field"))
	in.ExpInYear, _ = strconv.ParseInt(r.FormValue("exp_in_year"), 10, 64)
	in.ActiveFieldNumber, _ = strconv.ParseInt(r.FormValue("active_field_number"), 10, 64)
	in.SowSeedsPerCycle, _ = strconv.ParseInt(r.FormValue("sow_seeds_per_cycle"), 10, 64)
	in.NeededFertilizerPerCycleInKg, _ = strconv.ParseInt(r.FormValue("needed_fertilizer_per_cycle_in_kg"), 10, 64)
	in.EstimatedYieldInKg, _ = strconv.ParseInt(r.FormValue("estimated_yield_in_kg"), 10, 64)
	in.EstimatedPriceOfHarvestPerKg, _ = strconv.ParseInt(r.FormValue("estimated_price_of_harvest_per_kg"), 10, 64)
	in.HarvestCycleInMonths, _ = strconv.ParseInt(r.FormValue("harvest_cycle_in_months"), 10, 64)
	in.LoanApplicationInIdr, _ = strconv.ParseInt(r.FormValue("loan_application_in_idr"), 10, 64)
	in.BusinessIncomePerMonthInIdr, _ = strconv.ParseInt(r.FormValue("business_income_per_month_in_idr"), 10, 64)
	in.BusinessOutcomePerMonthInIdr, _ = strconv.ParseInt(r.FormValue("business_outcome_per_month_in_idr"), 10, 64)

	file, header, err := r.FormFile("id_card")
	if err != nil {
		resp.NewResponse(http.StatusInternalServerError, "", err).HttpJSON(w, nil)
		return
	}

	in.IdCard = FileHeader{
		Filename: header.Filename,
		File:     file,
	}

	userId := r.Header.Get("authorization")
	out := a.CreateLoan(r.Context(), userId, in)
	out.HttpJSON(w, resp.NewHttpBody(out.Res))
}

func (a *LoanApp) UpdateLoanPut(w http.ResponseWriter, r *http.Request) {
	loanId := r.URL.Query().Get("id")
	if loanId == "" {
		http.NotFound(w, r)
		return
	}

	if err := r.ParseMultipartForm(1024); err != nil {
		resp.NewResponse(http.StatusInternalServerError, "", err).HttpJSON(w, nil)
		return
	}

	in := UpdateLoanIn{
		FullName:       r.FormValue("full_name"),
		BirthDate:      r.FormValue("birth_date"),
		FullAddress:    r.FormValue("full_address"),
		Phone:          r.FormValue("phone"),
		OtherBussiness: r.FormValue("other_business"),
	}

	in.IsPrivateField, _ = strconv.ParseBool(r.FormValue("is_private_field"))
	in.ExpInYear, _ = strconv.ParseInt(r.FormValue("exp_in_year"), 10, 64)
	in.ActiveFieldNumber, _ = strconv.ParseInt(r.FormValue("active_field_number"), 10, 64)
	in.SowSeedsPerCycle, _ = strconv.ParseInt(r.FormValue("sow_seeds_per_cycle"), 10, 64)
	in.NeededFertilizerPerCycleInKg, _ = strconv.ParseInt(r.FormValue("needed_fertilizer_per_cycle_in_kg"), 10, 64)
	in.EstimatedYieldInKg, _ = strconv.ParseInt(r.FormValue("estimated_yield_in_kg"), 10, 64)
	in.EstimatedPriceOfHarvestPerKg, _ = strconv.ParseInt(r.FormValue("estimated_price_of_harvest_per_kg"), 10, 64)
	in.HarvestCycleInMonths, _ = strconv.ParseInt(r.FormValue("harvest_cycle_in_months"), 10, 64)
	in.LoanApplicationInIdr, _ = strconv.ParseInt(r.FormValue("loan_application_in_idr"), 10, 64)
	in.BusinessIncomePerMonthInIdr, _ = strconv.ParseInt(r.FormValue("business_income_per_month_in_idr"), 10, 64)
	in.BusinessOutcomePerMonthInIdr, _ = strconv.ParseInt(r.FormValue("business_outcome_per_month_in_idr"), 10, 64)

	file, header, err := r.FormFile("id_card")
	if err != nil {
		resp.NewResponse(http.StatusInternalServerError, "", err).HttpJSON(w, nil)
		return
	}

	in.IdCard = FileHeader{
		Filename: header.Filename,
		File:     file,
	}

	userId := r.Header.Get("authorization")
	out := a.UpdateLoan(r.Context(), loanId, userId, in)
	out.HttpJSON(w, resp.NewHttpBody(out.Res))
}

func (a *LoanApp) UserLoanDelete(w http.ResponseWriter, r *http.Request) {
	loanId := r.URL.Query().Get("id")
	if loanId == "" {
		http.NotFound(w, r)
		return
	}

	userId := r.Header.Get("authorization")
	out := a.DeleteLoan(r.Context(), loanId, userId)
	out.HttpJSON(w, resp.NewHttpBody(out.Res))
}

func (a *LoanApp) LoansGet(w http.ResponseWriter, r *http.Request) {
	out := a.GetLoans(r.Context())
	out.HttpJSON(w, resp.NewHttpBody(out.Res))
}

func (a *LoanApp) LoanDetailGet(w http.ResponseWriter, r *http.Request) {
	loanId := r.URL.Query().Get("id")
	if loanId == "" {
		http.NotFound(w, r)
		return
	}

	out := a.GetLoanDetail(r.Context(), loanId)
	out.HttpJSON(w, resp.NewHttpBody(out.Res))
}

func (a *LoanApp) ProceedLoanPatch(w http.ResponseWriter, r *http.Request) {
	loanId := r.URL.Query().Get("id")
	if loanId == "" {
		http.NotFound(w, r)
		return
	}

	out := a.ProceedLoan(r.Context(), loanId)
	out.HttpJSON(w, resp.NewHttpBody(out.Res))
}

func (a *LoanApp) ApproveLoanPatch(w http.ResponseWriter, r *http.Request) {
	loanId := r.URL.Query().Get("id")
	if loanId == "" {
		http.NotFound(w, r)
		return
	}

	var in ApproveLoanIn
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		resp.NewResponse(http.StatusInternalServerError, "", err).HttpJSON(w, nil)
		return
	}

	out := a.ApproveLoan(r.Context(), loanId, in)
	out.HttpJSON(w, resp.NewHttpBody(out.Res))
}
