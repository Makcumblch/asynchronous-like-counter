package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	incrementservice "github.com/Makcumblch/asynchronous-like-counter/internal/app/incrementService"
	"github.com/Makcumblch/asynchronous-like-counter/internal/infra/http"
	incrementhandler "github.com/Makcumblch/asynchronous-like-counter/internal/infra/http/incrementHandler"
	"github.com/Makcumblch/asynchronous-like-counter/internal/infra/rabbitmq"
	"github.com/Makcumblch/asynchronous-like-counter/internal/util/config"
)

func main() {
	config := config.LoadConfig()

	rabbitmq := rabbitmq.NewRabbitmq(config.GetRabbit())

	incService := incrementservice.NewIncrementService(rabbitmq)

	handler := incrementhandler.NewIncrementHandler(incService)

	routes := incrementhandler.InitIncrementRoutes(handler)

	server := http.NewServer(config.GetHttp("HTTP_PORT_PRODUCER"), routes)
	server.Run()

	log.Print("--- Producer service started ---")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	if err := rabbitmq.Close(); err != nil {
		log.Printf("Ошибка завершения работы RabbitMQ: %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("error occred on producer service shutting down: %s", err.Error())
	}

	log.Print("--- Producer service shutdown ---")
}
