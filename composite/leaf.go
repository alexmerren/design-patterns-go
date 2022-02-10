package main

import "fmt"

type File struct {
	name string
}

func NewFile(name string) *File {
	return &File{
		name: name,
	}
}

func (f *File) Print() {
	fmt.Printf("\tFile '%s'\n", f.name)
}
