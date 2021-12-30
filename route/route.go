package route

import (
	"github.com/gorilla/mux"
	"rest-api/controller"
	"rest-api/middleware"
)

var router = route()

func route() *mux.Router {
	return mux.NewRouter()
}

func RegisterRoutes() *mux.Router {

	authorizationRouter := router.PathPrefix("/api/v1/auth").Subrouter()
	authorizationRouter.HandleFunc("/token", controller.CreateToken).Methods("POST")

	authenticatedRouter := router.PathPrefix("/api/v1").Subrouter()
	authenticatedRouter.Use(middleware.ContentType, middleware.Accept, middleware.Auth)

	authenticatedRouter.HandleFunc("/users", controller.CreateUser).Methods("POST")
	authenticatedRouter.HandleFunc("/users", controller.GetUsers).Methods("GET")
	authenticatedRouter.HandleFunc("/users/{id}", controller.ShowUser).Methods("GET")
	authenticatedRouter.HandleFunc("/users/{id}", controller.DeleteUser).Methods("DELETE")
	authenticatedRouter.HandleFunc("/users/{id}", controller.UpdateUser).Methods("PUT")

	authenticatedRouter.HandleFunc("/tasks", controller.CreateTask).Methods("POST")
	authenticatedRouter.HandleFunc("/tasks", controller.GetTasks).Methods("GET")
	authenticatedRouter.HandleFunc("/tasks/{id}", controller.ShowTask).Methods("GET")
	authenticatedRouter.HandleFunc("/tasks/{id}", controller.DeleteTask).Methods("DELETE")
	authenticatedRouter.HandleFunc("/tasks/{id}", controller.UpdateTask).Methods("PUT")

	return router
}
