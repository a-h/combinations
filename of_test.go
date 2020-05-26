package combinations

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOfInts(t *testing.T) {
	var tests = []struct {
		n        int
		r        []int
		expected [][]int
	}{
		{
			n: 1,
			r: []int{1, 2, 3},
			expected: [][]int{
				{1},
				{2},
				{3},
			},
		},
		{
			n: 2,
			r: []int{1, 2, 3, 4},
			expected: [][]int{
				{1, 2},
				{1, 3},
				{2, 3},
				{1, 4},
				{2, 4},
				{3, 4},
			},
		},
		{
			n: -1,
			r: []int{1, 2, 3},
			expected: [][]int{
				{},
				{1},
				{2},
				{1, 2},
				{3},
				{1, 3},
				{2, 3},
				{1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprint(tt.n), func(t *testing.T) {
			var actual [][]int
			OfInts(tt.n, tt.r, func(combination []int) (stop bool) {
				actual = append(actual, combination)
				return false
			})
			if diff := cmp.Diff(actual, tt.expected); diff != "" {
				t.Errorf("(%d of %v): expected %v, actual %v", tt.n, tt.r, tt.expected, actual)
				t.Errorf(diff)
			}
		})
	}
}

func TestOfStrings(t *testing.T) {
	var tests = []struct {
		n        int
		r        []string
		expected [][]string
	}{
		{
			n: 1,
			r: []string{"a", "b", "c"},
			expected: [][]string{
				{"a"},
				{"b"},
				{"c"},
			},
		},
		{
			n: 2,
			r: []string{"a", "b", "c"},
			expected: [][]string{
				{"a", "b"},
				{"a", "c"},
				{"b", "c"},
			},
		},
		{
			n: 3,
			r: []string{"a", "b", "c"},
			expected: [][]string{
				{"a", "b", "c"},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%d of %v", tt.n, tt.r), func(t *testing.T) {
			var actual [][]string
			OfStrings(tt.n, tt.r, func(combination []string) (stop bool) {
				actual = append(actual, combination)
				return false
			})
			if diff := cmp.Diff(actual, tt.expected); diff != "" {
				t.Errorf("expected %v, actual %v", tt.expected, actual)
				t.Errorf(diff)
			}
		})
	}
}

func TestStop(t *testing.T) {
	var count int
	All(-1, 8, func(combination []int) (stop bool) {
		count++
		if count == 4 {
			return true
		}
		return false
	})
	if count != 4 {
		t.Errorf("expected to be able to stop, but all combinations were processed")
	}
}

func TestBigNumbers(t *testing.T) {
	count := new(big.Int)
	one := big.NewInt(1)
	oneMillion := big.NewInt(1000000)
	All(-1, 64, func(combination []int) (stop bool) {
		count = count.Add(count, one)
		if result := count.Mod(count, oneMillion); result == big.NewInt(0) {
			fmt.Println(count)
		}
		return false
	})
	expected := new(big.Int).Exp(big.NewInt(2), big.NewInt(64), nil)
	if expected.Cmp(count) != 0 {
		t.Errorf("expected %v, got %v", expected, count)
	}
}
