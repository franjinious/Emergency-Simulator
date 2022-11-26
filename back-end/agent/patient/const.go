package patient

type patient_status int8

const (
	Waiting_for_nurse      patient_status = 1
	Waiting_for_register   patient_status = 2
	Waiting_for_treat      patient_status = 3
	Being_treated_now      patient_status = 4
	Being_treated_urgently patient_status = 5
)