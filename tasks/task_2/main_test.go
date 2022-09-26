package main

import "testing"

func TestAdd(t *testing.T) {
	m := Map[string, int]{
		elems: map[string]int{"first": 1, "second": 2},
	}
	err := m.Add("hello", 1)

	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if len(m.elems) != 3 {
		t.Fatalf("wrong length: %d", len(m.elems))
	}

	if m.elems["hello"] != 1 {
		t.Fatal("wrong element", m.elems["hello"])
	}

	err = m.Add("first", 8)

	if err == nil {
		t.Fatal("expected error")
	}

	if m.elems["first"] != 1 {
		t.Fatal("wrong element", m.elems["first"])
	}

}

func TestUpdate(t *testing.T) {
	m := Map[string, int]{
		elems: map[string]int{"first": 1, "second": 2},
	}
	err := m.Update("first", 2)

	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if len(m.elems) != 2 {
		t.Fatalf("wrong length: %d", len(m.elems))
	}

	if m.elems["first"] != 2 {
		t.Fatal("wrong element", m.elems["first"])
	}

	err = m.Update("qwerty", 8)

	if err == nil {
		t.Fatal("expected error")
	}

	if _, ok := m.elems["qwerty"]; ok {
		t.Fatal("wrong element", m.elems["qwerty"])
	}

}

func TestDelete(t *testing.T) {
	m := Map[string, int]{
		elems: map[string]int{"first": 1, "second": 2},
	}
	err := m.Delete("first")

	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if len(m.elems) != 1 {
		t.Fatalf("wrong length: %d", len(m.elems))
	}

	if _, ok := m.elems["first"]; ok {
		t.Fatal("wrong element", m.elems["first"])
	}

	err = m.Delete("qwerty")

	if err == nil {
		t.Fatal("expected error")
	}

	if _, ok := m.elems["qwerty"]; ok {
		t.Fatal("wrong element", m.elems["qwerty"])
	}

}

func TestGet(t *testing.T) {
	m := Map[string, int]{
		elems: map[string]int{"first": 1, "second": 2},
	}
	el, err := m.Get("first")

	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if ex, ok := m.elems["first"]; !ok || el != ex {
		t.Fatal("wrong element", m.elems["first"])
	}

	_, err = m.Get("qwerty")

	if err == nil {
		t.Fatal("expected error")
	}

	if _, ok := m.elems["qwerty"]; ok {
		t.Fatal("wrong element", m.elems["qwerty"])
	}

}
