package main

import "errors"

type ApprovalRequired struct {
	application *AccountApplication
}

func (s *ApprovalRequired) FillOutPaperwork() error {
	return errors.New("The paperwork has already been filled out, please wait for approval")
}

func (s *ApprovalRequired) SendInformation() error {
	return errors.New("You have already sent the required information, please wait for approval")
}

func (s *ApprovalRequired) AssignAppointment() error {
	return errors.New("You have already had an appointment, please wait for approval")
}

func (s *ApprovalRequired) ApproveApplication() error {
	s.application.SetState(s.application.Confirmed)
	return nil
}

func (s *ApprovalRequired) String() string {
	return "approval is required"
}
