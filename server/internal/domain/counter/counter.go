package counter

type Counter struct {
	Likes uint
}

type ICounterRepository interface {
	Get() (uint, error)
	Increment() error
}
