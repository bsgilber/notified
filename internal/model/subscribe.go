package model

type Subscriber struct {
	ID            string  `json:"id,omitempty" msgpack:"id,omitempty"`
	Subscriptions []Topic `json:"subscriptions,omitempty" msgpack:"subscriptions,omitempty"`
	Created       int64   `json:"created,omitempty" msgpack:"created,omitempty"`
	Updated       int64   `json:"updated,omitempty" msgpack:"updated,omitempty"`
}

func (t *Topic) AddSubscriber(subscriber Subscriber) {
	t.Subscribers = append(t.Subscribers, subscriber)
}

func (t *Topic) DeleteSubscriber(subscriber Subscriber) {
	t.Subscribers = append(t.Subscribers, subscriber)
}

func (t *Topic) ListSubscribers() []Subscriber {
	return t.Subscribers
}
