package controller

import (
	"github.com/bbriggs/quotes/model"
	"github.com/bbriggs/quotes/quotes"
	"github.com/gin-gonic/gin"
)

// raiderGET godoc
// @Summary Raider quotes
// @Description Get a random quote from a Raider NPC
// @ID raider
// @Produce json
// @Success 200 {object} model.Response
// @router /fallout/raider [get]
func (s *Server) raiderGET(c *gin.Context) {
	c.JSON(200, model.Response{
		Quote: quotes.Raider[s.Random.Intn(len(quotes.Raider))],
	})
}
