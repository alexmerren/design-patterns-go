package main

import "errors"

type AppointmentRequired struct {
	application *AccountApplication
}

func (s *AppointmentRequired) FillOutPaperwork() error {
	return errors.New("The paperwork has already been filled out, please request an appointment")
}

func (s *AppointmentRequired) SendInformation() error {
	return errors.New("You have already sent the required information, please request an appointment")
}

func (s *AppointmentRequired) AssignAppointment() error {
	s.application.SetState(s.application.ApprovalRequired)
	return nil
}

func (s *AppointmentRequired) ApproveApplication() error {
	return errors.New("You must first request an appointment before getting approval")
}

func (s *AppointmentRequired) String() string {
	return "an Appointment is required"
}
