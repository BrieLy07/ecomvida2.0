package controllers

import (
	"address-service/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ObtenerDirecciones(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	usuarioID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "ID de usuario inválido", http.StatusBadRequest)
		return
	}

	repo := repository.NuevoAddressRepo()
	direcciones, err := repo.ListarPorUsuario(usuarioID)
	if err != nil {
		http.Error(w, "Error al obtener direcciones", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(direcciones)
}
