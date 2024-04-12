package models

type Establishment struct {
	ID            int64
	Name          string
	Phone         string
	Logo          string
	Background    string
	WifiPassword  string
	CanMakeOrders bool
	Country       string
	City          string
	Address       string
	Text          string
	ProfileID     int64
	CurrencyID    int64
	Link          string
}
