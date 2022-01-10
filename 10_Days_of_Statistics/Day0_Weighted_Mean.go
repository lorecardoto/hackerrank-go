
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "math"
    //"sort"
)

func main() {
    var values []int
    var values3 []int
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        p := scanner.Text()
        for _, v := range strings.Split(p, " ") {
            numb, _ := strconv.Atoi(v)
            values = append(values, int(numb))
            values3 = append(values3, int(numb))
        
    }

    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "Error:", err)
    }
    wValues := []int{10, 40, 30, 50, 20}
    weights := []int{1, 2, 3, 4, 5}

    res := weightedMean(wValues, weights)

    fmt.Printf("%.1f",res)
    //fmt.Printf("%.1f",math.Round(res*10) / 10)
}

func weightedMean(values []int, weights []int) float64 {
    lenValues := len(values)
    lenWeights := len(weights)

    if lenValues != lenWeights {
        panic("cannot be diff length between values and weights")
    }
    counter := 0
    sumWeight := 0
    for i := 0; i < lenValues; i++ {
        counter += values[i] * weights[i]
        sumWeight += weights[i]
    }

    return math.Round((float64(counter)/float64(sumWeight))*10) / 10
}

