package main

import (
	"log"
	"pizza/db"
	// "pizza/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"os"
	"io"
	// "time"
	// "fmt"
	"pizza/utils"
)	
var Dbs *gorm.DB
func main() {
	config := utils.ReadConfig()
	f, _ := os.OpenFile(config.LogName+".log",os.O_WRONLY,0666)
	log.SetOutput(f)
	gin.DefaultWriter = io.MultiWriter(f)
	logger := logrus.New()
	logger.Level = logrus.TraceLevel
	logger.SetOutput(gin.DefaultWriter)
	Dbs = db.Open(config.DbURI,logger)
	utils.Dbs=Dbs
	r := gin.Default()
	r.Use(gin.Recovery())
	r.GET("/menu", utils.GetMenu)
	r.POST("/buy", utils.BuyPizza)
	r.GET("/trans", utils.Trans)
	r.Run(":"+config.Port) 
}
