package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitMysql(user, pass, host, dbname string, port int) error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, pass, host, port, dbname)
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
