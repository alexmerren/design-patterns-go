package main

import "errors"

type Confirmed struct {
	application *AccountApplication
}

func (s *Confirmed) FillOutPaperwork() error {
	return errors.New("error")
}

func (s *Confirmed) SendInformation() error {
	return errors.New("error")
}

func (s *Confirmed) AssignAppointment() error {
	return errors.New("error")
}

func (s *Confirmed) ApproveApplication() error {
	return errors.New("error")
}

func (s *Confirmed) String() string {
	return "Application has been approved!"
}
