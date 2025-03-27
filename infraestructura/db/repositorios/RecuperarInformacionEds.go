package repositorios_infraestruture

import (
	"ms-sincronizador-tienda/dominio/entidades"
	dominio_repositorios "ms-sincronizador-tienda/dominio/repositorios/db"
)

type RecuperarInformacionEds struct {
	Cliente dominio_repositorios.IClienteDB
}

func (RIE *RecuperarInformacionEds) Ejecutar() (*entidades.ParametrosEds, error) {

	respuesta, err := RIE.Cliente.Select(`
		select e.id , e2.id
		from empresas e
		inner join equipos e2 on e2.empresas_id = e.id
		limit 1`, []any{})

	if err != nil {
		return nil, err
	}

	parametro := &entidades.ParametrosEds{}
	for _, valor := range respuesta {
		parametro.EdsId = valor[0].(int64)
		parametro.EquipoId = valor[1].(int64)
	}

	return parametro, nil
}
