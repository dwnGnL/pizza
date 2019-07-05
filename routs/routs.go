package routs


import(
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"pizza/models"
	"net/http"
	"log"
	"time"
)
var Dbs *gorm.DB
var menu []models.Menu
func GetMenu(c *gin.Context)  {
	Dbs.Find(&menu)
	c.JSON(http.StatusOK,menu)
}

func BuyPizza(c *gin.Context){
	var request *models.Request
	var trans models.Trans
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Cannot decode " + err.Error()})
		log.Panic(err.Error())
		return
	}
	Dbs.Find(&menu)
	log.Println(menu)
	for _,row:= range menu{
		for j,column:=range request.Pizzas{
			if row.Name == column.Name{
				if column.Count<=0{
					c.JSON(http.StatusBadGateway, gin.H{"message": "Please choose right count"})
					return
				}
				request.Pizzas[j].Sum=column.Count*row.Price
			}
		}
	}
	now := time.Now()
	timef := now.Format("02.01.2006 15:04:05")
	request.Date=timef
	log.Println(request)
	for _,column:=range request.Pizzas{
		if column.Sum==0{
			c.JSON(http.StatusBadGateway, gin.H{"message": "We are don`t heve this pizzas"})
			return
		}
		request.TotalSum+=column.Sum

	}
	c.JSON(http.StatusOK, request)

	for _,column:=range request.Pizzas{
		trans.Name=column.Name
		trans.Count=column.Count
		trans.Sum=request.TotalSum
		trans.Date=request.Date
		Dbs.Create(trans) 
	}
	
}

func Trans(c *gin.Context){
	var trans []models.Trans
	Dbs.Find(&trans)
	c.JSON(http.StatusOK,trans)
}
