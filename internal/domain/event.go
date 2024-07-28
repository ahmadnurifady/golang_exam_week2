package domain

type Event struct {
	Id        string   `json:"id" valo:"notblank,sizeMin=2,sizeMax=10"`
	EventName string   `json:"event_name" valo:"notblank,sizeMin=2,sizeMax=10"`
	Location  string   `json:"location" valo:"notblank,sizeMin=2,sizeMax=10"`
	Date      string   `json:"date" valo:"notblank,sizeMin=2,sizeMax=10"`
	Ticket    []Ticket `json:"ticket" valo"valid,notnil"`
}
