package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var values []int

	//var values2 []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		p := scanner.Text()
		for _, v := range strings.Split(p, " ") {
			numb, _ := strconv.Atoi(v)
			values = append(values, int(numb))

			//values = append(remove(values,0), int(numb))
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	}
	values = remove(values, 0)
	//fmt.Println(values,"valuesini")
	fmt.Println(mean(values))
	fmt.Println(median(values))
	fmt.Println(mode(values))

}

func mean(values []int) float32 {
	//fmt.Println(values,"valuesmean")
	//fmt.Println(values2,"values")
	var count int
	for _, v := range values {
		count += v
	}
	return float32(count) / float32(len(values))
}

func median(values []int) float32 {
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})

	if len(values)%2 == 0 {
		l := float32(values[len(values)/2-1])
		r := float32(values[len(values)/2])
		return (r + l) / 2
	}

	return float32(values[+len(values)/21])
}

func mode(values []int) float64 {
	hashmap := make(map[int]int)
	repeated := -1
	modeValue := float64(99999)
	for _, key := range values {
		if v, ok := hashmap[key]; ok {
			hashmap[key] = v + 1
		} else {
			hashmap[key] = 1
		}
	}

	for k, v := range hashmap {
		if repeated <= v {
			if repeated == v {
				modeValue = math.Min(float64(k), float64(modeValue))
				repeated = v
			}
			if repeated < v {
				repeated = v
				modeValue = float64(k)
			}
		}
	}
	return modeValue
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
