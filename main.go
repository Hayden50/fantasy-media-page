package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {

	res, httpErr := http.Get("https://api.sleeper.app/v1/league/1182033902772338688")
	if httpErr != nil {
		logrus.Error("error fetching content from sleeper app")
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logrus.Error("error reading response body: ", err)
		return
	}

	fmt.Printf("HTTP Status: %s\n", res.Status)
	fmt.Printf("Response Body:\n%s\n", body)
}
