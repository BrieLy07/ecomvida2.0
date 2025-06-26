package strategies

import (
	"role-permission-service/config"
	"role-permission-service/models"
)

type RolEstrategia struct{}

func (e *RolEstrategia) Ejecutar(data interface{}) (interface{}, error) {
	rol, ok := data.(models.Rol)
	if !ok {
		return nil, ErrTipoInvalido
	}

	db := config.ConectarDB()
	query := `INSERT INTO roles (nombre) VALUES ($1) RETURNING id`
	err := db.QueryRow(query, rol.Nombre).Scan(&rol.ID)
	if err != nil {
		return nil, err
	}

	return rol, nil
}
