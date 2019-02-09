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

	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create("./result/result.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var largest int64 = 0
	arr := make([]int64, n)
	for i := 0; i < int(m); i++ {
		queriesRowTemp := strings.Split(readLine(reader), " ")
		s, err := strconv.ParseInt(queriesRowTemp[0], 10, 64)
		checkError(err)
		e, err := strconv.ParseInt(queriesRowTemp[1], 10, 64)
		checkError(err)
		v, err := strconv.ParseInt(queriesRowTemp[2], 10, 64)
		checkError(err)
		for i := int(s - 1); i <= int(e-1); i++ {
			arr[i] += v
			if largest < arr[i] {
				largest = arr[i]
			}
		}
	}
	fmt.Fprintf(writer, "%d\n", largest)
	writer.Flush()
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
