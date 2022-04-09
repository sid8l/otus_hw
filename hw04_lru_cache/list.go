package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	List
	len   int
	front *ListItem
	back  *ListItem
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *ListItem {
	return l.front
}

func (l list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v, Next: l.front}
	if l.len == 0 {
		l.back = newItem
	} else {
		l.front.Prev = newItem
	}
	l.front = newItem
	l.len++
	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v, Prev: l.back}
	if l.len == 0 {
		l.front = newItem
	} else {
		l.back.Next = newItem
	}
	l.back = newItem
	l.len++
	return newItem
}

func (l *list) Remove(i *ListItem) {
	if l.front == i && l.back == i {
		l.front, l.back = nil, nil
	}
	switch i {
	case l.front:
		l.front = i.Next
		i.Next.Prev = nil
	case l.back:
		l.back = i.Prev
		i.Prev.Next = nil
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	switch i {
	case l.front:
		return
	case l.back:
		l.back = i.Prev
	default:
		i.Next.Prev = i.Prev
	}
	i.Prev.Next = i.Next
	i.Next, i.Prev = l.front, nil
	l.front.Prev, l.front = i, i
}

func NewList() List {
	return new(list)
}
