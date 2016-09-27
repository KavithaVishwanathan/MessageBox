package main

import (
	"net/http"
	"testing"
	"encoding/json"
	"reflect"
	"github.com/KavithaVishwanathan/messagebox/model"
)

func TestGetUser(t *testing.T) {
	res, err := http.Get("https://infinite-chamber-34739.herokuapp.com/user/2")
    if err != nil {
        t.Fatal(err)
    }
    defer res.Body.Close()
    
    var actualUser model.User
    json.NewDecoder(res.Body).Decode(&actualUser)
    expectedUser := model.User{2,"Mary","mary@gmail.com"}
    if !reflect.DeepEqual(actualUser, expectedUser) {
    	t.Error("Two Users are not equal")
    }
}

func TestGetUser_invalidId(t *testing.T) {
	res, err := http.Get("https://infinite-chamber-34739.herokuapp.com/user/12")
    if err != nil {
        t.Fatal(err)
    }
    defer res.Body.Close()
    
    var actualUser model.User
    json.NewDecoder(res.Body).Decode(&actualUser)
    expectedUser := model.User{0,"",""}
    if !reflect.DeepEqual(actualUser, expectedUser) {
    	t.Error("Two Users are not equal")
    }
}

func TestGetMessage(t *testing.T) {
	res, err := http.Get("https://infinite-chamber-34739.herokuapp.com/message/1")
    if err != nil {
        t.Fatal(err)
    }
    defer res.Body.Close()
    
    var actualMessage model.Message
    json.NewDecoder(res.Body).Decode(&actualMessage)
    expectedMessage := model.Message{Id: 1, Subject: "Re: Intro", Body:"Nice to meet you", FromUser: 1, ToUser: 2}
    if !reflect.DeepEqual(actualMessage, expectedMessage) {
    	t.Error("Two Messages are not equal")
    }
}

func TestGetMessage_invalidId(t *testing.T) {
	res, err := http.Get("https://infinite-chamber-34739.herokuapp.com/message/100")
    if err != nil {
        t.Fatal(err)
    }
    defer res.Body.Close()
    
    var actualMessage model.Message
    json.NewDecoder(res.Body).Decode(&actualMessage)
    expectedMessage := model.Message{Id: 0, Subject: "", Body:"", FromUser: 0, ToUser: 0}
    if !reflect.DeepEqual(actualMessage, expectedMessage) {
    	t.Error("Two Messages are not equal")
    }
}

func TestGetLabel(t *testing.T) {
	res, err := http.Get("https://infinite-chamber-34739.herokuapp.com/label/1")
    if err != nil {
        t.Fatal(err)
    }
    defer res.Body.Close()
    
    var actualLabel model.Label
    json.NewDecoder(res.Body).Decode(&actualLabel)
    expectedLabel := model.Label{1,"Purchases"}
    if !reflect.DeepEqual(actualLabel, expectedLabel) {
    	t.Error("Two Labels are not equal")
    }
}

func TestGetLabel_invalidId(t *testing.T) {
	res, err := http.Get("https://infinite-chamber-34739.herokuapp.com/label/11")
    if err != nil {
        t.Fatal(err)
    }
    defer res.Body.Close()
    
    var actualLabel model.Label
    json.NewDecoder(res.Body).Decode(&actualLabel)
    expectedLabel := model.Label{0,""}
    if !reflect.DeepEqual(actualLabel, expectedLabel) {
    	t.Error("Two Labels are not equal")
    }
}

func TestGetLabelByMessage(t *testing.T) {
	res, err := http.Get("https://infinite-chamber-34739.herokuapp.com/message/1/labels")
    if err != nil {
        t.Fatal(err)
    }
    defer res.Body.Close()
    
    var actualLabels []model.Label
    json.NewDecoder(res.Body).Decode(&actualLabels)
    expectedLabel := []model.Label{
    	model.Label{0,"Personal"},
    }
    if !reflect.DeepEqual(actualLabels, expectedLabel) {
    	t.Error("Two Labels are not equal %s %s", res.Body, expectedLabel)
    }
}

func TestGetLabelByMessage_invalidId(t *testing.T) {
	res, err := http.Get("https://infinite-chamber-34739.herokuapp.com/message/11/labels")
    if err != nil {
        t.Fatal(err)
    }
    defer res.Body.Close()
    
    var actualLabel interface{}
    json.NewDecoder(res.Body).Decode(&actualLabel)

    if !(reflect.DeepEqual(actualLabel,nil)) {
    	t.Error("Label is not null", actualLabel)
    }
}




