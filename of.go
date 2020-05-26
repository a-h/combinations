package combinations

import (
	"errors"
	"math/bits"
)

// OfStrings returns combinations of n values out of the r options.
// Pass n < 0 to return all possible combinations (including no items).
// If len(r) > 63, this function will panic, because that number of combinations
// would overflow an uint64.
func OfStrings(n int, r []string, f func(combination []string) (stop bool)) {
	All(n, len(r), func(permutation []int) (stop bool) {
		op := make([]string, len(permutation))
		for i := 0; i < len(permutation); i++ {
			op[i] = r[permutation[i]]
		}
		return f(op)
	})
}

// All returns combinations of r values.
// Pass n < 0 to return all possible combinations (including no items).
// If len(r) > 63, this function will panic, because that number of combinations
// would overflow an uint64.
func All(n int, r int, f func(combination []int) (stop bool)) {
	values := make([]int, r)
	for i := 0; i < r; i++ {
		values[i] = i
	}
	OfInts(n, values, f)
}

func pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

var ErrCombinationOverflow = errors.New("combinations: r must be < 64")

// OfInts returns combinations of n values out of the r options.
// Pass n < 0 to return all possible combinations (including no items).
// If len(r) > 63, this function will panic, because that number of combinations
// would overflow an uint64.
func OfInts(n int, r []int, f func(combination []int) (stop bool)) {
	if len(r) > 63 {
		panic(ErrCombinationOverflow)
	}
	max := pow(2, len(r))
	for c := uint(0); c < uint(max); c++ {
		size := bits.OnesCount(uint(c))
		if n < 0 || size == n {
			combination := make([]int, size)
			var j int
			for i, rr := range r {
				if (c & (1 << i)) > 0 {
					combination[j] = rr
					j++
				}
			}
			if stop := f(combination); stop {
				return
			}
		}
	}
	return
}
