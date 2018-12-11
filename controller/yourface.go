package controller

import (
	"fmt"

	"github.com/bbriggs/quotes/model"
	"github.com/gin-gonic/gin"
)

// yourFaceGET godoc
// @Summary Patrick Carver Simulation Machine
// @Description Find out what your face is
// @ID your-face
// @Produce json
// @Success 200 {object} model.Response
// @router /trd/pikachu [get]
func (s *Server) yourFaceGET(c *gin.Context) {
	adj1, err := s.getRandomPOS("adjective")
	adj2, err := s.getRandomPOS("adjective")
	noun, err := s.getRandomPOS("noun")

	if err != nil {
		c.JSON(200, model.Error{
			Error: err.Error(),
		})
		return
	}

	plot := fmt.Sprintf("Your face is a %s %s %s", adj1, adj2, noun)

	c.JSON(200, model.Response{
		Quote: plot,
	})
}
