package models

type Menu struct {
	ID              int64
	Name            string
	IsVisible       bool
	EstablishmentID int64
	OrderParam      int64
}
