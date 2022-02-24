package main_test

import (
	"fmt"
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
