package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	var err error
	file, err := os.Open("./data/input.txt")
	checkError(err)
	defer file.Close()
	reader := bufio.NewReader(file)
	line, _, err := reader.ReadLine()
	checkError(err)
	nm := strings.Split(string(line[:]), " ")
	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)
	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)
	var largest, sum int64
	arr := make([]int64, n+1)
	for i := 0; i < int(m); i++ {
		line, _, err = reader.ReadLine()
		checkError(err)
		queriesRowTemp := strings.Split(string(line[:]), " ")
		s, err := strconv.ParseInt(queriesRowTemp[0], 10, 64)
		checkError(err)
		e, err := strconv.ParseInt(queriesRowTemp[1], 10, 64)
		checkError(err)
		v, err := strconv.ParseInt(queriesRowTemp[2], 10, 64)
		checkError(err)
		arr[int(s)-1] += v
		arr[int(e)] -= v
	}
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if largest < sum {
			largest = sum
		}
	}
	fmt.Println(largest)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
