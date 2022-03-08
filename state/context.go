package main

type AccountApplication struct {
	NewApplication        State
	ProvideSupportingInfo State
	AppointmentRequired   State
	ApprovalRequired      State
	Confirmed             State
	CurrentState          State
	name                  string
}

func NewAccountApplication(n string) *AccountApplication {
	application := &AccountApplication{name: n}

	newApplicationState := &NewApplication{application: application}
	provideInfoState := &ProvideSupportingInfo{application: application}
	appointmentRequiredState := &AppointmentRequired{application: application}
	approvalRequiredState := &ApprovalRequired{application: application}
	confirmedState := &Confirmed{application: application}

	application.SetState(newApplicationState)
	application.NewApplication = newApplicationState
	application.ProvideSupportingInfo = provideInfoState
	application.AppointmentRequired = appointmentRequiredState
	application.ApprovalRequired = approvalRequiredState
	application.Confirmed = confirmedState

	return application
}

func (a *AccountApplication) FillOutPaperwork() error {
	return a.CurrentState.FillOutPaperwork()
}

func (a *AccountApplication) SendInformation() error {
	return a.CurrentState.SendInformation()
}

func (a *AccountApplication) AssignAppointment() error {
	return a.CurrentState.AssignAppointment()
}

func (a *AccountApplication) ApproveApplication() error {
	return a.CurrentState.ApproveApplication()
}

func (a *AccountApplication) SetState(s State) {
	a.CurrentState = s
}
