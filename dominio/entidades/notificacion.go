package entidades

type Notificacion struct {
	ID               int    `json:"id"`
	TipoNotificacion int    `json:"tipo_notificacion"`
	Data             string `json:"data"`
	Prioridad        bool   `json:"prioridad"`
	Procesada        bool   `json:"procesada"`
	FechaRecibido    string `json:"fecha_recibido"`
	FechaCompletado  string `json:"fecha_completado"`
}

func NewNotificacion(id int, tipoNotificacion int, data string, prioridad bool, procesada bool, fechaRecibido string, fechaCompletado string) *Notificacion {
	return &Notificacion{
		ID:               id,
		TipoNotificacion: tipoNotificacion,
		Data:             data,
		Prioridad:        prioridad,
		Procesada:        procesada,
		FechaRecibido:    fechaRecibido,
		FechaCompletado:  fechaCompletado,
	}
}
