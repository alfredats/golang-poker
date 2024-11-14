package main 

import "fmt"
import "alfredats/golang-poker/src/classes"

func Hello() {
	fmt.Println("Hello, World!")
  d := classes.Dealer{}
  d.Init()
  fmt.Println(d.Deal(2))
}
