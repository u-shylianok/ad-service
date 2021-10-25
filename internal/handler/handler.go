package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/u-shylianok/ad-service/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode) // set if release
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	//api := router.Group("/api", h.userIdentity)
	//{
	ads := router.Group("/ads")
	{
		ads.POST("/", h.createAd)
		ads.GET("/", h.listAds)
		//ads.GET("/find", h.findAds)
		ads.GET("/:id", h.getAd)
		//ads.PUT("/:id", h.updateAd)
		//ads.DELETE("/:id", h.deleteAd)
	}
	//}

	return router
}

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
