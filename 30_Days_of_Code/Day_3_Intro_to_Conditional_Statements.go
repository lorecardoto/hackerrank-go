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

	NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	N := int32(NTemp)
	//If  is odd, print Weird AND If  is even and in the inclusive range of  to , print Weird
	if N%2 != 0 || (N >= 5 && N <= 20) {
		fmt.Println("Weird")
	} else {
		//If  is even and in the inclusive range of  to , print Not Weird AND If  is even and greater than , print Not Weird
		if (N >= 2 && N <= 5) || N >= 20 {
			fmt.Println("Not Weird")
		}
	}
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
