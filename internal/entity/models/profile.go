package models

type Profile struct {
	ID               int64
	Name             string
	Email            string
	Password         string
	NewPassword      string
	Verified         bool
	VerificationCode string
}
