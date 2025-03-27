package servicios_observadores_productos

import (
	"fmt"
	"log"
	casosusos_sincronizacion "ms-sincronizador-tienda/aplicacion/casosusos/sincronizacion"
	casosusos_sincronizacion_productos "ms-sincronizador-tienda/aplicacion/casosusos/sincronizacion/productos"
	"ms-sincronizador-tienda/dominio/constantes"
	entidades_sincronizacion_productos "ms-sincronizador-tienda/dominio/entidades/sincronizacion/productos"
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

	// _, err := OP.ConsultarProductos.Ejecutar(peticion)
	// if err != nil {
	// 	log.Println("No se pudieron obtener los productos Cloud")
	// 	return err
	// }
	tiposProductos := []entidades_sincronizacion_productos.TipoProducto{}
	tipoProducto := entidades_sincronizacion_productos.TipoProducto{
		Id:          1,
		Estado:      "A",
		Descripcion: "PARAMETROS",
	}
	tiposProductos = append(tiposProductos, tipoProducto)

	codigosBarras := []entidades_sincronizacion_productos.CodigosBarras{}
	codigoBarra := entidades_sincronizacion_productos.CodigosBarras{
		Codigo:         "ASDASD",
		ProductoID:     27224,
		CodigoBarrasID: 100,
	}
	codigosBarras = append(codigosBarras, codigoBarra)

	impuestosDetalles := []entidades_sincronizacion_productos.ImpuestosDetalles{}
	impuestoDetalle := entidades_sincronizacion_productos.ImpuestosDetalles{
		ImpuestosID:              4,
		Valor:                    19.000,
		TipoImpuestoID:           1,
		EstadoID:                 1,
		Descripcion:              "IVA 19%",
		ProductosImpuestosID:     1,
		Tipo:                     1,
		ClasificacionID:          1,
		TipoImpuestoDescripcion:  "PORCENTAJES",
		PorcentajeValor:          "%",
		ClasificacionEstadoID:    1,
		ProductoID:               27224,
		ClasificacionDescripcion: "GRAVADO",
	}
	impuestosDetalles = append(impuestosDetalles, impuestoDetalle)

	unidadesMedidasVentas := entidades_sincronizacion_productos.UnidadMedida{
		ID:          1,
		Alias:       "UN",
		Valor:       1,
		Estado:      "A",
		Descripcion: "UNIDAD",
		EmpresasID:  4,
		Atributos: map[string]interface{}{
			"Alias": "UN",
		},
	}
	unidadesMedidasCompras := entidades_sincronizacion_productos.UnidadMedida{
		ID:          1,
		Alias:       "UN",
		Valor:       1,
		Estado:      "A",
		Descripcion: "UNIDAD",
		EmpresasID:  4,
		Atributos: map[string]interface{}{
			"Alias": "UN",
		},
	}

	SubcategoriaProducto := entidades_sincronizacion_productos.SubcategoriaProducto{
		ID:             3,
		Estado:         1,
		Descripcion:    "BEBIDAS FRIAS",
		CategoriaID:    3,
		TiempoCreacion: "2025-03-20T11:51:05.693218",
	}

	productosFake := []entidades_sincronizacion_productos.ProductosEds{}
	prducto := &entidades_sincronizacion_productos.ProductosEds{
		Producto: entidades_sincronizacion_productos.Producto{
			ID:                     27224,
			SKU:                    "12312312",
			Estado:                 "A",
			Precio:                 12.000,
			Atributos:              map[string]interface{}{},
			EmpresaID:              914,
			Descripcion:            "TIENDA PRB",
			PrecioMaximo:           2222.000,
			PrecioMinimo:           123.000,
			TipoNegocioID:          12,
			TipoProductoID:         1,
			ImpuestoOperacion:      "1",
			SubcategoriaProductoID: 3,
		},
		Negocio: entidades_sincronizacion_productos.Negocio{
			Valor:         "T",
			Estado:        true,
			Descripcion:   "TIENDA",
			IDTipoNegocio: 12,
		},
		TipoProducto:         tiposProductos,
		CodigosBarras:        codigosBarras,
		ImpuestosDetalles:    impuestosDetalles,
		UnidadMedidaCompra:   unidadesMedidasCompras,
		UnidadMedidaVenta:    unidadesMedidasVentas,
		SubcategoriaProducto: SubcategoriaProducto,
	}
	productosFake = append(productosFake, *prducto)

	if len(productosFake) > 0 {
		for _, dato := range productosFake {
			OP.ProcesarInformacion.Ejecutar(dato.SubcategoriaProducto, constantes.PROCESAR_SUBCATEGORIA)
			OP.ProcesarInformacion.Ejecutar(dato.UnidadMedidaCompra, constantes.PROCESAR_UNIDAD_MEDIDA)
			OP.ProcesarInformacion.Ejecutar(dato.UnidadMedidaVenta, constantes.PROCESAR_UNIDAD_MEDIDA)
			OP.ProcesarInformacion.Ejecutar(dato.Negocio, constantes.PROCESAR_NEGOCIO)
			OP.ProcesarInformacion.Ejecutar(dato.TipoProducto, constantes.PROCESAR_TIPO_PRODUCTO)
			OP.ProcesarInformacion.Ejecutar(dato.Producto, constantes.PROCESAR_PRODUCTO)
			OP.ProcesarInformacion.Ejecutar(dato.CodigosBarras, constantes.PROCESAR_CODIGOBARRAS)
			OP.ProcesarInformacion.Ejecutar(dato.ImpuestosDetalles, constantes.PROCESAR_IMPUESTOS)
		}
	}

	return nil
}

func (o *ObservadorSincronizarProductosTienda) ObtenerTipo() dominio_notificacion.TipoNotificacion {
	return dominio_notificacion.TIPO_PRODUCTOS
}
