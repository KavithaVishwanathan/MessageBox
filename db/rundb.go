package db

import (
	"database/sql"
    "github.com/KavithaVishwanathan/messagebox/model"
)

func Run() {
	dbpath := "./sample.db"
	db := InitDB(dbpath)
	defer db.Close()
	createTables(db)
	populateUsers(db)
	populateMessages(db)
	populateLabels(db)
	populateLabelMsgMap(db)
}

func populateUsers(db *sql.DB) {
	users := []model.User{
		model.User{Id: 1, Name:"John", EmailId: "john@gmail.com",},
		model.User{Id: 2, Name:"Mary", EmailId: "mary@gmail.com",},
		model.User{Id: 3, Name:"Alan", EmailId: "alan@gmail.com",},
	}
	InsertUsers(db,users)
}

func populateMessages(db *sql.DB) {
	msgs := []model.Message{
		model.Message{Id: 0, Subject: "Intro", Body:"Hi John", FromUser: 2, ToUser: 1},
		model.Message{Id: 1, Subject: "Re: Intro", Body:"Nice to meet you", FromUser: 1, ToUser: 2},
		model.Message{Id: 2, Subject: "Bill", Body:"Gym Renewal", FromUser: 3, ToUser: 2},
		model.Message{Id: 3, Subject: "New Phone", Body:"Offer 20 percent", FromUser: 2, ToUser: 3},
	}
	InsertMessages(db,msgs)
}

func populateLabels(db *sql.DB) {
	labels := model.CreateLabels()
	InsertLabels(db,labels)
}

func populateLabelMsgMap(db *sql.DB) {
	labelMaps := []model.LabelMessageMapper{
		model.LabelMessageMapper{LabelId:0,MessageId:0},
		model.LabelMessageMapper{LabelId:0,MessageId:1},
		model.LabelMessageMapper{LabelId:1,MessageId:2},
		model.LabelMessageMapper{LabelId:2,MessageId:3},
		model.LabelMessageMapper{LabelId:0,MessageId:2},
	}
	InsertLabelMsgMap(db,labelMaps)
}

