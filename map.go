package piscine

func Map(f interface{}, a interface{}) interface{} {
	switch fn := f.(type) {
	case func(int) int:
		arr := a.([]int)
		res := make([]int, len(arr))
		for i, v := range arr {
			res[i] = fn(v)
		}
		return res
	case func(string) string:
		arr := a.([]string)
		res := make([]string, len(arr))
		for i, v := range arr {
			res[i] = fn(v)
		}
		return res
	case func(int) bool:
		arr := a.([]int)
		res := make([]bool, len(arr))
		for i, v := range arr {
			res[i] = fn(v)
		}
		return res
	}
	return nil
}
