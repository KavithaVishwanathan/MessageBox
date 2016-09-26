package db

import (
    "testing"
)

func Test_InsertUsers(t *testing.T) {
	db := InitDB("test.db")
	createTables(db)
	InsertUsers(db,nil)
}

func Test_InitDB(t *testing.T) {
	db := InitDB("test.db");
	if (db == nil) {
		t.Errorf("Test Fail")
	}
}

func Test_InitDBEmptyFile(t *testing.T) {
	db := InitDB("");
	if (db != nil) {
		t.Errorf("Test Fail")
	}
}





