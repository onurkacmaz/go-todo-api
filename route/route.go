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

	router.Use(middleware.ContentType, middleware.Accept, middleware.BasicAuth)

	usersRouter := router.PathPrefix("/api/v1").Subrouter()
	usersRouter.HandleFunc("/users", controller.CreateUser).Methods("POST")
	usersRouter.HandleFunc("/users", controller.GetUsers).Methods("GET")
	usersRouter.HandleFunc("/users/{id}", controller.ShowUser).Methods("GET")
	usersRouter.HandleFunc("/users/{id}", controller.DeleteUser).Methods("DELETE")
	usersRouter.HandleFunc("/users/{id}", controller.UpdateUser).Methods("PUT")

	tasksRouter := router.PathPrefix("/api/v1").Subrouter()
	tasksRouter.HandleFunc("/tasks", controller.CreateTask).Methods("POST")
	tasksRouter.HandleFunc("/tasks", controller.GetTasks).Methods("GET")
	tasksRouter.HandleFunc("/tasks/{id}", controller.ShowTask).Methods("GET")
	tasksRouter.HandleFunc("/tasks/{id}", controller.DeleteTask).Methods("DELETE")
	tasksRouter.HandleFunc("/tasks/{id}", controller.UpdateTask).Methods("PUT")
	return router
}
