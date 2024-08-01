package domain

type Ticket struct {
	Id        string  `json:"id" valo:"notblank,sizeMin=5,sizeMax=50"`
	Type      string  `json:"type" valo:"notblank,sizeMin=5,sizeMax=50"`
	Price     float64 `json:"price" valo:"min=50,max=10000"`
	Stock     int     `json:"stock" valo:"min=10,max=10000"`
	EventName string  `json:"event_name" valo:"notblank,sizeMin=5,sizeMax=50"`
}

type TicketForUser struct {
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}

type TicketForFindAll struct {
	Type  string  `json:"type" valo:"notblank,sizeMin=5,sizeMax=50"`
	Price float64 `json:"price" valo:"min=50,max=10000"`
	Stock int     `json:"stock" valo:"min=10,max=10000"`
}
