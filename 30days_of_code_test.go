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

func TestArrays(t *testing.T) {
	f := func(arr []int32) []int32 {
		l := len(arr) - 1
		for i := 0; i < len(arr)/2; i++ {
			arr[i], arr[l-i] = arr[l-i], arr[i]
		}

		return arr
	}

	cases := []struct {
		in   []int32
		want []int32
	}{
		{[]int32{1, 4, 3, 2}, []int32{2, 3, 4, 1}},
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

func TestDictionariesAndMaps(t *testing.T) {
	f := func(r []string, k []string) []string {
		dic := make(map[string]string)
		for i := range r {
			v := strings.Split(r[i], " ")
			dic[v[0]] = v[1]
		}

		out := make([]string, 0)
		for i := range k {
			if v, ok := dic[k[i]]; ok {
				out = append(out, fmt.Sprintf("%v=%v", k[i], v))
				continue
			}

			out = append(out, "Not found")
		}

		return out
	}

	cases := []struct {
		r, k []string
		want []string
	}{
		{
			[]string{"sam 99912222", "tom 11122222", "harry 12299933"},
			[]string{"sam", "edward", "harry"},
			[]string{"sam=99912222", "Not found", "harry=12299933"},
		},
	}

	for _, c := range cases {
		got := f(c.r, c.k)
		for i := range c.want {
			if got[i] == c.want[i] {
				continue
			}

			t.Errorf("want=%v, got=%v", c.want, got)
		}
	}
}

func factorial(n int32) int32 {
	if n < 2 {
		return 1
	}

	return factorial(n-1) * n
}

func TestRecursion3(t *testing.T) {
	cases := []struct {
		in   int32
		want int32
	}{
		{3, 3 * 2 * 1},
	}

	for _, c := range cases {
		got := factorial(c.in)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

func TestBinaryNumbers(t *testing.T) {
	f := func(i int32) int32 {
		s := strconv.FormatInt(int64(i), 2)
		fmt.Println(s)

		var max, count int32
		for _, r := range s {
			if r == '1' {
				count++
				continue
			}

			if count > max {
				max = count
			}

			count = 0
		}

		if count > max {
			max = count
		}

		return max
	}

	cases := []struct {
		in   int32
		want int32
	}{
		{5, 1},
		{13, 2},
		{65535, 16},
	}

	for _, c := range cases {
		got := f(c.in)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

func Test2DArrays(t *testing.T) {
	f := func(A [][]int32) int32 {
		if len(A) != 6 {
			panic("invalid array length")
		}

		if len(A[0]) != 6 {
			panic("invalid array length")
		}

		max := int32(-1 << 31)
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				a := A[i][j] + A[i][j+1] + A[i][j+2]
				b := A[i+1][j+1]
				c := A[i+2][j] + A[i+2][j+1] + A[i+2][j+2]

				sum := a + b + c
				if sum > max {
					max = sum
				}
			}
		}

		return max
	}

	cases := []struct {
		in   [][]int32
		want int32
	}{
		{
			[][]int32{
				{1, 1, 1, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{1, 1, 1, 0, 0, 0},
				{0, 0, 2, 4, 4, 0},
				{0, 0, 0, 2, 0, 0},
				{0, 0, 1, 2, 4, 0},
			},
			19,
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

func TestInheritance(t *testing.T) {
	// Not provided for Go
	//
	// Java 8
	// class Student extends Person{
	// 		private int[] testScores;
	//
	// 	Student(String first, String last, int id, int [] s){
	// 		super(first, last, id);
	// 		this.testScores=s;
	// 	}
	//
	// 	public String calculate() {
	// 		int sum =0;
	// 		for(int i : this.testScores) {
	// 			sum+=i;
	// 		}
	//
	// 		double avg= sum/testScores.length;
	// 		if (avg<=100 && avg>=90) {
	// 			return "O";
	// 		}
	//
	// 		if(avg<90 && avg>=80) {
	// 			 return "E";
	// 		}
	//
	// 		if(avg<80 && avg>=70) {
	// 			 return "A";
	// 		}
	//
	// 		if(avg<70 && avg>=55) {
	// 			 return "P";
	// 		}
	//
	// 		if(avg<55 && avg>=40) {
	// 			 return "D";
	// 		}
	//
	// 		return "T";
	// 	}
	// }
}

func TestAbstractClasses(t *testing.T) {
	// Not provided for Go
	//
	// Java 8
	// class MyBook extends Book {
	//		int price;
	//
	// 		MyBook(String title, String author, int price){
	// 			super(title, author);
	// 			this.price = price;
	// 		}
	//
	// 		@Override
	// 		void display(){
	// 			System.out.println("Title: " + title);
	// 			System.out.println("Author: " + author);
	// 			System.out.println("Price: " + price);
	// 		}
	// }
}

func TestScope(t *testing.T) {
	// Not provided for Go
	//
	// Java8
	// class Difference {
	// 		private int[] elements;
	// 		public int maximumDifference;
	//
	// 		public Difference(int[] nums) {
	// 			elements = nums;
	// 		}
	//
	// 		public void computeDifference() {
	// 			Arrays.sort(elements);
	// 			maximumDifference = elements[elements.length - 1] - elements[0];
	// 		}
	// }
}

func TestLinkedList(t *testing.T) {
	// Not provided for Go
	//
	// Java8
	// public static Node insert(Node head,int data) {
	// 		if(head == null){
	// 			return new Node(data);
	// 		}
	//
	// 		Node n = head;
	// 		while(n.next != null){
	// 			n = n.next;
	// 		}
	// 		n.next = new Node(data);
	//
	// 		return head;
	// }
}

func TestExceptions(t *testing.T) {
	// Not provided for Go
	//
	// Java8
	// String S = bufferedReader.readLine();
	// try {
	// 		int input = Integer.parseInt(S);
	// 		System.out.println(input);
	// } catch(Exception e) {
	// 		System.out.println("Bad String");
	// }
}

func TestMoreExceptions(t *testing.T) {
	// Not provided for Go
	//
	// Java8
	// class Calculator{
	// 	public int power(int n, int p) throws Exception{
	// 		if(n < 0 || p < 0){
	// 			throw new Exception("n and p should be non-negative");
	// 		}
	//
	// 		return (int)Math.pow(n,p);
	// 	}
	// }
}
