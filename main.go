package main

import (
	"fmt"
	"net/http"
	"log"
	"os"

    "github.com/aws/aws-lambda-go/lambda"
    "github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

const port = ":8080"

func main() {
    mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)

	if os.Getenv("LOCAL") == "true" {
		fmt.Println("Handling request locally on port", port)	
		log.Fatal(http.ListenAndServe(port, mux))
	} else {
		fmt.Println("Starting Lambda")

		// Create a lambda adapter to allow for normal HTTP traffic flow
		adapter := httpadapter.New(mux)
		lambda.Start(adapter.ProxyWithContext) // blocking function
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page")
}

