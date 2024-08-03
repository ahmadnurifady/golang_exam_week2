package domain

import "time"

type Transaction struct {
	Id        string    `json:"id"`
	User      User      `json:"user"`
	Event     Event     `json:"event"`
	Ticket    Ticket    `json:"ticket"`
	Create_at time.Time `json:"create_at"`
	Update_at time.Time `json:"update_at"`
}
