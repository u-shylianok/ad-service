package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/u-shylianok/ad-service/api-gateway/domain/model"
)

func (h *Handler) signUp(c *gin.Context) {
	var log = handlerLogger.WithFields(logrus.Fields{
		"method": "signUp",
	})

	var input model.SignUpRequest
	if err := c.BindJSON(&input); err != nil {
		log.WithError(err).Error("failed to bind request JSON to struct")
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	log.WithField("input", input).Debug("input bound successfully")

	response, err := h.clients.AuthService.SignUp(context.Background(), input.ToPb())
	if err != nil {
		log.WithError(err).Error("failed to create user")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	log.WithField("id", response.Id).Debug("user created successfully")

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": response.Id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var log = handlerLogger.WithFields(logrus.Fields{
		"method": "signIn",
	})

	var input model.SignInRequest
	if err := c.BindJSON(&input); err != nil {
		log.WithError(err).Error("failed to bind request JSON to struct")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	log.WithField("input", input).Debug("input bound successfully")

	response, err := h.clients.AuthService.SignIn(context.Background(), input.ToPb())
	if err != nil {
		log.WithError(err).Error("failed to sign in")
		newErrorResponse(c, http.StatusUnauthorized, "failed to sign in:"+err.Error())
		return
	}
	log.WithField("token", response.Token).Debug("token generated successfully")

	c.JSON(http.StatusOK, map[string]interface{}{
		"token":      response.Token,
		"expires_at": response.ExpiresAt,
	})
}
