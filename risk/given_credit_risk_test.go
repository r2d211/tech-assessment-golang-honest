package risk

import (
	"testing"
)

func TestCalculateCreditRisk(t *testing.T) {
	age := 18
	numberOfCC := 0
	got := CalculateCreditRisk(age, numberOfCC)
	want := "LOW"
	if got != want {
		t.Errorf("for age %d numberOfCC %d got %s, wanted %s", age, numberOfCC, got, want)
	}

	age = 19
	numberOfCC = 1
	got = CalculateCreditRisk(age, numberOfCC)
	want = "HIGH"
	if got != want {
		t.Errorf("for age %d numberOfCC %d got %s, wanted %s", age, numberOfCC, got, want)
	}

	age = 20
	numberOfCC = 2
	got = CalculateCreditRisk(age, numberOfCC)
	want = "MEDIUM"
	if got != want {
		t.Errorf("for age %d numberOfCC %d got %s, wanted %s", age, numberOfCC, got, want)
	}
}
