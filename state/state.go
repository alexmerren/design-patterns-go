package main

import "errors"

type State interface {
	FillOutPaperwork() error
	SendInformation() error
	AssignAppointment() error
	ApproveApplication() error
	String() string
}

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
