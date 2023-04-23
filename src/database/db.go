package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBold() *gorm.DB {
	dsn := "root:beLL1923@tcp(127.0.0.1:3306)/prep50?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func DBnew() *gorm.DB {
	// dsn := "root:beLL1923@tcp(127.0.0.1:3306)/klass?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "magakfcswc1mohmz:ssnuwtx9cqu61kk8@tcp(iu51mf0q32fkhfpl.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306)/qbqg0wy70bpgbczb"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db.Set("gorm:table_options", "ENGINE=InnoDB")
}
