package main

func main() {
	file1 := NewFile("help.md")
	file2 := NewFile("test.md")
	file3 := NewFile("quarter_notes.md")
	notes := NewFolder("Notes")
	notes.AddChild(file1, file2, file3)

	file4 := NewFile("Why_I_Should_Get_A_Promotion.md")
	secretNotes := NewFolder("Secret Notes")
	secretNotes.AddChild(file4)

	notes.AddChild(secretNotes)

	notes.Print()
	//secretNotes.Print()
}
