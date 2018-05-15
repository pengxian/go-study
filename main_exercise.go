package main

import "fmt"

var a string = "G"

func main() {
  n()
  m()
  n()
}

func n() { fmt.Println(a) }

func m() {
   a := "O"
   fmt.Println(a)
}
