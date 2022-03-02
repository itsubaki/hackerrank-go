package main_test

import (
	"bufio"
	"fmt"
	"os"
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

func TestStdIn(t *testing.T) {
	f := func(in []string) []string {
		return in
	}

	var n int
	fmt.Scan(&n)

	var sc = bufio.NewScanner(os.Stdin)
	var in []string
	for i := 0; i < n; i++ {
		if sc.Scan() {
			in = append(in, sc.Text())
		}
	}

	for _, o := range f(in) {
		fmt.Println(o)
	}
}
