package main

import (
	"address-service/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando .env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3003"
	}

	r := routes.CargarRutas()
	fmt.Printf("✅ Servidor corriendo en el puerto %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
