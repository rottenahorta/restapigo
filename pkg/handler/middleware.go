package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) setUID(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerSplit := strings.Split(header, " ")
	if len(headerSplit) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	UID, err := h.services.Auth.ParseJWT(headerSplit[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set("UID", UID)
}

func (h *Handler) getUID(c *gin.Context) (int) {
	uid,ok := c.Get("UID")
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "cant get uid")
		return 0
	}
	uidint, ok := uid.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "uid not valid type")
		return 0
	}
	return uidint
}