package utils

import (
	m "gqlgen-subscriptions/graph/model"
	"sync"
)

type Pubsub struct {
	mutex       sync.RWMutex
	subscribers map[string][]chan *m.Event
}

func NewPubsub() *Pubsub {
	return &Pubsub{
		mutex:       sync.RWMutex{},
		subscribers: make(map[string][]chan *m.Event),
	}
}

func (pubsub *Pubsub) Publish(eventName string, e *m.Event) {
	pubsub.mutex.RLock()
	defer pubsub.mutex.RUnlock()

	for _, channel := range pubsub.subscribers[eventName] {
		channel <- e
	}
}

func (pubsub *Pubsub) Subscribe(eventName string) <-chan *m.Event {
	pubsub.mutex.Lock()
	defer pubsub.mutex.Unlock()

	channel := make(chan *m.Event, 1)
	pubsub.subscribers[eventName] = append(pubsub.subscribers[eventName], channel)

	return channel
}

func (pubsub *Pubsub) Unsubscribe(eventName string, channel <-chan *m.Event) {
	pubsub.mutex.Lock()
	defer pubsub.mutex.Unlock()

	for i, currentChannel := range pubsub.subscribers[eventName] {
		if currentChannel == channel {
			pubsub.subscribers[eventName] = append(pubsub.subscribers[eventName][:i], pubsub.subscribers[eventName][i+1:]...)
			return
		}
	}
}
