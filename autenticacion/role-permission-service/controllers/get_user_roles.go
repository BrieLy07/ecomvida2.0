package controllers

import (
	"encoding/json"
	"net/http"
	"role-permission-service/config"
	"role-permission-service/models"
	"strconv"

	"github.com/gorilla/mux"
)

func ObtenerRolesUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usuarioID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	db := config.ConectarDB()
	query := `
		SELECT r.id, r.nombre
		FROM roles r
		INNER JOIN usuario_rol ur ON r.id = ur.rol_id
		WHERE ur.usuario_id = $1
	`

	rows, err := db.Query(query, usuarioID)
	if err != nil {
		http.Error(w, "Error al obtener roles del usuario", http.StatusInternalServerError)
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
