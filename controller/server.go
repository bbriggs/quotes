package controller

import (
	"math/rand"
	"time"

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

type Server struct {
	Random *rand.Rand  // Initialized PRNG
	API    *gin.Engine // Gin router
}

func Run() {
	s := &Server{
		Random: rand.New(rand.NewSource(time.Now().UnixNano())),
		API:    gin.Default(),
	}
	s.API.GET("/", func(c *gin.Context) {
		c.String(200, "Quotes\n")
	})
	s.API.GET("/fallout/raider", s.raiderGET)
	s.API.GET("/status", s.statusGET)
	s.API.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.API.Run() // listen and serve on 0.0.0.0:8080
}

// statusGET godoc
// @Summary Health check URL
// @Description Status check
// @ID status
// @Produce json
// @Success 200
// @router /status [get]
func (s *Server) statusGET(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
