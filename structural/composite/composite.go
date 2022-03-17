package main

import "fmt"

type Composite interface {
	AddChild()
	RemoveChild()
	GetChild()
}

type Folder struct {
	contents []Component
	name     string
}

func NewFolder(name string) *Folder {
	return &Folder{
		name: name,
	}
}

func (f *Folder) AddChild(newComponent ...Component) {
	f.contents = append(f.contents, newComponent...)
}

func (f *Folder) Print() {
	fmt.Printf("Folder '%s':\n", f.name)
	for _, file := range f.contents {
		file.Print()
	}
}
