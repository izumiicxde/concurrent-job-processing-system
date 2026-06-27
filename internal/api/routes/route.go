package routes

import (
	"log"
	"net/http"
)

func HandleHealthRoute(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("System up and working...."))
	if err != nil {
		log.Println("Error writing to route: ", err)
	}
}
