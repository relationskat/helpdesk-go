package domain

type Ticket struct {
	ID     int
	Title  string
	Status string
}

const StatusOpen = "open"
