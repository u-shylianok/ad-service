package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	pbAds "github.com/u-shylianok/ad-service/svc-ads/client/ads"
)

func (h *Handler) listAdPhotos(c *gin.Context) {
	var log = handlerLogger.WithFields(logrus.Fields{
		"method": "listAdPhotos",
	})

	adID, err := getUint32(c.Param("id"))
	if err != nil {
		log.WithError(err).Error("failed to read ad id URL param")
		newErrorResponse(c, http.StatusBadRequest, "invalid ad id param")
		return
	}
	log.WithField("adID", adID).Debug("ad id param read successfully")

	photos, err := h.clients.AdsService.ListPhotos(context.Background(), &pbAds.ListPhotosRequest{AdId: adID})
	if err != nil {
		log.WithError(err).Error("failed to get ads photos")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	log.WithField("photos", photos).Debug("ads photos read successfully")

	c.JSON(http.StatusOK, photos)
}

func (h *Handler) listPhotos(c *gin.Context) {
	var log = handlerLogger.WithFields(logrus.Fields{
		"method": "listPhotos",
	})

	photos, err := h.clients.AdsService.ListPhotos(context.Background(), &pbAds.ListPhotosRequest{})
	if err != nil {
		log.WithError(err).Error("failed to get photos")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	log.WithField("photos", photos).Debug("photos read successfully")

	c.JSON(http.StatusOK, photos)
}
