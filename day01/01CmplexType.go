package main

import (
	"container/list"
	"fmt"
)

func PrintStr(s string) {
	fmt.Println(s)
}

func main() {
	// 数组
	var arr = [...]int{1, 3, 5, 7, 9}
	// 切片
	slc := arr[1:5]
	slc = append(slc, 999)
	fmt.Println(arr, slc)

	// map
	mp1 := make(map[int][]int, 3)
	mp1[2] = slc
	mp1[7] = []int{2, 3}
	mp1[3] = arr[1:3]
	delete(mp1, 0)
	for k, v := range mp1 {
		v = append(v, -1)
		fmt.Println(k, v)
	}
	delete(mp1, 7)
	fmt.Println(mp1)

	// list
	lst := list.New()
	lst.PushBack("string")
	lst.PushBack(233333)
	lst.PushFront(mp1)
	fmt.Println(lst) // &{{0xc0000745d0 0xc0000745a0 <nil> <nil>} 3}
	for i := lst.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	// 函数变量
	var f func(s string)
	f = PrintStr
	f("Test func var")

	// 匿名函数变量
	fn := func(x int, y int) int {
		return x + y
	}

	z := fn(123, 321)
	fmt.Println(z)

	// interface{}
	var it interface{}
	it = "sdwr"
	it = 23
	it = "fer"
	fmt.Printf("type(it): %T\tit: %v\n", it, it)
	s := "string" + it.(string)
	fmt.Println(s)
}
