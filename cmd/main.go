package main

import (
	"fmt"
	"grafana-example/internal/router"
	"log"
	"net/http"
)

func main() {
	r := router.Init()
	fmt.Println("Serving requests on port 3000")
	err := http.ListenAndServe(":3000", r)
	log.Fatal(err)
}
