package finding

import (
	"encoding/json"
	"testing"
)

func TestCreateFinding(t *testing.T) {
	want := Finding{
		Message:  "Samplemessage",
		Location: [2]int{0, 10},
		Severity: -1,
	}

	find := createFinding(
		"Samplemessage",
		[2]int{0, 10},
		-1,
	)
	findJSON, _ := json.Marshal(find)
	wantJSON, _ := json.Marshal(want)
	if find != want {
		t.Errorf(
			"Got: %v, Wanted %v",
			string(findJSON),
			string(wantJSON),
		)
	}

}

func TestAdjustLocationPlusFive(t *testing.T) {
	want := createFinding(
		"Samplemessage shifted by five",
		[2]int{7, 25},
		0,
	)

	find := createFinding(
		"Samplemessage shifted by five",
		[2]int{2, 20},
		0,
	)
	find.AdjustLocation(5) // here we test moving the finding by five

	findJSON, _ := json.Marshal(find)
	wantJSON, _ := json.Marshal(want)
	if find != want {
		t.Errorf(
			"Got: %v, Wanted %v",
			string(findJSON),
			string(wantJSON),
		)
	}
}

func TestAdjustLocationMinusFive(t *testing.T) {
	want := createFinding(
		"Samplemessage shifted by five",
		[2]int{2, 20},
		0,
	)

	find := createFinding(
		"Samplemessage shifted by five",
		[2]int{7, 25},
		0,
	)
	find.AdjustLocation(-5) // here we test moving the finding by five

	findJSON, _ := json.Marshal(find)
	wantJSON, _ := json.Marshal(want)
	if find != want {
		t.Errorf(
			"Got: %v, Wanted %v",
			string(findJSON),
			string(wantJSON),
		)
	}
}

func TestAdjustLocations(t *testing.T) {
	want := []Finding{
		createFinding(
			"Samplemessage1",
			[2]int{7, 25},
			0,
		),
		createFinding(
			"Samplemessage2",
			[2]int{60, 64},
			1,
		),
		createFinding(
			"Samplemessage3",
			[2]int{107, 116},
			2,
		),
	}

	find := []Finding{
		createFinding(
			"Samplemessage1",
			[2]int{2, 20},
			0,
		),
		createFinding(
			"Samplemessage2",
			[2]int{55, 59},
			1,
		),
		createFinding(
			"Samplemessage3",
			[2]int{102, 111},
			2,
		),
	}

	find = AdjustLocations(find, 5) // here we test moving the finding by five

	findJSON, _ := json.Marshal(find)
	wantJSON, _ := json.Marshal(want)
	if string(findJSON) != string(wantJSON) {
		t.Errorf(
			"Got: %v, Wanted %v",
			string(findJSON),
			string(wantJSON),
		)
	}
}
