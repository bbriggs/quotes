package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/bbriggs/quotes/model"
	"github.com/bbriggs/quotes/util"
	"github.com/gin-gonic/gin"
)

// christmasPlotsGET godoc
// @Summary Randomized Christmas movie plots
// @Description Generate a random, whacky Christmas movie plot
// @ID christmas-plots
// @Produce json
// @Success 200 {object} model.Response
// @router /plots/christmas [get]
func (s *Server) christmasPlotsGET(c *gin.Context) {
	heroes := []string{"girl", "boy", "man", "woman", "elf"}
	actions := []string{"save", "ruin"}
	hero := heroes[s.Random.Intn(len(heroes))]
	action := actions[s.Random.Intn(len(actions))]

	adj1, err := s.getRandomPOS("adjective")
	adj2, err := s.getRandomPOS("adjective")
	verb, err := s.getRandomPOS("verb")
	noun, err := s.getRandomPOS("noun")

	if err != nil {
		c.JSON(200, model.Error{
			Error: err.Error(),
		})
		return
	}

	plot := fmt.Sprintf("%s %s %s tries to %s %s to %s Christmas", adj1, adj2, hero, verb, noun, action)

	c.JSON(200, model.Response{
		Quote: plot,
	})
}

func (s *Server) getRandomPOS(p string) (string, error) {
	parts := []string{"noun", "verb", "adjective"}
	if !(util.StringInSlice(p, parts)) {
		return "", fmt.Errorf("Incorrect part of speech supplied")
	}

	resp, err := http.Get(fmt.Sprintf("https://raw.githubusercontent.com/aaronbassett/Pass-phrase/master/%ss.txt", p))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	words := strings.Split(string(b), "\n")
	return words[s.Random.Intn(len(words))], nil
}
