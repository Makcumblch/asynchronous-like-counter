package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type HttpConfig struct {
	Port string
}

type RabbitConfig struct {
	User string
	Pass string
	IP   string
	Port string
}

type MongoConfig struct {
	User string
	Pass string
	IP   string
	Port string
}

type IConfig interface {
	GetHttp(key string) HttpConfig
	GetRabbit() RabbitConfig
	GetMongo() MongoConfig
}

type Config struct{}

func (c *Config) GetHttp(key string) HttpConfig {
	port, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Not found env \"%s\"", key)
	}

	httpConfig := HttpConfig{
		Port: port,
	}
	return httpConfig
}

func (c *Config) GetRabbit() RabbitConfig {
	user, ok := os.LookupEnv("RABBITMQ_DEFAULT_USER")
	if !ok {
		log.Fatal("Not found env \"RABBITMQ_DEFAULT_USER\"")
	}
	pass, ok := os.LookupEnv("RABBITMQ_DEFAULT_PASS")
	if !ok {
		log.Fatal("Not found env \"RABBITMQ_DEFAULT_PASS\"")
	}
	ip, ok := os.LookupEnv("RABBITMQ_IP")
	if !ok {
		log.Fatal("Not found env \"RABBITMQ_IP\"")
	}
	port, ok := os.LookupEnv("RABBITMQ_PORT")
	if !ok {
		log.Fatal("Not found env \"RABBITMQ_PORT\"")
	}
	rabbitConfig := RabbitConfig{
		User: user,
		Pass: pass,
		IP:   ip,
		Port: port,
	}
	return rabbitConfig
}

func (c *Config) GetMongo() MongoConfig {
	user, ok := os.LookupEnv("MONGO_ROOT_USER")
	if !ok {
		log.Fatal("Not found env \"MONGO_ROOT_USER\"")
	}
	pass, ok := os.LookupEnv("MONGO_ROOT_PASSWORD")
	if !ok {
		log.Fatal("Not found env \"MONGO_ROOT_PASSWORD\"")
	}
	ip, ok := os.LookupEnv("MONGO_IP")
	if !ok {
		log.Fatal("Not found env \"MONGO_IP\"")
	}
	port, ok := os.LookupEnv("MONGO_PORT")
	if !ok {
		log.Fatal("Not found env \"MONGO_PORT\"")
	}
	mongoConfig := MongoConfig{
		User: user,
		Pass: pass,
		IP:   ip,
		Port: port,
	}
	return mongoConfig
}

func LoadConfig() IConfig {
	if err := godotenv.Load("../../../.env"); err != nil {
		log.Println(".env file not found")
	}

	config := &Config{}

	return config
}
