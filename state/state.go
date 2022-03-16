package main

type State interface {
	FillOutPaperwork() error
	SendInformation() error
	AssignAppointment() error
	ApproveApplication() error
	String() string
}
