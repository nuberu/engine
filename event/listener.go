package event

type strategy uint8

const (
	strategyAlways strategy = 0
	strategyOnce strategy = 1
)

type Listener struct {
	strategy strategy
	callback func(interface{}, *Args)
}