package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	fmt.Println("Starting EchoServer on port", port, "of the container...")
	err := http.ListenAndServe(":"+port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := map[string]any{
			"Proto": r.Proto,
			"Host":  r.Host,
		}
		for key, value := range r.Header {
			if len(value) <= 1 {
				headers[key] = r.Header.Get(key)
				continue
			}
			headers[key] = r.Header[key]
		}
		b, err := json.Marshal(headers)
		if err != nil {
			log.Printf("unable to marshal headers as JSON: %s\n", err)
		}

		log := fmt.Sprintf("received request: %s %s%s\nwith headers: %s\n", r.Method, r.Host, r.URL, string(b))

		fmt.Print(log)
		fmt.Fprint(w, log)
	}))
	if err != nil {
		log.Printf("error running EchoServer: %s\n", err)
	}
}
