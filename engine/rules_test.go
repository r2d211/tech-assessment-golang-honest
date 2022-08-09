package engine

import (
	"strconv"
	"testing"
)

func TestPoliticallyExposed(t *testing.T) {

	exposedRecord := RecordData{
		PoliticallyExposed: true,
	}
	got := PoliticallyExposedCheck(exposedRecord)
	want := true
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

}

func TestAreaCodeCheck(t *testing.T) {

	allowed := [4]int{0, 2, 5, 8}
	disallowed := [6]int{1, 3, 4, 6, 7, 9}
	for _, v := range allowed {
		num := strconv.Itoa(v) + "12345678"
		record := RecordData{
			PhoneNumber: num,
		}
		got := AreaCodeCheck(record)
		want := true

		if got != want {
			t.Errorf(" for num %s got %t, wanted %t", num, got, want)
		}

	}
	for _, v := range disallowed {
		num := strconv.Itoa(v) + "12345678"
		record := RecordData{
			PhoneNumber: num,
		}
		got := AreaCodeCheck(record)
		want := false

		if got != want {
			t.Errorf(" for num %s got %t, wanted %t", num, got, want)
		}

	}
}

func TestIncomeCheck(t *testing.T) {

	underIncomeRecord := RecordData{
		Income: 10000,
	}
	got := IncomeCheck(underIncomeRecord)
	want := false
	if got != want {
		t.Errorf("for income %d got %t, wanted %t", underIncomeRecord.Income, got, want)
	}

	healthyIncomeRecord := RecordData{
		Income: 100001,
	}

	got = IncomeCheck(healthyIncomeRecord)
	want = true
	if got != want {
		t.Errorf("for income %d got %t, wanted %t", healthyIncomeRecord.Income, got, want)
	}
}

func TestCreditCardNumCheck(t *testing.T) {

	LowCreditCardRiskRecords := []RecordData{
		{
			NumberOfCreditCards: 1,
			Age:                 23,
		},
		{
			NumberOfCreditCards: 2,
			Age:                 52,
		},
		{
			NumberOfCreditCards: 3,
			Age:                 48,
		},
	}
	for _, record := range LowCreditCardRiskRecords {
		got := CreditCardNumCheck(record)
		want := true
		if got != want {
			t.Errorf("for age %v got %t, wanted %t", record, got, want)
		}
	}

	HighCreditCardRiskRecords := []RecordData{
		{
			NumberOfCreditCards: 1,
			Age:                 22,
		},
		{
			NumberOfCreditCards: 4,
			Age:                 50,
		},
		{
			NumberOfCreditCards: 3,
			Age:                 47,
		},
	}
	for _, record := range HighCreditCardRiskRecords {
		got := CreditCardNumCheck(record)
		want := false
		if got != want {
			t.Errorf("for age %v got %t, wanted %t", record, got, want)
		}
	}
}

func TestAgeCheck(t *testing.T) {

	underAgeRecord := RecordData{
		Age: 10,
	}
	got := AgeCheck(underAgeRecord)
	want := false
	if got != want {
		t.Errorf("for age %d got %t, wanted %t", underAgeRecord.Age, got, want)
	}

	healthyAgeRecord := RecordData{
		Age: 18,
	}

	got = AgeCheck(healthyAgeRecord)
	want = true
	if got != want {
		t.Errorf("for age %d got %t, wanted %t", healthyAgeRecord.Age, got, want)
	}
}
