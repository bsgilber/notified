package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

var subscribers = map[string]*Subscriber{}

type Subscriber struct {
	ID            string   `json:"id,omitempty" msgpack:"id,omitempty"`
	Subscriptions []*Topic `json:"subscriptions,omitempty" msgpack:"subscriptions,omitempty"`
	Created       int64    `json:"created,omitempty" msgpack:"created,omitempty"`
	Updated       int64    `json:"updated,omitempty" msgpack:"updated,omitempty"`
}

type CreateSubscriptionRequest struct {
	SubscriptionId string `json:"subscription_id,omitempty" msgpack:"subscription_id,omitempty"`
	TopicName      string `json:"topic_name,omitempty" msgpack:"topic_name,omitempty"`
}

func CreateSubscriber() *Subscriber {
	id := uuid.New().String()

	subscribers[id] = &Subscriber{
		id,
		[]*Topic{},
		time.Now().UnixMilli(),
		time.Now().UnixMilli(),
	}

	return subscribers[id]
}

func GetSubscriber(SubscriptionId string) *Subscriber {
	return subscribers[SubscriptionId]
}

func DeleteSubscriber(SubscriptionId string) error {
	subscriber, exists := subscribers[SubscriptionId]

	if exists {
		return fmt.Errorf("subscriber with id %s doesn't exist", SubscriptionId)
	}

	for _, j := range subscriber.Subscriptions {
		j.RemoveSubscription(*subscriber)
	}

	for i := range subscribers {
		if subscribers[i].ID == subscriber.ID {
			delete(subscribers, subscriber.ID)
			break
		}
	}

	return nil
}

func FlushSubscribers() {
	subscribers = map[string]*Subscriber{}
}

func ListSubscribers() map[string]*Subscriber {
	return subscribers
}
