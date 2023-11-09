package catchcalculator

import "testing"

func TestCalculateCatchChance(t *testing.T) {
	expected := 99
	testVal := 36

	actual := calculateCatchChance(testVal)

	if actual != expected {
		t.Errorf("expected base experience 36 to calculate 99 catch chance, got %d", actual)
	}

	expected = 1
	testVal = 608

	actual = calculateCatchChance(testVal)

	if actual != expected {
		t.Errorf("expected base experience 608 to calculate catch chance of 1, got %d", actual)
	}
}

func TestEvaluateCatchOutcome(t *testing.T) {
	//perform catch outcome on base experience 36 100 times, require that 75 of outcomes are true.. well below the 99 that is expected
	baseXp := 36
	count := 0
	for i := 0; i < 100; i++ {
		res := EvaluateCatchOutcome(baseXp)
		if res {
			count++
		}
	}

	if count < 75 {
		t.Errorf("Expected base xp of 36 to be true at least 75 times, got %d", count)
	}

	//perform catch outcome on base experience of 608 100 times, require that 75 of outcomes are false.. well below the 99 that is expected
	baseXp = 608
	count = 0
	for i := 0; i < 100; i++ {
		res := EvaluateCatchOutcome(baseXp)
		if !res {
			count++
		}
	}

	if count < 75 {
		t.Errorf("Expected base xp of 608 to be false at least 75 times, got %d", count)
	}
}
