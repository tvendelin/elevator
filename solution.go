package main

type Person struct {
	From, To int
}

type LinkedPerson struct {
	From, To   int
	Prev, Next *LinkedPerson
}

type Elevator struct {
	Dest       int
	Passengers map[int][]*LinkedPerson
	Pos        int
	Log        []int
}

func (e *Elevator) Travel(q *LinkedPerson, f [][]*LinkedPerson) []int {
	for len(e.Passengers) != 0 && q != nil {
	}
	return e.Log
}

func linkThem(q []Person) (*LinkedPerson, [][]*LinkedPerson) {
	floors := make([][]*LinkedPerson, 0, len(q))
	prev := &LinkedPerson{From: q[0].From, To: q[0].To}
	head := prev
	for _, p := range q[1:] {
		l := &LinkedPerson{From: p.From, To: p.To, Prev: prev}
		floors[p.From] = append(floors[p.From], l)
		prev.Next = l
		prev = l
	}
	return head, floors
}

func Order(level int, queue []Person) []int {
	q, floors := linkThem(queue)
	el := &Elevator{Dest: q.From, Pos: level}
	return el.Travel(q, floors)
}
