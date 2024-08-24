package incrementservice

type IQueue interface {
	Send() error
}

type IncrementService struct {
	queue IQueue
}

func NewIncrementService(queue IQueue) *IncrementService {
	return &IncrementService{
		queue: queue,
	}
}

func (s *IncrementService) Increment() error {
	return s.queue.Send()
}
