package main_test

import (
	"fmt"
	"math/big"
	"testing"
)

func TestNonDivisibleSubset(t *testing.T) {
	f := func(k int32, s []int32) int32 {
		min := func(a, b int32) int32 {
			if a < b {
				return a
			}

			return b
		}

		max := func(a, b int32) int32 {
			if a > b {
				return a
			}

			return b
		}

		count := make([]int32, k)
		for _, e := range s {
			count[e%k]++
		}

		a := min(count[0], 1)
		if k%2 == 0 {
			a = a + min(count[k/2], 1)
		}

		for i := int32(1); i < k/2+1; i++ {
			if i*2 == k {
				continue
			}

			a = a + max(count[i], count[k-i])
		}

		return a
	}

	cases := []struct {
		k    int32
		s    []int32
		want int32
	}{
		{3, []int32{1, 7, 2, 4}, 3},
	}

	for _, c := range cases {
		got := f(c.k, c.s)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

func bigF(n *big.Int, k int64) *big.Int {
	if k == 1 {
		return n
	}

	return bigF(n.Mul(n, big.NewInt(k-1)), k-1)
}

func TestExtraLongFactorials(t *testing.T) {
	f := func(n int32) string {
		return fmt.Sprintf("%s", bigF(big.NewInt(int64(n)), int64(n)))
	}

	cases := []struct {
		in   int32
		want string
	}{
		{25, "15511210043330985984000000"},
	}

	for _, c := range cases {
		got := f(c.in)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}
