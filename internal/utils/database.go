package utils

import (
	"fmt"

	"gorm.io/gorm"
)

func SetSchema(db *gorm.DB, schema string) *gorm.DB {
	db.Exec(`set search_path='` + schema + `'`)
	return db
}

func CleanupDBConnection(DB *gorm.DB) {
	var stmtManger, ok = DB.ConnPool.(*gorm.PreparedStmtDB)
	if ok {
		defer stmtManger.Close()
	} else {
		fmt.Println(stmtManger)
	}
}
