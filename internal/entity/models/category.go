package models

type Category struct {
	ID         int64
	Name       string
	Background string
	MenuID     int64
	IsVisible  bool
	OrderParam int64
}
