package combinations

import (
	"math/big"
)

// OfStrings returns combinations of n values out of the r options.
// Pass n < 0 to return all possible combinations (including no items).
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
func All(n int, r int, f func(combination []int) (stop bool)) {
	values := make([]int, r)
	for i := 0; i < r; i++ {
		values[i] = i
	}
	OfInts(n, values, f)
}

// OfInts returns combinations of n values out of the r options.
// Pass n < 0 to return all possible combinations (including no items).
func OfInts(n int, r []int, f func(combination []int) (stop bool)) {
	max := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(len(r))), nil)
	one := big.NewInt(1)
	for c := new(big.Int); c.Cmp(max) == -1; c = c.Add(c, one) {
		onesCount := onesCount(c)
		if n < 0 || onesCount == n {
			combination := make([]int, onesCount)
			var j int
			for i := 0; i < len(r); i++ {
				if c.Bit(i) == 1 {
					combination[j] = r[i]
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

func onesCount(v *big.Int) (op int) {
	for i := 0; i < v.BitLen(); i++ {
		op += int(v.Bit(i))
	}
	return op
}
