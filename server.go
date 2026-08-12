package main

import (
	"os"
	"log"
	"fmt"
	"net/http"
	"encoding/json"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(
		map[string]string{
			"body": fmt.Sprintf("Reached server listening on port: %s", os.Getenv("PORT")),
		},
	)

	w.Write(res)
}


func main() {
	os.Setenv("PORT", "80")

	http.HandleFunc("/", requestHandler)

	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
	if err != nil {
		log.Panic("Server listening and serving error")
	}
}