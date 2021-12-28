package main

import (
	"fmt"
	"time"

	uilive "github.com/gosuri/uilive"
)

func main() {
	output := uilive.New()

	// List in reverse order for correct display.
	output.Start()

	for i := 0; i <= 1000; i++ {
		timestamp := getTime()

		timeOutput := fmt.Sprintf("Last updated: %s\n", timestamp)
		headers := fmt.Sprintf("Product\t\t\t| Availability\n")
		productAvailability := getProductAvailability()

		outputString := timeOutput + headers + productAvailability

		fmt.Fprintf(output, outputString)

		time.Sleep(time.Second * 5)
	}

	// flush and stop rendering
	output.Stop()
}

func getTime() string {
	t := time.Now()
	timestamp := fmt.Sprintf("%02d:%02d:%02d",
		t.Hour(), t.Minute(), t.Second())

	return timestamp
}

func getProductAvailability() string {
	return "iPhone 13 Pro Max \t| ✅\t\niPhone 13 \t\t| ✅\n"
}
