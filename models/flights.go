package models

type Plane struct {
	Type       string `json:"type"`
	TotalSeats int    `json:"totalSeats"`
}
type Flight struct {
	Id            int     `json:"id"`
	Code          string  `json:"code"`
	Price         float32 `json:"price"`
	DepartureDate string  `json:"departureDate"`
	Origin        string  `json:"origin"`
	Destination   string  `json:"destination"`
	EmptySeats    int     `json:"emptySeats"`
	Plane         Plane   `json:"plane"`
}

var Flights = []Flight{
	{
		Id:            1,
		Code:          "ER38sd",
		Price:         400,
		DepartureDate: "2017/07/26",
		Origin:        "CLE",
		Destination:   "SFO",
		EmptySeats:    0,
		Plane: Plane{
			Type:       "Boeing 737",
			TotalSeats: 150,
		},
	},
	{
		Id:            2,
		Code:          "ER45if",
		Price:         540.99,
		DepartureDate: "2017/07/27",
		Origin:        "SFO",
		Destination:   "ORD",
		EmptySeats:    54,
		Plane: Plane{
			Type:       "Boeing 737",
			TotalSeats: 300,
		},
	},
}
