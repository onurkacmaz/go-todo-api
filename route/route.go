package route

import (
	"github.com/gorilla/mux"
	"rest-api/controllers"
)

var router = route()

func route() *mux.Router {
	return mux.NewRouter()
}

func ConfigureRoutes() *mux.Router {
	usersRouter := router.PathPrefix("/api/v1").Subrouter()
	usersRouter.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	usersRouter.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	usersRouter.HandleFunc("/users/{id}", controllers.ShowUser).Methods("GET")
	usersRouter.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
	usersRouter.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	return router
}
