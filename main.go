package main

import (
	"log"
	"pizza/db"
	// "pizza/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"os"
	"io"
	"time"
	"fmt"
	"pizza/utils"
)	
var Dbs *gorm.DB
func main() {
	config := utils.ReadConfig()
	Dbs = db.Open(config.DbURI)
	utils.Dbs=Dbs
	f, _ := os.OpenFile(config.LogName+".log",os.O_WRONLY,0666)
	log.SetOutput(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())
	log.Println(Dbs)
	r.GET("/menu", utils.GetMenu)
	r.POST("/buy", utils.BuyPizza)
	r.GET("/trans", utils.Trans)
	r.Run(":"+config.Port) 
}
