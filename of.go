package combinations

import (
	"errors"
	"math/bits"
)

// String returns combinations of n values out of the r options.
// Pass n < 0 to return all possible combinations (including no items).
// The maximum length of r is 63.
func Strings(n int, r []string, f func(combination []string) (stop bool)) {
	values := make([]int, len(r))
	for i := 0; i < len(values); i++ {
		values[i] = i
	}
	N(n, values, func(permutation []int) (stop bool) {
		op := make([]string, len(permutation))
		for i := 0; i < len(permutation); i++ {
			op[i] = r[permutation[i]]
		}
		return f(op)
	})
}

func hasBit(n uint, pos uint) bool {
	return (n & (1 << pos)) > 0
}

func pow(a, b uint64) uint64 {
	p := uint64(1)
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

var ErrOverflow = errors.New("combinations: the maximum length of r is 63")

// N returns combinations of n values out of the r options.
// Pass n < 0 to return all possible combinations (including no items).
// The maximum length of r is 63.
func N(n int, r []int, f func(combination []int) (stop bool)) error {
	if len(r) > 63 {
		return ErrOverflow
	}
	max := pow(2, uint64(len(r)))
	for c := uint64(0); c < max; c++ {
		onesCount := bits.OnesCount(uint(c))
		if n < 0 || onesCount == n {
			combination := make([]int, onesCount)
			var j int
			for i := 0; i < len(r); i++ {
				if hasBit(uint(c), uint(i)) {
					combination[j] = r[i]
					j++
				}
			}
			if stop := f(combination); stop {
				return nil
			}
		}
	}
	return nil
}
