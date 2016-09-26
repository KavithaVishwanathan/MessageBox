package main

import (
	"log"
    "net/http"
    _ "fmt"
    "encoding/json"
    "database/sql"
    "github.com/gorilla/mux"
    "github.com/KavithaVishwanathan/messagebox/db"
)

var conn *sql.DB 

func users(rw http.ResponseWriter, request *http.Request) {
	users := db.ListUsers(conn)
	json.NewEncoder(rw).Encode(users)
    //rw.Write([]byte("Hello world."))
}

func userById(rw http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	user := db.GetUserById(conn, id)
	json.NewEncoder(rw).Encode(user)
    //rw.Write([]byte("Hello world."))
}

func messages(rw http.ResponseWriter, request *http.Request) {
	messages := db.ListMessages(conn)
	json.NewEncoder(rw).Encode(messages)
    //rw.Write([]byte("Hello world."))
}

func messageById(rw http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	message := db.GetMessageById(conn, id)
	json.NewEncoder(rw).Encode(message)
    //rw.Write([]byte("Hello world."))
}

func label(rw http.ResponseWriter, request *http.Request) {
	labels := db.ListLabels(conn)
	json.NewEncoder(rw).Encode(labels)
    //rw.Write([]byte("Hello world."))
}

func labelById(rw http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	label := db.GetLabelById(conn, id)
	json.NewEncoder(rw).Encode(label)
    //rw.Write([]byte("Hello world."))
}

func labelByMessageId(rw http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	label := db.GetLabelByMessageId(conn, id)
	json.NewEncoder(rw).Encode(label)
    //rw.Write([]byte("Hello world."))
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
    
    router.HandleFunc("/user", users)
    router.HandleFunc("/user/{id}", userById)
    router.HandleFunc("/message", messages)
    router.HandleFunc("/message/{id}", messageById)
    router.HandleFunc("/message/{id}/labels", labelByMessageId)
    router.HandleFunc("/label", label)
    router.HandleFunc("/label/{id}", labelById)

    //Render the html
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
    http.Handle("/", router)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	//db.Run()
	conn = db.InitDB("./sample.db")
    handleRequests()
}