package main

import "fmt"

type Counter[I ~int8 | ~int16 | ~int32 | ~int64] struct {
	i I
}

func (c *Counter[I]) Get() I {
	return c.i
}

func (c *Counter[I]) Next() {
	c.i++
}

type MyInt int64

// Task: create Counter with method Next(). Counter expected to support all kinds of int and its inherit types
func main() {
	a := Counter[MyInt]{
		i: 123,
	}
	fmt.Println(a.Get())
	a.Next()
	fmt.Println(a.Get())
}
