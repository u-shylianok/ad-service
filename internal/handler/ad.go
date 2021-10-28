package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/u-shylianok/ad-service/internal/model"
)

func (h *Handler) createAd(c *gin.Context) {
	var log = handlerLogger.WithFields(logrus.Fields{
		"method": "createAd",
	})

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

	id, err := h.services.Ad.CreateAd(input)
	if err != nil {
		log.WithError(err).Error("failed to create ad")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	log.WithField("id", id).Debug("ad created successfully")

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) listAds(c *gin.Context) {
	var log = handlerLogger.WithFields(logrus.Fields{
		"method": "listAds",
	})

	sortingParams := model.ListAdsSortingParamsFromURL(c.Request.URL.Query())
	log.WithField("sorting params", sortingParams).Debug("sorting params was formed")

	ads, err := h.services.Ad.ListAds(sortingParams)
	if err != nil {
		log.WithError(err).Error("failed to get ads")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	log.WithField("ads", ads).Debug("ads read successfully")

	c.JSON(http.StatusOK, ads)
}

func (h *Handler) getAd(c *gin.Context) {
	var log = handlerLogger.WithFields(logrus.Fields{
		"method": "getAd",
	})

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.WithError(err).Error("failed to read id URL param")
		newErrorResponse(c, http.StatusBadRequest, "invalid ad id param")
		return
	}
	log.WithField("id", id).Debug("id param read successfully")

	fields := model.GetAdOptionalFieldsFromURL(c.Request.URL.Query())
	log.WithField("fields", fields).Debug("optional fields was formed")

	ad, err := h.services.Ad.GetAd(id, fields)
	if err != nil {
		log.WithError(err).Error("failed to get ad")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	log.WithField("ad", ad).Debug("ad read successfully")

	c.JSON(http.StatusOK, ad)
}

func (h *Handler) updateAd(c *gin.Context) {
	var log = handlerLogger.WithFields(logrus.Fields{
		"method": "updateAd",
	})

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.WithError(err).Error("failed to read id URL param")
		newErrorResponse(c, http.StatusBadRequest, "invalid ad id param")
		return
	}
	log.WithField("id", id).Debug("id param read successfully")

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

	if err := h.services.Ad.UpdateAd(id, input); err != nil {
		log.WithError(err).Error("failed to update ad")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	log.WithField("id", id).Debug("ad updated successfully")

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteAd(c *gin.Context) {
	var log = handlerLogger.WithFields(logrus.Fields{
		"method": "deleteAd",
	})

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.WithError(err).Error("failed to read id URL param")
		newErrorResponse(c, http.StatusBadRequest, "invalid ad id param")
		return
	}
	log.WithField("id", id).Debug("id param read successfully")

	if err := h.services.Ad.DeleteAd(id); err != nil {
		log.WithError(err).Error("failed to delete ad")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	log.WithField("id", id).Debug("ad deleted successfully")

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
