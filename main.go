package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	uilive "github.com/gosuri/uilive"
)

func main() {
	output := uilive.New()
	output.Start()

	for i := 0; i <= 1000; i++ {
		timestamp := getTime()

		timeOutput := fmt.Sprintf("Last updated: %s\n", timestamp)
		headers := fmt.Sprintf("Product\t\t\t| Availability\n")
		productAvailability := getProductAvailability("MLLA3B/A")

		outputString := timeOutput + headers + productAvailability

		fmt.Fprintf(output, outputString)

		time.Sleep(time.Second * 5)
	}

	output.Stop()
}

func getTime() string {
	t := time.Now()
	timestamp := fmt.Sprintf("%02d:%02d:%02d",
		t.Hour(), t.Minute(), t.Second())

	return timestamp
}

func getProductAvailability(productCode string) string {
	var productOutput string
	var j interface{}

	postcode := os.Getenv("POSTCODE")

	resp, err := http.Get("https://www.apple.com/uk/shop/pickup-message-recommendations?location=" + postcode + "&product=MLLA3B/A")
	catchError(err)

	err = json.NewDecoder(resp.Body).Decode(&j)
	catchError(err)

	// productOutput = ""

	return productOutput
}

func catchError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
