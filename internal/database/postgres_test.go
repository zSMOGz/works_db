package database

import (
	"testing"
	"works_db/internal/messages"
)

func TestConnectSQL(t *testing.T) {
	db, err := ConnectSQL()
	if err != nil {
		t.Fatalf(messages.ErrorConnect, err)
	}
	defer db.Close()
}

func TestConnectGORM(t *testing.T) {
	db, err := ConnectGORM()
	if err != nil {
		t.Fatalf(messages.ErrorConnect, err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf(messages.ErrorGetSQLDB, err)
	}
	defer sqlDB.Close()
}
