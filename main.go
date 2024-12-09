package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// Adaugă ruta pentru handler
	router.Get("/hello", basicHandler)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	fmt.Println("Server starting on port 3000...")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to listen to server:", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
