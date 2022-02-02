package main_test

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	f := func(r io.Reader) string {
		var sc = bufio.NewScanner(r)

		var in string
		if sc.Scan() {
			in = sc.Text()
		}

		return fmt.Sprintf("Hello, World.\n%v", in)
	}
	// fmt.Println(f(os.Stdin))

	cases := []struct {
		in   io.Reader
		want string
	}{
		{
			strings.NewReader("Welcome to 30 Days of Code!"),
			"Hello, World.\nWelcome to 30 Days of Code!",
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

func TestDataTypes(t *testing.T) {
	f := func(r io.Reader) string {
		var i uint64 = 4
		var d float64 = 4.0
		var s string = "HackerRank "

		var sc = bufio.NewScanner(r)

		sc.Scan()
		si, _ := strconv.ParseUint(sc.Text(), 10, 64)

		sc.Scan()
		sf, _ := strconv.ParseFloat(sc.Text(), 64)

		sc.Scan()
		st := sc.Text()

		return fmt.Sprintf("%d\n%.1f\n%s\n", i+si, d+sf, s+st)
	}

	cases := []struct {
		in   io.Reader
		want string
	}{
		{
			strings.NewReader("12\n4.0\nis the best place to learn and practice coding!"),
			"16\n8.0\nHackerRank is the best place to learn and practice coding!\n",
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

func TestOperators(t *testing.T) {
	f := func(meal_cost float64, tip_percent int32, tax_percent int32) int32 {
		tip := meal_cost * float64(tip_percent) / 100.0
		tax := meal_cost * float64(tax_percent) / 100.0
		ret := math.Round(meal_cost + tip + tax)
		return int32(ret)
	}

	cases := []struct {
		cost     float64
		tip, tax int32
		want     int32
	}{}

	for _, c := range cases {
		got := f(c.cost, c.tip, c.tax)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

func TestConditionalStatements(t *testing.T) {
	f := func(n int32) string {
		if n%2 == 1 {
			return "Weird"
		}

		if n >= 2 && n <= 5 {
			return "Not Weird"
		}

		if n >= 6 && n <= 20 {
			return "Weird"
		}

		return "Not Weird"
	}

	cases := []struct {
		in   int32
		want string
	}{
		{3, "Weird"},
	}

	for _, c := range cases {
		got := f(c.in)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

// func TestClassVsInstance(t *testing.T) {}
type person struct {
	age int
}

func (p person) NewPerson(initialAge int) person {
	if initialAge < 0 {
		fmt.Println("Age is not valid, setting age to 0.")
		initialAge = 0
	}

	return person{age: initialAge}
}

func (p person) amIOld() {
	if p.age < 13 {
		fmt.Println("You are young.")
		return
	}

	if p.age < 18 {
		fmt.Println("You are a teenager.")
		return
	}

	fmt.Println("You are old.")
}

func (p person) yearPasses() person {
	p.age++
	return p
}

func TestLoops(t *testing.T) {
	f := func(n int32) []string {
		out := make([]string, 0)

		for i := 1; i < 11; i++ {
			out = append(out, fmt.Sprintf("%v x %v = %v", n, i, n*int32(i)))
		}

		return out
	}

	cases := []struct {
		in   int32
		want []string
	}{
		{
			3,
			[]string{
				"3 x 1 = 3",
				"3 x 2 = 6",
				"3 x 3 = 9",
				"3 x 4 = 12",
				"3 x 5 = 15",
				"3 x 6 = 18",
				"3 x 7 = 21",
				"3 x 8 = 24",
				"3 x 9 = 27",
				"3 x 10 = 30",
			},
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

func TestLetsReview(t *testing.T) {
	f := func(r io.Reader) []string {
		var sc = bufio.NewScanner(r)

		sc.Scan()
		n, _ := strconv.ParseInt(sc.Text(), 10, 64)

		out := make([]string, 0)
		for i := int64(0); i < n; i++ {
			sc.Scan()
			str := sc.Text()

			var odd, even string
			for i := range str {
				if i%2 == 0 {
					even = even + string(str[i])
					continue
				}

				odd = odd + string(str[i])
			}

			out = append(out, fmt.Sprintf("%v %v", even, odd))
		}

		return out
	}

	// for _, o := range f(os.Stdin) {
	// 	fmt.Println(o)
	// }

	cases := []struct {
		in   io.Reader
		want []string
	}{
		{
			strings.NewReader("1\nadbecf"),
			[]string{"abc def"},
		},
		{
			strings.NewReader("2\nHacker\nRank\n"),
			[]string{"Hce akr", "Rn ak"},
		},
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
