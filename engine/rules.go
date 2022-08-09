package engine

import (
	"github.com/honestbank/tech-assignment-backend-engineer/risk"
)

type RecordData struct {
	Income              int    `json:"income"`
	NumberOfCreditCards int    `json:"number_of_credit_cards"`
	Age                 int    `json:"age"`
	PoliticallyExposed  bool   `json:"politically_exposed"`
	JobIndustryCode     string `json:"job_industry_code"`
	PhoneNumber         string `json:"phone_number"`
}

var preApprovedPhoneNumbers = make(map[string]struct{})

var allowedAreaCode = make(map[rune]struct{})

var exists = struct{}{}

var checks []func(RecordData) bool

func init() {
	allowedAreaCode['0'] = exists
	allowedAreaCode['2'] = exists
	allowedAreaCode['5'] = exists
	allowedAreaCode['8'] = exists
	checks = append(checks, IncomeCheck, AgeCheck,
		CreditCardNumCheck, PoliticallyExposedCheck, AreaCodeCheck)

}

func AddApprovedPhoneNumber(number string) {

	preApprovedPhoneNumbers[number] = exists
}

func RemovePreapprovedPhoneNumber(number string) {
	delete(preApprovedPhoneNumbers, number)
}

func CheckApproved(applicationRecord RecordData) bool {
	if PreApprovedCheck(applicationRecord) {
		return true
	} else {
		result := true
		for _, check := range checks {
			result = result && check(applicationRecord)
		}
		return result
	}
}

func IncomeCheck(applicationRecord RecordData) bool {
	if applicationRecord.Income > 100000 {
		return true
	} else {
		return false
	}
}

func AgeCheck(applicationRecord RecordData) bool {
	if applicationRecord.Age >= 18 {
		return true
	} else {
		return false
	}
}

func CreditCardNumCheck(applicationRecord RecordData) bool {
	if applicationRecord.NumberOfCreditCards <= 3 &&
		risk.CalculateCreditRisk(applicationRecord.Age, applicationRecord.NumberOfCreditCards) == "LOW" {
		return true
	} else {
		return false
	}
}

func PoliticallyExposedCheck(applicationRecord RecordData) bool {
	return !applicationRecord.PoliticallyExposed
}

func AreaCodeCheck(applicationRecord RecordData) bool {
	first := rune(applicationRecord.PhoneNumber[0])
	_, check := allowedAreaCode[first]
	return check

}

func PreApprovedCheck(applicationRecord RecordData) bool {
	_, check := preApprovedPhoneNumbers[applicationRecord.PhoneNumber]
	return check
}
