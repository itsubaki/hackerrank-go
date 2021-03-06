package main_test

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func factorial(n int32) int32 {
	if n < 2 {
		return 1
	}

	return factorial(n-1) * n
}

func comb(n, r int32) int32 {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func binomial(n int32, p float64, begin, end int32) float64 {
	var a float64
	for i := begin; i < end; i++ {
		a = a + float64(comb(n, i))*math.Pow(p, float64(i))*math.Pow(1-p, float64(n-i))
	}

	return a
}

func poisson(lambda float64, k int32) float64 {
	return math.Pow(lambda, float64(k)) * math.Exp(-lambda) / float64(factorial(k))
}

// The cumulative distribution function for a function with normal distribution
func cdf(mean, stddev, x float64) float64 {
	return 0.5 * (1.0 + math.Erf((x-mean)/(stddev*math.Sqrt2)))
}

func TestMeanMedianAndMode(t *testing.T) {
	f := func(n []int32) (float64, float64, int32) {
		var sum int32
		for i := range n {
			sum = sum + n[i]
		}
		mean := float64(sum) / float64(len(n))

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

		return mean, median, mode
	}

	cases := []struct {
		in           []int32
		mean, median float64
		mode         int32
	}{
		{
			[]int32{64630, 11735, 14216, 99233, 14470, 4978, 73429, 38120, 51135, 67060},
			43900.6, 44627.5, 4978,
		},
	}

	for _, c := range cases {
		mean, median, mode := f(c.in)
		if mean != c.mean {
			t.Errorf("want=%v, got=%v", c.mean, mean)
		}
		if median != c.median {
			t.Errorf("want=%v, got=%v", c.median, median)
		}
		if mode != c.mode {
			t.Errorf("want=%v, got=%v", c.mode, mode)
		}
	}
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

func TestNormalDistribution1(t *testing.T) {
	got := []string{
		fmt.Sprintf("%.3f", cdf(20, 2, 19.5)),
		fmt.Sprintf("%.3f", cdf(20, 2, 22)-cdf(20, 2, 20)),
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

func TestNormalDistribution2(t *testing.T) {
	got := []string{
		fmt.Sprintf("%.2f", (1.0-cdf(70, 10, 80))*100),
		fmt.Sprintf("%.2f", (1.0-cdf(70, 10, 60))*100),
		fmt.Sprintf("%.2f", (cdf(70, 10, 60))*100),
	}

	want := []string{
		"15.87",
		"84.13",
		"15.87",
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("want=%v, got=%v", want, got)
		}
	}
}

func TestTheCentralLimitTheorem1(t *testing.T) {
	f := func(x, n, mean, stddev float64) float64 {
		return cdf(mean, stddev/math.Sqrt(n), x/n)
	}

	cases := []struct {
		x, n, mean, stddev float64
		want               float64
	}{
		{9800, 49, 205, 15, 0.00981532862864537},
	}

	for _, c := range cases {
		got := f(c.x, c.n, c.mean, c.stddev)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

func TestTheCentralLimitTheorem2(t *testing.T) {
	f := func(x, n, mu, sigma float64) float64 {
		return cdf(n*mu, sigma*math.Sqrt(n), x)
	}

	cases := []struct {
		x, n, mu, sigma float64
		want            float64
	}{
		{250, 100, 2.4, 2.0, 0.6914624612740131},
	}

	for _, c := range cases {
		got := f(c.x, c.n, c.mu, c.sigma)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

func TestTheCentralLimitTheorem3(t *testing.T) {
	f := func(n, mean, std, interval, z float64) []float64 {
		return []float64{
			mean - (std/math.Sqrt(n))*z,
			mean + (std/math.Sqrt(n))*z,
		}
	}

	cases := []struct {
		n, mean, std, interval, z float64
		want                      []float64
	}{
		{100, 500, 80, 0.95, 1.96, []float64{484.32, 515.68}},
	}

	for _, c := range cases {
		got := f(c.n, c.mean, c.std, c.interval, c.z)
		for i := range got {
			if got[i] == c.want[i] {
				continue
			}

			t.Errorf("want=%v, got=%v", c.want, got)
		}
	}
}

func TestPearsonCorrelationCoefficient1(t *testing.T) {
	f := func(n int, x, y []float64) float64 {
		var xsum, ysum float64
		for i := 0; i < n; i++ {
			xsum = xsum + x[i]
			ysum = ysum + y[i]
		}

		xmean := xsum / float64(n)
		ymean := ysum / float64(n)

		var xm2sum, ym2sum float64
		for i := 0; i < n; i++ {
			xm2sum = xm2sum + math.Pow(x[i]-xmean, 2.0)
			ym2sum = ym2sum + math.Pow(y[i]-ymean, 2.0)
		}

		stdx := math.Sqrt(xm2sum / float64(n))
		stdy := math.Sqrt(ym2sum / float64(n))

		var cov float64
		for i := 0; i < n; i++ {
			cov = cov + (x[i]-xmean)*(y[i]-ymean)
		}

		return cov / (float64(n) * stdx * stdy)
	}

	cases := []struct {
		n    int
		x, y []float64
		want float64
	}{
		{
			10,
			[]float64{10, 9.8, 8, 7.8, 7.7, 7, 6, 5, 4, 2},
			[]float64{200, 44, 32, 24, 22, 17, 15, 12, 8, 4},
			0.6124721937208479,
		},
	}

	for _, c := range cases {
		got := f(c.n, c.x, c.y)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

func TestSpearmansRankCorrelationCoefficient(t *testing.T) {
	rank := func(n int, arr []float64) []int {
		type pair struct {
			index int
			value float64
		}

		var p []pair
		for i := 0; i < n; i++ {
			p = append(p, pair{
				index: i,
				value: arr[i],
			})
		}
		sort.Slice(p, func(i, j int) bool { return p[i].value < p[j].value })

		rnk := 1
		out := make([]int, n)
		for _, e := range p {
			rnk++
			out[e.index] = rnk
		}

		return out
	}

	f := func(n int, x, y []float64) float64 {
		rx := rank(n, x)
		ry := rank(n, y)

		var d int
		for i := 0; i < n; i++ {
			d = d + (rx[i]-ry[i])*(rx[i]-ry[i])
		}

		return 1.0 - (6.0 * float64(d) / float64(n*(n*n-1)))
	}

	cases := []struct {
		n    int
		x, y []float64
		want float64
	}{
		{
			10,
			[]float64{10, 9.8, 8, 7.8, 7.7, 1.7, 6, 5, 1.4, 2},
			[]float64{200, 44, 32, 24, 22, 17, 15, 12, 8, 4},
			0.9030303030303031,
		},
	}

	for _, c := range cases {
		got := f(c.n, c.x, c.y)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

func TestLeastSquareRegressionLine(t *testing.T) {
	f := func(n int, x, y []int, in int) float64 {
		var xsum, ysum, x2sum, xyprd int
		for i := 0; i < n; i++ {
			xsum = xsum + x[i]
			ysum = ysum + y[i]
			x2sum = x2sum + (x[i] * x[i])
			xyprd = xyprd + (x[i] * y[i])
		}

		xmean := float64(xsum) / float64(n)
		ymean := float64(ysum) / float64(n)

		b := float64(n*xyprd-xsum*ysum) / float64(n*x2sum-xsum*xsum)
		a := ymean - b*xmean

		return a + b*float64(in)
	}

	cases := []struct {
		n    int
		x, y []int
		in   int
		want float64
	}{
		{
			5, []int{95, 85, 80, 70, 60}, []int{85, 95, 70, 65, 70}, 80,
			78.2876712328767,
		},
	}

	for _, c := range cases {
		got := f(c.n, c.x, c.y, c.in)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

func TestMultipleLinearRegression(t *testing.T) {
	transpose := func(m [][]float64) [][]float64 {
		out := make([][]float64, 0)
		for i := 0; i < len(m[i]); i++ {
			v := make([]float64, 0)
			for j := 0; j < len(m); j++ {
				v = append(v, m[j][i])
			}

			out = append(out, v)
		}

		return out
	}

	inverse := func(m [][]float64) [][]float64 {
		out := make([][]float64, 0)
		for i := 0; i < len(m); i++ {
			v := make([]float64, 0)
			for j := 0; j < len(m[i]); j++ {
				if i == j {
					v = append(v, 1)
					continue
				}
				v = append(v, 0)
			}
			out = append(out, v)
		}

		for i := 0; i < len(m); i++ {
			c := 1 / m[i][i]
			for j := 0; j < len(m[i]); j++ {
				m[i][j] = c * m[i][j]
				out[i][j] = c * out[i][j]
			}
			for j := 0; j < len(m[i]); j++ {
				if i == j {
					continue
				}

				c := m[j][i]
				for k := 0; k < len(m[i]); k++ {
					m[j][k] = m[j][k] - c*m[i][k]
					out[j][k] = out[j][k] - c*out[i][k]
				}
			}
		}

		return out
	}

	dot := func(m, n [][]float64) [][]float64 {
		out := make([][]float64, 0)
		for i := 0; i < len(n); i++ {
			v := make([]float64, 0)
			for j := 0; j < len(m[i]); j++ {
				var c float64
				for k := 0; k < len(m); k++ {
					c = c + n[i][k]*m[k][j]
				}
				v = append(v, c)
			}

			out = append(out, v)
		}

		return out
	}

	apply := func(v []float64, m [][]float64) []float64 {
		out := make([]float64, 0)
		for i := 0; i < len(m); i++ {
			var tmp float64
			for j := 0; j < len(m[i]); j++ {
				tmp = tmp + v[j]*m[i][j]
			}
			out = append(out, tmp)
		}

		return out
	}

	XY := func(input [][]float64) ([][]float64, []float64) {
		X, Y := make([][]float64, 0), make([]float64, 0)
		for i := 0; i < len(input); i++ {
			row := []float64{1.0}
			for j := 0; j < len(input[i]); j++ {
				if j == len(input[i])-1 {
					Y = append(Y, input[i][j])
					continue
				}
				row = append(row, input[i][j])
			}
			X = append(X, row)
		}

		return X, Y
	}

	f := func(input, q [][]float64) []float64 {
		X, Y := XY(input)
		Xinv := dot(transpose(X), inverse(dot(X, transpose(X))))
		B := apply(Y, Xinv)
		return apply(B, q)
	}

	cases := []struct {
		in, q [][]float64
		want  []float64
	}{
		{
			[][]float64{
				{0.18, 0.89, 109.85},
				{1.0, 0.26, 155.72},
				{0.92, 0.11, 137.66},
				{0.07, 0.37, 76.17},
				{0.85, 0.16, 139.75},
				{0.99, 0.41, 162.6},
				{0.87, 0.47, 151.77},
			},
			[][]float64{
				{1, 0.49, 0.18},
				{1, 0.57, 0.83},
				{1, 0.56, 0.64},
				{1, 0.76, 0.18},
			},
			[]float64{
				105.2145583510693,
				142.67095130729913,
				132.93605469124682,
				129.70175404502442,
			},
		},
		{
			[][]float64{
				{5, 7, 10},
				{6, 6, 20},
				{7, 4, 60},
				{8, 5, 40},
				{9, 6, 50},
			},
			[][]float64{
				{1, 5, 5},
			},
			[]float64{29.395348837209703},
		},
	}

	for _, c := range cases {
		got := f(c.in, c.q)
		for i := range got {
			if got[i] == c.want[i] {
				continue
			}

			t.Errorf("want=%v, got=%v", c.want, got)
		}
	}

	// var sc = bufio.NewScanner(os.Stdin)
	// sc.Scan()
	// mn := strings.Split(sc.Text(), " ")
	// n, _ := strconv.ParseInt(mn[1], 10, 64)
	//
	// input := make([][]float64, 0)
	// for i := int64(0); i < n; i++ {
	// 	sc.Scan()
	// 	row := make([]float64, 0)
	// 	for _, e := range strings.Split(sc.Text(), " ") {
	// 		v, _ := strconv.ParseFloat(e, 64)
	// 		row = append(row, v)
	// 	}
	// 	input = append(input, row)
	// }
	//
	// sc.Scan()
	// qn, _ := strconv.ParseInt(sc.Text(), 10, 64)
	//
	// q := make([][]float64, 0)
	// for i := int64(0); i < qn; i++ {
	// 	sc.Scan()
	// 	row := []float64{1.0}
	// 	for _, e := range strings.Split(sc.Text(), " ") {
	// 		v, _ := strconv.ParseFloat(e, 64)
	// 		row = append(row, v)
	// 	}
	// 	q = append(q, row)
	// }
	//
	// for _, e := range f(input, q) {
	// 	fmt.Printf("%.2f\n", e)
	// }
}
