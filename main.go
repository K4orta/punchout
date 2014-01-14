package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	sse "github.com/k4orta/punchout/sse"
	"net/http"
)

func punchRecieve(status string) string {
	// clockIn -> mealBreak -> clockIn -> clockOut
	switch status {
	case "clockedOut":
		return "clockedIn"
	case "clockedIn":
		return "mealBreak"
	case "mealBreak":
		return "clockedIn"
	}
	return ""
}

func main() {

	status := "clockedOut"
	update := make(chan string)
	waiting := false
	m := martini.Classic()

	// Middleware to set responses Content-Type to application/json
	m.Use(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
	})

	// Sends a command to the js extension and outputs which command was sent.
	m.Get("/press", func(res http.ResponseWriter, req *http.Request) {
		newStatus := punchRecieve(status)
		sse.Message("punchEvent", newStatus)
		res.Write([]byte(fmt.Sprintf("{\"punch\": \"%s\"}", newStatus)))
	})

	// Blocks until a status arrives through the update channel.
	m.Get("/confirm", func(res http.ResponseWriter, req *http.Request) {
		waiting = true
		status = <-update
		waiting = false
		res.Write([]byte(fmt.Sprintf("{\"done\": \"%s\"}", status)))
	})

	//
	m.Get("/update/:status", func(params martini.Params) string {
		changed := false
		if waiting && params["status"] != status {
			changed = true
			update <- params["status"]
		}
		return "{\"statusChaged\": " + fmt.Sprint(changed) + "}"
	})

	http.Handle("/", m)

	http.Handle("/commands", sse.New())
	fmt.Println("Starting server")
	http.ListenAndServe(":8081", nil)
}
