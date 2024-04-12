package models

type Profile struct {
	ID               int64
	Name             string
	Email            string
	Password         string
	Verified         bool
	VerificationCode string
}
