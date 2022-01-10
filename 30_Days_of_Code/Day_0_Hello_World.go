package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {

    message := bufio.NewReader(os.Stdin)

    variable, _ := message.ReadString('\n')

    fmt.Println("Hello, World.")
    fmt.Println(variable)
}
