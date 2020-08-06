package problem22

import "testing"

func TestGenerateParenthesis(t *testing.T) {
	t.Logf("%v\n", generateParenthesis(0))
	t.Logf("%v\n", generateParenthesis(1))
	t.Logf("%v\n", generateParenthesis(2))
	t.Logf("%v\n", generateParenthesis(3))
}
