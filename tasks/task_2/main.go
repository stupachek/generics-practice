package main

import (
	"errors"
	"fmt"
)

type Map[K comparable, V any] struct {
	elems map[K]V
}

func (m *Map[K, V]) Add(k K, v V) error {
	if _, ok := m.elems[k]; ok {
		return errors.New("have a value for the key")
	}
	m.elems[k] = v
	return nil

}

func (m *Map[K, V]) Delete(k K) error {
	if _, ok := m.elems[k]; !ok {
		return errors.New("don't have a value for the key")
	}
	delete(m.elems, k)
	return nil
}

func (m *Map[K, V]) Update(k K, v V) error {
	if _, ok := m.elems[k]; !ok {
		return errors.New("don't have a value for the key")
	}
	m.elems[k] = v
	return nil
}

func (m *Map[K, V]) Get(k K) (V, error) {
	if _, ok := m.elems[k]; !ok {
		var zero V
		return zero, errors.New("don't have a value for the key")
	}
	return m.elems[k], nil

}

// Task: create Map[K, V] with methods Add(K, V), Delete(K), Update(K, V), Get(K)
func main() {
	m := Map[string, int]{
		elems: map[string]int{"first": 1, "second": 2},
	}
	fmt.Println(m.elems)

	err := m.Add("third", 3)

	if err != nil {
		fmt.Println("\n", err)
	} else {
		fmt.Println("\n", m.elems)
	}

	err = m.Add("third", 2)
	if err != nil {
		fmt.Println("\n", err)
	} else {
		fmt.Println("\n", m.elems)
	}

	err = m.Delete("third")
	if err != nil {
		fmt.Println("\n", err)
	} else {
		fmt.Println("\n", m.elems)
	}

	err = m.Delete("123123")
	if err != nil {
		fmt.Println("\n", err)
	} else {
		fmt.Println("\n", m.elems)
	}

	err = m.Update("second", 123)
	if err != nil {
		fmt.Println("\n", err)
	} else {
		fmt.Println("\n", m.elems)
	}

	err = m.Update("123123", 123123)
	if err != nil {
		fmt.Println("\n", err)
	} else {
		fmt.Println("\n", m.elems)
	}

	v, err := m.Get("first")
	if err != nil {
		fmt.Println("\n", err)
	} else {
		fmt.Println("\n", v)
	}

	v, err = m.Get("fasdfsdfsdfds")
	if err != nil {
		fmt.Println("\n", err)
	} else {
		fmt.Println("\n", v)
	}
}
