package funcs

type Calc struct {}

func (c Calc) Add(nums ...int) int {
	res := 0
	
	for _, num := range nums {
		res += num
	}
	
	return res
}


func (c Calc) Subtract(nums ...int) int {
	res := 0
	
	for i, num := range nums {
		if i == 0 { res = num } else { res -= num }
	}
	
	return res
}

