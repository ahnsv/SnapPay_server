package db

import (
	"testing"
)

func TestDBInit(t *testing.T) {
	db := GetDB()
	if db == nil {
		t.Error("db init failed")
	}
}
