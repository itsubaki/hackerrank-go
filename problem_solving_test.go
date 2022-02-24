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

func TestGradingStudents(t *testing.T) {
	var grade = map[int32]int32{
		38: 40, 39: 40, 40: 40,
		41: 41, 42: 42, 43: 45, 44: 45, 45: 45, 46: 46, 47: 47, 48: 50, 49: 50, 50: 50,
		51: 51, 52: 52, 53: 55, 54: 55, 55: 55, 56: 56, 57: 57, 58: 60, 59: 60, 60: 60,
		61: 61, 62: 62, 63: 65, 64: 65, 65: 65, 66: 66, 67: 67, 68: 70, 69: 70, 70: 70,
		71: 71, 72: 72, 73: 75, 74: 75, 75: 75, 76: 76, 77: 77, 78: 80, 79: 80, 80: 80,
		81: 81, 82: 82, 83: 85, 84: 85, 85: 85, 86: 86, 87: 87, 88: 90, 89: 90, 90: 90,
		91: 91, 92: 92, 93: 95, 94: 95, 95: 95, 96: 96, 97: 97, 98: 100, 99: 100, 100: 100,
	}

	f := func(grades []int32) []int32 {
		out := make([]int32, 0)
		for _, g := range grades {
			if g < 38 {
				out = append(out, g)
				continue
			}

			out = append(out, grade[g])
		}

		return out
	}

	cases := []struct {
		in   []int32
		want []int32
	}{
		{[]int32{73, 67, 38, 33}, []int32{75, 67, 40, 33}},
	}

	for _, c := range cases {
		got := f(c.in)
		for i := range got {
			if got[i] == c.want[i] {
				continue
			}

			t.Errorf("want=%v, got=%v", c.want, got)
		}
	}
}

func TestCompareTheTriplets(t *testing.T) {
	f := func(a []int32, b []int32) []int32 {
		out := make([]int32, 2)
		for i := range a {
			if a[i] > b[i] {
				out[0]++
			}

			if a[i] < b[i] {
				out[1]++
			}
		}

		return out
	}

	cases := []struct {
		a, b []int32
		want []int32
	}{
		{[]int32{5, 6, 7}, []int32{3, 6, 10}, []int32{1, 1}},
	}

	for _, c := range cases {
		got := f(c.a, c.b)
		for i := range got {
			if got[i] == c.want[i] {
				continue
			}

			t.Errorf("want=%v, got=%v", c.want, got)
		}
	}
}

func TestStaircase(t *testing.T) {
	f := func(n int32) []string {
		out := make([]string, 0)
		for i := int32(1); i < n+1; i++ {
			var v string
			for j := int32(0); j < n-i; j++ {
				v = v + fmt.Sprintf(" ")
			}

			for j := int32(0); j < i; j++ {
				v = v + fmt.Sprintf("#")
			}

			out = append(out, v)
		}

		return out
	}

	cases := []struct {
		in   int32
		want []string
	}{
		{6, []string{
			"     #",
			"    ##",
			"   ###",
			"  ####",
			" #####",
			"######",
		}},
	}

	for _, c := range cases {
		got := f(c.in)
		for i := range got {
			if got[i] == c.want[i] {
				continue
			}

			t.Errorf("want=%v, got=%v", c.want, got)
		}
	}
}

func TestBirthdayCakeCandles(t *testing.T) {
	f := func(candles []int32) int32 {
		count := make([]int32, 100000000)
		var max int32
		for _, c := range candles {
			if c > max {
				max = c
			}

			count[c]++
		}

		return count[max]
	}

	cases := []struct {
		in   []int32
		want int32
	}{
		{[]int32{4, 4, 1, 3}, 2},
		{[]int32{3, 2, 1, 3}, 2},
	}

	for _, c := range cases {
		got := f(c.in)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

func TestBreakingTheRecords(t *testing.T) {
	f := func(scores []int32) []int32 {
		var maxc, minc int32

		max, min := scores[0], scores[0]
		for _, s := range scores {
			if s > max {
				max = s
				maxc++
			}

			if s < min {
				min = s
				minc++
			}
		}

		return []int32{maxc, minc}
	}

	cases := []struct {
		in   []int32
		want []int32
	}{
		{[]int32{10, 5, 20, 20, 4, 5, 2, 25, 1}, []int32{2, 4}},
		{[]int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}, []int32{4, 0}},
	}

	for _, c := range cases {
		got := f(c.in)
		for i := range got {
			if got[i] == c.want[i] {
				continue
			}

			t.Errorf("want=%v, got=%v", c.want, got)
		}
	}
}

func TestCountingSort2(t *testing.T) {
	f := func(arr []int32) []int32 {
		count := make([]int32, 100)
		for i := range arr {
			count[arr[i]] += 1
		}

		out := make([]int32, 0)
		for i := range count {
			for j := int32(0); j < count[i]; j++ {
				out = append(out, int32(i))
			}
		}

		return out
	}

	cases := []struct {
		in   []int32
		want []int32
	}{}

	for _, c := range cases {
		got := f(c.in)
		for i := range got {
			if got[i] == c.want[i] {
				continue
			}

			t.Errorf("want=%v, got=%v", c.want, got)
		}
	}
}
