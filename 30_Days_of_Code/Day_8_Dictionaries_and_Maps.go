package main

import "fmt"


func scanf(name *string) bool {
  readItems, _ := fmt.Scanf("%s", name)
  return readItems > 0
}

func main() {
 
  nameToNumber := make(map[string]uint)
  
  var N uint
  fmt.Scanln(&N)
  var name string
  var phone uint
  for N > 0 {
    fmt.Scanf("%s %d", &name, &phone)
    nameToNumber[name] = phone
    N--
  }

 
  var queryName string

  for scanf(&queryName) {
    phone, found := nameToNumber[queryName]
    if found {
     fmt.Printf("%s=%d\n", queryName, phone)
    } else {
     fmt.Printf("Not found\n")
    }
  }
}
