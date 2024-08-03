package dto

import "time"

type CreateTxDto struct {
	UserId     string `json:"user_id" valo:"notblank,sizeMin=2,sizeMax=10"`
	EventId    string `json:"event_id" valo:"notblank,sizeMin=2,sizeMax=10"`
	TicketType string `json:"ticket_type" valo:"notblank,sizeMin=2,sizeMax=10"`
}

type CreateManyTxDto struct {
	UserId             string                                     `json:"user_id" valo:"notblank,sizeMin=2,sizeMax=20"`
	RequestEventTicket []RequestCreateManyTxButThisForEventTicket `json:"request_event"`
}

type RequestCreateManyTxButThisForEventTicket struct {
	EventId string                 `json:"event_id" valo:"notblank,sizeMin=2,sizeMax=20"`
	Ticket  []RequestForManyTicket `json:"ticket"`
}

type RequestForManyTicket struct {
	TicketType   string `json:"ticket_type" valo:"notblank,sizeMin=2,sizeMax=20"`
	StockRequest int    `json:"stock_request"`
}

type CreateTxDatabase struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user"`
	EventId   string    `json:"event"`
	TicketId  string    `json:"ticket"`
	Create_at time.Time `json:"create_at"`
	Update_at time.Time `json:"update_at"`
}

type FindAllEventDto struct {
	Id        string
	EventName string
	Location  string
	Date      string
	Type      string
	Price     float64
	Stock     int
}
