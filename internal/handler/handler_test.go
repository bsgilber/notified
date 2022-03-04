package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealth(t *testing.T) {
	tests := []struct {
		name       string
		want       string
		method     string
		statusCode int
	}{
		{
			name:       "validate /health endpoint",
			want:       "{\"ResponseMessage\":\"Healthy\",\"ResponseCode\":200}",
			method:     http.MethodGet,
			statusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, "/health", nil)
			responseRecorder := httptest.NewRecorder()

			Health(responseRecorder, request)
			if responseRecorder.Code != tt.statusCode {
				t.Errorf("Want status '%d', got '%d'", tt.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tt.want {
				t.Errorf("Want '%s', got '%s'", tt.want, responseRecorder.Body)
			}

		})
	}
}

func TestSubscribe(t *testing.T) {
	tests := []struct {
		name       string
		want       string
		method     string
		statusCode int
	}{
		{
			name:       "validate /subscribe endpoint",
			want:       "",
			method:     http.MethodGet,
			statusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, "/subscribe", nil)
			responseRecorder := httptest.NewRecorder()

			Subscribe(responseRecorder, request)
		})
	}
}

func TestUnsubscribe(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Unsubscribe(tt.args.w, tt.args.req)
		})
	}
}

func TestPublish(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Publish(tt.args.w, tt.args.req)
		})
	}
}

func TestConsume(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Consume(tt.args.w, tt.args.req)
		})
	}
}

func TestTopicList(t *testing.T) {
	tests := []struct {
		name       string
		want       string
		method     string
		statusCode int
	}{
		{
			name:       "list empty topic list",
			want:       "{}",
			method:     http.MethodGet,
			statusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, "/topic/list", nil)
			responseRecorder := httptest.NewRecorder()

			TopicList(responseRecorder, request)

			if responseRecorder.Code != tt.statusCode {
				t.Errorf("Want status '%d', got '%d'", tt.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tt.want {
				t.Errorf("Want '%s', got '%s'", tt.want, responseRecorder.Body)
			}
		})
	}
}

func TestTopic(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		want       string
		method     string
		statusCode int
	}{
		{
			name:       "create topic test",
			body:       `{"topicName":"testTopic"}`,
			method:     http.MethodPost,
			statusCode: http.StatusOK,
		},
		{
			name:       "fail create topic with bad payload",
			body:       `{"badTopicName":"badTopic"}`,
			method:     http.MethodPost,
			statusCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, "/topic", strings.NewReader(tt.body))
			responseRecorder := httptest.NewRecorder()

			Topic(responseRecorder, request)

			if responseRecorder.Code != tt.statusCode {
				t.Errorf("Want status '%d', got '%d'", tt.statusCode, responseRecorder.Code)
			}
		})
	}
}
