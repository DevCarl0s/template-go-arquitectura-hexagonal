package constantes

var ID_POS int64 = 0

const (
	DB_CON      string = "postgres://pos_master:24002A8%C3%B1E5E588D09F41E3A230%C3%9101F7E8@localhost:5432/pos_transacciones?sslmode=allow"
	HOST_PORT   string = "10665"
	HOST_IP     string = "0.0.0.0"
	EFECTIVO    string = "EFECTIVO"
	COMBUSTIBLE string = "combustible"

	PATH_PRINCIPAL           string = "/api"
	PATH_SINCRONIZACION      string = "/sincronizacion"
	SINCRONIZACION_PRODUCTOS string = PATH_SINCRONIZACION + "/productos"

	//Cloud
	MS_CLOUD_SINCRONIZADOR string = ":10665/api"
	API_PRODUCTOS_TIENDA   string = "/tienda/productos/eds"
	API_CONSECUTIVO_TIENDA string = "/tienda/consecutivos"

	//TIPOS SINCRONIZACION
	PROCESAR_PRODUCTO                string = "producto"
	PROCESAR_SUBCATEGORIA            string = "subcategoria_producto"
	PROCESAR_CATEGORIA               string = "categoria"
	PROCESAR_CATEGORIA_TIPO_NEGOCIO  string = "categoria_tipo_negocio"
	PROCESAR_UNIDAD_MEDIDA           string = "unidad_medida"
	PROCESAR_NEGOCIO                 string = "negocio"
	PROCESAR_TIPO_PRODUCTO           string = "tipo_producto"
	PROCESAR_CODIGOBARRAS            string = "codigos_barras"
	PROCESAR_RESOLUCIONES            string = "resoluciones"
	PROCESAR_CATEGORIAS              string = "categorias"
	PROCESAR_IMPUESTOS               string = "impuestos"
	PROCESAR_IMPUESTOS_CLASIFICACION string = "impuestos_clasificacion"
	PROCESAR_PRODUCTOS_IMPUESTOS     string = "productos_impuestos"
	PROCESAR_IMPUESTOS_TIPOS         string = "impuestos_tipos"
)
