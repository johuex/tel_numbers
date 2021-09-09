package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("gen_numbers.txt"); err == nil {
		fmt.Println("Numbers have already generated!")
	} else if os.IsNotExist(err) {
		//does *not* exist
		gen()
	}
	fmt.Println()
	fmt.Println("Sorting with memory limit (1 Gb)")
	sorter_1()
	//fmt.Println()
	//fmt.Println("Sorting with no memory limit")
	//sorter_2()
}
