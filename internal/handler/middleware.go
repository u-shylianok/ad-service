package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AUTH_HEADER = "Authorization"
	USER_CTX    = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(AUTH_HEADER)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, err := h.services.Auth.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(USER_CTX, userId)
}

func getUserId(c *gin.Context) (int, error) {
	userCtx, ok := c.Get(USER_CTX)
	if !ok {
		return 0, fmt.Errorf("user id not found")
	}

	id, ok := userCtx.(int)
	if !ok {
		return 0, fmt.Errorf("user id is of invalid type")
	}

	return id, nil
}
