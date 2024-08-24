package incrementhandler

import (
	"net/http"

	"github.com/Makcumblch/asynchronous-like-counter/internal/infra/http/middleware"
)

func InitIncrementRoutes(handler *IncrementHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/like/increment", middleware.ErrorMW(handler.Increment))

	return mux
}
