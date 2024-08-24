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
}

type IConfig interface {
	GetHttp() HttpConfig
	GetRabbit() RabbitConfig
	GetMongo() MongoConfig
}

type Config struct{}

func (c *Config) GetHttp() HttpConfig {
	port, ok := os.LookupEnv("HTTP_PORT")
	if !ok {
		log.Fatal("Not found env \"HTTP_PORT\"")
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
	mongoConfig := MongoConfig{}
	return mongoConfig
}

func LoadConfig() IConfig {
	if err := godotenv.Load("../../../.env"); err != nil {
		log.Println(".env file not found")
	}

	config := &Config{}

	return config
}
