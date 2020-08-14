package problem227

import "testing"

func TestCalculate(t *testing.T) {
	r1 := calculate("3+2*2")
	if r1 != 7 {
		t.Errorf("Expect 7 but received %v\n", r1)
	}
	r2 := calculate(" ")
	if r2 != 0 {
		t.Errorf("Expect 0 but received %v\n", r2)
	}
	r3 := calculate(" 3+5 / 2 ")
	if r3 != 5 {
		t.Errorf("Expect 5 but received %v\n", r3)
	}
	r4 := calculate2("1+2*5/3+6/4*2")
	if r4 != 6 {
		t.Errorf("Expect 6 but received %v\n", r4)
	}
	r5 := calculate("2*3*4")
	if r5 != 24 {
		t.Errorf("Expect 24 but received %v\n", r4)
	}
}
