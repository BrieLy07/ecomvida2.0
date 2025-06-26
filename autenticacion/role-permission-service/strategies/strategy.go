package strategies

type Estrategia interface {
	Ejecutar(data interface{}) (interface{}, error)
}
