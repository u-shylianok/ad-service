package handler

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/sirupsen/logrus"
// )

// func (h *Handler) signUp(c *gin.Context) {
// 	var log = handlerLogger.WithFields(logrus.Fields{
// 		"method": "signUp",
// 	})

// 	var input model.User
// 	if err := c.BindJSON(&input); err != nil {
// 		log.WithError(err).Error("failed to bind request JSON to struct")
// 		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
// 		return
// 	}
// 	log.WithField("input", input).Debug("input bound successfully")

// 	id, err := h.services.Auth.CreateUser(input)
// 	if err != nil {
// 		log.WithError(err).Error("failed to create user")
// 		newErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	log.WithField("id", id).Debug("user created successfully")

// 	c.JSON(http.StatusOK, map[string]interface{}{
// 		"id": id,
// 	})
// }

// func (h *Handler) signIn(c *gin.Context) {
// 	var log = handlerLogger.WithFields(logrus.Fields{
// 		"method": "signIn",
// 	})

// 	var input model.SignInRequest
// 	if err := c.BindJSON(&input); err != nil {
// 		log.WithError(err).Error("failed to bind request JSON to struct")
// 		newErrorResponse(c, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	log.WithField("input", input).Debug("input bound successfully")

// 	id, err := h.services.Auth.CheckUser(input.Username, input.Password)
// 	if err != nil {
// 		log.WithError(err).Error("failed to check user by username and password")
// 		newErrorResponse(c, http.StatusUnauthorized, "incorrect username or password")
// 		return
// 	}
// 	log.WithField("userID", id).Debug("user verified successfully")

// 	token, expiresAt, err := h.services.Auth.GenerateToken(id)
// 	if err != nil {
// 		log.WithError(err).Error("failed to generate token")
// 		newErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	log.WithField("token", token).Debug("token generated successfully")

// 	c.JSON(http.StatusOK, map[string]interface{}{
// 		"token":      token,
// 		"expires_at": expiresAt,
// 	})
// }
