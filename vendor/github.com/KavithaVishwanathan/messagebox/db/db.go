package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "github.com/KavithaVishwanathan/messagebox/model"
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	checkErr(err)

	if db == nil { panic("db nil") }
	return db
}

func createTables(db *sql.DB) {
	user_table := "CREATE TABLE IF NOT EXISTS user(id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, emailId TEXT);"
	message_table := "CREATE TABLE IF NOT EXISTS message(id INTEGER PRIMARY KEY AUTOINCREMENT, subject TEXT, body TEXT, from_id INTEGER, to_id INTEGER, FOREIGN KEY(from_id) REFERENCES user(id), FOREIGN KEY(to_id) REFERENCES user(id));"
	label_table := "CREATE TABLE IF NOT EXISTS label(id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT);"
	label_msg_table := "CREATE TABLE IF NOT EXISTS labelMessage(label_id INTEGER, message_id INTEGER, FOREIGN KEY(label_id) REFERENCES label(id), FOREIGN KEY(message_id) REFERENCES message(id));"

	_, err := db.Exec(user_table)
	checkErr(err)

	_, err1 := db.Exec(message_table)
	checkErr(err1)

	_, err2 := db.Exec(label_table)
	checkErr(err2)

	_, err3 := db.Exec(label_msg_table)
	checkErr(err3)
}

func InsertUsers(db *sql.DB, users []model.User) {
	sql_additem := "INSERT INTO user(id,name,emailId) values(?, ?, ?)"

	stmt, err := db.Prepare(sql_additem)
	checkErr(err)

	for _, user := range users {
		_, err := stmt.Exec(user.Id, user.Name, user.EmailId)
		checkErr(err)
	}
}

func InsertMessages(db *sql.DB, msgs []model.Message) {
	sql_additem := "INSERT INTO message(id,subject,body,from_id,to_id) values(?, ?, ?, ?, ?)"

	stmt, err := db.Prepare(sql_additem)
	checkErr(err)

	for _, msg := range msgs {
		_, err := stmt.Exec(msg.Id, msg.Subject, msg.Body, msg.FromUser, msg.ToUser)
		checkErr(err)
	}
}

func InsertLabels(db *sql.DB, labels []model.Label) {
	sql_additem := "INSERT INTO label(id,name) values(?, ?)"

	stmt, err := db.Prepare(sql_additem)
	checkErr(err)

	for _, label := range labels {
		_, err := stmt.Exec(label.Id, label.Name)
		checkErr(err)
	}
}

func InsertLabelMsgMap(db *sql.DB, LabelMessageMaps []model.LabelMessageMapper) {
	sql_additem := "INSERT INTO labelMessage(label_id,message_id) values(?, ?)"

	stmt, err := db.Prepare(sql_additem)
	checkErr(err)

	for _, labelmsgmap := range LabelMessageMaps {
		_, err := stmt.Exec(labelmsgmap.LabelId, labelmsgmap.MessageId)
		checkErr(err)
	}
}

func ListUsers(db *sql.DB) []model.User {
	sql_readall := `SELECT * FROM user`
	rows, err := db.Query(sql_readall)
	checkErr(err)
	defer rows.Close()

	var result []model.User
	for rows.Next() {
		user := model.User{}
		err2 := rows.Scan(&user.Id, &user.Name, &user.EmailId)
		checkErr(err2)
		result = append(result, user)
	}
	err3 := rows.Err()
	checkErr(err3)
	return result
}

func GetUserById(db *sql.DB, id string) model.User {
	sql_get := `SELECT * FROM user where id = ?`
	row, err := db.Query(sql_get,id)
	checkErr(err)
	defer row.Close()
	
	user := model.User{}
	for row.Next() {
		err2 := row.Scan(&user.Id, &user.Name, &user.EmailId)
		checkErr(err2)
	}
	
	err3 := row.Err()
	checkErr(err3)
	return user
}

func ListMessages(db *sql.DB) []model.Message {
	sql_readall := `SELECT * FROM message`
	rows, err := db.Query(sql_readall)
	checkErr(err)
	defer rows.Close()

	var result []model.Message
	for rows.Next() {
		msg := model.Message{}
		err2 := rows.Scan(&msg.Id, &msg.Subject, &msg.Body, &msg.FromUser, &msg.ToUser)
		checkErr(err2)
		result = append(result, msg)
	}
	err3 := rows.Err()
	checkErr(err3)
	return result
}

func GetMessageById(db *sql.DB, id string) model.Message {
	sql_get := `SELECT * FROM message where id = ?`
	row, err := db.Query(sql_get,id)
	checkErr(err)
	defer row.Close()

	msg := model.Message{}
	for row.Next() {
		err2 := row.Scan(&msg.Id, &msg.Subject, &msg.Body, &msg.FromUser, &msg.ToUser)
		checkErr(err2)
	}
	
	err3 := row.Err()
	checkErr(err3)
	return msg
}

func ListLabels(db *sql.DB) []model.Label {
	sql_readall := `SELECT * FROM label`
	rows, err := db.Query(sql_readall)
	checkErr(err)
	defer rows.Close()

	var result []model.Label
	for rows.Next() {
		label := model.Label{}
		err2 := rows.Scan(&label.Id, &label.Name)
		checkErr(err2)
		result = append(result, label)
	}
	err3 := rows.Err()
	checkErr(err3)
	return result
}

func GetLabelById(db *sql.DB, id string) model.Label {
	sql_get := `SELECT * FROM label where id = ?`
	row, err := db.Query(sql_get,id)
	checkErr(err)
	defer row.Close()
	
	label := model.Label{}
	for row.Next() {
		err2 := row.Scan(&label.Id, &label.Name)
		checkErr(err2)
	}
	
	err3 := row.Err()
	checkErr(err3)
	return label
}

func GetLabelByMessageId(db *sql.DB, id string) []model.Label {
	sql_get := `Select label.id,label.name From message INNER JOIN labelMessage ON message.id = labelMessage.message_id INNER JOIN label ON labelMessage.label_id = label.id WHERE message.id = ?`
	rows, err := db.Query(sql_get,id)
	checkErr(err)
	defer rows.Close()
	
	var labels []model.Label
	for rows.Next() {
		label := model.Label{}
		err2 := rows.Scan(&label.Id, &label.Name)
		checkErr(err2)
		labels = append(labels,label)
	}
	
	err3 := rows.Err()
	checkErr(err3)
	return labels
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
