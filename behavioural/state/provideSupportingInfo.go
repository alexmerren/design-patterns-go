package main

import "errors"

type ProvideSupportingInfo struct {
	application *AccountApplication
}

func (s *ProvideSupportingInfo) FillOutPaperwork() error {
	return errors.New("The paperwork has already been filled out, please provide required information")
}

func (s *ProvideSupportingInfo) SendInformation() error {
	s.application.SetState(s.application.AppointmentRequired)
	return nil
}

func (s *ProvideSupportingInfo) AssignAppointment() error {
	return errors.New("You must first provide supporting information before getting an appointment")
}

func (s *ProvideSupportingInfo) ApproveApplication() error {
	return errors.New("You must first provide support information before getting approval")
}

func (s *ProvideSupportingInfo) String() string {
	return "provide supporting information"
}
