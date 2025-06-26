package routes

import (
	"role-permission-service/controllers"

	"github.com/gorilla/mux"
)

func CargarRutas() *mux.Router {
	r := mux.NewRouter()

	// Roles
	r.HandleFunc("/roles", controllers.ObtenerRoles).Methods("GET")
	r.HandleFunc("/roles", controllers.CrearRol).Methods("POST")

	// Permisos
	r.HandleFunc("/permissions", controllers.ObtenerPermisos).Methods("GET")
	r.HandleFunc("/permissions", controllers.CrearPermiso).Methods("POST")

	// Usuario → Roles
	r.HandleFunc("/users/{id}/roles", controllers.ObtenerRolesUsuario).Methods("GET")
	r.HandleFunc("/users/{id}/roles", controllers.AsignarRolesUsuario).Methods("POST")

	return r
}
