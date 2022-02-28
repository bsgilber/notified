package main

import (
	"log"
	"net/http"

	"github.com/bsgilber/notified/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	healthCheck := http.HandlerFunc(handler.Health)
	subscribe := http.HandlerFunc(handler.Subscribe)
	unsubscribe := http.HandlerFunc(handler.Unsubscribe)
	publish := http.HandlerFunc(handler.Publish)
	consume := http.HandlerFunc(handler.Consume)
	topic := http.HandlerFunc(handler.Topic)
	topicList := http.HandlerFunc(handler.TopicList)

	mux.Handle("/", healthCheck)
	mux.Handle("/health", healthCheck)
	mux.Handle("/subscribe", subscribe)
	mux.Handle("/unsubscribe", unsubscribe)
	mux.Handle("/publish", publish)
	mux.Handle("/consume", consume)
	mux.Handle("/topic", topic)
	mux.Handle("/topic/list", topicList)

	log.Fatal(http.ListenAndServe(":7077", mux))
}
