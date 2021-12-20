package main

import "fmt"

// ** CLIENT CODE **

func main() {
	entityFactory := &EntityFactory{}
	user := entityFactory.Create("user")
	user.DoSomething()

	admin := entityFactory.Create("admin")
	admin.DoSomething()
}

// ** PRODUCT INTERFACE **

type ProductInterface interface {
	DoSomething()
}

// ** CONCRETE PRODUCT **

type User struct {
	ProductInterface
}

func (u *User) DoSomething() {
	fmt.Println("I am doing something normal user-y")
}

type Admin struct {
	ProductInterface
}

func (a *Admin) DoSomething() {
	fmt.Println("I am doing something admin-y")
}

// ** ABSTRACT FACTORY **

type FactoryInterface interface {
	Create(string) ProductInterface
}

// ** CONCRETE FACTORY **

type EntityFactory struct{}

func (ef *EntityFactory) Create(input string) ProductInterface {
	switch input {
	case "user":
		return &User{}
	case "admin":
		return &Admin{}
	default:
		return nil
	}
}
