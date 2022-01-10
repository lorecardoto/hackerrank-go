package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)



/*
 * Complete the 'longestSubarray' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func longestSubarray(arr []int32) int32 {
    // Write your code here
    if len(arr) == 0 {
        return 0
    }

    sums := make([]int32, len(arr))

    sums[0] = arr[0]
    for i := 1; i < len(arr); i++ {
        // If sums[i-1] < 0, nums[i] + sums[i-1] must < nums[i],
        // don't pick the sub-array.
        sums[i] = max(arr[i]+sums[i-1], arr[i])
    }

    // Find the maximum sum of sub-arrays.
    var maxVal int32 = math.MinInt32

    for i := 0; i < len(sums); i++ {
            maxVal = max(sums[i], maxVal)
        }

        return maxVal 
}

func max(x, y int32) int32 {
    if x >= y {
        return x
    }
    return y
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var arr []int32

    for i := 0; i < int(arrCount); i++ {
        arrItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := longestSubarray(arr)

    fmt.Fprintf(writer, "%d\n", result)

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
