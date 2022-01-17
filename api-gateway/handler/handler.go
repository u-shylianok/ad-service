package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/u-shylianok/ad-service/api-gateway/grpc/client"
)

type Handler struct {
	clients *client.Client
}

func NewHandler(clients *client.Client) *Handler {
	return &Handler{clients: clients}
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

	authorized := router.Group("/", h.userIdentity)
	{
		ads := authorized.Group("/ads")
		{
			ads.POST("/", h.createAd)
			ads.PUT("/:id", h.updateAd)
			ads.DELETE("/:id", h.deleteAd)
		}
	}

	ads := router.Group("/ads")
	{
		ads.GET("/", h.listAds)
		ads.GET("/search", h.searchAds)
		ads.GET("/:id", h.getAd)

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
