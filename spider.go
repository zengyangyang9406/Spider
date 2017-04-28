package main

import (
	"Spider/writeto"
	"fmt"
)

func main() {

	fmt.Println(writeto.Writetofile(writeto.GetJokes()))

	fmt.Println("golang")
}
