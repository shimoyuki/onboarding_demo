package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"onboarding-demo/service"
)

func main() {
	router := mux.NewRouter()
	userRouter := router.PathPrefix("/users").Subrouter()

	userRouter.HandleFunc("", service.GetWholeUsers).Methods("GET")
	userRouter.HandleFunc("", service.CreateUser).Methods("POST")

	userRouter.HandleFunc("/{user_id}/relationships", service.GetRelationshipsByUser).Methods("GET")
	userRouter.HandleFunc("/{user_id}/relationships/{other_user_id}", service.CreateOrUpdateRelationships).Methods("PUT")


	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
