package entidades_sincronizacion_productos

type Producto struct {
	ID                     int     `json:"id"`
	Descripcion            string  `json:"descripcion"`
	Estado                 string  `json:"estado"`
	Atributos              any     `json:"atributos"`
	Precio                 float64 `json:"precio"`
	PrecioMinimo           float64 `json:"precio_minimo"`
	PrecioMaximo           float64 `json:"precio_maximo"`
	ImpuestoOperacion      string  `json:"impuesto_operacion"`
	FechaCreacion          string  `json:"fecha_creacion"`
	FechaModificacion      string  `json:"fecha_modificacion"`
	UsuarioCreacion        string  `json:"usuario_creacion"`
	UsuarioModificacion    string  `json:"usuario_modificacion"`
	SKU                    string  `json:"sku"`
	TipoProductoID         int     `json:"tipo_producto_id"`
	SubcategoriaProductoID int     `json:"subcategoria_producto_id"`
	TipoNegocioID          int     `json:"id_tipo_negocio"`
	EmpresaID              int     `json:"empresa_id"`
}

type Empresa struct {
	EmpresasID int    `json:"empresas_id"`
	Nombre     string `json:"nombre"`
}

type Negocio struct {
	IDTipoNegocio int    `json:"id_tipo_negocio"`
	Descripcion   string `json:"descripcion"`
	Valor         string `json:"valor"`
	Estado        bool   `json:"estado"`
}

type UnidadMedida struct {
	ID          int     `json:"id"`
	Descripcion string  `json:"descripcion"`
	Valor       float64 `json:"valor"`
	Estado      string  `json:"estado"`
	Atributos   any     `json:"u_atributos"`
	EmpresasID  int     `json:"empresas_id"`
	Alias       string  `json:"alias"`
}

type SubcategoriaProducto struct {
	ID             int    `json:"subcategoria_producto_id"`
	Descripcion    string `json:"subcategoria_producto_descripcion"`
	Estado         int    `json:"subcategoria_producto_estado"`
	TiempoCreacion string `json:"subcategoria_producto_tiempo_creacion"`
	CategoriaID    int    `json:"subcategoria_producto_categoria_id"`
}

type CategoriaProducto struct {
	Id          int    `json:"id"`
	Descripcion string `json:"descripcion"`
}

type CodigosBarras struct {
	CodigoBarrasID int    `json:"codigo_barras_id"`
	ProductoID     int    `json:"producto_id"`
	Codigo         string `json:"codigo"`
}

type ImpuestosDetalles struct {
	ProductosImpuestoId              int     `json:"productos_impuestos_id"`
	ProductosImpuestoTipo            int     `json:"productos_impuestos_tipo"`
	ProductoImpuestoProductoId       int     `json:"productos_impuestos_producto_id"`
	ImpuestoId                       int     `json:"impuestos_id"`
	ImpuestoDescripcion              string  `json:"impuestos_descripcion"`
	ImpuestoPorcentajeValor          string  `json:"impuestos_porcentaje_valor"`
	ImpuestoValor                    float64 `json:"impuestos_valor"`
	ImpuestoEstado                   string  `json:"impuestos_estado"`
	ImpuestoEmpresasId               int     `json:"impuestos_empresas_id"`
	TipoImpuestoId                   int     `json:"tipo_impuesto_id"`
	TipoImpuestoDescripcion          string  `json:"tipo_impuesto_descripcion"`
	TipoImpuestoEstado               int     `json:"tipo_impuesto_estado"`
	ClasificacionImpuestoId          int     `json:"clasificacion_impuesto_id"`
	ClasificacionImpuestoDescripcion string  `json:"clasificacion_impuesto_descripcion"`
	ClasificacionImpuestoEstado      int     `json:"clasificacion_impuesto_estado"`
}

type TipoProducto struct {
	Id          int    `json:"id"`
	Descripcion string `json:"descripcion"`
	Estado      string `json:"estado"`
}

type CategoriaTipoNegocio struct {
	Id            int `json:"id"`
	CategoriaId   int `json:"categoria_id"`
	TipoNegocioId int `json:"tipo_negocio_id"`
}

type ProductosEds struct {
	Producto             Producto             `json:"producto"`
	Empresa              Empresa              `json:"empresa"`
	Negocio              Negocio              `json:"negocio"`
	UnidadMedidaCompra   UnidadMedida         `json:"unidad_medida_compra"`
	UnidadMedidaVenta    UnidadMedida         `json:"unidad_medida_venta"`
	SubcategoriaProducto SubcategoriaProducto `json:"subcategoria_producto"`
	CategoriaProducto    CategoriaProducto    `json:"categoria_producto"`
	CodigosBarras        []CodigosBarras      `json:"codigos_barras"`
	ImpuestosDetalles    []ImpuestosDetalles  `json:"impuestos_detalles"`
	TipoProducto         []TipoProducto       `json:"tipo_producto"`
	CategoriaTipoNegocio CategoriaTipoNegocio `json:"categoria_tipo_negocio"`
}

type RespuestaProductos struct {
	Datos   []ProductosEds `json:"datos"`
	Total   int            `json:"total"`
	Mensaje string         `json:"mensaje"`
}
