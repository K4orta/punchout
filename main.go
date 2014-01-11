package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	sse "github.com/k4orta/punchout/sse"
	"net/http"
	"time"
)

func main() {

	clockedIn := false
	lunchBreakTaken := false

	m := martini.Classic()

	m.Get("/press", func(res http.ResponseWriter, req *http.Request) {
		sse.Message("punchEvent", "message")

		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte(fmt.Sprintf("{\"punch\": \"%s\"}", "accepted")))
	})

	http.Handle("/", m)

	go func() {
		for {
			sse.Message("punchEvent", "message")
			time.Sleep(time.Second * 2)
		}
	}()

	http.Handle("/commands", sse.New())
	http.ListenAndServe(":8080", nil)
}
