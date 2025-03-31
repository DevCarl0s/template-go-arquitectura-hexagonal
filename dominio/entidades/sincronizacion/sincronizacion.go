package entidades_sincronizacion

type DataSync struct {
	CodigoRespuesta interface{} `json:"codigo_respuesta"`
	Estado          interface{} `json:"estado"`
}
type JsonRespuesta struct {
	Productos              DataSync `json:"productos"`
	UnidadMedida           DataSync `json:"unidad_medida"`
	TipoProducto           DataSync `json:"tipo_producto"`
	SubcategoriaProducto   DataSync `json:"subcategoria_producto"`
	CodigosBarras          DataSync `json:"codigos_barras"`
	Resoluciones           DataSync `json:"resoluciones"`
	CategoriaTipoNegocio   DataSync `json:"categoria_tipo_negocio"`
	Categoria              DataSync `json:"categoria"`
	Impuestos              DataSync `json:"impuestos"`
	ImpuestosClasificacion DataSync `json:"impuestos_clasificacion"`
	ProductosImpuestos     DataSync `json:"productos_impuestos"`
	ImpuestosTipos         DataSync `json:"impuestos_tipos"`
}

type RespuestaSincronizacion struct {
	CodigoRespuesta int           `json:"codigo_respuesta"`
	Mensaje         interface{}   `json:"estado"`
	JsonRespuesta   JsonRespuesta `json:"json_respuesta"`
}
