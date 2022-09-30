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
	first  *ListItem
	last   *ListItem
	length int
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	return l.first
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}
	if l.length != 0 {
		l.first.Prev, newItem.Next = newItem, l.first
	} else {
		l.last = newItem
	}
	l.first = newItem
	l.length++
	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}
	if l.length != 0 {
		l.last.Next, newItem.Prev = newItem, l.last
	} else {
		l.first = newItem
	}
	l.last = newItem
	l.length++
	return newItem
}

func (l *list) Remove(i *ListItem) {
	l.length--
	if l.first == i && l.last == i {
		l.first, l.last = nil, nil
		return
	}
	switch i {
	case l.first:
		l.first, i.Next.Prev = i.Next, nil
	case l.last:
		l.last, i.Prev.Next = i.Prev, nil
	default:
		i.Next.Prev, i.Prev.Next = i.Prev, i.Next
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if l.first == i {
		return
	}
	if l.last == i {
		i.Prev.Next = nil
		l.last = i.Prev
	} else {
		i.Next.Prev, i.Prev.Next = i.Prev, i.Next
	}
	l.first.Prev, i.Next = i, l.first
	l.first = i
}

func NewList() List {
	return new(list)
}
