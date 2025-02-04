package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
)

func main() {
	s := puppy.Bark()
	fmt.Println("Hello from Ben")
	fmt.Println("First " + s)
	fmt.Println(puppy.BigBark())
	puppy.From12()
}
