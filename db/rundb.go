package db

import (
    //"encoding/json"
    "github.com/KavithaVishwanathan/messagebox/model"
)

func init() {
	dbpath := "./sample.db"
	db := InitDB(dbpath)
	defer db.Close()
	createTables(db)

	//type Users []User
	users = []User{
		User{Id: 1, Name:"John", EmailId: "john@gmail.com",},
		User{Id: 2, Name:"Mary", EmailId: "mary@gmail.com",},
		User{Id: 3, Name:"Alan", EmailId: "alan@gmail.com",},
	}

	insertUser(db,users)
}




