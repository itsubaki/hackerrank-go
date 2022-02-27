package main_test

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"sort"
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

func TestQueuesAndStacks(t *testing.T) {
	// Not provided for Go
	//
	// Java8
	// LinkedList<Character> queue = new LinkedList();
	// LinkedList<Character> stack = new LinkedList();
	//
	// public void pushCharacter(char ch) {
	// 	stack.push(ch);
	// }
	//
	// public void enqueueCharacter(char ch) {
	// 	queue.add(ch);
	// }
	//
	// public char popCharacter() {
	// 	return stack.pop();
	// }
	//
	// public char dequeueCharacter() {
	// 	return queue.remove();
	// }
}

func TestInterfaces(t *testing.T) {
	// Not provided for Go
	//
	// Java8
	// public int divisorSum(int n) {
	// 	if(n == 1){
	// 		return 1;
	// 	}
	//
	// 	int half = n/2;
	// 	int sum = n;
	// 	do {
	// 		if(n % half == 0){
	// 			sum +=half;
	// 		}
	// 	} while( half-- > 1 );
	//
	// 	return sum;
	// }
}

func TestSorting(t *testing.T) {
	f := func(a []int32) (int32, int32, int32) {
		var swaps int32
		for i := 0; i < len(a)-1; i++ {
			for j := 0; j < len(a)-i-1; j++ {
				if a[j] > a[j+1] {
					a[j], a[j+1] = a[j+1], a[j]
					swaps++
				}
			}
		}

		return swaps, a[0], a[len(a)-1]
	}

	cases := []struct {
		in                 []int32
		swaps, first, last int32
	}{
		{[]int32{3, 2, 1}, 3, 1, 3},
	}

	for _, c := range cases {
		s, f, l := f(c.in)
		if s != c.swaps {
			t.Errorf("want=%v, got=%v", c.swaps, s)
		}
		if f != c.first {
			t.Errorf("want=%v, got=%v", c.swaps, s)
		}
		if l != c.last {
			t.Errorf("want=%v, got=%v", c.swaps, s)
		}
	}
}

func TestGenerics(t *testing.T) {
	// Not provided for Go
	//
	// Java8
	// public <T> void printArray(T[] a) {
	// 	for (T e : a) {
	// 		System.out.println(e);
	// 	}
	// }
}

func TestBSTLevelOrderTraversal(t *testing.T) {
	// Not provided for Go
	//
	// Java8
	// static void levelOrder(Node root){
	// 	LinkedList<Node> q = new LinkedList<Node>();
	// 	q.add(root);
	//
	// 	while(q.peek() != null) {
	// 		Node n = q.remove();
	// 		System.out.print(n.data);
	//
	// 		if(n.left != null) {
	// 			q.add(n.left);
	// 		}
	// 		if(n.right != null) {
	// 			q.add(n.right);
	// 		}
	//
	// 		if(q.peek() != null) {
	// 			System.out.print(" ");
	// 		}
	// 	}
	// }
}

func TestMoreLinkedList(t *testing.T) {
	// Not provided for Go
	//
	// Java8
	// public static Node removeDuplicates(Node head) {
	// 	Set<Integer> set = new HashSet<Integer>();
	// 	if (head == null) return head;
	// 	set.add(head.data);
	//
	// 	Node cur = head;
	// 	while (cur.next != null) {
	// 		if (set.contains(cur.next.data)) {
	// 			cur.next = cur.next.next;
	// 			continue;
	// 		}
	//
	// 		set.add(cur.next.data);
	// 		cur = cur.next;
	// 	}
	//
	// 	return head;
	// }
}

func TestRunningTimeAndComplexity(t *testing.T) {
	f := func(N int64) bool {
		if N < 2 {
			return false
		}

		if N == 2 {
			return true
		}

		if N%2 == 0 {
			return false
		}

		for i := int64(3); i < int64(math.Sqrt(float64(N)))+1; i = i + 2 {
			if N%i == 0 {
				return false
			}
		}

		return true
	}

	cases := []struct {
		in   int64
		want bool
	}{
		{3, true},
		{5, true},
		{7, true},
		{15, false},
	}

	for _, c := range cases {
		got := f(c.in)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}

	// var T int
	// fmt.Scan(&T)
	// for i := 0; i < T; i++ {
	// 	var N int64
	// 	fmt.Scan(&N)
	// 	if f(N) {
	// 		fmt.Println("Prime")
	// 	} else {
	// 		fmt.Println("Not prime")
	// 	}
	// }
}

func TestNestedLogic(t *testing.T) {
	f := func(aDay, aMon, aYear, eDay, eMon, eYear int) int {
		if aYear < eYear {
			return 0
		}

		if aYear > eYear {
			return 10000
		}

		if aMon < eMon {
			return 0
		}

		if aMon > eMon {
			return (aMon - eMon) * 500
		}

		if aDay < eDay {
			return 0
		}

		if aDay > eDay {
			return (aDay - eDay) * 15
		}

		return 0
	}

	cases := []struct {
		aDay, aMon, aYear int
		eDay, eMon, eYear int
		want              int
	}{
		{9, 6, 2015, 6, 6, 2015, 45},
	}

	for _, c := range cases {
		got := f(c.aDay, c.aMon, c.aYear, c.eDay, c.eMon, c.eYear)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}

	// var aDay, aMon, aYear int
	// fmt.Scan(&aDay)
	// fmt.Scan(&aMon)
	// fmt.Scan(&aYear)
	//
	// var eDay, eMon, eYear int
	// fmt.Scan(&eDay)
	// fmt.Scan(&eMon)
	// fmt.Scan(&eYear)
	//
	// fmt.Println(f(aDay, aMon, aYear, eDay, eMon, eYear))
}

func TestTesting(t *testing.T) {
	// Not provided for Go
}

func TestRegexPatternAndIntroToDatabase(t *testing.T) {
	f := func(in []string) []string {
		name, email := make([]string, 0), make([]string, 0)
		for _, e := range in {
			s := strings.Split(e, " ")
			name = append(name, s[0])
			email = append(email, s[1])
		}

		out := make([]string, 0)
		for i := range email {
			if strings.HasSuffix(email[i], "@gmail.com") {
				out = append(out, name[i])
			}
		}

		sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
		return out
	}

	cases := []struct {
		in   []string
		want []string
	}{
		{
			[]string{
				"riya riya@gmail.com",
				"julia julia@julia.me",
				"julia sjulia@gmail.com",
				"julia julia@gmail.com",
				"samantha samantha@gmail.com",
				"tanya tanya@gmail.com",
			},
			[]string{"julia", "julia", "riya", "samantha", "tanya"},
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

func TestBitwiseAND(t *testing.T) {
	f := func(N, K int32) int32 {
		var max int32
		for i := int32(1); i < N+1; i++ {
			for j := i + 1; j < N+1; j++ {
				h := i & j
				if h > max && h < K {
					max = h
				}
			}
		}

		return max
	}

	cases := []struct {
		N, K int32
		want int32
	}{
		{5, 2, 1},
		{8, 5, 4},
		{2, 2, 0},
	}

	for _, c := range cases {
		got := f(c.N, c.K)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}

}
