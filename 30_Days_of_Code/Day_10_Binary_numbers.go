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
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)
	fmt.Scanln(&n)
	fmt.Printf("%d\n", binaryNumbers(n))
}

func binaryNumbers(N int32) int32 {
	var maxCount int32 = 0
	var count int32 = 0
	for i := 0; i < 64; i++ {
		//https://stackoverflow.com/questions/53167273/what-does-0x1-in-golang-mean/53167489
		if (N>>i)&0x1 == 1 {
			count++
		} else {
			if count > maxCount {
				maxCount = count
			}
			count = 0
		}
	}

	return maxCount
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
