package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) listAdPhotos(c *gin.Context) {

	adID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid ad id param")
		return
	}

	photos, err := h.services.Photo.ListAdPhotos(adID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, photos)
}

func (h *Handler) listPhotos(c *gin.Context) {

	photos, err := h.services.Photo.ListPhotos()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, photos)
}
