package main

import "fmt"

// ** CLIENT CODE **

func main() {
    userFactory := &UserFactory{}
    user := userFactory.Create()
    user.DoSomething()

    adminFactory := &AdminFactory{}
    admin := adminFactory.Create()
    admin.DoSomething()
}

// ** PRODUCT INTERFACE **

type ProductInterface interface {
    DoSomething() 
}

// ** CONCRETE PRODUCT **

type User struct {}

func (u *User) DoSomething() {
    fmt.Println("I am doing something normal user-y")
}

type Admin struct {}

func (a *Admin) DoSomething() {
    fmt.Println("I am doing something admin-y")
}

// ** ABSTRACT FACTORY **

type FactoryInterface interface {
    Create() ProductInterface
}

// ** CONCRETE FACTORY **

type UserFactory struct {}

func (uf *UserFactory) Create() ProductInterface {
    return &User{}
}

type AdminFactory struct {}

func (af *AdminFactory) Create() ProductInterface {
    return &Admin{}
}

type EntityFactory struct {}

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
