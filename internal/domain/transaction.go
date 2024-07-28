package domain

import "time"

type Transaction struct {
	Id        string `json:"id"`
	User      User   `json:"user"`
	EventName string `json:"event_name"`
	Ticket    TicketForUser
	Create_at time.Time `json:"create_at"`
	Update_at time.Time `json:"update_at"`
}

// type Transaction struct {
// 	Id        string    `json:"id"`
// 	User      User      `json:"user"`
// 	Event     Event     `json:"event"`
// 	Create_at time.Time `json:"create_at"`
// 	Update_at time.Time `json:"update_at"`
// }
// type Event struct {
// 	Id        string
// 	EventName string
// 	Ticket    Ticket
// }

// type Ticket struct {
// 	Id    int
// 	Type  string
// 	Price float64
// 	Stock int
// }

// type User struct {
// 	Id   string
// 	Name string
// }
