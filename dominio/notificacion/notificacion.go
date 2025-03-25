package dominio_notificacion

type TipoNotificacion int

const (
	TIPO_PRODUCTOS TipoNotificacion = 1
	TipoOtro       TipoNotificacion = 8
)

type Notificacion struct {
	ID   int32
	Tipo TipoNotificacion
	Data interface{}
}

type ObservadorNotificacion interface {
	ProcesarNotificacion(notificacion Notificacion) error
	ObtenerTipo() TipoNotificacion
}

type GestorObservadores struct {
	observadores map[TipoNotificacion][]ObservadorNotificacion
}

func NuevoGestorObservadores() *GestorObservadores {
	return &GestorObservadores{
		observadores: make(map[TipoNotificacion][]ObservadorNotificacion),
	}
}

func (g *GestorObservadores) RegistrarObservador(observador ObservadorNotificacion) {
	tipo := observador.ObtenerTipo()
	g.observadores[tipo] = append(g.observadores[tipo], observador)
}

func (g *GestorObservadores) NotificarObservadores(notificacion Notificacion) error {
	observadores := g.observadores[notificacion.Tipo]
	for _, observador := range observadores {
		if err := observador.ProcesarNotificacion(notificacion); err != nil {
			return err
		}
	}
	return nil
}
