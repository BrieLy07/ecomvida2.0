package controllers

import (
	"encoding/json"
	"net/http"
	"role-permission-service/config"
	"role-permission-service/models"
)

func ObtenerRoles(w http.ResponseWriter, r *http.Request) {
	db := config.ConectarDB()
	rows, err := db.Query(`SELECT id, nombre FROM roles`)
	if err != nil {
		http.Error(w, "Error al obtener roles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var roles []models.Rol
	for rows.Next() {
		var rol models.Rol
		rows.Scan(&rol.ID, &rol.Nombre)
		roles = append(roles, rol)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roles)
}
