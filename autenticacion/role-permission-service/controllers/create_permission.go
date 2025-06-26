package controllers

import (
	"encoding/json"
	"net/http"
	"role-permission-service/models"
	"role-permission-service/strategies"
)

func CrearPermiso(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Permiso
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	estrategia := strategies.PermisoEstrategia{}
	resultado, err := estrategia.Ejecutar(nuevo)
	if err != nil {
		http.Error(w, "Error al crear permiso", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resultado)
}
