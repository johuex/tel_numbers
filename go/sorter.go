package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

//Returns the minimum element of an array and its index
func Min(array *[]uint64) (uint64, uint64) {
	min := (*array)[0]
	min_i := uint64(0)
	for i, value := range *array {
		if value < min {
			min = value
			min_i = uint64(i)
		}
	}
	return min, min_i
}

//sort with 1GB memory limit
func sorterWithRamLimit() {
	startTime := time.Now()
	allNumbers := 1000000000
	numFiles := 8
	bufferSize := allNumbers / numFiles                                                // 1'000'000'000 / 8
	fileIn, _ := os.OpenFile("gen_numbers.txt", os.O_RDONLY, 0666)                     //open for reading
	fileOut, _ := os.OpenFile("sortedWithRamLimit.txt", os.O_CREATE|os.O_WRONLY, 0666) //create and open
	var files []*os.File
	tempFileName := "temp_"
	writer := bufio.NewWriter(fileOut)
	defer fileIn.Close()
	defer fileOut.Close()

	//split and sort temp files
	splitSortTime := time.Now()
	numbers := make([]uint64, bufferSize) //numbers for one temp file
	q := 0                                //counting numbers for Buffer
	j := 0
	reader := bufio.NewReader(fileIn)
	for j != numFiles {
		elem, _ := reader.ReadString('\n')
		if elem != "" && elem != "8936" {
			/*
				"" - это знак окончания чтения исходного файла, на нем завершаем чтение.
				НО Почему 8936? Почему то в исходном файле находится данное число.
				Поймать его при записе в функции gen() не удалось. Оно появляется только в этом месте при чтении.
				Как решить эту проблему пока идей нет :(
			*/
			num, _ := strconv.ParseUint(elem[:len(elem)-1], 10, 64)
			numbers[q] = num
			q++
		} else {
			//shrink to fit our slice to avoid extra zeros if catch last empty element of file_in
			numbers = numbers[:q]
		}
		if q == bufferSize || elem == "" {
			//sort and write one part to temp file
			sort.Slice(numbers, func(i, j int) bool { return numbers[i] < numbers[j] }) //ascending sort
			tmpFile, _ := os.OpenFile(tempFileName+strconv.FormatInt(int64(j), 10), os.O_CREATE|os.O_RDWR, 0666)
			w := bufio.NewWriter(tmpFile)
			for i := 0; i < len(numbers); i++ {
				_, _ = fmt.Fprintln(w, numbers[i])
			}
			files = append(files, tmpFile)
			q = 0
			j++
			numbers = nil
			numbers = make([]uint64, bufferSize)
			w = nil
		}
	}
	reader = nil
	fmt.Printf("Split and sort temp files time: %d seconds\n", time.Now().Unix()-splitSortTime.Unix())

	//start sorting
	externalSortTime := time.Now()
	minimumNumbers := make([]uint64, len(files)) // list of minimum elements in temp files
	var buffer []uint64
	fReader := make([]*bufio.Reader, len(files))
	//fill in minimum elements of temp files
	for i := 0; i < len(files); i++ {
		_, _ = files[i].Seek(0, 0)             //cursor to begin of file
		fReader[i] = bufio.NewReader(files[i]) //scanner for each file
		elem, _ := fReader[i].ReadString('\n')
		tmpNum, _ := strconv.ParseUint(elem[:len(elem)-1], 10, 64) // minimum of temp file
		minimumNumbers[i] = tmpNum
	}
	//external sorting
	for {
		_, min_i := Min(&minimumNumbers) //search minimum element of all temp files
		if len(buffer) == bufferSize {   //wait for fill in buffer to write in file_out
			for i := 0; i < len(buffer); i++ {
				_, _ = fmt.Fprintln(writer, buffer[i])
			}
			buffer = nil
			buffer = append(buffer, minimumNumbers[min_i])
		} else {
			buffer = append(buffer, minimumNumbers[min_i])
		}
		k, _ := fReader[min_i].ReadString('\n')
		// delete minimum number if temp file is EOF else Scan new number of temp file
		if k == "" || k == "\n" {
			minimumNumbers = append(minimumNumbers[:min_i], minimumNumbers[min_i+1:]...)
			fReader = append(fReader[:min_i], fReader[min_i+1:]...)
		} else {
			minimumNumbers[min_i], _ = strconv.ParseUint(k[:len(k)-1], 10, 64)
		}
		//if no temp_files or minimum_numbers than end of sorting
		if len(fReader) == 0 || len(minimumNumbers) == 0 {
			if len(buffer) > 0 {
				// write in remaining buffer's elements in file_out
				for i := 0; i < len(buffer); i++ {
					_, _ = fmt.Fprintln(writer, buffer[i])
				}
				buffer = nil
			}
			break
		}
	}
	fmt.Printf("External sorting time: %d seconds\n", time.Now().Unix()-externalSortTime.Unix())

	for i := 0; i < len(files); i++ {
		_ = files[i].Close()
		os.Remove(files[i].Name())
	}
	fmt.Printf("All time: %d seconds\n", time.Now().Unix()-startTime.Unix())
}

func sorterNoLimit() {
	//sort with no memory limit
	startTime := time.Now()
	fileIn, _ := os.OpenFile("gen_numbers.txt", os.O_RDONLY, 0666)                   //open for reading
	fileOut, _ := os.OpenFile("sortedNoRamLimit.txt", os.O_CREATE|os.O_WRONLY, 0666) //create and open
	defer fileIn.Close()
	defer fileOut.Close()

	//read from file
	readTime := time.Now()
	numbers := make([]uint64, 1000000000) // size = all phone numbers
	j := uint64(0)
	reader := bufio.NewReader(fileIn)
	for j != 1000000000 {
		elem, _ := reader.ReadString('\n')
		if elem != "" && elem != "8936" {
			num, _ := strconv.ParseUint(elem[:len(elem)-1], 10, 64)
			numbers[j] = num
			j++
		} else {
			//shrink to fit our slice to avoid extra zeros if catch last empty element of file_in
			numbers = numbers[:j]
		}
	}
	fmt.Printf("Read from file time: %d seconds\n", time.Now().Unix()-readTime.Unix())

	//sort uint64[]
	sortTime := time.Now()
	sort.Slice(numbers, func(i, j int) bool { return numbers[i] < numbers[j] }) //ascending sort
	fmt.Printf("Sort time: %d seconds\n", time.Now().Unix()-sortTime.Unix())

	//write to file
	writeTime := time.Now()
	lenNumbers := len(numbers)
	w := bufio.NewWriter(fileOut)
	for i := 0; i < lenNumbers; i++ {
		_, _ = fmt.Fprintln(w, numbers[i])
	}
	fmt.Printf("Write to file time: %d seconds\n", time.Now().Unix()-writeTime.Unix())

	fmt.Printf("All time: %d seconds\n", time.Now().Unix()-startTime.Unix())
}

func sqlSort() {
	allNumbers := 1000000000
	fileIn, _ := os.OpenFile("gen_numbers.txt", os.O_RDONLY, 0666)
	fileOut, _ := os.OpenFile("sortedSQL.txt", os.O_CREATE|os.O_WRONLY, 0666) //create and open
	defer fileIn.Close()
	defer fileOut.Close()
	result := make([]string, allNumbers)

	reader := bufio.NewReader(fileIn)
	writer := bufio.NewWriter(fileOut)

	_, _ = os.Create("numbers.db")
	db, err := sql.Open("sqlite3", "numbers.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// CREATE + INSERT
	insertTime := time.Now()
	_, err = db.Exec("CREATE TABLE numbers(phone VARCHAR(11) PRIMARY KEY);") // CREATE TABLE
	for i := 0; i < allNumbers; i++ {
		elem, _ := reader.ReadString('\n')
		_, err = db.Exec("INSERT INTO numbers (phone) values ($1)", elem[:len(elem)-1]) // INSERT VALUES
	}
	fmt.Printf("Inserting to DB: %d seconds\n", time.Now().Unix()-insertTime.Unix())

	//SELECT ASC
	selectTime := time.Now()
	rows, _ := db.Query("SELECT * FROM numbers ORDER BY phone ASC") // SELECT WITH ASC
	rows.Scan(result)                                               // from sql.Query to []string
	fmt.Printf("SELECT ASC + translating from sql.Query to []string: %d seconds\n", time.Now().Unix()-selectTime.Unix())

	//writing to file
	writeTime := time.Now()
	for _, value := range result {
		_, _ = fmt.Fprintln(writer, value)
	}
	fmt.Printf("Write to file time: %d seconds\n", time.Now().Unix()-writeTime.Unix())

}
