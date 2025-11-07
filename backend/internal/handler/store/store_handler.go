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

	srs := make([]Response, len(output.Stores))
	for i, s := range output.Stores {
		srs[i] = Response{
			ID:              s.ID,
			Name:            s.Name,
			RegularHolidays: s.RegularHolidays,
			CategoryNames:   s.CategoryNames,
			PaymentMethods:  s.PaymentMethods,
			WebProfiles:     s.WebProfiles,
		}
	}

	// レスポンスを作成
	res := ListResponse{
		Stores: srs,
	}
	c.JSON(http.StatusOK, res)
}
