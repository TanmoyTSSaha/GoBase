package main

import (
	"context"
	"log"
	"net/http"

	"github.com/TanmoyTSSaha/GoBase/configs"
	"github.com/TanmoyTSSaha/GoBase/internal/database/mongodb"
	"github.com/TanmoyTSSaha/GoBase/internal/gateway"
)

func main() {
	if err := configs.LoadConfig(); err != nil {
		log.Fatalf("ERROR LOADING CONFIG: %v", err)
	}
	
	mongoClient, err := mongodb.MongoConnect(configs.Config.MongoDB.URI, configs.Config.MongoDB.Database)
	if err != nil {
		log.Fatalf("ERROR CONNECTING MONGODB: %v", err)
	}

	defer mongoClient.Client.Disconnect(context.Background())

	r := gateway.InitRouter(mongoClient.Database)

	// SERVE STATIC FILES (CSS, JS & IMAGES)
	staticFileHandler := http.FileServer(http.Dir("pkg/templates/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFileHandler))

	http.Handle("/", r)
	log.Println("Server is starting on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
