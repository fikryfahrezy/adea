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
	OtherBussiness               string
	BusinessIncomePerMonthInIdr  int64
	BusinessOutcomePerMonthInIdr int64
	Id                           string
	UserId                       string
	OfficerId                    string
	FullName                     string
	BirtDate                     string
	FullAddress                  string
	Phone                        string
	IdCardUrl                    string
	Status                       string
	CreatedDate                  time.Time
	UpdatedDate                  time.Time
}
