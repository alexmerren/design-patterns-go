package main

import "fmt"

func main() {
	// Happy path of creating an account
	application := NewAccountApplication("")
	fmt.Println("--- Application Happy Path ---")
	application.FillOutPaperwork()
	application.SendInformation()
	application.AssignAppointment()
	application.ApproveApplication()
	fmt.Println(application.CurrentState.String())

	// Calling ApproveApplication when in the 'wrong' state
	application2 := NewAccountApplication("")
	fmt.Println("--- Application Error ---")
	application2.FillOutPaperwork()
	if err := application2.ApproveApplication(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(application2.CurrentState.String())
}
