package booking

type Booking struct {
	Uid       string `json:"Uid"`
	Name      string `json:"Name"`
	Session   string `json:"Session"`
	SessionID string `json:"SessionID"`
}

var Bookings []Booking
