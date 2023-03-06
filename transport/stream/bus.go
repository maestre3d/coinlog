package stream

import "github.com/maestre3d/coinlog/messaging"

func NewBus(w messaging.Writer, r messaging.Reader, mapper *ControllerMapper) *messaging.Bus {
	bus := messaging.NewBus(w, r)
	mapper.RegisterSubscribers(bus)
	return bus
}
