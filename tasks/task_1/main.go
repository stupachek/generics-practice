package main

type List[T any] struct {
	elems []T
}

func (l *List[T]) Add(t T) {
	l.elems = append(l.elems, t)
}

func (l *List[T]) Update(i int, t T) {
	l.elems[i] = t
}

func (l *List[T]) Pop() {

}

// Task: create chaining methods Add, Update, Pop, Delete
func main() {

}
