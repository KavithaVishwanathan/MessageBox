package model

import (
    _"time"
)

type Message struct {
	Id	int	`json:"id"`
	Subject	string	`json:"subject"`
	Body	string	`json:"body"`
	//Time	time.Time `json:"time"`
	FromUser	int `json:"fromuser"`
	ToUser	int	`json:"touser"`
}

type LabelMessageMapper struct {
	LabelId	int	`json:"labelid"`
	MessageId	int	`json:"messageid"`
}

