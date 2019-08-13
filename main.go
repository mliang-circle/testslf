package main

import (
	"fmt"
)

func Initialize() int {
	a := 1
	b := 2
	c := a + b
	fmt.Print("running init\n")
	return c
}

func main() {

	c := Initialize()
	fmt.Printf("init result %d\n", c)
	fmt.Print("finish main")


}

