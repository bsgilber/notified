package model

type Consumer struct {
	ID             string `json:"id,omitempty" msgpack:"id,omitempty"`
	SubscriptionId string `json:"subscriptions,omitempty" msgpack:"subscriptions,omitempty"`
	Created        int64  `json:"created,omitempty" msgpack:"created,omitempty"`
	Updated        int64  `json:"updated,omitempty" msgpack:"updated,omitempty"`
}
