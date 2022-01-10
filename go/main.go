package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("gen_numbers.txt"); err == nil {
		fmt.Println("Numbers have already generated!")
	} else if os.IsNotExist(err) {
		//file doesn't exist
		gen()
	}
	fmt.Println()
	fmt.Println("Sorting with memory limit (1 Gb)")
	sorterWithRamLimit()
	fmt.Println()
	fmt.Println("Sorting with no memory limit")
	sorterNoLimit()
	//fmt.Println()
	//fmt.Println("Sorting with SQL")
	//sqlSort()
}
