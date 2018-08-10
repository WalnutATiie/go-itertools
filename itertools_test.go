package go_itertools

import (
	"fmt"
	"testing"
)

func TestNext(t *testing.T) {
	i := New(1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	defer Finalizer(i)
	fmt.Println(i.Next())
	fmt.Println(i.Next())
	fmt.Println(i.Next())
	fmt.Println(i.Next())
	fmt.Println(i.Next())
	fmt.Println(i.Next())
	fmt.Println(i.Next())
	fmt.Println(i.Next())
	fmt.Println(i.Next())
	fmt.Println(i.Next())
}

func TestCountInt(t *testing.T) {
	i := Count(0, 10)
	defer Finalizer(i)
	m := 0
	for value := range Iter(i) {
		fmt.Println(value)
		m += 1
		if m > 10 {
			break
		}
	}
}

func TestCountFloat32(t *testing.T) {
	i := Count(float32(0), float32(0.1))
	defer Finalizer(i)
	m := 0
	for value := range Iter(i) {
		fmt.Println(value)
		m += 1
		if m > 10 {
			break
		}
	}
}

func TestCycleString(t *testing.T) {
	i := Cycle("abcd")
	defer Finalizer(i)
	m := 0
	for value := range Iter(i) {
		fmt.Println(value)
		m += 1
		if m > 10 {
			break
		}
	}
}

func TestCycleSlice(t *testing.T) {
	i := Cycle([]string{"a", "b", "c", "d"})
	defer Finalizer(i)
	m := 0
	for value := range Iter(i) {
		fmt.Println(value)
		m += 1
		if m > 10 {
			break
		}
	}
}

func TestCycleMap(t *testing.T) {
	i := Cycle(map[string]interface{}{"a": 1, "b": 2})
	defer Finalizer(i)
	m := 0
	for value := range Iter(i) {
		fmt.Println(value)
		m += 1
		if m > 10 {
			break
		}
	}
}

func TestDropWhile(t *testing.T) {
	a := New(1, 11, 12, 3, 14, 2, 45, 2)
	defer Finalizer(a)
	i := DropWhile(func(i interface{}) bool {
		return i.(int) > 10
	}, a)
	defer Finalizer(i)
	for value := range Iter(i) {
		fmt.Println(value)
	}
}

func TestRepeat(t *testing.T) {
	i := Repeat(3, 10)
	defer Finalizer(i)
	for value := range Iter(i) {
		fmt.Println(value)
	}
}

func TestImap(t *testing.T) {
	a := New(1, 2, 3)
	defer Finalizer(a)
	b := New(4, 5, 6)
	defer Finalizer(b)
	c := Imap(func(a interface{}, b interface{}) interface{} {
		return a.(int) + b.(int)
	}, a, b)
	defer Finalizer(c)
	for value := range Iter(c) {
		fmt.Println(value)
	}

}
