package servicios_observadores_productos

import (
	"fmt"
	"log"
	casosusos_sincronizacion "ms-sincronizador-tienda/aplicacion/casosusos/sincronizacion"
	casosusos_sincronizacion_productos "ms-sincronizador-tienda/aplicacion/casosusos/sincronizacion/productos"
	"ms-sincronizador-tienda/dominio/constantes"
	dominio_notificacion "ms-sincronizador-tienda/dominio/notificacion"
)

type ObservadorSincronizarProductosTienda struct {
	RecuperarPeticionCloud *casosusos_sincronizacion_productos.RecuperarPeticion
	ConsultarProductos     *casosusos_sincronizacion_productos.ConsultarProductos
	ProcesarInformacion    *casosusos_sincronizacion.ProcesarInformacion
}

func (OP *ObservadorSincronizarProductosTienda) ProcesarNotificacion(notificacion dominio_notificacion.Notificacion) error {
	fmt.Printf("Procesando notificación de productos: %+v\n", notificacion)
	peticion, err := OP.RecuperarPeticionCloud.Ejecutar()
	if err != nil {
		return err
	}
	log.Println("peticion: ", peticion)

	productos, err := OP.ConsultarProductos.Ejecutar(peticion)
	if err != nil {
		log.Println("No se pudieron obtener los productos Cloud")
		return err
	}

	if len(productos.Datos) > 0 {
		for _, dato := range productos.Datos {
			err = OP.ProcesarInformacion.Ejecutar(dato.SubcategoriaProducto, constantes.PROCESAR_SUBCATEGORIA)
			if err != nil {
				return err
			}
			err = OP.ProcesarInformacion.Ejecutar(dato.UnidadMedidaCompra, constantes.PROCESAR_UNIDAD_MEDIDA)
			if err != nil {
				return err
			}
			err = OP.ProcesarInformacion.Ejecutar(dato.UnidadMedidaVenta, constantes.PROCESAR_UNIDAD_MEDIDA)
			if err != nil {
				return err
			}
			err = OP.ProcesarInformacion.Ejecutar(dato.Negocio, constantes.PROCESAR_NEGOCIO)
			if err != nil {
				return err
			}
			err = OP.ProcesarInformacion.Ejecutar(dato.TipoProducto, constantes.PROCESAR_TIPO_PRODUCTO)
			if err != nil {
				return err
			}
			err = OP.ProcesarInformacion.Ejecutar(dato.Producto, constantes.PROCESAR_PRODUCTO)
			if err != nil {
				return err
			}
			err = OP.ProcesarInformacion.Ejecutar(dato.CodigosBarras, constantes.PROCESAR_CODIGOBARRAS)
			if err != nil {
				return err
			}
			err = OP.ProcesarInformacion.Ejecutar(dato.ImpuestosDetalles, constantes.PROCESAR_IMPUESTOS)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (o *ObservadorSincronizarProductosTienda) ObtenerTipo() dominio_notificacion.TipoNotificacion {
	return dominio_notificacion.TIPO_PRODUCTOS
}
