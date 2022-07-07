package loan

import (
	"errors"
	"strconv"
	"time"
	"unicode/utf8"
)

var (
	ErrFullNameRequired         = errors.New("full name required")
	ErrBirthDateRequired        = errors.New("birth date required")
	ErrBirthDateNotValidDate    = errors.New("birth date not valid date")
	ErrFullAddressRequired      = errors.New("full address required")
	ErrPhoneRequired            = errors.New("phone required")
	ErrPhoneMin10               = errors.New("phone min 10 characters")
	ErrPhoneMax15               = errors.New("phone max 15 characters")
	ErrPhoneNotNumbers          = errors.New("phone should only contain numbers")
	ErrIdCardRequired           = errors.New("id card image required")
	ErrExpInYearRequired        = errors.New("experieance year required")
	ErrExpInYearLtZero          = errors.New("experieance should greater than zero")
	ErrActiveFieldRequired      = errors.New("active fields required")
	ErrActiveFieldLtZero        = errors.New("active should greater than zero")
	ErrSowSeedsPerCycleRequired = errors.New("sow seends per cycle required")
	ErrSowSeedsPerCycleLtZero   = errors.New("sow seends per cycle should greater than zero")
	ErrNeededFertilizerRequired = errors.New("needed fertilizer per cycle in kg required")
	ErrNeededFertilizerLtZero   = errors.New("needed fertilizer per cycle in kg shoud greater than zero")
	ErrEstimatedYieldRequired   = errors.New("estimated yield in kg required")
	ErrEstimatedYieldLtZero     = errors.New("estimated yield in kg should greater than zero")
	ErrEstimatedPriceRequired   = errors.New("estimated price of harvest per kg required")
	ErrEstimatedPriceLtZero     = errors.New("estimated price of harvest per kg should greater than zero")
	ErrHarvestCycleRequired     = errors.New("harvest cycle in monts required")
	ErrHarvestCycleLtZero       = errors.New("harvest cycle in monts should greater than zero")
	ErrLoanIdrRequired          = errors.New("loan application in idr required")
	ErrLoanIdrLtZero            = errors.New("loan application in idr should greater than zero")
	ErrIncomePerMonthRequired   = errors.New("business income per month required")
	ErrIncomePerMonthLtZero     = errors.New("business income per month should greater than zero")
	ErrOutcomePerMonthRequired  = errors.New("business outcome per month required")
	ErrOutcomePerMonthLtZero    = errors.New("business outcome per month should greater than zero")
)

func validateCreateLoan(in CreateLoanIn) error {
	if utf8.RuneCountInString(in.FullName) == 0 {
		return ErrFullNameRequired
	}
	if utf8.RuneCountInString(in.BirthDate) == 0 {
		return ErrBirthDateRequired
	}
	if _, err := time.Parse(in.BirthDate, "2006-01-02"); err != nil {
		return ErrBirthDateNotValidDate
	}
	if utf8.RuneCountInString(in.FullAddress) == 0 {
		return ErrFullAddressRequired
	}
	if utf8.RuneCountInString(in.Phone) == 0 {
		return ErrPhoneRequired
	}
	if utf8.RuneCountInString(in.Phone) < 10 {
		return ErrPhoneMin10
	}
	if utf8.RuneCountInString(in.Phone) > 15 {
		return ErrPhoneMax15
	}
	if _, err := strconv.Atoi(in.Phone); err != nil {
		return ErrPhoneNotNumbers
	}
	if _, err := strconv.Atoi(in.Phone); err != nil {
		return ErrPhoneNotNumbers
	}
	if in.ExpInYear == 0 {
		return ErrExpInYearRequired
	}
	if in.ExpInYear < 0 {
		return ErrExpInYearRequired
	}
	if in.ActiveFieldNumber == 0 {
		return ErrActiveFieldRequired
	}
	if in.ActiveFieldNumber < 0 {
		return ErrActiveFieldLtZero
	}
	if in.SowSeedsPerCycle == 0 {
		return ErrSowSeedsPerCycleRequired
	}
	if in.SowSeedsPerCycle <= 0 {
		return ErrSowSeedsPerCycleLtZero
	}
	if in.NeededFertilizerPerCycleInKg == 0 {
		return ErrNeededFertilizerRequired
	}
	if in.NeededFertilizerPerCycleInKg == 0 {
		return ErrNeededFertilizerRequired
	}
	if in.NeededFertilizerPerCycleInKg <= 0 {
		return ErrNeededFertilizerLtZero
	}
	if in.EstimatedYieldInKg == 0 {
		return ErrEstimatedYieldRequired
	}
	if in.EstimatedYieldInKg <= 0 {
		return ErrEstimatedYieldLtZero
	}
	if in.EstimatedPriceOfHarvestPerKg == 0 {
		return ErrEstimatedPriceRequired
	}
	if in.EstimatedPriceOfHarvestPerKg <= 0 {
		return ErrEstimatedPriceLtZero
	}
	if in.HarvestCycleInMonths == 0 {
		return ErrHarvestCycleRequired
	}
	if in.HarvestCycleInMonths <= 0 {
		return ErrHarvestCycleLtZero
	}
	if in.LoanApplicationInIdr == 0 {
		return ErrLoanIdrRequired
	}
	if in.LoanApplicationInIdr <= 0 {
		return ErrLoanIdrLtZero
	}
	if in.BusinessIncomePerMonthInIdr == 0 {
		return ErrIncomePerMonthRequired
	}
	if in.BusinessIncomePerMonthInIdr <= 0 {
		return ErrIncomePerMonthLtZero
	}
	if in.BusinessOutcomePerMonthInIdr == 0 {
		return ErrOutcomePerMonthRequired
	}
	if in.BusinessOutcomePerMonthInIdr <= 0 {
		return ErrOutcomePerMonthLtZero
	}
	if in.IdCard.File == nil {
		return ErrIdCardRequired
	}

	return nil
}

func validateUpdateLoan(in UpdateLoanIn) error {
	if utf8.RuneCountInString(in.FullName) == 0 {
		return ErrFullNameRequired
	}
	if utf8.RuneCountInString(in.BirthDate) == 0 {
		return ErrBirthDateRequired
	}
	if _, err := time.Parse(in.BirthDate, "2006-01-02"); err != nil {
		return ErrBirthDateNotValidDate
	}
	if utf8.RuneCountInString(in.FullAddress) == 0 {
		return ErrFullAddressRequired
	}
	if utf8.RuneCountInString(in.Phone) == 0 {
		return ErrPhoneRequired
	}
	if utf8.RuneCountInString(in.Phone) < 10 {
		return ErrPhoneMin10
	}
	if utf8.RuneCountInString(in.Phone) > 15 {
		return ErrPhoneMax15
	}
	if _, err := strconv.Atoi(in.Phone); err != nil {
		return ErrPhoneNotNumbers
	}
	if _, err := strconv.Atoi(in.Phone); err != nil {
		return ErrPhoneNotNumbers
	}
	if in.ExpInYear == 0 {
		return ErrExpInYearRequired
	}
	if in.ExpInYear < 0 {
		return ErrExpInYearRequired
	}
	if in.ActiveFieldNumber == 0 {
		return ErrActiveFieldRequired
	}
	if in.ActiveFieldNumber < 0 {
		return ErrActiveFieldLtZero
	}
	if in.SowSeedsPerCycle == 0 {
		return ErrSowSeedsPerCycleRequired
	}
	if in.SowSeedsPerCycle <= 0 {
		return ErrSowSeedsPerCycleLtZero
	}
	if in.NeededFertilizerPerCycleInKg == 0 {
		return ErrNeededFertilizerRequired
	}
	if in.NeededFertilizerPerCycleInKg == 0 {
		return ErrNeededFertilizerRequired
	}
	if in.NeededFertilizerPerCycleInKg <= 0 {
		return ErrNeededFertilizerLtZero
	}
	if in.EstimatedYieldInKg == 0 {
		return ErrEstimatedYieldRequired
	}
	if in.EstimatedYieldInKg <= 0 {
		return ErrEstimatedYieldLtZero
	}
	if in.EstimatedPriceOfHarvestPerKg == 0 {
		return ErrEstimatedPriceRequired
	}
	if in.EstimatedPriceOfHarvestPerKg <= 0 {
		return ErrEstimatedPriceLtZero
	}
	if in.HarvestCycleInMonths == 0 {
		return ErrHarvestCycleRequired
	}
	if in.HarvestCycleInMonths <= 0 {
		return ErrHarvestCycleLtZero
	}
	if in.LoanApplicationInIdr == 0 {
		return ErrLoanIdrRequired
	}
	if in.LoanApplicationInIdr <= 0 {
		return ErrLoanIdrLtZero
	}
	if in.BusinessIncomePerMonthInIdr == 0 {
		return ErrIncomePerMonthRequired
	}
	if in.BusinessIncomePerMonthInIdr <= 0 {
		return ErrIncomePerMonthLtZero
	}
	if in.BusinessOutcomePerMonthInIdr == 0 {
		return ErrOutcomePerMonthRequired
	}
	if in.BusinessOutcomePerMonthInIdr <= 0 {
		return ErrOutcomePerMonthLtZero
	}
	if in.IdCard.File == nil {
		return ErrIdCardRequired
	}

	return nil
}
