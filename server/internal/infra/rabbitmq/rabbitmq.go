package rabbitmq

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Makcumblch/asynchronous-like-counter/internal/app/consumer"
	"github.com/Makcumblch/asynchronous-like-counter/internal/util/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbitmq struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      *amqp.Queue
}

func (r *Rabbitmq) Send() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.channel.PublishWithContext(ctx,
		"",           // exchange
		r.queue.Name, // routing key
		false,        // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
		})
	if err != nil {
		log.Println("Ошибка отправки сообщения в очередь")
		return err
	}

	return nil
}

func NewRabbitmq(config config.RabbitConfig) *Rabbitmq {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", config.User, config.Pass, config.IP, config.Port))
	if err != nil {
		log.Fatal("Ошибка подключения к RabbitMQ: ", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Ошибка открытия канала", err)
	}

	q, err := ch.QueueDeclare(
		"increment", // name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		log.Fatal("Ошибка создания очереди", err)
	}

	return &Rabbitmq{
		connection: conn,
		channel:    ch,
		queue:      &q,
	}
}

func (r *Rabbitmq) StartConsume(app *consumer.ConsumerService) error {
	msgs, err := r.channel.Consume(
		r.queue.Name, // queue
		"",           // consumer
		false,        // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)

	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			err := app.Increment()
			if err == nil {
				err := d.Ack(false)
				if err != nil {
					log.Printf("Ошибка подтверждения обработку сообщения: %s", err.Error())
				}
			} else {
				log.Printf("Не удалось увеличить число лайков: %s", err.Error())
			}
		}
	}()

	return nil
}

func (r *Rabbitmq) Close() error {
	err := r.channel.Close()
	if err != nil {
		return err
	}
	err = r.connection.Close()
	if err != nil {
		return err
	}

	return nil
}
