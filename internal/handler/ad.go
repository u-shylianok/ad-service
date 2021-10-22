package handler

import (
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/u-shylianok/ad-service/internal/model"
)

func (h *Handler) createAd(c *gin.Context) {

	var input model.AdRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if utf8.RuneCountInString(input.Name) > 200 {
		newErrorResponse(c, http.StatusBadRequest, "name should be no more than 200 symbols")
		return
	}

	if input.Description == "" || utf8.RuneCountInString(input.Description) > 1000 {
		newErrorResponse(c, http.StatusBadRequest, "description should be no more than 1000 symbols")
		return
	}

	if input.MainPhoto.Link == "" {
		newErrorResponse(c, http.StatusBadRequest, "must be at least 1 photo")
		return
	}

	if input.OtherPhotos == nil || len(*input.OtherPhotos) > 2 {
		newErrorResponse(c, http.StatusBadRequest, "should be no more than 3 photos")
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
		if strings.ToLower(order) != "asc" && order != "dsc" {
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
