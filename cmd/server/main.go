package main

import (
	api "concurrent-job-processing-system/internal/api/routes"
	"fmt"
	"log"
)

func main() {
	server := api.InitializeAPI()

	fmt.Println("Server running on port 8000")
	log.Fatal(server.ListenAndServe())
}
