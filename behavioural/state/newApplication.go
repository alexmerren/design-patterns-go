package main

import "errors"

type NewApplication struct {
	application *AccountApplication
}

func (s *NewApplication) FillOutPaperwork() error {
	s.application.SetState(s.application.ProvideSupportingInfo)
	return nil
}

func (s *NewApplication) SendInformation() error {
	return errors.New("You must first fill out the required paperwork for a new application")
}

func (s *NewApplication) AssignAppointment() error {
	return errors.New("You must first fill out the required paperwork for a new application")
}

func (s *NewApplication) ApproveApplication() error {
	return errors.New("You must first fill out the required paperwork for a new application")
}

func (s *NewApplication) String() string {
	return "new Application"
}
