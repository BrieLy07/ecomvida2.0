package controllers

import (
	"address-service/repository"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func EliminarDireccion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	addressID, err := strconv.ParseInt(vars["addressId"], 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	repo := repository.NuevoAddressRepo()
	err = repo.Eliminar(addressID)
	if err != nil {
		http.Error(w, "No se pudo eliminar la dirección", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"mensaje":"Dirección eliminada correctamente"}`))
}
