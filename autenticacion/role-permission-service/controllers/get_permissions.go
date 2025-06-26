package controllers

import (
	"encoding/json"
	"net/http"
	"role-permission-service/config"
	"role-permission-service/models"
)

func ObtenerPermisos(w http.ResponseWriter, r *http.Request) {
	db := config.ConectarDB()
	rows, err := db.Query(`SELECT id, nombre FROM permisos`)
	if err != nil {
		http.Error(w, "Error al obtener permisos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var permisos []models.Permiso
	for rows.Next() {
		var permiso models.Permiso
		rows.Scan(&permiso.ID, &permiso.Nombre)
		permisos = append(permisos, permiso)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(permisos)
}
