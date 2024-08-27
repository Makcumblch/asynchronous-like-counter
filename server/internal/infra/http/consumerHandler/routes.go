package consumerhandler

import (
	"net/http"

	"github.com/Makcumblch/asynchronous-like-counter/internal/infra/http/middleware"
)

func InitConsumerRoutes(handler *ConsumerHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/s2/likes", middleware.ErrorMW(handler.Get))

	return mux
}
