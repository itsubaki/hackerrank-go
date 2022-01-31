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

	// fmt.Println(f(os.Stdin))
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

func TestLoop(t *testing.T) {
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
