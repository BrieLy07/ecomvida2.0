package repository

import (
	"address-service/models"
)

type AddressRepository interface {
	Crear(address models.Address) (int64, error)
	ListarPorUsuario(usuarioID int64) ([]models.Address, error)
	Actualizar(address models.Address) error
	Eliminar(addressID int64) error
}
