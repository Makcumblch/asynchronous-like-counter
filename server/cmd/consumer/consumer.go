package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Makcumblch/asynchronous-like-counter/internal/app/consumer"
	"github.com/Makcumblch/asynchronous-like-counter/internal/infra/http"
	consumerhandler "github.com/Makcumblch/asynchronous-like-counter/internal/infra/http/consumerHandler"
	"github.com/Makcumblch/asynchronous-like-counter/internal/infra/mongo"
	"github.com/Makcumblch/asynchronous-like-counter/internal/infra/mongo/counter"
	"github.com/Makcumblch/asynchronous-like-counter/internal/infra/rabbitmq"
	"github.com/Makcumblch/asynchronous-like-counter/internal/util/config"
)

func main() {
	config := config.LoadConfig()

	mongo := mongo.NewMongo(config.GetMongo())

	repo := counter.NewCounterRepository(mongo)

	recService := consumer.NewConsumerService(repo)

	rabbitmq := rabbitmq.NewRabbitmq(config.GetRabbit())

	rabbitmq.StartConsume(recService)

	handler := consumerhandler.NewConsumerHandler(recService)

	routes := consumerhandler.InitConsumerRoutes(handler)

	server := http.NewServer(config.GetHttp("HTTP_PORT_CONSUMER"), routes)
	server.Run()

	log.Print("--- Consumer service started ---")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	if err := rabbitmq.Close(); err != nil {
		log.Printf("Ошибка завершения работы RabbitMQ: %s", err.Error())
	}

	ctxMongo, cancelMongo := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelMongo()
	if err := mongo.Close(ctxMongo); err != nil {
		log.Printf("Ошибка завершения работы MongoDB: %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("error occred on producer service shutting down: %s", err.Error())
	}

	log.Print("--- Consumer service shutdown ---")
}
