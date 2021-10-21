package handler

import (
	"net/http"
	"strconv"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/u-shylianok/ad-service/internal/model"
)

func (h *Handler) createAd(c *gin.Context) {

	var input model.Ad
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

	if len(*input.OtherPhotos) > 2 {
		newErrorResponse(c, http.StatusBadRequest, "should be no more than 3 photos")
		return
	}

	id, err := h.services.Ad.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) listAds(c *gin.Context) {

	var order, sortBy string

	sortBy = c.Query("sort_by")
	if sortBy == "price" || sortBy == "date" {
		order = c.Query("order")
		if order != "asc" && order != "dsc" {
			order = "asc"
		}
	} else {
		sortBy = ""
	}

	ads, err := h.services.Ad.List(sortBy, order)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, ads)
}

func (h *Handler) getAd(c *gin.Context) {

	adId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid ad id param")
		return
	}

	fields := c.Request.URL.Query()["fields"]

	item, err := h.services.Ad.Get(adId, fields)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}
