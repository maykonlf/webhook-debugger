package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL.String())
	for k, v := range r.Header {
		fmt.Printf("%s: %v\n", k, v)
	}

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("ERR:", err)
	}
	fmt.Println(string(payload))
	fmt.Println("=================================================================================================")

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", handler)

	panic(http.ListenAndServe(":8080", nil))
}
