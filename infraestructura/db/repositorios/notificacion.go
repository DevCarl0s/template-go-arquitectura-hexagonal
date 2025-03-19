package repositorios_infraestruture

import (
	comunes_db_clientes "ms-sincronizador-tienda/comunes/dominio/adaptadores/clientes/db"
	"ms-sincronizador-tienda/dominio/entidades"
)

type Notificacion struct {
	Cliente comunes_db_clientes.IClienteDB
}

func (r *Notificacion) ObtenerPendientes() ([]*entidades.Notificacion, error) {
	query := `
		SELECT id, tipo_notificacion, data, prioridad, procesada, fecha_recibido, fecha_completado
		FROM sincronizacion.sincronizacion_frontal
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
			valor[0].(int),
			valor[1].(int),
			valor[2].(string),
			valor[3].(bool),
			valor[4].(bool),
			valor[5].(string),
			valor[6].(string),
		),
		)
	}

	return notificaciones, nil
}

func (r *Notificacion) MarcarProcesada(id int) error {
	query := `
		UPDATE sincronizacion.sincronizacion_frontal 
		SET procesada = true
		WHERE id = ?
	`

	_, err := r.Cliente.Select(query, []any{id})
	return err
}
