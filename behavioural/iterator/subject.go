package main

import "fmt"

type Subject interface {
	Print()
}

type User struct {
	name string
	age  uint
}

func (u *User) Print() {
	fmt.Printf("%s - %d\n", u.name, u.age)
}
