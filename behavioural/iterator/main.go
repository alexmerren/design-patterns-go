package main

func main() {
	User1 := &User{
		name: "Alex",
		age:  21,
	}
	User2 := &User{
		name: "Jeremy",
		age:  22,
	}
	User3 := &User{
		name: "Sam",
		age:  28,
	}

	Collection := &UserCollection{
		users: []*User{User1, User2, User3},
	}

	Iterator := Collection.CreateIterator()
	for Iterator.HasNext() {
		Iterator.GetNext().Print()
	}
}
