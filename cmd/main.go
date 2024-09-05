package main

import (
    "log"
    "net/http"
    "103HW/config"
    "103HW/handlers"
    "103HW/middleware"
    "github.com/gorilla/mux"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    r := mux.NewRouter()

    r.HandleFunc("/users", handlers.CreateUserHandler).Methods("POST")
    r.HandleFunc("/users/{id}", handlers.GetUserHandler).Methods("GET")
    r.HandleFunc("/users/{id}", handlers.UpdateUserHandler).Methods("PUT")
    r.HandleFunc("/users/{id}", handlers.DeleteUserHandler).Methods("DELETE")

    r.Use(middleware.LoggingMiddleware)

    log.Printf("Server running on port %s", cfg.DBPort)
    log.Fatal(http.ListenAndServe(":8080", r))
}
