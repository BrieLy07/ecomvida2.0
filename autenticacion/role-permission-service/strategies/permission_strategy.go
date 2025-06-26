package strategies

import (
	"role-permission-service/config"
	"role-permission-service/models"
)

type PermisoEstrategia struct{}

func (e *PermisoEstrategia) Ejecutar(data interface{}) (interface{}, error) {
	permiso, ok := data.(models.Permiso)
	if !ok {
		return nil, ErrTipoInvalido
	}

	db := config.ConectarDB()
	query := `INSERT INTO permisos (nombre) VALUES ($1) RETURNING id`
	err := db.QueryRow(query, permiso.Nombre).Scan(&permiso.ID)
	if err != nil {
		return nil, err
	}

	return permiso, nil
}
