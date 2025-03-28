package app

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"zoob-back/internal/server"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Run() {
	port := 8080
	s := server.New(fmt.Sprintf(":%d", port))
	log.Println("Server starting on port", port)
	err := s.Run()
	if err != nil {
		log.Fatal("There's an error occurred running s:", err)
		return
	}
}
