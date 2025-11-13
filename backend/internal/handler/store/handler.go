package store

import (
	"net/http"
	"store-review/internal/usecase/store"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	listUseCase *store.ListUseCase
}

func NewHandler(listUseCase *store.ListUseCase) *Handler {
	return &Handler{
		listUseCase: listUseCase,
	}
}

func (h *Handler) GetList(c *gin.Context) {
	output, err := h.listUseCase.Execute(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := make([]Response, len(output.Stores))
	for i, s := range output.Stores {
		res[i] = newResponse(s)
	}

	c.JSON(http.StatusOK, res)
}
