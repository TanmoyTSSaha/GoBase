package routes

import (
	"github.com/gorilla/mux"
	"github.com/TanmoyTSSaha/GoBase/pkg/project"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/", project.RenderHomePage).Methods("GET")
	router.HandleFunc("/login", project.LoginPage).Methods("GET")
}