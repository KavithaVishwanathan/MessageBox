package model

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	checkErr(err)
	if db == nil { panic("db nil") }
	return db
}

func createSitesTable(db *sql.DB) {
	user_table := "
	CREATE TABLE user(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		userName TEXT,
		emailId TEXT
	);"

	message_table := "
	CREATE TABLE message(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		subject TEXT,
		body TEXT,
		FOREIGN KEY(from_id) REFERENCES user(id),
		FOREIGN KEY(to_id) REFERENCES user(id)
	);"

	label_table := "
	CREATE TABLE label(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		labelName TEXT
	);"

	label_msg_table := "
	CREATE TABLE labelMessage(
		FOREIGN KEY(label_id) REFERENCES label(id),
		FOREIGN KEY(message_id) REFERENCES message(id)
	);"

	_, err := db.Exec(user_table)
	checkErr(err)

	_, err := db.Exec(message_table)
	checkErr(err)

	_, err := db.Exec(label_table)
	checkErr(err)

	_, err := db.Exec(label_msg_table)
	checkErr(err)
}

func insertUser(db *sql.DB, users []User) {
	sql_additem := "
	INSERT INTO user(
		id,
		name,
		emailId
	) values(?, ?, ?)"

	stmt, err := db.Prepare(sql_additem)
	checkErr(err)

	for _, user := range users {
		_, err := stmt.Exec(user.id, user.name, user.emailId)
		checkErr(err)
	}
}

func insertMessage(db *sql.DB, msgs []message) {
	sql_additem := "
	INSERT INTO message(
		id,
		subject,
		body,
		from_id,
		to_id
	) values(?, ?, ?, ?, ?)"

	stmt, err := db.Prepare(sql_additem)
	checkErr(err)

	for _, msg := range msgs {
		_, err := stmt.Exec(msg.id, msg.subject, msg.body, msg.from_id, msg.to_id)
		checkErr(err)
	}
}

func insertLabel(db *sql.DB, labels []Label) {
	sql_additem := "
	INSERT INTO label(
		id,
		name
	) values(?, ?)"

	stmt, err := db.Prepare(sql_additem)
	checkErr(err)

	for _, label := range labels {
		_, err := stmt.Exec(label.id, label.name)
		checkErr(err)
	}
}

func insertLabelMsg(db *sql.DB, labels []Label) {
	sql_additem := "
	INSERT INTO label(
		id,
		name
	) values(?, ?)"

	stmt, err := db.Prepare(sql_additem)
	checkErr(err)

	for _, label := range labels {
		_, err := stmt.Exec(label.id, label.name)
		checkErr(err)
	}
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}