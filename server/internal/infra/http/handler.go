package http

import (
	"net/http"

	"github.com/Makcumblch/asynchronous-like-counter/internal/app"
)

type HandlerFunc func(http.ResponseWriter, *http.Request) app.Error
