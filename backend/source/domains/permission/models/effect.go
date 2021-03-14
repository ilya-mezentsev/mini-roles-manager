package models

type PermitEffect struct{}

func (p PermitEffect) IsPermit() bool {
	return true
}

func (p PermitEffect) IsEmpty() bool {
	return false
}

type DenyEffect struct{}

func (d DenyEffect) IsPermit() bool {
	return false
}

func (d DenyEffect) IsEmpty() bool {
	return false
}

type MissedEffect struct{}

func (m MissedEffect) IsPermit() bool {
	return false
}

func (m MissedEffect) IsEmpty() bool {
	return true
}
