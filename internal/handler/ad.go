package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/u-shylianok/ad-service/internal/model"
)

func (h *Handler) createAd(c *gin.Context) {

	var input model.AdRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if err := input.Validate(); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Ad.CreateAd(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) listAds(c *gin.Context) {

	var order, sortBy string

	sortBy = c.Query("sort_by")
	if strings.ToLower(sortBy) == "price" || strings.ToLower(sortBy) == "date" {
		order = c.Query("order")
		if strings.ToLower(order) != "asc" && strings.ToLower(order) != "dsc" {
			order = "asc"
		}
	}

	ads, err := h.services.Ad.ListAds(sortBy, order)
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

	fields := c.Request.URL.Query()["fields"]

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
