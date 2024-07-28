package dto

type CreateTxDto struct {
	UserId     string `json:"user_id" valo:"notblank,sizeMin=2,sizeMax=10"`
	EventId    string `json:"event_id" valo:"notblank,sizeMin=2,sizeMax=10"`
	TicketType string `json:"ticket_type" valo:"notblank,sizeMin=2,sizeMax=10"`
}
