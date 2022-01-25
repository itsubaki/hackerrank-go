package main_test

import (
	"fmt"
	"testing"
)

func TestTemplate(t *testing.T) {
	f := func(s string) string {
		return s
	}

	cases := []struct {
		in   string
		want string
	}{
		{"foo", "foo"},
	}

	for _, c := range cases {
		got := f(c.in)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

func TestPlusMinus(t *testing.T) {
	f := func(arr []int32) []string {
		p, n, z := 0.0, 0.0, 0.0
		for i := range arr {
			if arr[i] > 0 {
				p++
			} else if arr[i] < 0 {
				n++
			} else {
				z++
			}
		}

		out := make([]string, 0)
		out = append(out, fmt.Sprintf("%.6f", p/float64(len(arr))))
		out = append(out, fmt.Sprintf("%.6f", n/float64(len(arr))))
		out = append(out, fmt.Sprintf("%.6f", z/float64(len(arr))))

		for _, o := range out {
			fmt.Println(o)
		}

		return out
	}

	cases := []struct {
		in   []int32
		want []string
	}{
		{
			[]int32{1, 1, 0, -1, -1},
			[]string{"0.400000", "0.400000", "0.200000"},
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

func TestMiniMaxSum(t *testing.T) {
	f := func(arr []int32) []int64 {
		if len(arr) != 5 {
			panic(fmt.Sprintf("invalid arr=%v", arr))
		}

		ans := make([]int64, 0)
		for _, skip := range []int{0, 1, 2, 3, 4} {
			var sum int64
			for i, a := range arr {
				if i == skip {
					continue
				}

				sum = sum + int64(a)
			}

			ans = append(ans, sum)
		}

		min, max := int64(10000000000), int64(0)
		for i := range ans {
			if ans[i] > max {
				max = ans[i]
			}

			if ans[i] < min {
				min = ans[i]
			}
		}

		out := make([]int64, 0)
		out = append(out, min)
		out = append(out, max)

		fmt.Printf("%v %v\n", out[0], out[1])
		return out
	}

	cases := []struct {
		in   []int32
		want []int64
	}{
		{
			[]int32{1, 3, 5, 7, 9},
			[]int64{16, 24},
		},
		{
			[]int32{1, 2, 3, 4, 5},
			[]int64{10, 14},
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

func TestTimeConversion(t *testing.T) {
	f := func(s string) string {
		hour, ampm := s[:2], s[len(s)-2:]
		if ampm == "AM" && hour == "12" {
			hour = "00"
		}
		if ampm == "PM" {
			switch hour {
			case "01":
				hour = "13"
			case "02":
				hour = "14"
			case "03":
				hour = "15"
			case "04":
				hour = "16"
			case "05":
				hour = "17"
			case "06":
				hour = "18"
			case "07":
				hour = "19"
			case "08":
				hour = "20"
			case "09":
				hour = "21"
			case "10":
				hour = "22"
			case "11":
				hour = "23"
			}
		}

		return fmt.Sprintf("%v%v", hour, s[2:len(s)-2])
	}

	cases := []struct {
		in   string
		want string
	}{
		{"12:01:00PM", "12:01:00"},
		{"12:01:00AM", "00:01:00"},
		{"07:05:45PM", "19:05:45"},
	}

	for _, c := range cases {
		got := f(c.in)
		if got != c.want {
			t.Errorf("want=%v, got=%v", c.want, got)
		}
	}
}

func TestLonelyInteger(t *testing.T) {
	f := func(a []int32) int32 {
		imap := make(map[int32]int32)
		for _, i := range a {
			v, ok := imap[i]
			if !ok {
				imap[i] = 1
				continue
			}

			imap[i] = v + 1
		}

		for k, v := range imap {
			if v == 1 {
				return k
			}
		}

		panic("invalid array")
	}

	cases := []struct {
		in   []int32
		want int32
	}{
		{[]int32{1, 2, 3, 4, 3, 2, 1}, 4},
	}

	for _, c := range cases {
		got := f(c.in)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

func TestDiagonalDifference(t *testing.T) {
	f := func(arr [][]int32) int32 {
		r := make([][]int32, 0)
		for i := range arr {
			v := make([]int32, 0)
			for j := range arr[i] {
				v = append(v, arr[i][len(arr)-1-j])
			}
			r = append(r, v)
		}

		var lr int32
		for i := range arr {
			lr = lr + arr[i][i]
		}

		var rl int32
		for i := range r {
			rl = rl + r[i][i]
		}

		diff := lr - rl
		if diff < 0 {
			diff = -1 * diff
		}

		return diff
	}

	cases := []struct {
		in   [][]int32
		want int32
	}{
		{
			[][]int32{
				{1, 2, 3},
				{4, 5, 6},
				{9, 8, 9},
			},
			2,
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

func TestCountingSort(t *testing.T) {
	f := func(arr []int32) []int32 {
		out := make([]int32, 100)
		for i := range arr {
			out[arr[i]] += 1
		}

		return out
	}

	cases := []struct {
		in   []int32
		want []int32
	}{
		{[]int32{1, 1, 3, 2, 1}, []int32{0, 3, 1, 1}},
	}

	for _, c := range cases {
		got := f(c.in)
		for i := range c.want {
			if got[i] == c.want[i] {
				continue
			}

			t.Errorf("want=%v, got=%v", c.want, got)
		}
	}
}
