package funcs

import (
	"errors"
	"strconv"
)

func StrListToInt(strings []string) ([]int, error) {
	result := []int{}
	
	for _, str := range strings {
		int, err := strconv.Atoi(str)
		if err != nil { return nil, errors.New("invalid argument(s)") }
		result = append(result, int)
	}
	
	return result, nil
}