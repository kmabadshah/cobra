package tests

import (
	"cobra/funcs"
	"reflect"
	"testing"
)

func TestStrListToInt(t *testing.T) {
	tests := []struct{
		Name string
		Input []string
		Output []int
	} {
		{
			"3 string list as input",
			[]string{"1", "2", "3"},
			[]int{1, 2, 3},
		},
		
		{
			"4 string list as input",
			[]string{"1", "2", "3", "4", "5"},
			[]int{1, 2, 3, 4, 5},
		},
	}
	
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got, err := funcs.StrListToInt(test.Input)
			if err != nil { t.Fatalf("Expected no error but got: %#v", err) }
			
			want := test.Output
			
			if !reflect.DeepEqual(got, want) {
				t.Errorf("WANT %#v, GOT %#v", want, got)
			}
		})
	}
	
	t.Run("Invalid strings as input", func(t *testing.T) {
		input := []string{"1", "2", "3", "4e", "5"}
		_, err := funcs.StrListToInt(input)
		
		if err == nil {
			t.Errorf("wanted error but got nil")
		}
	})
}
