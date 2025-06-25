package controllers

import (
	"address-service/models"
	"address-service/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ActualizarDireccion(w http.ResponseWriter, r *http.Request) {
	var actualizada models.Address
	err := json.NewDecoder(r.Body).Decode(&actualizada)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	addressID, err := strconv.ParseInt(vars["addressId"], 10, 64)
	if err != nil {
		http.Error(w, "ID de dirección inválido", http.StatusBadRequest)
		return
	}

	actualizada.ID = addressID
	repo := repository.NuevoAddressRepo()
	err = repo.Actualizar(actualizada)
	if err != nil {
		http.Error(w, "No se pudo actualizar la dirección", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(actualizada)
}
