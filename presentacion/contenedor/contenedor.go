package contenedor

import (
	"log"
	"ms-sincronizador-tienda/aplicacion/casosusos"
	casosusos_sincronizacion "ms-sincronizador-tienda/aplicacion/casosusos/sincronizacion"
	casosusos_sincronizacion_productos "ms-sincronizador-tienda/aplicacion/casosusos/sincronizacion/productos"
	"ms-sincronizador-tienda/aplicacion/servicios"
	servicios_sincronizacion "ms-sincronizador-tienda/aplicacion/servicios/sincronizacion"
	comunes_db_clientes "ms-sincronizador-tienda/comunes/dominio/adaptadores/clientes/db"
	comunes_http_clientes "ms-sincronizador-tienda/comunes/dominio/adaptadores/clientes/http"
	"ms-sincronizador-tienda/comunes/dominio/adaptadores/mapeadores"
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
var SincronizacionProductos servicios_sincronizacion.Productos
var GestorObservadores *dominio_notificacion.GestorObservadores

// casosusos
var GestionarNotificaciones *casosusos.GestionarNotificaciones

var RecuperarPeticionProductos *casosusos_sincronizacion_productos.RecuperarPeticion
var ConsultarProductos *casosusos_sincronizacion_productos.ConsultarProductos
var ProcesarInformacion *casosusos_sincronizacion.ProcesarInformacion

// repository
var NotificacionRepository dominio_repositorios.INotificacion
var ConsultarProductosRepository dominio_repositorios_http.IConsultarProductos
var RecuperarWacherRepository dominio_repositorios.IRecuperarWacher
var InformacionEdsRepository dominio_repositorios.IInformacionEds
var ProcesarInformacionRepository dominio_repositorios.IProcesarInformacion

// Clientes
var ClienteDB comunes_db_clientes.IClienteDB
var ClienteHttp comunes_http_clientes.IClienteHttp

// MAPPER
var MapPosDatos mapeadores.MapearDatosPos

func InicializarContenedor() error {
	canalNotificaciones := make(chan *entidades.Notificacion)

	// Inicializar el gestor de observadores
	GestorObservadores = dominio_notificacion.NuevoGestorObservadores()
	GestorObservadores.RegistrarObservador()

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

	//casosusos
	GestionarNotificaciones = &casosusos.GestionarNotificaciones{
		CanalEventos: canalNotificaciones,
	}
	RecuperarPeticionProductos = &casosusos_sincronizacion_productos.RecuperarPeticion{
		InformacionEds:  InformacionEdsRepository,
		WacherParametro: RecuperarWacherRepository,
	}
	ConsultarProductos = &casosusos_sincronizacion_productos.ConsultarProductos{
		Cliente: ConsultarProductosRepository,
	}
	ProcesarInformacion = &casosusos_sincronizacion.ProcesarInformacion{
		Procesar: ProcesarInformacionRepository,
	}

	//Servicios
	SupervisarNotificaciones = servicios.SupervisarNotificaciones{
		Notificaciones:      NotificacionRepository,
		CanalNotificaciones: canalNotificaciones,
	}
	SincronizacionProductos = servicios_sincronizacion.Productos{
		RecuperarPeticionCloud: RecuperarPeticionProductos,
		ConsultarProductos:     ConsultarProductos,
		ProcesarInformacion:    ProcesarInformacion,
	}

	return nil
}
