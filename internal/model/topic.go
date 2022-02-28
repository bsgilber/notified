package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var topics = map[string]*Topic{}

type Topic struct {
	ID          string       `json:"id,omitempty" msgpack:"id,omitempty"`
	TopicName   string       `json:"topicName,omitempty" msgpack:"topicName,omitempty"`
	Subscribers []Subscriber `json:"subscribers,omitempty" msgpack:"subscribers,omitempty"`
	Created     int64        `json:"created,omitempty" msgpack:"created,omitempty"`
	Updated     int64        `json:"updated,omitempty" msgpack:"updated,omitempty"`
}

type CreateTopicRequest struct {
	TopicName string `json:"topicName,omitempty" msgpack:"topicName,omitempty"`
}

func Create(topicName string) ([]byte, error) {
	if topic, exists := topics[topicName]; exists {
		return []byte(""), fmt.Errorf("topic with name %s already exists", topic.TopicName)
	}

	id, err := uuid.New().MarshalBinary()
	if err != nil {
		return nil, err
	}

	topics[topicName] = &Topic{
		uuid.New().String(),
		topicName,
		[]Subscriber{},
		time.Now().UnixMilli(),
		time.Now().UnixMilli(),
	}

	return id, nil
}

func Delete(topicName string) ([]byte, error) {
	if topic, exists := topics[topicName]; exists {
		return []byte(topic.ID), nil
	}

	return []byte(""), fmt.Errorf("topic with name %s does not exist", topicName)
}

func Flush() {
	topics = map[string]*Topic{}
	return
}

func List() []byte {
	if len(topics) == 0 {
		return []byte("{}")
	}

	out, err := json.Marshal(topics)
	if err != nil {
		panic(err)
	}

	return out
}
