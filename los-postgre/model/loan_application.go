package model

import "time"

type LoanApplication struct {
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
	Id                           string
	UserId                       string
	OfficerId                    string
	FullName                     string
	BirthDate                    string
	FullAddress                  string
	Phone                        string
	IdCardUrl                    string
	OtherBussiness               string
	Status                       string
	CreatedDate                  time.Time
	UpdatedDate                  time.Time
}
