package main

import (
	"fmt"
	"sync"
)

var uniqueInstance *singleton
var so sync.Once
var sm sync.Mutex

func main() {
	GetInstance()
	GetInstance()
	fmt.Println(uniqueInstance.count)
	return
}

type singleton struct {
	count int
}

//func GetInstance() *singleton {
//so.Do(func() {
//uniqueInstance = new(singleton)
//})
//uniqueInstance.count++
//return uniqueInstance
//}

func GetInstance() *singleton {
	sm.Lock()
	defer sm.Unlock()
	if uniqueInstance == nil {
		uniqueInstance = new(singleton)
	}
	uniqueInstance.count++
	return uniqueInstance
}
