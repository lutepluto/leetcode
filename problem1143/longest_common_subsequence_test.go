package problem1143

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	l1 := longestCommonSubsequence("abcde", "ace")
	if l1 != 3 {
		t.Errorf("expect value 3 but %v returned", l1)
	}
	l2 := longestCommonSubsequence("abc", "abc")
	if l2 != 3 {
		t.Errorf("expect value 3 but %v returned", l2)
	}
	l3 := longestCommonSubsequence("abc", "def")
	if l3 != 0 {
		t.Errorf("expect value 0 but %v returned", l3)
	}
}

func TestLcs(t *testing.T) {
	l1 := lcs("abcde", "ace")
	if l1 != 3 {
		t.Errorf("expect value 3 but %v returned", l1)
	}
	l2 := lcs("abc", "abc")
	if l2 != 3 {
		t.Errorf("expect value 3 but %v returned", l2)
	}
	l3 := lcs("abc", "def")
	if l3 != 0 {
		t.Errorf("expect value 0 but %v returned", l3)
	}
}
