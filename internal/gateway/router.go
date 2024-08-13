package gateway

import (
	"github.com/TanmoyTSSaha/GoBase/api/v1/middleware"
	"github.com/TanmoyTSSaha/GoBase/api/v1/routes"
	"github.com/TanmoyTSSaha/GoBase/pkg/services"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRouter(db *mongo.Database) *mux.Router {
	r := mux.NewRouter()

	// INITIALIZING LOG SERVICES
	logService := services.NewLogService(db)

	// GLOBAL MIDDLEWARE
	r.Use(middleware.LoggingMiddleware(logService))

	// API V1 ROUTES
	v1Routes := r.PathPrefix("/api/v1").Subrouter()
	routes.SetupRoutes(v1Routes)

	return r
}