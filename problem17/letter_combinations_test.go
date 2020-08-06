package problem17

import (
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	r1 := letterCombinations("")
	if len(r1) != 0 {
		t.Errorf("expect [] but received %v\n", r1)
	}
	r2 := letterCombinations("23")
	t.Logf("%v\n", r2)
	if len(r2) != 9 {
		t.Error(r2)
	}
}
