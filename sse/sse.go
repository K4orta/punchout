package sse

import (
	eventsource "github.com/antage/eventsource/http"
	"net/http"
	"strconv"
)

var Src eventsource.EventSource
var eventId = 0

func New() eventsource.EventSource {
	Src = eventsource.New(eventsource.DefaultSettings(), func(req *http.Request) [][]byte {
		return [][]byte{
			[]byte("X-Accel-Buffering: no"),
			[]byte("Access-Control-Allow-Origin: " + req.Header.Get("Origin")),
			[]byte("Access-Control-Allow-Credentials: true"),
		}
	})
	return Src
}

func Message(kind string, message string) {
	eventId += 1
	if Src != nil {
		Src.SendMessage(message, kind, strconv.Itoa(int(eventId)))
	}
}
