package main

import (
	"log"
	"os"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/go-chi/cors"
	"github.com/go-chi/chi"
)

func main() {

	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found in variable")
	}
	router := chi.NewRouter()

router.Use(cors.Handle(cors.Options{
	AllowedOrigins: []string{"https://*","http://"},
	AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowedHeaders: []string{"*"},
	AllowCredentials: false,
	MaxAge: 300, // Maximum value not ignored by any of major browsers
}))

	srv := &http.Server{
		Handler: router,
		Addr: ":"+ portString,
	}

	log.Printf("serving passing throught that port %v:",portString)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}


}