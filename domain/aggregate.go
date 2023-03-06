package domain

type Aggregate struct {
	events []Event
}

func (a *Aggregate) PushEvents(events ...Event) {
	a.events = append(a.events, events...)
}

func (a *Aggregate) PullEvents() []Event {
	snapshot := a.events
	a.events = a.events[:0] // avoids extra malloc
	return snapshot
}
