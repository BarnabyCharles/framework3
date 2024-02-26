package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitMysql() error {
	var err error
	dsn := fmt.Sprintf("root:yuling@tcp(127.0.0.1:3306)/zg5?parseTime=true")
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return err
}

func WithFx(txF func(tx *gorm.DB) error) {
	var err error
	tx := Db.Begin()
	err = txF(tx)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
}
