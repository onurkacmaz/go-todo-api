package route

import (
	"github.com/gorilla/mux"
	"rest-api/controller"
)

var router = route()

func route() *mux.Router {
	return mux.NewRouter()
}

func RegisterRoutes() *mux.Router {
	usersRouter := router.PathPrefix("/api/v1").Subrouter()
	usersRouter.HandleFunc("/users", controller.CreateUser).Methods("POST")
	usersRouter.HandleFunc("/users", controller.GetUsers).Methods("GET")
	usersRouter.HandleFunc("/users/{id}", controller.ShowUser).Methods("GET")
	usersRouter.HandleFunc("/users/{id}", controller.DeleteUser).Methods("DELETE")
	usersRouter.HandleFunc("/users/{id}", controller.UpdateUser).Methods("PUT")

	usersRouter.HandleFunc("/tasks", controller.CreateTask).Methods("POST")
	usersRouter.HandleFunc("/tasks", controller.GetTasks).Methods("GET")
	usersRouter.HandleFunc("/tasks/{id}", controller.ShowTask).Methods("GET")
	usersRouter.HandleFunc("/tasks/{id}", controller.DeleteTask).Methods("DELETE")
	usersRouter.HandleFunc("/tasks/{id}", controller.UpdateTask).Methods("PUT")
	return router
}
