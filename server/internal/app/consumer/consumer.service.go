package consumer

import (
	"github.com/Makcumblch/asynchronous-like-counter/internal/domain/counter"
)

type ConsumerService struct {
	repo counter.ICounterRepository
}

func NewConsumerService(repo counter.ICounterRepository) *ConsumerService {
	consumerService := &ConsumerService{
		repo: repo,
	}
	return consumerService
}

func (s *ConsumerService) Get() (uint, error) {
	return s.repo.Get()
}

func (s *ConsumerService) Increment() error {
	return s.repo.Increment()
}
