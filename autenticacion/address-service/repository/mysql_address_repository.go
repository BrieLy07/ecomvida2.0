package repository

import (
	"address-service/config"
	"address-service/models"
	"database/sql"
	"log"
)

type MySQLAddressRepository struct {
	db *sql.DB
}

func NuevoAddressRepo() AddressRepository {
	return &MySQLAddressRepository{
		db: config.ConectarDB(),
	}
}

func (r *MySQLAddressRepository) Crear(address models.Address) (int64, error) {
	query := `
		INSERT INTO direcciones (usuario_id, calle, ciudad, provincia, codigo_postal, pais)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	resultado, err := r.db.Exec(query, address.UsuarioID, address.Calle, address.Ciudad, address.Provincia, address.CodigoPostal, address.Pais)
	if err != nil {
		return 0, err
	}

	id, _ := resultado.LastInsertId()
	return id, nil
}

func (r *MySQLAddressRepository) ListarPorUsuario(usuarioID int64) ([]models.Address, error) {
	query := `SELECT id, usuario_id, calle, ciudad, provincia, codigo_postal, pais FROM direcciones WHERE usuario_id = ?`

	rows, err := r.db.Query(query, usuarioID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var direcciones []models.Address

	for rows.Next() {
		var a models.Address
		err := rows.Scan(&a.ID, &a.UsuarioID, &a.Calle, &a.Ciudad, &a.Provincia, &a.CodigoPostal, &a.Pais)
		if err != nil {
			log.Println("Error al escanear dirección:", err)
			continue
		}
		direcciones = append(direcciones, a)
	}

	return direcciones, nil
}

func (r *MySQLAddressRepository) Actualizar(address models.Address) error {
	query := `
		UPDATE direcciones SET
			calle = ?, ciudad = ?, provincia = ?, codigo_postal = ?, pais = ?
		WHERE id = ?
	`
	_, err := r.db.Exec(query, address.Calle, address.Ciudad, address.Provincia, address.CodigoPostal, address.Pais, address.ID)
	return err
}

func (r *MySQLAddressRepository) Eliminar(addressID int64) error {
	query := `DELETE FROM direcciones WHERE id = ?`
	_, err := r.db.Exec(query, addressID)
	return err
}
