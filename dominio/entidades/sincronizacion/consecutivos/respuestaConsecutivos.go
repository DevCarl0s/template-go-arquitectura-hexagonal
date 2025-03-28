package entidades_sincronizacion_consecutivos

type RespuestaConsecutivos struct {
	Datos   []Consecutivos `json:"datos"`
	Total   int            `json:"total"`
	Mensaje string         `json:"mensaje"`
}

type Consecutivos struct {
	ID                 int64  `json:"id"`
	EmpresasID         int64  `json:"empresas_id"`
	TipoDocumento      string `json:"tipo_documento"`
	Prefijo            string `json:"prefijo"`
	FechaInicio        string `json:"fecha_inicio"`
	FechaFin           string `json:"fecha_fin"`
	ConsecutivoInicial int64  `json:"consecutivo_inicial"`
	ConsecutivoActual  *int64 `json:"consecutivo_actual"`
	ConsecutivoFinal   *int64 `json:"consecutivo_final"`
	Resolucion         string `json:"resolucion"`
	Observaciones      string `json:"observaciones"`
	EquiposID          *int64 `json:"equipos_id"`
	CsAtributos        *JSON  `json:"cs_atributos"`
	Origen             string `json:"origen"`
	ClaveTecnica       string `json:"clave_tecnica"`
	EstadoResolucionID int16  `json:"estado_resolucion_id"`
}

type JSON map[string]interface{}
