package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

func sorter_1() {
	//sort with memory limit (for example, 1 Gb)
	// 8(or 16????)*10^9 / 1024^3 = 7,45 or 8 temp files
	startTime := time.Now()
	BufferSize := 125000000
	file_in, _ := os.OpenFile("gen_numbers.txt", os.O_RDONLY, 0666) //open for reading
	file_out, _ := os.OpenFile("sorter_1.txt", os.O_CREATE|os.O_WRONLY, 0666) //create and open
	var files[8]*os.File
	tempFileName := "temp_"
	defer file_in.Close()
	defer file_out.Close()

	//split and sort temp
	split_sort_time := time.Now()
	var numbers []uint64
	i := 0 //for Buffer
	j := int64(0) //for itt files
	scanner := bufio.NewScanner(file_in)
	for scanner.Scan() {
		num,_ := strconv.ParseUint(scanner.Text(), 10, 64)
		numbers = append(numbers, num)
		i++
		if i == BufferSize {
			//sort and write 1/8 part
			sort.Slice(numbers, func(i, j int) bool { return numbers[i] < numbers[j] }) //ascending sort
			tmp_file, _ := os.OpenFile(tempFileName+strconv.FormatInt(j,10), os.O_CREATE|os.O_RDWR, 0666)
			files[j] = tmp_file
			w := bufio.NewWriter(files[j])
			for i := 0; i < BufferSize; i++ {
				_, _ = fmt.Fprintln(w, numbers[i])
			}
			_,_ = files[j].Seek(0,0) //cursor to begin of file
			i = 0
			j++
		}
	}
	fmt.Printf("Split and sort temp files time: %d seconds\n", time.Now().Unix()-split_sort_time.Unix())

	//external sorting
	external_sort_time := time.Now()

	fmt.Printf("Split and sort temp files time: %d seconds\n", time.Now().Unix()-external_sort_time.Unix())

	for i := 0; i < len(files); i++ {
		_ = files[i].Close()
	}
	fmt.Printf("All time: %d seconds\n", time.Now().Unix()-startTime.Unix())
}

func sorter_2() {
	//sort with no memory limit
	startTime := time.Now()
	file_in, _ := os.OpenFile("gen_numbers.txt", os.O_RDONLY, 0666) //open for reading
	file_out, _ := os.OpenFile("sorter_2.txt", os.O_CREATE|os.O_WRONLY, 0666) //create and open
	defer file_in.Close()
	defer file_out.Close()

	//read from file
	readTime := time.Now()
	var numbers []uint64
	scanner := bufio.NewScanner(file_in)
	for scanner.Scan() {
		num, _ := strconv.ParseUint(scanner.Text(), 10, 64)
		numbers = append(numbers, num)
	}
	fmt.Printf("Read from file time: %d seconds\n", time.Now().Unix()-readTime.Unix())

	//sort uint64[]
	sortTime := time.Now()
	sort.Slice(numbers, func(i, j int) bool { return numbers[i] < numbers[j] }) //ascending sort
	fmt.Printf("Sort time: %d seconds\n", time.Now().Unix()-sortTime.Unix())

	writeTime := time.Now()
	lenNumbers := len(numbers)
	w := bufio.NewWriter(file_out)
	for i := 0; i < lenNumbers; i++ {
		_, _ = fmt.Fprintln(w, numbers[i])
	}
	fmt.Printf("Write to file time: %d seconds\n", time.Now().Unix()-writeTime.Unix())

	fmt.Printf("All time: %d seconds\n", time.Now().Unix()-startTime.Unix())
}
