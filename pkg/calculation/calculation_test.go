package calculation

import (
	"errors"
	"testing"
)

type test struct {
	expression string
	answer     float64
	err        error
}

func TestCalc(t *testing.T) {
	cases := []test{
		{"1 +2-3*17 +2*(9+2)", -26, nil},
		{"31-8+9*3", 50, nil},
		{"-7 +5", -2, nil},
		{"0/2", 0, nil},
		{"(2+3) - (7-9)", 7, nil},
		{"1.2 + 2.3", 3.5, nil},
		{"-(-1+2)+7", 6, nil},
		{"1+1", 2, nil},
		{"1+1++", 0, ErrStackOverflow}, // broken
		{"(1+1)/0", 0, ErrDivisionByZero},
		{"1+1*", 0, ErrStackOverflow},
	}
	for _, c := range cases {
		t.Run(c.expression, func(t *testing.T) {
			got, err := Calc(c.expression)
			if !errors.Is(err, c.err) {
				t.Errorf("Calc(%q): err %v; want %v", c.expression, err, c.err)
			}
			if got != c.answer {
				t.Errorf("Calc(%v) = %v; want %v", c.expression, got, c.answer)
			}
		})
	}
}
