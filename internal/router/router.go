package router

import (
	v1 "github.com/TanmoyTSSaha/GoBase/api/v1"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()
	v1.SetupRoutes(r)

	return r
}
