package http

import (
	"net/http"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
	"github.com/billysutomo/chocolate-waffle/internal/middleware"
	"github.com/gin-gonic/gin"
)

// ResponseError ResponseError
type ResponseError struct {
	Message string `json:"message"`
}

// RequestCreateBlock RequestCreateBlock
type RequestCreateBlock struct {
	IDProject int    `json:"id_project"`
	Ordernum  int    `json:"ordernum"`
	Type      string `json:"type"`
	URL       string `json:"url"`
	Icon      string `json:"icon"`
	Title     string `json:"title"`
}

// BlockHandler BlockHandler
type BlockHandler struct {
	BlockUsecase domain.BlockUsecase
}

// NewBlockHandler NewBlockHandler
func NewBlockHandler(r *gin.Engine, mid *middleware.MainMiddleware, do domain.BlockUsecase) {
	handler := &BlockHandler{
		BlockUsecase: do,
	}

	r.POST("/block", handler.CreateBlock)
}

// CreateBlock CreateBlock
func (a *BlockHandler) CreateBlock(r *gin.Context) {
	r.JSON(http.StatusCreated, map[string]string{
		"message": "block created",
	})
}
