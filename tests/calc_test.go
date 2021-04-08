package tests

import (
	"testing"
	
	"cobra/funcs"
)

var calc = funcs.Calc{}

func TestAdd(t *testing.T) {
	tests := []struct{
		Name string
		Input []int
		Output int
	} {
		{
			"Add two numbers",
			[]int{2, 3}, 5,
		},
		
		{
			"Add three numbers",
			[]int{2, 3, 5}, 10,
		},
	}
	
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := calc.Add(test.Input...)
			want := test.Output
			
			assertError(t, got, want)
		})
	}
}


func TestSubtract(t *testing.T) {
	tests := []struct{
		Name string
		Input []int
		Output int
	} {
		{
			"Subtract two numbers",
			[]int{2, 3}, -1,
		},
		
		{
			"Subtract three numbers",
			[]int{2, 3, 5}, -6,
		},
	}
	
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := calc.Subtract(test.Input...)
			want := test.Output
			
			assertError(t, got, want)
		})
	}
}


func assertError(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("WANT %v, GOT %v", want, got)
	}
}