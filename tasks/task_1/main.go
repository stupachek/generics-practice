package main

import "fmt"

type List[T any] struct {
	elems []T
}

func (l *List[T]) Add(t T) {
	l.elems = append(l.elems, t)
}

func (l *List[T]) Update(i int, t T) {
	l.elems[i] = t
}

func (l *List[T]) Pop() T {
	last := l.elems[len(l.elems)-1]
	l.elems = l.elems[:len(l.elems)]
	return last
}

func (l *List[T]) Delete(i int) {
	l.elems = append(l.elems[:i], l.elems[i+1:]...)
}

// Task: create chaining methods Add, Update, Pop, Delete
func main() {
	list := List[string]{
		elems: []string{"qwe", "rty", "tyu"},
	}
	fmt.Println(list.elems)

	list.Add("dfgdfgdfgfd")
	fmt.Println("\n", list.elems)

	list.Delete(2)
	fmt.Println("\n", list.elems)

	l := list.Pop()
	fmt.Println("\n", l, list.elems)

	list.Update(0, "123")
	fmt.Println("\n", list.elems)

}
