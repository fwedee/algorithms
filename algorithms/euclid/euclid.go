package main

import (
	"fmt"
	"log"
	"strconv"
)

func euclid(x int, y int) int {
	if x == y {
		return x
	} else if x > y {
		return euclid(y, x)
	} else {
		return euclid(x, y-x)
	}

}

func main() {
	fmt.Println("Greatest common devisor")
	fmt.Print("Input two numbers: ")
	var a, b string
	_, err := fmt.Scanln(&a, &b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	x, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(b)
	if err != nil {
		log.Fatal()
	}

	result := euclid(x, y)
	fmt.Printf(`ggT(%d, %d) = %d`, x, y, result)

}
