package controllers

import (
	"address-service/models"
	"address-service/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CrearDireccion(w http.ResponseWriter, r *http.Request) {
	var nueva models.Address
	err := json.NewDecoder(r.Body).Decode(&nueva)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	usuarioID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "ID de usuario inválido", http.StatusBadRequest)
		return
	}

	nueva.UsuarioID = usuarioID
	repo := repository.NuevoAddressRepo()
	id, err := repo.Crear(nueva)
	if err != nil {
		http.Error(w, "No se pudo guardar la dirección", http.StatusInternalServerError)
		return
	}

	nueva.ID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nueva)
}
