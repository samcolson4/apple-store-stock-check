package main

import (
	"fmt"
	"time"

	uilive "github.com/gosuri/uilive"
)

func main() {
	headers := uilive.New()
	timeline := uilive.New()

	// start listening for updates and render
	headers.Start()
	timeline.Start()

	for i := 0; i <= 1000; i++ {
		timestamp := getTime()
		fmt.Fprintf(timeline, "Last updated: %s\n", timestamp)
		fmt.Fprintf(headers, "Product\t| Availability\n")
		time.Sleep(time.Second * 5)
	}

	// flush and stop rendering
	headers.Stop()
	timeline.Stop()
}

func getTime() string {
	t := time.Now()
	timestamp := fmt.Sprintf("%02d:%02d:%02d",
		t.Hour(), t.Minute(), t.Second())

	return timestamp
}
