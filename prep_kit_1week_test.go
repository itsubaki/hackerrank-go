package main_test

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestPlusMinus(t *testing.T) {
	f := func(arr []int32) (float64, float64, float64) {
		var p, n, z float64
		for i := range arr {
			if arr[i] > 0 {
				p++
			} else if arr[i] < 0 {
				n++
			} else {
				z++
			}
		}

		s := float64(len(arr))
		return p / s, n / s, z / s
	}

	cases := []struct {
		in      []int32
		p, n, z float64
	}{
		{
			[]int32{1, 1, 0, -1, -1},
			0.400000, 0.400000, 0.200000,
		},
	}

	for _, c := range cases {
		p, n, z := f(c.in)
		if p != c.p {
			t.Errorf("want=%v, got=%v", c.p, p)
		}
		if n != c.n {
			t.Errorf("want=%v, got=%v", c.n, n)
		}
		if z != c.z {
			t.Errorf("want=%v, got=%v", c.z, z)
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

		min, max := int64(1<<63-1), int64(-1<<63)
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
			if v, ok := imap[i]; ok {
				imap[i] = v + 1
				continue
			}

			imap[i] = 1
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

func TestCountingSort1(t *testing.T) {
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

func TestFindZigZagSequence(t *testing.T) {
	// Not provided for Go
	//
	// Java8
	// Note: You can modify at most three lines in the given code. You cannot add or remove lines of code.
	// public static void findZigZagSequence(int [] a, int n){
	// 	Arrays.sort(a);
	// 	int mid = n/2; // one
	// 	int temp = a[mid];
	// 	a[mid] = a[n - 1];
	// 	a[n - 1] = temp;
	//
	// 	int st = mid + 1;
	// 	int ed = n - 2; // two
	// 	while(st <= ed){
	// 		temp = a[st];
	// 		a[st] = a[ed];
	// 		a[ed] = temp;
	// 		st = st + 1;
	// 		ed = ed - 1; // three
	// 	}
	// 	for(int i = 0; i < n; i++){
	// 		if(i > 0) System.out.print(" ");
	// 		System.out.print(a[i]);
	// 	}
	// 	System.out.println();
	// }
}

func TestTowerBreakers(t *testing.T) {
	f := func(n, m int32) int32 {
		if m == 1 {
			return 2
		}

		if n%2 == 1 {
			return 1
		}

		return 2
	}

	cases := []struct {
		n, m int32
		want int32
	}{
		{2, 6, 2},
	}

	for _, c := range cases {
		got := f(c.n, c.m)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

func TestCaesarCipher(t *testing.T) {
	f := func(s string, k int32) string {
		var buf strings.Builder
		for _, r := range s {
			if r >= 'a' && r <= 'z' {
				a := r + k
				for a > rune('z') {
					a = a - 26
				}

				buf.WriteRune(a)
				continue
			}

			if r >= 'A' && r <= 'Z' {
				a := r + k
				for a > rune('Z') {
					a = a - 26
				}

				buf.WriteRune(a)
				continue
			}

			buf.WriteRune(r)
		}

		return buf.String()
	}

	cases := []struct {
		in   string
		k    int32
		want string
	}{
		{"abc-z", 3, "def-c"},
		{"ABC-Z", 3, "DEF-C"},
		{"abc-zZ", 3, "def-cC"},
		{"www.abc.xy", 87, "fff.jkl.gh"},
	}

	for _, c := range cases {
		got := f(c.in, c.k)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

func TestGridChallenge(t *testing.T) {
	f := func(grid []string) string {
		for i := range grid {
			s := strings.Split(grid[i], "")
			sort.Strings(s)
			grid[i] = strings.Join(s, "")
		}

		for i := 0; i < len(grid); i++ {
			for j := 1; j < len(grid[i]); j++ {
				if grid[i][j] >= grid[i][j-1] {
					continue
				}

				return "NO"
			}
		}

		for j := 0; j < len(grid[0]); j++ {
			for i := 1; i < len(grid); i++ {
				if grid[i][j] >= grid[i-1][j] {
					continue
				}

				return "NO"
			}
		}

		return "YES"
	}

	cases := []struct {
		in   []string
		want string
	}{
		{
			[]string{
				"abc",
				"ade",
				"efg",
			},
			"YES",
		},
		{
			[]string{
				"eabcd",
				"fghij",
				"olkmn",
				"trpqs",
				"xywuv",
			},
			"YES",
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

func superDigit(n string, k int32) int32 {
	if len(n) == 1 && k == 1 {
		v, _ := strconv.ParseInt(n, 10, 64)
		return int32(v)
	}

	var sum int64
	for i := range n {
		v, _ := strconv.ParseInt(string(n[i]), 10, 64)
		sum += v
	}

	s := strconv.FormatInt(sum*int64(k), 10)
	return superDigit(s, 1)
}

func TestSuperDigit(t *testing.T) {
	cases := []struct {
		n    string
		k    int32
		want int32
	}{
		{"9875", 4, 8},
	}

	for _, c := range cases {
		got := superDigit(c.n, c.k)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

func TestMinimumBribes(t *testing.T) {
	f := func(q []int32) string {
		var sum int32
		for i := len(q) - 1; i > -1; i-- {
			if q[i]-int32(i+1) > 2 {
				return "Too chaotic"
			}

			var max int32
			if q[i]-2 > max {
				max = q[i] - 2
			}

			for j := max; j < int32(i); j++ {
				if q[j] > q[i] {
					sum++
				}
			}
		}

		return strconv.FormatInt(int64(sum), 10)
	}

	cases := []struct {
		in   []int32
		want string
	}{
		{[]int32{1, 2, 3, 5, 4, 6, 7, 8}, "1"},
		{[]int32{4, 1, 2, 3}, "Too chaotic"},
		{[]int32{2, 1, 5, 3, 4}, "3"},
		{[]int32{2, 5, 1, 3, 4}, "Too chaotic"},
		{[]int32{1, 2, 5, 3, 7, 8, 6, 4}, "7"},
	}

	for _, c := range cases {
		got := f(c.in)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

func TestMergeTwoSortedLinkedLists(t *testing.T) {
	// Not provided for Go
	//
	// Java8
	// static SinglyLinkedListNode mergeLists(SinglyLinkedListNode head1, SinglyLinkedListNode head2) {
	// 	if(head1 == null) {
	// 		return head2;
	// 	}
	// 	if(head2 == null) {
	// 		return head1;
	// 	}

	// 	if(head1.data < head2.data) {
	// 		head1.next = mergeLists(head1.next, head2);
	// 		return head1;
	// 	} else {
	// 		head2.next = mergeLists(head1, head2.next);
	// 		return head2;
	// 	}
	// }
}

type stack struct {
	values []string
}

func (s *stack) Push(v string) {
	s.values = append(s.values, v)
}

func (s *stack) Pop() string {
	ret := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return ret
}

func (s *stack) Top() string {
	return s.values[len(s.values)-1]
}

func (s *stack) IsEmpty() bool {
	return len(s.values) == 0
}

type queue struct {
	in  *stack
	out *stack
}

func (q *queue) Enq(s string) {
	q.in.Push(s)
}

func (q *queue) Deq() string {
	if q.out.IsEmpty() {
		q.shift()
	}

	return q.out.Pop()
}

func (q *queue) Front() string {
	if q.out.IsEmpty() {
		q.shift()
	}

	return q.out.Top()
}

func (q *queue) shift() {
	for {
		if q.in.IsEmpty() {
			break
		}

		v := q.in.Pop()
		q.out.Push(v)
	}
}

func TestQueueUsingTwoStacks(t *testing.T) {
	f := func(ops []string) []string {
		q := &queue{
			in:  &stack{values: make([]string, 0)},
			out: &stack{values: make([]string, 0)},
		}

		out := make([]string, 0)
		for _, o := range ops {
			sp := strings.Split(o, " ")
			switch sp[0] {
			case "1":
				q.Enq(sp[1])
			case "2":
				q.Deq()
			case "3":
				out = append(out, q.Front())
			}
		}

		return out
	}

	cases := []struct {
		in   []string
		want []string
	}{
		{
			[]string{
				"1 42",
				"2",
				"1 14",
				"3",
				"1 28",
				"3",
				"1 60",
				"1 78",
				"2",
				"2",
			},
			[]string{"14", "14"},
		},
		{
			[]string{
				"1 92118642",
				"2",
				"1 107858633",
				"1 110186788",
				"1 883309178",
				"1 430939631",
				"2",
				"1 739711408",
				"1 803703507",
				"1 643797161",
				"1 538560826",
				"3",
				"1 595864615",
				"1 490282285",
				"1 558095366",
				"1 893666727",
				"1 595679828",
				"3",
			},
			[]string{"110186788", "110186788"},
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

	{
		read := func(reader *bufio.Reader) string {
			str, _, err := reader.ReadLine()
			if err == io.EOF {
				return ""
			}

			return strings.TrimRight(string(str), "\r\n")
		}

		stdout, _ := os.Create(os.Getenv("OUTPUT_PATH"))
		defer stdout.Close()

		r := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
		w := bufio.NewWriterSize(stdout, 16*1024*1024)
		n, _ := strconv.ParseInt(strings.TrimSpace(read(r)), 10, 64)

		q := &queue{
			in:  &stack{values: make([]string, 0)},
			out: &stack{values: make([]string, 0)},
		}

		for i := int64(0); i < n; i++ {
			s := strings.Split(strings.TrimSpace(read(r)), " ")

			switch s[0] {
			case "1":
				q.Enq(s[1])
			case "2":
				q.Deq()
			case "3":
				fmt.Fprintf(w, "%s\n", q.Front())
			}
		}

		w.Flush()
	}
}

func TestIsBalanced(t *testing.T) {
	f := func(s string) string {
		if len(s)%2 != 0 {
			return "NO"
		}

		q := make([]rune, 0)
		for _, r := range s {
			if strings.ContainsRune("{[(", r) {
				q = append(q, r)
				continue
			}

			// r is ")]}"
			if len(q) == 0 {
				return "NO"
			}

			var bra rune
			switch r {
			case ')':
				bra = '('
			case '}':
				bra = '{'
			case ']':
				bra = '['
			}

			if q[len(q)-1] != bra {
				return "NO"
			}

			q = q[:len(q)-1] // pop
		}

		if len(q) != 0 {
			return "NO"
		}

		return "YES"
	}

	cases := []struct {
		in   string
		want string
	}{
		{"{[()]}", "YES"},
		{"{[(])}", "NO"},
		{"([[)", "NO"},
		{"}][}}(}][))]", "NO"},
		{"[](){()}", "YES"},
		{"({}([][]))[]()", "YES"},
		{"{)[](}]}]}))}(())(", "NO"},
	}

	for _, c := range cases {
		got := f(c.in)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

func TestSimpleTextEditor(t *testing.T) {
	f := func(s string, ops []string) []string {
		out := make([]string, 0)
		prev := make([]string, 0)
		for _, o := range ops {
			sp := strings.Split(o, " ")
			switch sp[0] {
			case "1":
				prev = append(prev, s)
				s = strings.Join([]string{s, sp[1]}, "")
			case "2":
				prev = append(prev, s)
				k, _ := strconv.Atoi(sp[1])
				s = s[:len(s)-k]
			case "3":
				k, _ := strconv.Atoi(sp[1])
				out = append(out, string(s[k-1]))
			case "4":
				s = prev[len(prev)-1]
				prev = prev[:len(prev)-1]
			}
		}

		return out
	}

	cases := []struct {
		s    string
		ops  []string
		want []string
	}{
		{
			"abcde",
			[]string{"1 fg", "3 6", "2 5", "4", "3 7", "4", "3 4"},
			[]string{"f", "g", "d"},
		},
	}

	for _, c := range cases {
		got := f(c.s, c.ops)
		for i := range c.want {
			if got[i] == c.want[i] {
				continue
			}

			t.Errorf("want=%v, got=%v", c.want, got)
		}
	}
}

type LegoBlocks struct {
	ways     [][]int64
	restrict [][]int64
	mod      int64
}

func NewLegoBlocks() *LegoBlocks {
	ways, restrict := make([][]int64, 0), make([][]int64, 0)
	for i := 0; i < 1001; i++ {
		v, rv := make([]int64, 0), make([]int64, 0)
		for j := 0; j < 1001; j++ {
			v, rv = append(v, -1), append(rv, -1)
		}
		ways, restrict = append(ways, v), append(restrict, rv)
	}

	return &LegoBlocks{
		mod:      1000000007,
		ways:     ways,
		restrict: restrict,
	}
}

func (b *LegoBlocks) solve(n, m int32) int64 {
	if b.ways[n][m] != -1 {
		return b.ways[n][m]
	}

	if m == 1 {
		b.ways[n][m] = 1
		return 1
	}

	if n == 1 {
		if m < 5 {
			v := 2 * b.solve(1, m-1)
			b.ways[n][m] = v
			return v
		}

		v := b.solve(1, m-1)
		v = (v + b.solve(1, m-2)) % b.mod
		v = (v + b.solve(1, m-3)) % b.mod
		v = (v + b.solve(1, m-4)) % b.mod
		b.ways[n][m] = v
		return v
	}

	v := int64(1)
	o := b.solve(1, m)
	for i := int32(0); i < n; i++ {
		v = (v * o) % b.mod
	}

	b.ways[n][m] = v
	return v
}

func (b *LegoBlocks) Solve(n, m int32) int64 {
	if b.restrict[n][m] != -1 {
		return b.restrict[n][m]
	}

	if m == 1 {
		b.restrict[n][m] = 1
		return 1
	}

	v := b.solve(n, m)
	for i := int32(1); i < m; i++ {
		v = v - (b.solve(n, i)*b.Solve(n, m-i))%b.mod
		if v < 0 {
			v = v + b.mod
		}
	}

	b.restrict[n][m] = v
	return v
}

func TestLegoBlocks(t *testing.T) {
	cases := []struct {
		n, m int32
		want int64
	}{
		{2, 2, 3},
		{3, 2, 7},
		{2, 3, 9},
		{4, 4, 3375},
	}

	for _, c := range cases {
		got := NewLegoBlocks().Solve(c.n, c.m)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

// https://pkg.go.dev/container/heap
type int32Heap []int32

func (h int32Heap) Len() int            { return len(h) }
func (h int32Heap) Less(i, j int) bool  { return h[i] < h[j] }
func (h int32Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *int32Heap) Push(x interface{}) { *h = append(*h, x.(int32)) }
func (h *int32Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TestFindMedian(t *testing.T) {
	f := func(a []int32) int32 {
		h := &int32Heap{}
		heap.Init(h)

		for i := range a {
			heap.Push(h, a[i])
		}

		var v interface{}
		for i := 0; i < len(a)/2+1; i++ {
			v = heap.Pop(h)
		}

		return v.(int32)
	}

	cases := []struct {
		in   []int32
		want int32
	}{
		{[]int32{1, 2, 3, 4, 5}, 3},
	}

	for _, c := range cases {
		got := f(c.in)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

func TestJesseAndCookies(t *testing.T) {
	f := func(k int32, A []int32) int32 {
		h := &int32Heap{}
		heap.Init(h)

		for i := range A {
			heap.Push(h, A[i])
		}

		var count int32
		for {
			v0 := heap.Pop(h)
			if v0.(int32) >= k {
				return count
			}

			if h.Len() == 0 {
				return -1
			}

			v1 := heap.Pop(h)
			heap.Push(h, v0.(int32)+v1.(int32)*2)

			count++
		}
	}

	n := 100000
	testA := make([]int32, n)
	for i := 0; i < n; i++ {
		testA[i] = 1
	}

	cases := []struct {
		k    int32
		A    []int32
		want int32
	}{
		{7, []int32{1, 2, 3, 9, 10, 12}, 2},
		{105823341, testA, 99999},
	}

	for _, c := range cases {
		got := f(c.k, c.A)
		if got == c.want {
			continue
		}

		t.Errorf("want=%v, got=%v", c.want, got)
	}
}

func TestTreePreorderTraversal(t *testing.T) {
	// Not provided for Go
	//
	// Java8
	// public static void preOrder(Node root) {
	// 	System.out.print(root.data + " ");
	// 	if (root.left != null) {
	// 		preOrder(root.left);
	// 	}

	// 	if (root.right != null) {
	// 		preOrder(root.right);
	// 	}
	// }
}

func TestTreeHuffmanDecoding(t *testing.T) {
	// Not provided for Go
	//
	// Java8
	// void decode(String S, Node root) {
	// 	StringBuilder output = new StringBuilder();
	// 	Node base = root;
	// 	while (!S.isEmpty()){
	// 		if (S.charAt(0) == '1'){
	// 			base = base.right;
	// 			S = S.substring(1);
	// 		} else {
	// 			base = base.left;
	// 			S = S.substring(1);
	// 		}

	// 		if (base.left == null && base.right == null){
	// 			output.append(base.data);
	// 			base = root;
	// 		}
	// 	}

	// 	System.out.println(output.toString());
	// }
}

type trienode struct {
	key         rune
	children    map[rune]*trienode
	isCompleted bool
	count       int32
}

func newTrieNode(key rune) *trienode {
	return &trienode{
		key:      key,
		children: make(map[rune]*trienode),
	}
}

func add(root *trienode, str string) bool {
	cur := root
	for i := 0; i < len(str); i++ {
		c := rune(str[i])
		index := c - 'a'
		if _, ok := cur.children[index]; !ok {
			cur.children[index] = newTrieNode(c)
		}

		if cur.isCompleted {
			return false
		}

		cur.count++
		cur = cur.children[index]
	}

	cur.isCompleted = true
	cur.count++
	return cur.count < 2
}

func TestNoPrefixSet(t *testing.T) {
	// Use trie tree to reduce complexity
	f := func(words []string) []string {
		root := newTrieNode(' ')
		for i := range words {
			if add(root, words[i]) {
				continue
			}

			return []string{"BAD SET", words[i]}
		}

		return []string{"GOOD SET"}
	}

	cases := []struct {
		in   []string
		want []string
	}{
		{
			[]string{"abcd", "bcd", "abcde", "bcde"},
			[]string{"BAD SET", "abcde"},
		},
		{
			[]string{"ab", "bc", "cd"},
			[]string{"GOOD SET"},
		},
		{
			[]string{"aab", "defgab", "abcde", "aabcde", "cedaaa", "bbbbbbbbbb", "jabjjjad"},
			[]string{"BAD SET", "aabcde"},
		},
		{
			[]string{"aab", "aac", "aacghgh", "aabghgh"},
			[]string{"BAD SET", "aacghgh"},
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
