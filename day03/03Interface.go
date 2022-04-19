package main

import "fmt"

type A interface {
	PrintWord()
}

type B interface {
	JoinStr(s string)
}

type C interface {
	A
	B
	SubStr() string
}

type str string

func (s str) PrintWord() {
	fmt.Println("Hello world!")
}

func (s str) JoinStr(ss string) {
	// s += str	// 报错，str和string类型不同，有点奇怪，明明是个别名
	s += str(ss) //强制类型转换
}

func (s str) SubStr() string {
	return string(s)[2:]
}

func main() {
	var cclass C
	var s str
	cclass = s
	cclass.PrintWord()
	cclass.JoinStr("hhh")
	cclass = str("hello hei")
	fmt.Println(cclass.SubStr())
}
