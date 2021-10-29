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

var handlerLogger = logrus.WithFields(logrus.Fields{
	"package": "internal-handler",
})

func (h *Handler) InitRoutes() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode) // set if release
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/signup", h.signUp)
		auth.POST("/signin", h.signIn)
	}

	//api := router.Group("/api", h.userIdentity)
	//{
	ads := router.Group("/ads")
	{
		ads.POST("/", h.createAd)
		ads.GET("/", h.listAds)
		ads.GET("/search", h.searchAds)
		ads.GET("/:id", h.getAd)
		ads.PUT("/:id", h.updateAd)
		ads.DELETE("/:id", h.deleteAd)

		photos := ads.Group(":id/photos")
		{
			photos.GET("/", h.listAdPhotos)
		}

		tags := ads.Group(":id/tags")
		{
			tags.GET("/", h.listAdTags)
		}
	}

	tags := router.Group("/tags")
	{
		tags.GET("/", h.listTags)
	}

	photos := router.Group("/photos")
	{
		photos.GET("/", h.listPhotos)
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
