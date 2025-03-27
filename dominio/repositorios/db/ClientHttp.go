package dominio_repositorios

type IClienteDB interface {
	Select(query string, argumentos []any) ([][]interface{}, error)
}
