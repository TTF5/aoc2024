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
