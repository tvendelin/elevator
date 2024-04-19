package main

import (
	"testing"
)

type testcase struct {
	P [][]int
	L int
	E []int
}

var cases []testcase = []testcase{
	{P: [][]int{{3, 2}, {5, 2}, {2, 1}, {2, 5}, {4, 3}}, L: 1, E: []int{2, 5, 4, 3, 2, 1}},
}

func TestT(t *testing.T) {
	for i, tc := range cases {
		pq := QueueOfPersons(tc.P)
		actual := Order(tc.L, pq)
		if !compareIntSlice(actual, tc.E) { // replace with actual test
			t.Errorf("Test # %d: for level %d and queue %v expected %v, got %v\n", i, tc.L, tc.P, tc.E, actual)
		}
	}
}

func compareIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, n := range a {
		if n != b[i] {
			return false
		}
	}
	return true
}

func QueueOfPersons(q [][]int) []Person {
	qp := make([]Person, 0, len(q))
	for _, p := range q {
		qp = append(qp, Person{From: p[0], To: p[1]})
	}
	return qp
}

/* uncomment for tasks with linked lists
func toLinkedList(s []int) *ListNode {
	var last *ListNode
	var node *ListNode
	for i := len(s) - 1; i >= 0; i-- {
		node = &ListNode{Val: s[i], Next: last}
		last = node
	}
	return node
}
*/
