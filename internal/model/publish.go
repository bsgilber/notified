package model

type Publisher struct {
	ID      string  `json:"id,omitempty" msgpack:"id,omitempty"`
	Topics  []Topic `json:"topics,omitempty" msgpack:"topics,omitempty"`
	Created int64   `json:"created,omitempty" msgpack:"created,omitempty"`
	Updated int64   `json:"updated,omitempty" msgpack:"updated,omitempty"`
}
