package api

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"

	"github.com/gin-gonic/gin"

	"github.com/jfgrea27/loco-paster/internal/models"
	"github.com/jfgrea27/loco-paster/internal/utils"
	"github.com/rs/zerolog/log"
)

var pasteObjs = []models.PasteObj{}

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
	idx := utils.FindPasteObjIndex(id, pasteObjs)

	// found -> delete
	if idx == -1 {
		log.Warn().Msg(fmt.Sprintf("PasteObj %v not found.", id_str))
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Id %v not found", id)})
		return
	}

	deletePaste := pasteObjs[idx]

	log.Debug().Msg(fmt.Sprintf("PasteObj %v found, deleting.", id_str))
	pasteObjs = append(pasteObjs[idx:], pasteObjs[idx+1:]...)
	log.Debug().Msg(fmt.Sprintf("Current paste count: %v", len(pasteObjs)))

	c.IndentedJSON(http.StatusOK, deletePaste)
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

	router = setupCors(router)

	if os.Getenv("API_ONLY") != "true" {
		router = setupStaticRouter(router)
	}
	router = setupApiRouter(router)

	return router
}

func setupStaticRouter(router *gin.Engine) *gin.Engine {
	router.Use(static.Serve("/", static.LocalFile("./dist", true)))

	return router
}

func setupCors(router *gin.Engine) *gin.Engine {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"*"}
	config.AllowHeaders = []string{"*"}

	router.Use(cors.New(config))
	return router
}

func setupApiRouter(router *gin.Engine) *gin.Engine {
	api := router.Group("/api/v1")

	api.GET("/pastes/:id", getPasteById)
	api.GET("/pastes/", getPaste)
	api.POST("/pastes/", postPaste)
	api.DELETE("/pastes/:id", deletePasteById)

	return router
}

func RunServer(endpoint string) {
	router := setupRouter()
	log.Info().Msg(fmt.Sprintf("Running with %v", endpoint))

	router.Run(endpoint)
}
