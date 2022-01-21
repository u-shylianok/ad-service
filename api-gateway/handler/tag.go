package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	pbAds "github.com/u-shylianok/ad-service/svc-ads/client/ads"
)

func (h *Handler) listAdTags(c *gin.Context) {
	var log = handlerLogger.WithFields(logrus.Fields{
		"method": "listAdTags",
	})

	adID, err := getUint32(c.Param("id"))
	if err != nil {
		log.WithError(err).Error("failed to read ad id URL param")
		newErrorResponse(c, http.StatusBadRequest, "invalid ad id param")
		return
	}
	log.WithField("adID", adID).Debug("ad id param read successfully")

	tags, err := h.clients.AdsService.ListTags(context.Background(), &pbAds.ListTagsRequest{AdId: adID})
	if err != nil {
		log.WithError(err).Error("failed to get ads tags")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	log.WithField("tags", tags).Debug("ads tags read successfully")

	c.JSON(http.StatusOK, tags)
}

func (h *Handler) listTags(c *gin.Context) {
	var log = handlerLogger.WithFields(logrus.Fields{
		"method": "listTags",
	})

	tags, err := h.clients.AdsService.ListTags(context.Background(), &pbAds.ListTagsRequest{})
	if err != nil {
		log.WithError(err).Error("failed to get tags")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	log.WithField("tags", tags).Debug("tags read successfully")

	c.JSON(http.StatusOK, tags)
}
