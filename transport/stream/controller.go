package stream

import (
	"github.com/maestre3d/coinlog/messaging"
)

type Controller interface {
	MapStreams(b *messaging.Bus)
}

type ControllerMapper struct {
	controllers []Controller
}

func NewControllerMapper() *ControllerMapper {
	return &ControllerMapper{}
}

func (h *ControllerMapper) Add(cc ...Controller) {
	h.controllers = append(h.controllers, cc...)
}

func (h *ControllerMapper) RegisterSubscribers(b *messaging.Bus) {
	for _, ctrl := range h.controllers {
		ctrl.MapStreams(b)
	}
}
