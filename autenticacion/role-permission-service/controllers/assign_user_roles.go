package controllers

import (
	"encoding/json"
	"net/http"
	"role-permission-service/config"
	"strconv"

	"github.com/gorilla/mux"
)

func AsignarRolesUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usuarioID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var entrada struct {
		Roles []int64 `json:"roles"`
	}

	if err := json.NewDecoder(r.Body).Decode(&entrada); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	db := config.ConectarDB()
	for _, rolID := range entrada.Roles {
		_, err := db.Exec(`INSERT INTO usuario_rol (usuario_id, rol_id) VALUES ($1, $2)`, usuarioID, rolID)
		if err != nil {
			http.Error(w, "Error al asignar rol", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"mensaje":"Roles asignados correctamente"}`))
}
