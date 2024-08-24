package counter

type Counter struct {
	likes uint
}

type ICounterRepository interface {
	Get() uint
	Increment() uint
}
