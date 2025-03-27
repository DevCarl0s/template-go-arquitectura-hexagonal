package repositorios_infraestruture

import (
	"ms-sincronizador-tienda/dominio/entidades"
	dominio_repositorios "ms-sincronizador-tienda/dominio/repositorios/db"
	"ms-sincronizador-tienda/infraestructura/db/repositorios/mappers"
)

type Notificacion struct {
	Cliente dominio_repositorios.IClienteDB
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

func (r *Notificacion) MarcarProcesada(id int32) error {
	query := `
		UPDATE sincronizacion.sincronizacion_tienda 
		SET procesada = true , fecha_completado = NOW()
		WHERE id = $1
	`

	_, err := r.Cliente.Select(query, []any{int(id)})
	return err
}
