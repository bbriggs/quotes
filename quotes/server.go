package quotes

import (
	_ "github.com/bbriggs/quotes/docs"
	"github.com/gin-gonic/gin"
	//"github.com/sirupsen/logrus"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Quotes API
// @version 0.1.0
// @description This is a quotes response service

// @contact.url https://github.com/bbriggs/quotes

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

func Run() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Quotes\n")
	})
	r.GET("/status", statusGET)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run() // listen and serve on 0.0.0.0:8080
}

// statusGET godoc
// @Summary Health check URL
// @Description Status check
// @ID status
// @Produce json
// @Success 200
// @router /status [get]
func statusGET(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
