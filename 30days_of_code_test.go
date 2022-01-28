package main_test

import (
	"bufio"
	"fmt"
	"io"
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
