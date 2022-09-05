package domain

import "fmt"

type List struct {
	head *Input
}

var file int

func (l *List) Insert(lines []string) {
	list := &Input{name: fmt.Sprintf("file%d", file+1), lines: lines, start: make(chan bool, 1),
		currentLine: make(chan parse, 1), next: nil}
	if l.head == nil {
		l.head = list
	} else {
		in := l.head
		for in.next != nil {
			in = in.next
		}
		in.next = list
	}
	file++
}

func (l *List) Head() *Input {
	return l.head
}

func ConvertSinglyToCircular(l *List) {
	p := l.head
	for p.next != nil {
		p = p.next
	}
	p.next = l.head
}
