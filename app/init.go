package app

import "github.com/BarnabyCharles/framework3/mysql"

func Init(user, pass, host, dbname string, port int, str ...string) error {
	var err error
	for _, val := range str {
		switch val {
		case "mysql":
			err = mysql.InitMysql(user, pass, host, dbname, port)
		}
	}

	return err
}
