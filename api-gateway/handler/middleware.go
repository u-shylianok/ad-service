package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/u-shylianok/ad-service/api-gateway/grpc/dto"
	pbAuth "github.com/u-shylianok/ad-service/svc-auth/client/auth"
)

const (
	AuthHeader   = "Authorization"
	UserIDCtxKey = "user_id"
)

func (h *Handler) userIdentity(c *gin.Context) {
	var log = handlerLogger.WithFields(logrus.Fields{
		"method": "userIdentity",
	})

	header := c.GetHeader(AuthHeader)
	if header == "" {
		log.Error("empty auth header")
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		log.Error("invalid auth header")
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		log.Error("token is empty")
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	res, err := h.clients.AuthService.ParseToken(context.Background(), &pbAuth.ParseTokenRequest{Token: headerParts[1]})
	if err != nil {
		log.WithError(err).Error("failed to parse token")
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	userID := dto.PbAuth.FromParseTokenResponse(res)

	c.Set(UserIDCtxKey, userID)
}

func getUserID(c *gin.Context) (uint32, error) {
	var log = handlerLogger.WithFields(logrus.Fields{
		"method": "getUserID",
	})

	// log.Infof("%#v", c)
	userCtx, ok := c.Get(UserIDCtxKey)
	if !ok {
		log.Error("failed to parse token")
		return 0, fmt.Errorf("user id not found")
	}
	// log.Error(userCtx)

	id, ok := userCtx.(uint32)
	if !ok {
		log.Error("user id is of invalid type")
		return 0, fmt.Errorf("user id is of invalid type")
	}

	return id, nil
}
