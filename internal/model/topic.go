package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var topics = map[string]*Topic{}

type Topic struct {
	ID          string        `json:"id,omitempty" msgpack:"id,omitempty"`
	TopicName   string        `json:"topicName,omitempty" msgpack:"topicName,omitempty"`
	Subscribers []*Subscriber `json:"subscribers,omitempty" msgpack:"subscribers,omitempty"`
	Created     int64         `json:"created,omitempty" msgpack:"created,omitempty"`
	Updated     int64         `json:"updated,omitempty" msgpack:"updated,omitempty"`
}

type CreateTopicRequest struct {
	TopicName string `json:"topicName,omitempty"`
}

func CreateTopic(topicName string) (*Topic, error) {
	if topicName == "" {
		return nil, fmt.Errorf("empty topic name not allowed")
	}

	if topic, exists := topics[topicName]; exists {
		return nil, fmt.Errorf("topic with name %s already exists", topic.TopicName)
	}

	id := uuid.New().String()

	topics[topicName] = &Topic{
		id,
		topicName,
		[]*Subscriber{},
		time.Now().UnixMilli(),
		time.Now().UnixMilli(),
	}

	return topics[topicName], nil
}

func GetTopic(topicName string) *Topic {
	return topics[topicName]
}

func DeleteTopic(topicName string) ([]byte, error) {
	if topic, exists := topics[topicName]; exists {
		return []byte(topic.ID), nil
	}

	return []byte(""), fmt.Errorf("topic with name %s does not exist", topicName)
}

func FlushTopics() {
	topics = map[string]*Topic{}
}

func ListTopics() []byte {
	if len(topics) == 0 {
		return []byte("{}")
	}

	out, err := json.Marshal(topics)
	if err != nil {
		panic(err)
	}

	return out
}

func (t *Topic) AddSubscription(subscriber Subscriber) {
	t.Subscribers = append(t.Subscribers, &subscriber)
	subscriber.Subscriptions = append(subscriber.Subscriptions, *t)
}

func (t *Topic) RemoveSubscription(subscriber Subscriber) {
	for i := range t.Subscribers {
		if t.Subscribers[i].ID == subscriber.ID {
			t.Subscribers = append(t.Subscribers[:i], t.Subscribers[i+1:]...)
			break
		}
	}
}

func (t *Topic) ListSubscribers() []*Subscriber {
	return t.Subscribers
}
