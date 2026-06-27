package main

import "concurrent-job-processing-system/internal/api"

func main() {
	server := api.New()
	server.Run()
}
