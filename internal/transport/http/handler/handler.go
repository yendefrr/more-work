package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yendefrr/more-work/internal/service"
)

type Handler struct {
	jobsService      service.Jobs
	employersService service.Employers
	seekersService   service.Seekers
}

func NewHandler(jobsService service.Jobs, employersService service.Employers, seekersService service.Seekers) *Handler {
	return &Handler{
		jobsService:      jobsService,
		employersService: employersService,
		seekersService:   seekersService,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	h.SetRoutes(router)

	return router
}

func (h *Handler) SetRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
