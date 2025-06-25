package routes

import (
	"address-service/controllers"

	"github.com/gorilla/mux"
)

func CargarRutas() *mux.Router {
	r := mux.NewRouter()

	// Rutas para direcciones
	r.HandleFunc("/users/{id}/addresses", controllers.CrearDireccion).Methods("POST")
	r.HandleFunc("/users/{id}/addresses", controllers.ObtenerDirecciones).Methods("GET")
	r.HandleFunc("/addresses/{addressId}", controllers.ActualizarDireccion).Methods("PUT")
	r.HandleFunc("/addresses/{addressId}", controllers.EliminarDireccion).Methods("DELETE")

	return r
}
