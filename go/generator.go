package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

//var numbers [1000000000]int

func gen() {
	//8-9**-***-**-**
	lenNumbers := 1000000000
	numbers := make([]uint64, lenNumbers)
	startTime := time.Now()
	file, _ := os.OpenFile("gen_numbers.txt", os.O_CREATE|os.O_WRONLY, 0666) //create and open
	defer file.Close() //close at the end of func

	for i := 0; i < lenNumbers; i++ {
		numbers[i] = 89000000000 + uint64(i)
	}
	fmt.Printf("Generation time: %d seconds\n", time.Now().Unix()-startTime.Unix())

	shuffleTime := time.Now()
	halfLenNumbers := (lenNumbers - 1) / 2 //get int, not float
	for i:=0; i < halfLenNumbers; i++ {
		end_i := lenNumbers - 1 - i
		rand_i := rand.Intn(end_i/2 - 1)
		numbers[end_i], numbers[rand_i] = numbers[rand_i], numbers[end_i]
	}
	fmt.Printf("Shuffle time: %d seconds\n", time.Now().Unix()-shuffleTime.Unix())

	writeTime := time.Now()
	w := bufio.NewWriter(file)
	for i := 0; i < lenNumbers; i++ {
		_, _ = fmt.Fprintln(w, numbers[i])
	}
	fmt.Printf("Write to file time: %d seconds\n", time.Now().Unix()-writeTime.Unix())

	fmt.Printf("All time: %d seconds\n", time.Now().Unix()-startTime.Unix())
}


