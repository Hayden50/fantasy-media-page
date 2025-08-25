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
	defineRoutes(mux)

	if os.Getenv("LOCAL") == "true" {
		fmt.Println("Handling request locally on port", port)	
		log.Fatal(http.ListenAndServe(port, mux))
	} else {
		fmt.Println("Starting Lambda")
		adapter := httpadapter.New(mux)
		lambda.Start(adapter.ProxyWithContext) // blocking function
	}
}

func defineRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", homeHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request")
	fmt.Fprintln(w, "Welcome to dynasty fantasy coverage")
}

