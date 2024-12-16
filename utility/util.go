package utility

type Point struct {
	X, Y int
}

func IAbs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func SliceMap[Src any, Dest any](s []Src, f func(Src) Dest) []Dest {
	result := make([]Dest, len(s))
	for idx, val := range s {
		result[idx] = f(val)
	}
	return result
}

func GenerateCombinations[Value any](alphabet []Value, length int) <-chan []Value {
	ch := make(chan []Value)
	go func() {
		defer close(ch)
		if length == 0 {
			ch <- make([]Value, 0)
		} else {
			for _, v := range alphabet {
				for comb := range GenerateCombinations(alphabet, length-1) {
					ch <- append([]Value{v}, comb...)
				}
			}
		}
	}()
	return ch
}
