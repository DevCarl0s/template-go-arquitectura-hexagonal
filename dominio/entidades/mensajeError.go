package entidades

import "time"

type MensajeError struct {
	Mensaje      string    `json:"mensaje"`
	Data         any       `json:"data"`
	FechaProceso time.Time `json:"fechaProceso"`
}

func GetMensajeError(mensaje string, data any) *MensajeError {
	return &MensajeError{
		Mensaje:      mensaje,
		Data:         data,
		FechaProceso: time.Now(),
	}
}
