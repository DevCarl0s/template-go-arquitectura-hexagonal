package constantes

var ID_POS int64 = 0

const (
	DB_CON      string = "postgres://pos_master:24002A8%C3%B1E5E588D09F41E3A230%C3%9101F7E8@localhost:8891/pos_transacciones?sslmode=allow"
	HOST_PORT   string = "10665"
	HOST_IP     string = "0.0.0.0"
	EFECTIVO    string = "EFECTIVO"
	COMBUSTIBLE string = "combustible"

	PATH_PRINCIPAL           string = "/api"
	PATH_SINCRONIZACION      string = "/sincronizacion"
	SINCRONIZACION_PRODUCTOS string = PATH_SINCRONIZACION + "/productos"

	//Cloud
	MS_CLOUD_SINCRONIZADOR string = ":10665/api/productos"

	SINCRONIZAR_PRODUCTOS string = "productos"
)
