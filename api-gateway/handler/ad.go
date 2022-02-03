package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/u-shylianok/ad-service/api-gateway/domain/model"
	"github.com/u-shylianok/ad-service/api-gateway/grpc/dto"
)

func (h *Handler) getAd(c *gin.Context) {
	var log = handlerLogger.WithFields(logrus.Fields{
		"method": "getAd",
	})

	adID, err := getUint32(c.Param("id"))
	if err != nil {
		log.WithError(err).Error("failed to read id URL param")
		newErrorResponse(c, http.StatusBadRequest, "invalid ad id param")
		return
	}
	log.WithField("id", adID).Debug("id param read successfully")

	optional := model.GetAdsOptionalFromURL(c.Request.URL.Query())
	log.WithField("fields", optional).Debug("optional fields was formed")

	ad, err := h.clients.AdsService.GetAd(context.Background(), dto.ToPbAds_GetAdRequest(adID, &optional))
	if err != nil {
		log.WithError(err).Error("failed to get ad")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	log.WithField("ad", ad).Debug("ad read successfully")

	c.JSON(http.StatusOK, dto.FromPbAds_GetAdResponse(ad))
}

func (h *Handler) listAds(c *gin.Context) {
	var log = handlerLogger.WithFields(logrus.Fields{
		"method": "listAds",
	})

	sortingParams := model.GetAdsSortingParamsFromURL(c.Request.URL.Query())
	log.WithField("sorting params", sortingParams).Debug("sorting params was formed")

	ads, err := h.clients.AdsService.ListAds(context.Background(), dto.ToPbAds_ListAdsRequest(sortingParams))
	if err != nil {
		log.WithError(err).Error("failed to get ads")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	log.WithField("ads", ads).Debug("ads read successfully")

	c.JSON(http.StatusOK, dto.FromPbAds_ListAdsResponse(ads))
}

func (h *Handler) searchAds(c *gin.Context) {
	var log = handlerLogger.WithFields(logrus.Fields{
		"method": "searchAds",
	})

	filter := model.GetAdFilterFromURL(c.Request.URL.Query())
	log.WithField("filter", filter).Debug("sorting params was formed")

	ads, err := h.clients.AdsService.SearchAds(context.Background(), dto.ToPbAds_SearchAdsRequest(filter))
	if err != nil {
		log.WithError(err).Error("failed to get ads")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	log.WithField("ads", ads).Debug("ads read successfully")

	c.JSON(http.StatusOK, dto.FromPbAds_SearchAdsResponse(ads))
}

func (h *Handler) createAd(c *gin.Context) {
	var log = handlerLogger.WithFields(logrus.Fields{
		"method": "createAd",
	})

	userID, err := getUserID(c)
	if err != nil {
		log.WithError(err).Error("failed to get userID from context")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	log.WithField("userID", userID).Debug("userID getted successfully")

	var input model.AdRequest
	if err := c.BindJSON(&input); err != nil {
		log.WithError(err).Error("failed to bind request JSON to struct")
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	log.WithField("input", input).Debug("input bound successfully")

	if err := input.Validate(); err != nil {
		log.WithError(err).Error("invalid input")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	log.Debug("input validated successfully")

	ad, err := h.clients.AdsService.CreateAd(context.Background(), dto.ToPbAds_CreateAdRequest(userID, input))
	if err != nil {
		log.WithError(err).Error("failed to create ad")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	log.WithField("id", ad.Id).Debug("ad created successfully")

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": ad.Id,
	})
}

func (h *Handler) updateAd(c *gin.Context) {
	var log = handlerLogger.WithFields(logrus.Fields{
		"method": "updateAd",
	})

	userID, err := getUserID(c)
	if err != nil {
		log.WithError(err).Error("failed to get userID from context")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	log.WithField("userID", userID).Debug("userID getted successfully")

	adID, err := getUint32(c.Param("id"))
	if err != nil {
		log.WithError(err).Error("failed to read id URL param")
		newErrorResponse(c, http.StatusBadRequest, "invalid ad id param")
		return
	}
	log.WithField("id", adID).Debug("id param read successfully")

	var input model.AdRequest
	if err := c.BindJSON(&input); err != nil {
		log.WithError(err).Error("failed to bind request JSON to struct")
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	log.WithField("input", input).Debug("input bound successfully")

	if err := input.Validate(); err != nil {
		log.WithError(err).Error("invalid input")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	log.Debug("input validated successfully")

	if _, err := h.clients.AdsService.UpdateAd(context.Background(), dto.ToPbAds_UpdateAdRequest(userID, adID, input)); err != nil {
		log.WithError(err).Error("failed to update ad")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	log.WithField("id", adID).Debug("ad updated successfully")

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteAd(c *gin.Context) {
	var log = handlerLogger.WithFields(logrus.Fields{
		"method": "deleteAd",
	})

	userID, err := getUserID(c)
	if err != nil {
		log.WithError(err).Error("failed to get userID from context")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	log.WithField("userID", userID).Debug("userID getted successfully")

	adID, err := getUint32(c.Param("id"))
	if err != nil {
		log.WithError(err).Error("failed to read id URL param")
		newErrorResponse(c, http.StatusBadRequest, "invalid ad id param")
		return
	}
	log.WithField("id", adID).Debug("id param read successfully")

	if _, err := h.clients.AdsService.DeleteAd(context.Background(), dto.ToPbAds_DeleteAdRequest(userID, adID)); err != nil {
		log.WithError(err).Error("failed to delete ad")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	log.WithField("id", adID).Debug("ad deleted successfully")

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
