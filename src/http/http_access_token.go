package http

import (
	"net/http"

	atDomain "github.com/fdiaz7/bookstore_oauth-api/src/domain/access_token"
	"github.com/fdiaz7/bookstore_oauth-api/src/services/access_token"
	"github.com/fdiaz7/bookstore_oauth-api/src/utils/errors"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(ctx *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewAccessTokenHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	access_token, err := h.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, access_token)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var request atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		resErr := errors.NewBadRequestError("invalid json body")
		c.JSON(resErr.Status, resErr)
		return
	}
	accessToken, err := h.service.Create(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}
