package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	sse "github.com/k4orta/punchout/sse"
	"net/http"
)

func main() {

	status := "clockedOut"
	update := make(chan string)
	m := martini.Classic()

	m.Get("/press", func(res http.ResponseWriter, req *http.Request) {
		sse.Message("punchEvent", "message")
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte(fmt.Sprintf("{\"punch\": \"%s\"}", "clockIn")))
	})

	m.Get("/confirm", func(res http.ResponseWriter, req *http.Request) {
		status = <-update
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte(fmt.Sprintf("{\"done\": \"%s\"}", status)))
	})

	m.Get("/update/:status", func(params martini.Params) string {
		if params["status"] != status {
			update <- params["status"]
		}
		return params["status"]
	})

	http.Handle("/", m)

	http.Handle("/commands", sse.New())
	fmt.Println("Starting server")
	http.ListenAndServe(":8081", nil)
}
