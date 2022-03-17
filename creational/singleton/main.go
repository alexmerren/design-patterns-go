package main

import (
	"fmt"
)

func main() {
	GetInstance()
	GetInstance()
	fmt.Println(uniqueInstance.count)
	return
}
