package main

import "testing"

func TestAdd(t *testing.T) {
	list := List[string]{
		elems: []string{},
	}

	list.Add("hello")
	list.Add("world")

	if len(list.elems) != 2 {
		t.Fatalf("wrong length: %d", len(list.elems))
	}

	if list.elems[0] != "hello" {
		t.Fatal("wrong element", list.elems[0])
	}

	if list.elems[1] != "world" {
		t.Fatal("wrong element", list.elems[1])
	}
}

func TestPop(t *testing.T) {
	list := List[string]{
		elems: []string{},
	}

	list.Add("hello")
	list.Add("world")
	list.Add("and")
	list.Add("everyone")

	if len(list.elems) != 4 {
		t.Fatalf("wrong length: %d", len(list.elems))
	}

	list.Pop()
	list.Pop()

	if len(list.elems) != 2 {
		t.Fatalf("wrong length: %d", len(list.elems))
	}

	if list.elems[0] != "hello" {
		t.Fatal("wrong element", list.elems[0])
	}

	if list.elems[1] != "world" {
		t.Fatal("wrong element", list.elems[1])
	}
}

func TestDelete(t *testing.T) {
	list := List[string]{
		elems: []string{},
	}

	list.Add("hello")
	list.Add("world")
	list.Add("and")
	list.Add("everyone")

	if len(list.elems) != 4 {
		t.Fatalf("wrong length: %d", len(list.elems))
	}

	list.Delete(0)
	list.Delete(2)

	if len(list.elems) != 2 {
		t.Fatalf("wrong length: %d", len(list.elems))
	}

	if list.elems[0] != "world" {
		t.Fatal("wrong element", list.elems[0])
	}

	if list.elems[1] != "and" {
		t.Fatal("wrong element", list.elems[1])
	}
}

func TestUpdate(t *testing.T) {
	list := List[string]{
		elems: []string{},
	}

	list.Add("hello")
	list.Add("world")
	list.Add("and")
	list.Add("everyone")

	if len(list.elems) != 4 {
		t.Fatalf("wrong length: %d", len(list.elems))
	}

	list.Update(0, "qwerty")
	list.Update(3, "booboo")

	if len(list.elems) != 4 {
		t.Fatalf("wrong length: %d", len(list.elems))
	}

	if list.elems[0] != "qwerty" {
		t.Fatal("wrong element", list.elems[0])
	}

	if list.elems[3] != "booboo" {
		t.Fatal("wrong element", list.elems[1])
	}
}
