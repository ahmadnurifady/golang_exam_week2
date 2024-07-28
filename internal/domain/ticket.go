package domain

type Ticket struct {
	Id    int     `json:"id" valo:"min=1,max=10"`
	Type  string  `json:"type" valo:"notblank,sizeMin=10,sizeMax=50"`
	Price float64 `json:"price" valo:"min=50,max=10000"`
	Stock int     `json:"stock" valo:"min=10,max=10000"`
}

type TicketForUser struct {
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}
