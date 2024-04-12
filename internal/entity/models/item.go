package models

type Item struct {
	ID             int64
	Name           string
	Weight         string
	Price          float64
	AdditionalInfo string
	IsVisible      bool
	IsAvailable    bool
	Image          string
	CategoryID     int64
	OrderParam     int64
	SelectedAmount int64 `db:"-"`
}
