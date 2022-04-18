package main

import (
	"fmt"
	"time"
)

func PrintNum(tag string, n int) {
	for i := 0; i < n; i++ {
		fmt.Println("[", tag, "]: ", i)
	}
}
func main() {
	go PrintNum("Thread1", 10)
	go PrintNum("MainThread", 10)

	for i := 0; i < 10; i++ {
		fmt.Println("[ main ]: ", i)
	}
	time.Sleep(time.Second)
}
