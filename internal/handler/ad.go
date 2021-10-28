package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/u-shylianok/ad-service/internal/model"
)

var logger = logrus.WithFields(logrus.Fields{
	"package": "internal-handler",
})

func (h *Handler) createAd(c *gin.Context) {
	var log = logger.WithFields(logrus.Fields{
		"method": "createAd",
	})

	var input model.AdRequest
	if err := c.BindJSON(&input); err != nil {
		log.Errorf("bind JSON to struct:", err)
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	log.Trace("input bound successfully")

	if err := input.Validate(); err != nil {
		log.Errorf("validate input:", err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	log.Trace("input validated successfully")

	id, err := h.services.Ad.CreateAd(input)
	if err != nil {
		log.Errorf("service create ad:", err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	log.Tracef("service create ad with id = %d", id)

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) listAds(c *gin.Context) {

	sortingParams := model.ListAdsSortingParamsFromURL(c.Request.URL.Query())

	ads, err := h.services.Ad.ListAds(sortingParams)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, ads)
}

func (h *Handler) getAd(c *gin.Context) {

	adID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid ad id param")
		return
	}

	fields := model.GetAdOptionalFieldsFromURL(c.Request.URL.Query())

	item, err := h.services.Ad.GetAd(adID, fields)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) updateAd(c *gin.Context) {
	adID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid ad id param")
		return
	}

	var input model.AdRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if err := input.Validate(); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Ad.UpdateAd(adID, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteAd(c *gin.Context) {

	adID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid ad id param")
		return
	}

	if err := h.services.Ad.DeleteAd(adID); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
