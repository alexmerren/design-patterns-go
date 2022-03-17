package main

type Collection interface {
	CreateIterator() Iteratable
}

type UserCollection struct {
	users []*User
}

func (c *UserCollection) CreateIterator() Iteratable {
	return &UserIterator{
		users: c.users,
	}
}
