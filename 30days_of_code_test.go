package main_test

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	f := func(r io.Reader) string {
		var sc = bufio.NewScanner(r)
		sc.Scan()
		in := sc.Text()

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

}
