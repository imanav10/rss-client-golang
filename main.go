package main

import(
	"log"
	"os"
	"net/http"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/go-chi/cors"
	"github.com/go-chi/chi"
)


func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to Marshal json response %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Contant-Type","application/json")
    w.WriteHeader(code)
	w.Write(dat)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Responding with 5xx error:", message)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	responseWithJSON(w , code, errResponse{Error: message})
}


func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Internal server error")
}



func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, 200, struct{}{})
}

func main() {

	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found in variable")
	}
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*","http://"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: false,
		MaxAge: 300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)

	v1Router.Get("/err", handlerError)

	router.Mount("/v1", v1Router)



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