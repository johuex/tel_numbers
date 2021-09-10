package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

func Min(array []uint64) (uint64, uint64) {
	var min uint64
	var min_i uint64
	for i, value := range array {
		if min > value {
			min = value
			min_i = uint64(i)
		}
	}
	return min, min_i
}

func sorter_1() {
	//sort with memory limit (for example, 1 Gb)
	// 8(or 16????)*10^9 / 1024^3 = 7,45 or 8 temp files
	startTime := time.Now()
	buffer_size := 125000000
	file_in, _ := os.OpenFile("gen_numbers.txt", os.O_RDONLY, 0666)           //open for reading
	file_out, _ := os.OpenFile("sorter_1.txt", os.O_CREATE|os.O_WRONLY, 0666) //create and open
	var files []*os.File
	tempFileName := "temp_"
	defer file_in.Close()
	defer file_out.Close()

	//split and sort temp
	split_sort_time := time.Now()
	var numbers []uint64
	i := 0 //for Buffer
	j := int64(0)
	scanner := bufio.NewScanner(file_in)
	for scanner.Scan() {
		num, _ := strconv.ParseUint(scanner.Text(), 10, 64)
		numbers = append(numbers, num)
		i++
		if i == buffer_size {
			//sort and write 1/8 part
			sort.Slice(numbers, func(i, j int) bool { return numbers[i] < numbers[j] }) //ascending sort
			tmp_file, _ := os.OpenFile(tempFileName+strconv.FormatInt(j, 10), os.O_CREATE|os.O_RDWR, 0666)
			w := bufio.NewWriter(tmp_file)
			for i := 0; i < buffer_size; i++ {
				_, _ = fmt.Fprintln(w, numbers[i])
			}
			files = append(files, tmp_file)
			i = 0
			j++
		}
	}
	fmt.Printf("Split and sort temp files time: %d seconds\n", time.Now().Unix()-split_sort_time.Unix())

	//external sorting
	files_copy := files
	external_sort_time := time.Now()
	var temp_numbers []uint64
	var buffer []uint64
	for i := 0; i < len(files); i++ {
		_, _ = files[j].Seek(0, 0) //cursor to begin of file
		tmp_scanner := bufio.NewScanner(files[j])
		tmp_num, _ := strconv.ParseUint(tmp_scanner.Text(), 10, 64)
		temp_numbers = append(temp_numbers, tmp_num)
	}
	for {
		_, min_i := Min(temp_numbers)
		if len(buffer) == buffer_size {
			w := bufio.NewWriter(file_out)
			for i := 0; i < len(buffer); i++ {
				_, _ = fmt.Fprintln(w, buffer[i])
			}
			buffer = nil
			buffer = append(buffer, temp_numbers[min_i])
		} else {
			buffer = append(buffer, temp_numbers[min_i])
		}

		tmp_scanner := bufio.NewScanner(files[min_i])
		k := tmp_scanner.Text()
		if k == "" || k == "\n" {
			temp_numbers = append(temp_numbers[:min_i], temp_numbers[min_i+1:]...)
			files = append(files[:min_i], files[min_i+1:]...)
		} else {
			temp_numbers[min_i], _ = strconv.ParseUint(k, 10, 64)
		}
		if len(files) == 0 || len(temp_numbers) == 0 {
			if len(buffer) > 0 {
				w := bufio.NewWriter(file_out)
				for i := 0; i < len(buffer); i++ {
					_, _ = fmt.Fprintln(w, buffer[i])
				}
				buffer = nil
			}
			break
		}
	}
	fmt.Printf("Split and sort temp files time: %d seconds\n", time.Now().Unix()-external_sort_time.Unix())

	for i := 0; i < len(files); i++ {
		_ = files_copy[i].Close()
	}
	fmt.Printf("All time: %d seconds\n", time.Now().Unix()-startTime.Unix())
}

func sorter_2() {
	//sort with no memory limit
	startTime := time.Now()
	file_in, _ := os.OpenFile("gen_numbers.txt", os.O_RDONLY, 0666)           //open for reading
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
