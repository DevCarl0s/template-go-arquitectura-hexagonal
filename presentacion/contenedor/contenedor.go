package contenedor

import (
	"log"
	"ms-sincronizador-tienda/aplicacion/casosusos"
	casosusos_sincronizacion "ms-sincronizador-tienda/aplicacion/casosusos/sincronizacion"
	casosusos_sincronizacion_consecutivos "ms-sincronizador-tienda/aplicacion/casosusos/sincronizacion/consecutivos"
	casosusos_sincronizacion_productos "ms-sincronizador-tienda/aplicacion/casosusos/sincronizacion/productos"
	"ms-sincronizador-tienda/aplicacion/servicios"
	servicios_observadores_consecutivos "ms-sincronizador-tienda/aplicacion/servicios/observadores/consecutivos"
	servicios_observadores_productos "ms-sincronizador-tienda/aplicacion/servicios/observadores/productos"
	"ms-sincronizador-tienda/dominio/constantes"
	"ms-sincronizador-tienda/dominio/entidades"
	dominio_notificacion "ms-sincronizador-tienda/dominio/notificacion"
	dominio_repositorios "ms-sincronizador-tienda/dominio/repositorios/db"
	dominio_repositorios_http "ms-sincronizador-tienda/dominio/repositorios/http"
	infraestructura_db_cliente "ms-sincronizador-tienda/infraestructura/db/cliente"
	repositorios_infraestruture "ms-sincronizador-tienda/infraestructura/db/repositorios"
	cliente_infrastruc "ms-sincronizador-tienda/infraestructura/http/cliente"
	repositorios_infraestruture_http "ms-sincronizador-tienda/infraestructura/http/repositorios"
)

// servicios
var SupervisarNotificaciones servicios.SupervisarNotificaciones
var GestorObservadores *dominio_notificacion.GestorObservadores

// casosusos
var GestionarNotificaciones *casosusos.GestionarNotificaciones
var ProcesarInformacionCasoUso *casosusos_sincronizacion.ProcesarInformacion

var RecuperarPeticionProductosCasoUso *casosusos_sincronizacion_productos.RecuperarPeticion
var ConsultarProductosCasoUso *casosusos_sincronizacion_productos.ConsultarProductos

var RecuperarPeticionConsecutivosCasoUso *casosusos_sincronizacion_consecutivos.RecuperarPeticion
var ConsultarConsecutivosCasoUso *casosusos_sincronizacion_consecutivos.ConsultarConsecutivos

// repository
var ProcesarInformacionRepository dominio_repositorios.IProcesarInformacion
var NotificacionRepository dominio_repositorios.INotificacion
var InformacionEdsRepository dominio_repositorios.IInformacionEds
var RecuperarWacherRepository dominio_repositorios.IRecuperarWacher

var ConsultarProductosRepository dominio_repositorios_http.IConsultarProductos
var ConsultarConsecutivosRepository dominio_repositorios_http.IConsultarConsecutivos

// Clientes
var ClienteDB dominio_repositorios.IClienteDB
var ClienteHttp dominio_repositorios_http.IClienteHttp

// MAPPER

func InicializarContenedor() error {
	canalNotificaciones := make(chan *entidades.Notificacion)

	// Inicializar el gestor de observadores
	GestorObservadores = dominio_notificacion.NuevoGestorObservadores()

	//CLIENTES
	ClienteDB, err := infraestructura_db_cliente.InicializarCliente(constantes.DB_CON)
	if err != nil {
		log.Fatal(err)
		return err
	}
	ClienteHttp, err = cliente_infrastruc.InicializarCliente()
	if err != nil {
		log.Fatal(err)
		return err
	}

	//repository
	NotificacionRepository = &repositorios_infraestruture.Notificacion{Cliente: ClienteDB}
	InformacionEdsRepository = &repositorios_infraestruture.RecuperarInformacionEds{Cliente: ClienteDB}
	RecuperarWacherRepository = &repositorios_infraestruture.RecuperarWatcherParametors{Cliente: ClienteDB}

	ProcesarInformacionRepository = &repositorios_infraestruture.ProcesarInformacion{Cliente: ClienteDB}

	ConsultarProductosRepository = &repositorios_infraestruture_http.ConsultarProductos{Client: ClienteHttp}
	ConsultarConsecutivosRepository = &repositorios_infraestruture_http.ConsultarConsecutivos{Client: ClienteHttp}

	//casosusos
	GestionarNotificaciones = &casosusos.GestionarNotificaciones{
		CanalEventos:       canalNotificaciones,
		GestorObservadores: GestorObservadores,
		Notificaciones:     NotificacionRepository,
	}
	ProcesarInformacionCasoUso = &casosusos_sincronizacion.ProcesarInformacion{
		Procesar: ProcesarInformacionRepository,
	}

	RecuperarPeticionProductosCasoUso = &casosusos_sincronizacion_productos.RecuperarPeticion{
		InformacionEds:  InformacionEdsRepository,
		WacherParametro: RecuperarWacherRepository,
	}
	ConsultarProductosCasoUso = &casosusos_sincronizacion_productos.ConsultarProductos{
		Cliente: ConsultarProductosRepository,
	}

	RecuperarPeticionConsecutivosCasoUso = &casosusos_sincronizacion_consecutivos.RecuperarPeticion{
		InformacionEds:  InformacionEdsRepository,
		WacherParametro: RecuperarWacherRepository,
	}
	ConsultarConsecutivosCasoUso = &casosusos_sincronizacion_consecutivos.ConsultarConsecutivos{
		Cliente: ConsultarConsecutivosRepository,
	}

	//Servicios
	SupervisarNotificaciones = servicios.SupervisarNotificaciones{
		Notificaciones:      NotificacionRepository,
		CanalNotificaciones: canalNotificaciones,
	}

	observadorProductos := &servicios_observadores_productos.ObservadorSincronizarProductosTienda{
		RecuperarPeticionCloud: RecuperarPeticionProductosCasoUso,
		ConsultarProductos:     ConsultarProductosCasoUso,
		ProcesarInformacion:    ProcesarInformacionCasoUso,
	}
	observadorConsecutivos := &servicios_observadores_consecutivos.ObservadorSincronizarConsecutivosTienda{
		RecuperarPeticionCloud: RecuperarPeticionConsecutivosCasoUso,
		ConsultarConsecutivos:  ConsultarConsecutivosCasoUso,
		ProcesarInformacion:    ProcesarInformacionCasoUso,
	}
	GestorObservadores.RegistrarObservador(observadorProductos)
	GestorObservadores.RegistrarObservador(observadorConsecutivos)

	return nil
}
