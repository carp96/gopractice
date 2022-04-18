package main

import "fmt"

type cal interface {
	Add() interface{}
	Sub() interface{}
	Print()
}

type Cnt struct {
	a int
	b int
}

func (c *Cnt) Add() interface{} {
	return c.a + c.b
}

func (c *Cnt) Sub() interface{} {
	return c.a - c.b
}

func (c *Cnt) Print() {
	fmt.Println(c)
}

func CalInfo(i cal) {
	fmt.Println(i)
}

func main() {
	c := Cnt{
		a: 6,
		b: 2,
	}

	fmt.Printf("%T", c.Add())
	fmt.Println(c.Sub())
	(&c).Print() // 这里会做隐式转换，Cnt类型和 *Cnt类型都是可以直接调用的
	CalInfo(&c)
}
