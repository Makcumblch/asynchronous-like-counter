package consumerhandler

import (
	"encoding/json"
	"net/http"

	"github.com/Makcumblch/asynchronous-like-counter/internal/app/consumer"
)

type GetLikesResponse struct {
	Likes uint `json:"likes"`
}

type ConsumerHandler struct {
	ConsumerService *consumer.ConsumerService
}

func NewConsumerHandler(service *consumer.ConsumerService) *ConsumerHandler {
	return &ConsumerHandler{ConsumerService: service}
}

func (h *ConsumerHandler) Increment(w http.ResponseWriter, r *http.Request) error {
	err := h.ConsumerService.Increment()
	if err != nil {
		return err
	}

	return nil
}

func (h *ConsumerHandler) Get(w http.ResponseWriter, r *http.Request) error {
	count, err := h.ConsumerService.Get()
	if err != nil {
		return err
	}
	likes := GetLikesResponse{
		Likes: count,
	}
	res, err := json.Marshal(likes)
	if err != nil {
		return err
	}

	w.Write(res)

	return nil
}
