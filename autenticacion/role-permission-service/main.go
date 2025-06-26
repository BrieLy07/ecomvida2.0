package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"role-permission-service/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando .env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3005"
	}

	r := routes.CargarRutas()

	fmt.Printf("✅ role-permission-service corriendo en el puerto %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
