package repositorios_infraestruture

import (
	comunes_db_clientes "ms-sincronizador-tienda/comunes/dominio/adaptadores/clientes/db"
	"ms-sincronizador-tienda/dominio/entidades"
)

type RecuperarInformacionEds struct {
	Cliente comunes_db_clientes.IClienteDB
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
