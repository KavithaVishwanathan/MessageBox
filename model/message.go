package model

import (
    "time"
)

type Message struct {
	Id	int	`json:"id"`
	Subject	string	`json:"name"`
	Body	string	`json:"emailid"`
	Time	time.Time `json:"time"`
	FromUser	int `json:"fromuser`
	ToUser	int	`json:"touser`
}

