package db

import (
	"log"
	// "os"
	"pizza/logs"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Open function recieve an expr as argument
// and making connection to our database
func Open(expr string,logger *logrus.Logger) *gorm.DB {
	db, err := gorm.Open("mysql", expr)
	if err != nil {
		log.Panic("Couldn't opendatabase", err.Error())
	}

	db.LogMode(true)
	db.SetLogger(&logs.GormLogger{
		Name:   "db gorm logger",
		Logger: logger,
	})
	// db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	return db
}
