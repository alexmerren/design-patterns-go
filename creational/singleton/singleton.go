package main

import "sync"

var uniqueInstance *singleton
var so sync.Once
var sm sync.Mutex

type singleton struct {
	count int
}

func GetInstance() *singleton {
	sm.Lock()
	defer sm.Unlock()
	if uniqueInstance == nil {
		uniqueInstance = new(singleton)
	}
	uniqueInstance.count++
	return uniqueInstance
}
