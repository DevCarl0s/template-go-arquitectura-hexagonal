-- DROP PROCEDURE sincronizacion.prc_procesar_sincronizacion_tienda(in json, inout json);
CREATE OR REPLACE PROCEDURE sincronizacion.prc_procesar_sincronizacion_tienda(IN i_json_datos json, INOUT o_json_respuesta json)
 LANGUAGE plpgsql
AS $procedure$
	 -- !! ------------------------------------------------------------------------ !! --
     -- !! -- AUTOR:                       			CARLOS DE LAS SALAS   		   	!! --
     -- !! -- FECHA CREACION:               		19/03/2025            			!! --
     -- !! -- AUTOR ULT ACTUALIZACION:    				   		   					!! --
     -- !! -- FECHA ULTIMA ACTUALIZACION:  			        			   			!! --
     -- !! ---------------------------------------------------------------  		!! --
DECLARE 
	-- VARIABLES --
	v_status	integer;
	v_texto_log  text;
   	v_state text;
   	v_msg   text;
   	v_detail text;
   	v_hint  text;
   	v_contex  text;
   	v_error text;
    v_texto_respuesta  text;
   	v_codigo_respuesta integer 	:=-1;
    v_nombre_up  varchar(100) 	:= 'prc_procesar_sincronizacion_tienda';
    v_id_movimiento integer 	:=-1;
    v_fecha_hora_inicio  timestamp;
	v_ano  int8;
	v_mes int8;
	v_dia 	int8;
	property_name text;
    property_data json;
    element_data jsonb;
    v_id_logs bigint := -1;
    v_json_respuesta_tanque json := '{}' ;

    tipo_dato text;
        
BEGIN
        v_fecha_hora_inicio := clock_timestamp();
        v_ano   := COALESCE(to_char(v_fecha_hora_inicio,'YYYY'), '0' );
        v_mes   := COALESCE(to_char(v_fecha_hora_inicio,'MM') , '0');
        v_dia   := COALESCE(to_char(v_fecha_hora_inicio,'DD'),'0');  

        
        IF (  i_json_datos IS NOT NULL AND  (i_json_datos::text <> '[]'::text  AND  i_json_datos::text <> '{}'::text)  ) THEN
            FOR property_name IN SELECT json_object_keys(i_json_datos)
            LOOP
        
                IF v_id_logs = -1 THEN
                    CALL sincronizacion.prc_sincronizar_logs(
                        i_json_datos => i_json_datos, 
                        i_nombre_up => v_nombre_up, 
                        i_id_log => v_id_logs::bigint, 
                        i_estado => 'S', 
                        i_tabla => property_name, 
                        i_error => 'NA', 
                        o_json_respuesta => v_json_respuesta_logs::json
                    );
                    v_id_logs := (v_json_respuesta_logs ->> 'id_logs')::bigint;
                END IF;					   
                    
                    property_data := (i_json_datos->property_name);	
                    
                        FOR element_data IN SELECT * FROM json_array_elements(property_data)
                        LOOP
                            BEGIN
                                        --RAISE NOTICE 'element_data antes: %', element_data;
                                        tipo_dato := pg_typeof(element_data);
                                        --RAISE NOTICE 'pg_typeof(element_data): %', tipo_dato;
                                        element_data = jsonb_set(
                                                            element_data, 
                                                            '{id_logs}', 
                                                            to_jsonb(v_id_logs)::jsonb, 
                                                            true 
                                                        );
                                        --RAISE NOTICE 'element_data: %', element_data;		
                                        CASE
                                            WHEN property_name = 'surtidores' THEN
                
                                                    CALL sincronizacion.prc_sincronizar_surtidores(
                                                        i_json_datos => element_data::json, 
                                                        o_json_respuesta => v_json_respuesta_surtidores::json
                                                    );		
                                                    
                                            WHEN property_name = 'tanques' THEN
                
                                                    CALL sincronizacion.prc_sincronizar_tanque(
                                                        i_json_datos => element_data::json, 
                                                        o_json_respuesta => v_json_respuesta_tanque::json
                                                    );		
                                                
                                            WHEN property_name = 'unidades' THEN
                                                
                                                    CALL sincronizacion.prc_sincronizar_unidades(
                                                        i_json_datos => element_data::json, 
                                                        o_json_respuesta => v_json_respuesta_unidades::json
                                                    );		
                                                
                                            WHEN property_name = 'protocolos' THEN
                                                
                                                        CALL sincronizacion.prc_sincronizar_protocolos(
                                                        i_json_datos => element_data::json, 
                                                        o_json_respuesta => v_json_respuesta_protocolos::json
                                                    );	
                                                
                                            WHEN property_name = 'tipos_surtidores' THEN
            
                                                    CALL sincronizacion.prc_sincronizar_tipos(
                                                        i_json_datos => element_data::json, 
                                                        o_json_respuesta => v_json_respuesta_tipos::json
                                                    );	
                        
                                            WHEN property_name = 'productos_tipos' THEN
                                                
                                                    CALL sincronizacion.prc_sincronizar_productos_tipos(
                                                        i_json_datos => element_data::json, 
                                                        o_json_respuesta => v_json_respuesta_productos_tipos::json
                                                    );	
                                            
                                            WHEN property_name = 'productos_familias' THEN
                                                    
                                                    CALL sincronizacion.prc_sincronizar_productos_familias(
                                                        i_json_datos => element_data::json, 
                                                        o_json_respuesta => v_json_respuesta_productos_familias::json
                                                    );	
                                                
                                            WHEN property_name = 'productos_unidades' THEN
                        
                                                    CALL sincronizacion.prc_sincronizar_productos_unidades(
                                                        i_json_datos => element_data::json, 
                                                        o_json_respuesta => v_json_respuesta_productos_unidades::json
                                                    );
                                            WHEN property_name = 'productos_combustible' THEN
                                                    
                                                    CALL sincronizacion.prc_sincronizar_productos_combustible(
                                                        i_json_datos => element_data::json, 
                                                        o_json_respuesta => v_json_respuesta_productos_combustible::json
                                                    );
                                                    
                                            WHEN property_name = 'ct_consecutivos' THEN
            
                                                CALL sincronizacion.prc_sincronizar_consecutivo(
                                                        i_json_datos => element_data::json, 
                                                        o_json_respuesta => v_json_respuesta_consecutivos::json
                                                    );	
                                            
                                            WHEN property_name = 'lt_horarios' THEN
                                                    
                                                    CALL sincronizacion.prc_sincronizar_lt_horarios(
                                                        i_json_datos => element_data::json, 
                                                        o_json_respuesta => v_json_respuesta_lt_horarios::json
                                                    );	
                                                
                                            WHEN property_name = 'ct_bodegas' THEN
                                                    
                                                    CALL sincronizacion.prc_sincronizar_lt_horarios(
                                                        i_json_datos => element_data::json, 
                                                        o_json_respuesta => v_json_respuesta_ct_bodegas::json
                                                    );	
                                                
                                            WHEN property_name = 'wacher_parametros' THEN
                                            
                                                    CALL sincronizacion.prc_sincronizar_wacher_parametros(
                                                        i_json_datos => element_data::json, 
                                                        o_json_respuesta => v_json_respuesta_wacher_parametros::json
                                                    );

                                            
                                            WHEN property_name = 'proveedores_tecnologicos' THEN
                                            
                                                    CALL sincronizacion.prc_sincronizar_proveedores_tecnologicos(
                                                        i_json_datos => element_data::json, 
                                                        o_json_respuesta => v_json_respuesta_proveedores::json
                                                    );	
                                                
                                            WHEN property_name = 'empresas' THEN
                                            
                                                    CALL sincronizacion.prc_sincronizar_empresas(
                                                        i_json_datos => element_data::json, 
                                                        o_json_respuesta => v_json_respuesta_empresas::json
                                                    );
                                                
                                            WHEN property_name = 'cambio_precio' THEN
                                            
                                                CALL sincronizacion.prc_sincronizar_cambios_precios(
                                                    i_json_datos => element_data::json,
                                                    o_json_respuesta => v_json_respuesta_cambio_precio::json
                                                );
                                            
                                            WHEN property_name = 'dispositivos' THEN
                                            
                                                CALL sincronizacion.prc_sincronizar_dispositivos(
                                                    i_json_datos => element_data::json,
                                                    o_json_respuesta => v_json_respuesta_dispositivos::json
                                                );
                                            
                                            WHEN property_name = 'medios_pagos' THEN
                                            
                                                CALL sincronizacion.prc_sincronizar_medios_pagos(
                                                    i_json_datos => element_data::json,
                                                    o_json_respuesta => v_json_respuesta_medios_pagos::json
                                                );
                                            
                                            WHEN property_name = 'inventario' THEN
                                            
                                                CALL sincronizacion.prc_sincronizar_inventary(
                                                    i_json_datos => element_data::json,
                                                    o_json_respuesta => v_json_respuesta_inventario::json
                                                );

                                            WHEN property_name = 'tipos_identificaciones' THEN
                                            
                                                CALL sincronizacion.prc_sincronizar_tipos_identificadores(
                                                    i_json_datos => element_data::json,
                                                    o_json_respuesta => v_json_respuesta_tipo_identificaciones::json
                                                );
                                            
                                            WHEN property_name = 'personas' THEN
                                            
                                                CALL sincronizacion.prc_sincronizar_personas(
                                                    i_json_datos => element_data::json,
                                                    o_json_respuesta => v_json_respuesta_personas::json
                                                );

                                            WHEN property_name = 'clientes_gnv' THEN
                                            
                                                CALL sincronizacion.prc_sincronizar_clientes_gnv(
                                                    i_json_datos => element_data::json,
                                                    o_json_respuesta => v_json_respuesta_clientes_gnv::json
                                                );
                                        
                                                
                                        END CASE;

                                EXCEPTION
                                            WHEN OTHERS THEN
                                                        GET  STACKED DIAGNOSTICS  
                                                        v_state   = RETURNED_SQLSTATE,
                                                            v_msg     = MESSAGE_TEXT,
                                                            v_detail  = PG_EXCEPTION_DETAIL,
                                                            v_hint    = PG_EXCEPTION_HINT;
                                                            v_error   := v_state || ' ' || v_msg || ' ' || v_detail || ' ' || v_hint;
                                                            v_codigo_respuesta        := '501';
                                                            v_texto_respuesta        := v_detail;
                                                            --raise notice 'Error fnc_procesar_insertar_productos_registry %', v_error;	 
                                                                                                
                                                        CALL sincronizacion.prc_sincronizar_logs(
                                                                    i_json_datos => element_data::json, 
                                                                    i_nombre_up => v_nombre_up, 
                                                                    i_id_log => v_id_logs::bigint, 
                                                                    i_estado => 'R', 
                                                                    i_tabla => property_name, 
                                                                    i_error => v_error, 
                                                                    o_json_respuesta => v_json_respuesta_logs::json
                                                            );   
                                                        
                                CONTINUE;	
                    END;       
                                        
                    END LOOP;
                    
                    IF property_name <> 'productos_combustible' THEN   
                        CALL sincronizacion.prc_pg_notify( i_tipo => property_name::text, i_estado_prc => 'E'::text, i_error_prc => ''::text );		
                    END IF;  
                
                    UPDATE sincronizacion.sincronizacion_hist
                        SET fecha_completada=v_fecha_hora_inicio, estado= 'E'
                    WHERE id = v_id_logs;
                    
                    v_id_logs := -1;
                
            END LOOP;

            o_json_respuesta := json_build_object(
                                                                'codigo_respuesta', '201',
                                                                'estado', 'Sincronizacion realizada con exito',
                                                                'json_respuesta', json_build_object(
                                                                                                                                'surtidores', v_json_respuesta_surtidores,
                                                                                                                                'tanque', v_json_respuesta_tanque,
                                                                                                                                'unidades', v_json_respuesta_unidades,
                                                                                                                                'protocolos',v_json_respuesta_protocolos,
                                                                                                                                'tipos', v_json_respuesta_tipos,
                                                                                                                                'productos_tipos',v_json_respuesta_productos_tipos,
                                                                                                                                'productos_familias', v_json_respuesta_productos_familias,
                                                                                                                                'ct_consecutivos', v_json_respuesta_consecutivos,
                                                                                                                                'lt_horarios', v_json_respuesta_lt_horarios,
                                                                                                                                'wacher_parametros',v_json_respuesta_wacher_parametros,
                                                                                                                                'empresas',v_json_respuesta_empresas,
                                                                                                                                'cambio_precio', v_json_respuesta_cambio_precio,
                                                                                                                                'productos_combustible', v_json_respuesta_productos_combustible,
                                                                                                                                'productos_unidades', v_json_respuesta_productos_unidades,
                                                                                                                                'dispositivos', v_json_respuesta_dispositivos,
                                                                                                                                'medios_pagos', v_json_respuesta_medios_pagos,
                                                                                                                                'inventario', v_json_respuesta_inventario,
                                                                                                                                'proveedores_tecnologicos', v_json_respuesta_proveedores,
                                                                                                                                'tipos_identificaciones', v_json_respuesta_tipo_identificaciones,
                                                                                                                                'personas', v_json_respuesta_personas,
                                                                                                                                'clientes_gnv', v_json_respuesta_clientes_gnv
                                                                                                                            )	
                                                            );	
        
        ELSE   
		   					
					UPDATE sincronizacion.sincronizacion_hist
						SET fecha_completada=v_fecha_hora_inicio, estado= 'R'
					WHERE id = v_id_logs;
		   
		   			 o_json_respuesta := json_build_object(
						  											'codigo_respuesta', '501',
						  											'estado', 'Error: no hay data para ser procesadas',
						  											'json_respuesta', i_json_datos
						  										);	
           END IF;						  													  									
			exception  
			    when others then
			         get  STACKED diagnostics  
			         v_state   = RETURNED_SQLSTATE,
			              v_msg     = MESSAGE_TEXT,
			              v_detail  = PG_EXCEPTION_DETAIL,
			              v_hint    = PG_EXCEPTION_HINT;
			              v_error   := v_state || ' ' || v_msg || ' ' || v_detail || ' ' || v_hint;
			              v_codigo_respuesta        := '501';
			              v_texto_respuesta        := v_detail;
   			              --raise notice 'Error prc_procesar_sincronizacion %', v_error;						
						  o_json_respuesta := json_build_object(
						  											'codigo_respuesta', v_codigo_respuesta,
						  											'estado', v_error,
						  											'json_data', i_json_datos
						  										);								
			UPDATE sincronizacion.sincronizacion_hist
				SET fecha_completada=v_fecha_hora_inicio, estado= 'R'
			WHERE id = v_id_logs;
END;
$procedure$
;


CREATE TABLE sincronizacion.sincronizacion_tienda (
	id serial4 NOT NULL,
	tipo_notificacion int8 NULL,
	"data" text NULL,
	prioridad bool DEFAULT false NULL,
	procesada bool DEFAULT false NULL,
	fecha_recibido timestamp NULL,
	fecha_completado timestamp NULL,
	CONSTRAINT sincronizacion_tienda_pkey PRIMARY KEY (id)
);


CREATE OR REPLACE FUNCTION sincronizacion.delete_sync_tienda_success()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$
BEGIN
    -- Verifica si la columna fecha_Completado se ha llenado
    IF NEW.fecha_Completado IS NOT NULL THEN
        -- Elimina el registro
        DELETE FROM sincronizacion.sincronizacion_tienda WHERE id = NEW.id;
    END IF;
    RETURN NEW;
END;
$function$
;


create trigger trigger_check_sync_tienda_success after
update
    on
    sincronizacion.sincronizacion_tienda for each row execute function sincronizacion.delete_sync_tienda_success();


CREATE TABLE sincronizacion.sincronizacion_tienda_hist (
	id serial4 NOT NULL,
	tipo_notificacion int4 NULL,
	tipo_sincronizacion text NULL,
	fecha_creacion timestamp DEFAULT CURRENT_TIMESTAMP NULL,
	fecha_completada timestamp NULL,
	data_log text NULL,
	responsable varchar NULL,
	estado text NULL,
	sincronizado bool NULL,
	CONSTRAINT logs_id_logs_pk PRIMARY KEY (id)
);
CREATE INDEX sincronizacion_tienda_hist_tipo_notificacion_idx ON sincronizacion.sincronizacion_tienda_hist USING btree (tipo_notificacion, tipo_sincronizacion, estado);

