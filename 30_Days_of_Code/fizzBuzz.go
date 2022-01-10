package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'fizzBuzz' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func fizzBuzz(n int32) {
	// Write your code here
	p := fmt.Print
	var i int32
	for i = 1; i < n+1; i++ {
		if i%3 < 1 {
			p("Fizz")
		}
		if i%5 < 1 {
			p("Buzz")
		}
		if i%3 > 0 && i%5 > 0 {
			p(i)
		}
		p("\n")
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)
	//var i int32
	//for i = 1; i < n+1; i++{
	fizzBuzz(n)
	//}
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
