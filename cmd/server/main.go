package main

import (
	"log"
	"net/http"
	"github.com/TanmoyTSSaha/GoBase/internal/router"
)

func main()  {
	r := router.InitRouter()

	// SERVE STATIC FILES (CSS, JS & IMAGES)
	staticFileHandler := http.FileServer(http.Dir("pkg/templates/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFileHandler))

	http.Handle("/", r)
	log.Println("Server is starting on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

