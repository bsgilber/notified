package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bsgilber/notified/internal/model"
)

type HealthCheckResponse struct {
	ResponseMessage string
	ResponseCode    int
}

func Health(w http.ResponseWriter, req *http.Request) {
	// create and format healthy response
	resp := &HealthCheckResponse{"Healthy", 200}
	out, err := json.Marshal(resp)
	if err != nil {
		var msg string = "{\"ResponseMessage\": \"Failed to marshal health response.\", 500 }"
		out = []byte(msg)
	}

	// set headers
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(out)
	if err != nil {
		panic(err)
	}
}

func Subscribe(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
	case http.MethodPost:
		var request model.CreateSubscriptionRequest

		err := json.NewDecoder(req.Body).Decode(&request)
		if err != nil {
			panic(err)
		}

		if request.TopicName == "" || request.SubscriberId == "" {
			http.Error(w, fmt.Sprintf("Error: both 'topic_name' and 'subscription_id' are required fields in the RequestBody"), http.StatusBadRequest)
		}

		topic, err := model.GetTopic(request.TopicName)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusBadRequest)
		}

		subscriber, err := model.GetSubscriber(request.SubscriberId)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusBadRequest)
		}

		topic.AddSubscription(subscriber)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(topic.ID))
	case http.MethodOptions:
		w.Header().Set("Allow", "POST, OPTIONS")
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "POST, OPTIONS")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func Unsubscribe(w http.ResponseWriter, req *http.Request) {

}

func Publish(w http.ResponseWriter, req *http.Request) {

}

func Consume(w http.ResponseWriter, req *http.Request) {
	return
}

func Topic(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
	case http.MethodPost:
		var request model.CreateTopicRequest

		err := json.NewDecoder(req.Body).Decode(&request)
		if err != nil {
			panic(err)
		}

		if request.TopicName == "" {
			http.Error(w, fmt.Sprintf("Error: topicName is a required field in the RequestBody"), http.StatusBadRequest)
		}

		topic, err := model.CreateTopic(request.TopicName)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error [create topic]: %s", err), http.StatusBadRequest)
		}

		out, err := json.Marshal(topic)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusBadRequest)
		}

		w.Write(out)
	case http.MethodOptions:
		w.Header().Set("Allow", "GET, POST, OPTIONS")
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "GET, POST, OPTIONS")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func TopicList(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		topics := model.ListTopics()

		out, err := json.Marshal(topics)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		w.Write(out)
	case http.MethodPost:
	case http.MethodOptions:
		w.Header().Set("Allow", "POST, OPTIONS")
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "POST, OPTIONS")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
