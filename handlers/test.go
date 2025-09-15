package handlers

import (
	"log"
	"net/http"
)

func Test (w http.ResponseWriter, r *http.Request) {
	log.Println("ami handler : ami mid e print hobo")
}