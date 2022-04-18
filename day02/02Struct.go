package main

import "fmt"

type Cat struct {
	Color string
	Name  string
}

func (cat Cat) NewCatByName(name string) {
	fmt.Println("Set cat name:")
	cat.Name = name
	fmt.Println(cat)
}

func (cat Cat) NewCatByColor(color string) {
	fmt.Println("Set cat color:")
	cat.Color = color
	fmt.Println(cat)
}

func (cat *Cat) ModifyColor(color string) {
	fmt.Println("Modify color:")
	cat.Color = color
}

func main() {
	var cat Cat
	// cat := &Cat{}
	cat.NewCatByColor("balck")
	cat.NewCatByName("lk")
	fmt.Println(cat)

	cat.Color = "Yellow"
	cat.Name = "Golden"
	fmt.Println(cat)

	ModCat := Cat{
		Name:  "kitty",
		Color: "white",
	}
	fmt.Println(ModCat)
	ModCat.ModifyColor("blue")
	fmt.Println(ModCat)
}
