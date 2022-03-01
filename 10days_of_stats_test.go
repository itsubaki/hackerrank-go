package main_test

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func TestMeanMedianAndMode(t *testing.T) {
	f := func(n []int32) []string {
		var sum int32
		for i := range n {
			sum = sum + n[i]
		}

		sort.Slice(n, func(i, j int) bool { return n[i] < n[j] })
		median := float64(n[len(n)/2])
		if len(n)%2 == 0 {
			median = float64((n[len(n)/2-1] + n[len(n)/2])) / 2.0
		}

		count, max := 0, 0
		cur, mode := n[0], n[0]
		for i := range n {
			if n[i] == cur {
				count++
			} else {
				count = 1
				cur = n[i]
			}

			if count > max {
				max = count
				mode = n[i]
			}
		}

		return []string{
			fmt.Sprintf("%.1f", float64(sum)/float64(len(n))),
			fmt.Sprintf("%.1f", median),
			fmt.Sprintf("%v", mode),
		}
	}

	cases := []struct {
		in   []int32
		want []string
	}{
		{
			[]int32{64630, 11735, 14216, 99233, 14470, 4978, 73429, 38120, 51135, 67060},
			[]string{"43900.6", "44627.5", "4978"},
		},
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

	// var T int
	// fmt.Scan(&T)
	//
	// input := make([]int32, 0 )
	// for i := 0; i < T; i++ {
	// 	var N int32
	// 	fmt.Scan(&N)
	// 	input = append(input, N)
	// }
	//
	// for _, o := range f(input) {
	// 	fmt.Println(o)
	// }
}

func TestWeightedMean(t *testing.T) {
	f := func(X []int32, W []int32) float32 {
		var sum, sumW int32
		for i := range W {
			sum += X[i] * W[i]
			sumW += W[i]
		}

		return float32(sum) / float32(sumW)
	}

	cases := []struct {
		X, W []int32
		want float32
	}{
		{
			[]int32{10, 40, 30, 50, 20},
			[]int32{1, 2, 3, 4, 5},
			32.0,
		},
	}

	for _, c := range cases {
		got := f(c.X, c.W)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

func TestQuartiles(t *testing.T) {
	f := func(arr []int32) []int32 {
		med := func(n []int32) int32 {
			if len(n)%2 == 0 {
				return (n[len(n)/2] + n[len(n)/2-1]) / 2
			}

			return n[len(n)/2]
		}

		sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
		half := len(arr) / 2
		if len(arr)%2 == 0 {
			return []int32{med(arr[:half]), med(arr), med(arr[half:])}
		}

		return []int32{med(arr[:half]), med(arr), med(arr[half+1:])}
	}

	cases := []struct {
		in   []int32
		want []int32
	}{
		{
			[]int32{9, 5, 7, 1, 3},
			[]int32{2, 5, 8},
		},
		{
			[]int32{1, 3, 5, 7},
			[]int32{2, 4, 6},
		},
		{
			[]int32{3, 7, 8, 5, 12, 14, 21, 13, 18},
			[]int32{6, 12, 16},
		},
		{
			[]int32{3, 7, 8, 5, 12, 14, 21, 15, 18, 14},
			[]int32{7, 13, 15},
		},
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

func TestInterquartileRange(t *testing.T) {
	f := func(values []int32, freqs []int32) float64 {
		med := func(n []int32) float64 {
			if len(n)%2 == 0 {
				return float64(n[len(n)/2]+n[len(n)/2-1]) / 2.0
			}

			return float64(n[len(n)/2])
		}

		arr := make([]int32, 0)
		for i := range values {
			for j := int32(0); j < freqs[i]; j++ {
				arr = append(arr, values[i])
			}
		}

		sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
		half := len(arr) / 2

		if len(arr)%2 == 0 {
			return med(arr[half:]) - med(arr[:half])
		}

		return med(arr[half+1:]) - med(arr[:half])
	}

	cases := []struct {
		v, f []int32
		want float64
	}{
		{
			[]int32{6, 12, 8, 10, 20, 16},
			[]int32{5, 4, 3, 2, 1, 5},
			9.0,
		},
		{
			[]int32{10, 40, 30, 50, 20},
			[]int32{1, 2, 3, 4, 5},
			30.0,
		},
	}

	for _, c := range cases {
		got := f(c.v, c.f)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

func TestStandardDeviation(t *testing.T) {
	f := func(arr []int32) float64 {
		var sum int32
		for _, a := range arr {
			sum = sum + a
		}
		mean := float64(sum) / float64(len(arr))

		var s float64
		for _, a := range arr {
			s = s + math.Pow((float64(a)-mean), 2.0)
		}

		v := s / float64(len(arr))
		return math.Sqrt(v)
	}

	cases := []struct {
		in   []int32
		want float64
	}{
		{
			[]int32{10, 40, 30, 50, 20},
			14.142135623730951,
		},
	}

	for _, c := range cases {
		got := f(c.in)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

func TestBasicProbability(t *testing.T) {
	var count int
	for i := 1; i < 7; i++ {
		for j := 1; j < 7; j++ {
			if i+j < 10 {
				count++
			}
		}
	}

	if count != 30 {
		t.Errorf("want=30, got=%v", count)
	}

	// answer
	// 30/36 -> 5/6
}

func TestMoreDice(t *testing.T) {
	var count int
	for i := 1; i < 7; i++ {
		for j := 1; j < 7; j++ {
			if i == j {
				// the values rolled by each die will be different
				continue
			}

			if i+j == 6 {
				count++
			}
		}
	}

	if count != 4 {
		t.Errorf("want=4/36, got=%v/36", count)
	}

	// answer
	// 4/36 ->1/9
}

func TestCompoundEventProbability(t *testing.T) {
	frac := func(real float64) []int {
		list := make([]int, 0)
		r := real
		for {
			t := math.Trunc(r)
			list = append(list, int(t))

			diff := r - t
			if diff < 1e-3 {
				break
			}

			r = 1.0 / diff
		}

		return list
	}

	conv := func(cfx []int) (int, int, float64) {
		l := len(cfx)
		if l == 1 {
			return cfx[0], 1, float64(cfx[0])
		}

		s, r := 1, cfx[l-1]
		for i := 2; i < l; i++ {
			s, r = r, cfx[l-i]*r+s
		}
		s = s + cfx[0]*r

		return s, r, float64(s) / float64(r)
	}

	xred := 4.0 / 7.0
	yred := 5.0 / 9.0
	zred := 4.0 / 8.0

	p := (xred * yred * (1 - zred)) + (xred * (1 - yred) * zred) + ((1 - xred) * yred * zred)
	n, d, _ := conv(frac(p))

	if n != 17 || d != 42 {
		t.Errorf("want=17/42, got=%v/%v\n", n, d)
	}

	// answer
	// 17/42
}

func TestConditionalProbability(t *testing.T) {
	// answer
	// 1/3
}

func TestCardsOfTheSameSuit(t *testing.T) {
	// answer
	// 12/51
}

func TestDrawingMarbles(t *testing.T) {
	// answer
	// 2/3
}

func comb(n, r int32) float64 {
	a := factorial(n) / (factorial(r) * factorial(n-r))
	return float64(a)
}

func binomial(n int32, p float64, begin, end int32) float64 {
	var a float64
	for i := begin; i < end; i++ {
		a = a + comb(n, i)*math.Pow(p, float64(i))*math.Pow(1-p, float64(n-i))
	}

	return a
}

func TestBinomialDistribution1(t *testing.T) {
	n := int32(6)
	p := 1.09 / (1.09 + 1.0)

	got := fmt.Sprintf("%.3f", binomial(n, p, 3, 7))
	want := "0.696"
	if got != want {
		t.Errorf("want=%v, got=%v", want, got)
	}
}

func TestBinomialDistribution2(t *testing.T) {
	n := int32(10)
	p := 0.12

	got := []string{
		fmt.Sprintf("%.3f", binomial(n, p, 0, 3)),
		fmt.Sprintf("%.3f", binomial(n, p, 2, n+1)),
	}

	want := []string{
		"0.891",
		"0.342",
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("want=%v, got=%v", want, got)
		}
	}
}

func TestGeometricDistribution1(t *testing.T) {
	p := 1.0 / 3.0
	a := (1 - p) * (1 - p) * (1 - p) * (1 - p) * p

	got := fmt.Sprintf("%.3f", a)
	want := "0.066"
	if got != want {
		t.Errorf("want=%v, got=%v", want, got)
	}
}

func TestGeometricDistribution2(t *testing.T) {
	p := 1.0 / 3.0
	a := 1 - ((1 - p) * (1 - p) * (1 - p) * (1 - p) * (1 - p))

	got := fmt.Sprintf("%.3f", a)
	want := "0.868"
	if got != want {
		t.Errorf("want=%v, got=%v", want, got)
	}
}

func poisson(m float64, x int32) float64 {
	return math.Pow(m, float64(x)) * math.Exp(-m) / float64(factorial(x))
}

func TestPoissonDistribution1(t *testing.T) {
	f := func(m float64, x int32) string {
		return fmt.Sprintf("%.3f", poisson(m, x))
	}

	cases := []struct {
		m    float64
		x    int32
		want string
	}{
		{2.5, 5, "0.067"},
	}

	for _, c := range cases {
		got := f(c.m, c.x)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

func TestPoissonDistribution2(t *testing.T) {
	x := 0.88
	y := 1.55

	got := []string{
		fmt.Sprintf("%.3f", 160+40*(x+math.Pow(x, 2.0))),
		fmt.Sprintf("%.3f", 128+40*(y+math.Pow(y, 2.0))),
	}

	want := []string{
		"226.176",
		"286.100",
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("want=%v, got=%v", want, got)
		}
	}
}

// The cumulative distribution function for a function with normal distribution
func normal(m, s, x float64) float64 {
	return 0.5 * (1.0 + math.Erf((x-m)/(s*math.Sqrt2)))
}

func TestNormalDistribution1(t *testing.T) {
	got := []string{
		fmt.Sprintf("%.3f", normal(20, 2, 19.5)),
		fmt.Sprintf("%.3f", normal(20, 2, 22)-normal(20, 2, 20)),
	}

	want := []string{
		"0.401",
		"0.341",
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("want=%v, got=%v", want, got)
		}
	}
}
