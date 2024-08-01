package domain

type Event struct {
	Id        string   `json:"id" valo:"notblank,sizeMin=2,sizeMax=10"`
	EventName string   `json:"event_name" valo:"notblank,sizeMin=2,sizeMax=50"`
	Location  string   `json:"location" valo:"notblank,sizeMin=2,sizeMax=50"`
	Date      string   `json:"date" valo:"notblank,sizeMin=2,sizeMax=20"`
	Ticket    []Ticket `json:"ticket,omitempty" valo"valid,notnil"`
}

type EventForTx struct {
	EventName string `json:"event_name" valo:"notblank,sizeMin=2,sizeMax=50"`
	Location  string `json:"location" valo:"notblank,sizeMin=2,sizeMax=50"`
	Date      string `json:"date" valo:"notblank,sizeMin=2,sizeMax=20"`
}
