package routers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Routes() {
	router := mux.NewRouter()
	// router.HandleFunc("/health", (health.Health)).Methods("GET")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
	handler := cors.AllowAll().Handler(router)

	fmt.Println("Servidor escuchando en el puerto", PORT)
	log.Fatal(http.ListenAndServe("127.0.0.1:"+PORT, handler))

}