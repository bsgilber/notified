package handler

import (
	"encoding/json"
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
	// Common code for all requests can go here...

	switch req.Method {
	case http.MethodGet:
	case http.MethodPost:
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
	// Common code for all requests can go here...

	switch req.Method {
	case http.MethodGet:
	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)

		var request model.CreateTopicRequest

		err := decoder.Decode(&request)
		if err != nil {
			panic(err)
		}

		model.Create(request.TopicName)
	case http.MethodOptions:
		w.Header().Set("Allow", "GET, POST, OPTIONS")
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "GET, POST, OPTIONS")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func TopicList(w http.ResponseWriter, req *http.Request) {
	// Common code for all requests can go here...

	switch req.Method {
	case http.MethodGet:
		// set headers
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		w.Write(model.List())
	case http.MethodPost:
	case http.MethodOptions:
		w.Header().Set("Allow", "POST, OPTIONS")
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "POST, OPTIONS")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
