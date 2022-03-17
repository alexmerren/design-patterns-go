package main

type Iteratable interface {
	HasNext() bool
	GetNext() Subject
}

type UserIterator struct {
	index int
	users []*User
}

func (i *UserIterator) HasNext() bool {
	if i.index < len(i.users) {
		return true
	}
	return false
}

func (i *UserIterator) GetNext() Subject {
	if i.HasNext() {
		user := i.users[i.index]
		i.index++
		return user
	}
	return nil
}
