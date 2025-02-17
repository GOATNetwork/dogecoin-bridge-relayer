package eventbus

import (
	"sync"
)

type (
	EventUnknown    struct{}
	EventTssMsg     struct{}
	EventTask       struct{}
	EventSubmitTask struct{}
	EventSynced     struct{}
	EventTestTask   struct{}
)

type Event any

type Bus interface {
	SubscribeWithLen(event Event, channelLen int) <-chan any
	Subscribe(event Event) <-chan any
	SubscriberLen(event Event) int
	Publish(event Event, data any)
	Unsubscribe(event Event, ch <-chan any)
	UnsubscribeAll(event Event)
	Close()
}

type defaultBus struct {
	subscribers map[Event][]chan any
	mu          sync.RWMutex
}

func NewBus() Bus {
	return &defaultBus{
		subscribers: make(map[Event][]chan any),
		mu:          sync.RWMutex{},
	}
}

func (d *defaultBus) SubscriberLen(event Event) int {
	defer d.mu.Unlock()
	d.mu.Lock()
	if origin, ok := d.subscribers[event]; ok {
		return len(origin)
	}

	return 0
}

func (d *defaultBus) Subscribe(event Event) <-chan any {
	return d.SubscribeWithLen(event, 100000)
}

func (d *defaultBus) SubscribeWithLen(event Event, channelLen int) <-chan any {
	d.mu.Lock()
	defer d.mu.Unlock()

	ch := make(chan any, channelLen)
	origin, ok := d.subscribers[event]

	if ok {
		d.subscribers[event] = append(origin, ch)
	} else {
		d.subscribers[event] = []chan any{ch}
	}

	return ch
}

func (d *defaultBus) Publish(event Event, data any) {
	d.mu.RLock()

	subscribers, ok := d.subscribers[event]
	if !ok {
		d.mu.RUnlock()
		return
	}

	originLen := len(subscribers)
	removeIndexes := make(map[int]bool)

	for i := 0; i < originLen; i++ {
		ch := subscribers[i]
		select {
		case ch <- data:
			// Success
		default:
			// If cannot receive or closed, remove the subscriber
			removeIndexes[i] = true
		}
	}
	d.mu.RUnlock()

	if len(removeIndexes) > 0 {
		d.mu.Lock()
		if originLen == len(d.subscribers[event]) {
			var newSubscribers []chan any

			for index, ch := range d.subscribers[event] {
				if _, is := removeIndexes[index]; !is {
					newSubscribers = append(newSubscribers, ch)
				}
			}

			d.subscribers[event] = newSubscribers
		}
		d.mu.Unlock()
	}
}

func (d *defaultBus) Unsubscribe(event Event, ch <-chan any) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if subscribers, ok := d.subscribers[event]; ok {
		for i, subscriber := range subscribers {
			if subscriber == ch {
				if i == len(subscribers)-1 {
					d.subscribers[event] = subscribers[:i]
				} else {
					d.subscribers[event] = append(subscribers[:i], subscribers[i+1:]...)
				}

				close(subscriber)

				break
			}
		}

		if len(d.subscribers[event]) == 0 {
			delete(d.subscribers, event)
		}
	}
}

func (d *defaultBus) UnsubscribeAll(event Event) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.unsubscribeAll(event)
}

func (d *defaultBus) unsubscribeAll(event Event) {
	if subscribers, ok := d.subscribers[event]; ok {
		for _, subscriber := range subscribers {
			close(subscriber)
		}

		delete(d.subscribers, event)
	}
}

func (d *defaultBus) Close() {
	d.mu.Lock()
	defer d.mu.Unlock()

	for topic, subscribers := range d.subscribers {
		for _, subscriber := range subscribers {
			close(subscriber)
		}

		delete(d.subscribers, topic)
	}
}
