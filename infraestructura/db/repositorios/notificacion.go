package repositorios_infraestruture

import (
	comunes_db_clientes "ms-sincronizador-tienda/comunes/dominio/adaptadores/clientes/db"
	"ms-sincronizador-tienda/dominio/entidades"
	"ms-sincronizador-tienda/infraestructura/db/repositorios/mappers"
)

type Notificacion struct {
	Cliente comunes_db_clientes.IClienteDB
}

func (r *Notificacion) ObtenerPendientes() ([]*entidades.Notificacion, error) {
	query := `
		SELECT id, tipo_notificacion, data, prioridad, procesada, fecha_recibido, fecha_completado
		FROM sincronizacion.sincronizacion_tienda
		WHERE procesada = false AND fecha_completado IS NULL
		ORDER BY fecha_recibido ASC
	`

	resultado, err := r.Cliente.Select(query, []any{})
	if err != nil {
		return nil, err
	}

	var notificaciones []*entidades.Notificacion
	for _, valor := range resultado {
		notificaciones = append(notificaciones, entidades.NewNotificacion(
			valor[0].(int32),
			valor[1].(int64),
			mappers.GetStringPointer(valor[2]),
			valor[3].(bool),
			valor[4].(bool),
			mappers.GetStringPointer(valor[5]),
			mappers.GetStringPointer(valor[6]),
		),
		)
	}

	return notificaciones, nil
}

func (r *Notificacion) MarcarProcesada(id int) error {
	query := `
		UPDATE sincronizacion.sincronizacion_tienda 
		SET procesada = true
		WHERE id = ?
	`

	_, err := r.Cliente.Select(query, []any{id})
	return err
}
