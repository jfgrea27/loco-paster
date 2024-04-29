package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jfgrea27/loco-paster/internal/models"
	"github.com/rs/zerolog/log"
)

var pasteObjs = []models.PasteObj{
	{Id: 1, Blob: "foo"},
}

func getPaste(c *gin.Context) {
	log.Info().Msg("Handling GET /pastes")

	log.Debug().Msg(fmt.Sprintf("Current paste count: %v", len(pasteObjs)))

	c.IndentedJSON(http.StatusOK, pasteObjs)
}

func getPasteById(c *gin.Context) {
	id_str := c.Param("id")

	log.Info().Msg(fmt.Sprintf("Handling GET /pastes/%v", id_str))

	id, err := strconv.Atoi(id_str)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Invalid id %v", id_str)})
	}
	for _, o := range pasteObjs {
		if o.Id == id {
			c.IndentedJSON(http.StatusOK, o)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Id %v not found", id)})
}

func deletePasteById(c *gin.Context) {
	id_str := c.Param("id")

	log.Info().Msg(fmt.Sprintf("Handling DELETE /pastes/%v", id_str))

	id, err := strconv.Atoi(id_str)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Invalid id %v", id_str)})
	}

	idx := -1
	for i, o := range pasteObjs {
		if o.Id == id {
			idx = i
		}
	}
	// found -> delete
	if idx > -1 {
		log.Debug().Msg(fmt.Sprintf("PasteObj %v found, deleting.", id_str))
		pasteObjs = append(pasteObjs[:idx], pasteObjs[idx+1:]...)
	} else {
		log.Warn().Msg(fmt.Sprintf("PasteObj %v not found.", id_str))
	}
	log.Debug().Msg(fmt.Sprintf("Current paste count: %v", len(pasteObjs)))
	c.IndentedJSON(http.StatusOK, pasteObjs)
}

func postPaste(c *gin.Context) {

	log.Info().Msg("Handling POST /pastes/")

	var p models.PasteObj

	if err := c.BindJSON(&p); err != nil {
		msg := fmt.Sprintf("Could not process POST /pastes/ - invalid JSON, reason %v", err)
		log.Error().Msg(msg)

		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": msg})
		return
	}

	p.Id = len(pasteObjs) + 1
	pasteObjs = append(pasteObjs, p)
	c.IndentedJSON(http.StatusCreated, p)
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/api/v1/pastes/:id", getPasteById)
	router.GET("/api/v1/pastes/", getPaste)
	router.POST("/api/v1/pastes/", postPaste)
	router.DELETE("/api/v1/pastes/:id", deletePasteById)
	return router
}
func Run(endpoint string) {
	router := setupRouter()
	router.Run(endpoint)
}
