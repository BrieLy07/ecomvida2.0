package models

type Address struct {
	ID           int64  `json:"id"`
	UsuarioID    int64  `json:"usuario_id"`
	Calle        string `json:"calle"`
	Ciudad       string `json:"ciudad"`
	Provincia    string `json:"provincia"`
	CodigoPostal string `json:"codigo_postal"`
	Pais         string `json:"pais"`
}
