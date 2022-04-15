package main

import "fmt"

func main() {
	/********************go允许函数调用在定义之前********************/
	TestFunc()

	/**************************闭包********************************/
	cl1 := TestClosure("hei", 1, 2, 3, 4, 5)
	fmt.Println(cl1())
	fmt.Println(cl1())
	fmt.Println(cl1())

	n := 10
	cl2 := TestClosure("ha", n)
	fmt.Println(cl2())
	fmt.Println(cl2())

	fmt.Println(cl1())

	/********************任意类型可变长参数***************************/
	TestInterface("string", 23, 33.6, []string{"int", "float"})
}

func TestFunc() {
	fmt.Println("TEST!")
}

// 闭包函数捕捉参数测试
func TestClosure(name string, cnt int, slc ...int) func() (string, int, int, int) {
	age := 18
	return func() (string, int, int, int) {
		var sum = 0
		for _, v := range slc {
			v++
			sum += v
		}
		cnt++
		return name, age, cnt, sum
	}
}

// 任意类型可变长参数
func TestInterface(args ...interface{}) {
	for _, v := range args {
		fmt.Printf("Type: %T, Val: %v\n", v, v)
	}
}
