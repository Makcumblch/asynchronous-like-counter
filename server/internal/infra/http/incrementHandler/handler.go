package incrementhandler

import (
	"net/http"

	incrementservice "github.com/Makcumblch/asynchronous-like-counter/internal/app/incrementService"
)

type IncrementHandler struct {
	IncrementService *incrementservice.IncrementService
}

func NewIncrementHandler(service *incrementservice.IncrementService) *IncrementHandler {
	return &IncrementHandler{IncrementService: service}
}

func (h *IncrementHandler) Increment(w http.ResponseWriter, r *http.Request) error {
	err := h.IncrementService.Increment()
	if err != nil {
		return err
	}

	return nil
}
