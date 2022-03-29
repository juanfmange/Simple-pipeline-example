package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

//Funcion de la conexion de las bases de datos a Go
func conexionBD() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	contraseña := "260697"
	nombre := "prueba" //cambiar por pipeline

	conexion, err := sql.Open(Driver, Usuario+":"+contraseña+"@tcp(127.0.1.1)/"+nombre)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

// Declaracion de las variables de plantillas en .tmpl
var plantillas = template.Must(template.ParseGlob("plantillas/*"))

// funcion main como el front end con plantillas, estado impreso en consola y puerto http en lsiten and serve
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/crear", crear)
	http.HandleFunc("/insertar", insertar)
	http.HandleFunc("/borrar", Borrar)
	http.HandleFunc("/editar", Editar)
	http.HandleFunc("/actualizar", Actualizar)
	fmt.Println("Servidor Corriendo....")
	http.ListenAndServe(":8080", nil)
}

//Funcion editar del CRUD
func Editar(w http.ResponseWriter, r *http.Request) {
	idMovimiento := r.URL.Query().Get("id")
	fmt.Println(idMovimiento)
	conexionEstablecida := conexionBD()
	registro, err := conexionEstablecida.Query("SELECT * FROM pipeline WHERE MovId=?;", idMovimiento)

	data := Datos{}
	for registro.Next() {
		var porcentajederentabilidad float64
		var movid, semana, volumenmensual, volumenanual, tarifadecompra, profit, tarifadeventa, ingresopotencialanual, ingresopotencialmensual, numeroderutasofertadas, movimientosrechazados int
		var vendedor, tipodecliente, tipodeempresa, fecha, nombredelcliente, tamañodecliente, bussinesconection, tipodeindustria, tipodemercancia, tipodeventa, servicio, contacto, correoelectronico, telefono, trafico, tipodeservicio, tipodeequipounidad, tipodemodalidad, ciudadestadopaisorigen, ciudadestadopaisdestino, licitacion, target, ultimostatus, nocotizacionenfowardcomentariosop, comentariodelstatus, comentariosdelrechazo string
		err = registro.Scan(&movid, &vendedor, &fecha, &semana, &nombredelcliente, &tipodecliente, &tamañodecliente, &bussinesconection, &tipodeempresa, &tipodeindustria, &tipodemercancia, &tipodeventa, &servicio, &contacto, &correoelectronico, &telefono, &trafico, &tipodeservicio, &tipodeequipounidad, &tipodemodalidad, &ciudadestadopaisorigen, &ciudadestadopaisdestino, &licitacion, &target, &volumenmensual, &volumenanual, &tarifadecompra, &profit, &tarifadeventa, &porcentajederentabilidad, &ingresopotencialmensual, &ingresopotencialanual, &ultimostatus, &numeroderutasofertadas, &nocotizacionenfowardcomentariosop, &comentariodelstatus, &movimientosrechazados, &comentariosdelrechazo)
		if err != nil {
			panic(err.Error())
		}

		data.MovId = movid
		data.Vendedor = vendedor
		data.Fecha = fecha
		data.Semana = semana
		data.NombreDelCliente = nombredelcliente
		data.TipoDeCliente = tipodecliente
		data.TamañoDeCliente = tamañodecliente
		data.BussinesConection = bussinesconection
		data.TipoDeEmpresa = tipodeempresa
		data.TipoDeIndustria = tipodeindustria
		data.TipoDeMercancia = tipodemercancia
		data.TipoDeVenta = tipodeventa
		data.Servicio = servicio
		data.Contacto = contacto
		data.CorreoElectronico = correoelectronico
		data.Telefono = telefono
		data.Trafico = trafico
		data.TipoDeServicio = tipodeservicio
		data.TipoDeEquipoUnidad = tipodeequipounidad
		data.TipoDeModalidad = tipodemodalidad
		data.CiudadEstadoPaisOrigen = ciudadestadopaisorigen
		data.CiudadEstadoPaisDestino = ciudadestadopaisdestino
		data.Licitacion = licitacion
		data.Target = target
		data.VolumenMensual = volumenmensual
		data.VolumenAnual = volumenanual
		data.TarifaDeCompra = tarifadecompra
		data.Profit = profit
		data.TarifaDeVenta = tarifadeventa
		data.PorcentajeDeRentabilidad = porcentajederentabilidad
		data.IngresoPotencialMensual = ingresopotencialmensual
		data.IngresoPotencialAnual = ingresopotencialanual
		data.UltimoStatus = ultimostatus
		data.NumeroDeRutasOfertadas = numeroderutasofertadas
		data.NoCotizacionEnFowardComentariosOP = nocotizacionenfowardcomentariosop
		data.ComentariosDelStatus = comentariodelstatus
		data.MovimientosRechazados = movimientosrechazados
		data.ComentariosDelRechazo = comentariosdelrechazo

	}
	fmt.Println(data)
	plantillas.ExecuteTemplate(w, "editar", data)
}

// Borrar funciona normal
func Borrar(w http.ResponseWriter, r *http.Request) {
	idMovimiento := r.URL.Query().Get("id")
	fmt.Println(idMovimiento)

	conexionEstablecida := conexionBD()

	borrarRegistros, err := conexionEstablecida.Prepare("DELETE FROM pipeline WHERE MovId=?")
	if err != nil {
		panic(err.Error())
	}
	borrarRegistros.Exec(idMovimiento)

	http.Redirect(w, r, "/", 301)

}

// Estructura de los datos
type Datos struct {
	MovId                             int
	Vendedor                          string
	Fecha                             string
	Semana                            int
	NombreDelCliente                  string
	TipoDeCliente                     string
	TamañoDeCliente                   string
	BussinesConection                 string
	TipoDeEmpresa                     string
	TipoDeIndustria                   string
	TipoDeMercancia                   string
	TipoDeVenta                       string
	Servicio                          string
	Contacto                          string
	CorreoElectronico                 string
	Telefono                          string
	Trafico                           string
	TipoDeServicio                    string
	TipoDeEquipoUnidad                string
	TipoDeModalidad                   string
	CiudadEstadoPaisOrigen            string
	CiudadEstadoPaisDestino           string
	Licitacion                        string
	Target                            string
	VolumenMensual                    int
	VolumenAnual                      int
	TarifaDeCompra                    int
	Profit                            int
	TarifaDeVenta                     int
	PorcentajeDeRentabilidad          float64
	IngresoPotencialMensual           int
	IngresoPotencialAnual             int
	UltimoStatus                      string
	NumeroDeRutasOfertadas            int
	NoCotizacionEnFowardComentariosOP string
	ComentariosDelStatus              string
	MovimientosRechazados             int
	ComentariosDelRechazo             string
}

// Pantalla principla, donde se listan los elementos
func index(w http.ResponseWriter, r *http.Request) {
	conexionEstablecida := conexionBD()

	registros, err := conexionEstablecida.Query("SELECT * FROM pipeline;")

	if err != nil {
		panic(err.Error())
	}

	movimientos := Datos{}
	arregloMovimientos := []Datos{}

	for registros.Next() {
		var porcentajederentabilidad float64
		var movid, semana, volumenmensual, volumenanual, tarifadecompra, profit, tarifadeventa, ingresopotencialanual, ingresopotencialmensual, numeroderutasofertadas, movimientosrechazados int
		var vendedor, tipodecliente, tipodeempresa, fecha, nombredelcliente, tamañodecliente, bussinesconection, tipodeindustria, tipodemercancia, tipodeventa, servicio, contacto, correoelectronico, telefono, trafico, tipodeservicio, tipodeequipounidad, tipodemodalidad, ciudadestadopaisorigen, ciudadestadopaisdestino, licitacion, target, ultimostatus, nocotizacionenfowardcomentariosop, comentariodelstatus, comentariosdelrechazo string
		err = registros.Scan(&movid, &vendedor, &fecha, &semana, &nombredelcliente, &tipodecliente, &tamañodecliente, &bussinesconection, &tipodeempresa, &tipodeindustria, &tipodemercancia, &tipodeventa, &servicio, &contacto, &correoelectronico, &telefono, &trafico, &tipodeservicio, &tipodeequipounidad, &tipodemodalidad, &ciudadestadopaisorigen, &ciudadestadopaisdestino, &licitacion, &target, &volumenmensual, &volumenanual, &tarifadecompra, &profit, &tarifadeventa, &porcentajederentabilidad, &ingresopotencialmensual, &ingresopotencialanual, &ultimostatus, &numeroderutasofertadas, &nocotizacionenfowardcomentariosop, &comentariodelstatus, &movimientosrechazados, &comentariosdelrechazo)

		if err != nil {
			panic(err.Error())
		}

		movimientos.MovId = movid
		movimientos.Vendedor = vendedor
		movimientos.Fecha = fecha
		movimientos.Semana = semana
		movimientos.NombreDelCliente = nombredelcliente
		movimientos.TipoDeCliente = tipodecliente
		movimientos.TamañoDeCliente = tamañodecliente
		movimientos.BussinesConection = bussinesconection
		movimientos.TipoDeEmpresa = tipodeempresa
		movimientos.TipoDeIndustria = tipodeindustria
		movimientos.TipoDeMercancia = tipodemercancia
		movimientos.TipoDeVenta = tipodeventa
		movimientos.Servicio = servicio
		movimientos.Contacto = contacto
		movimientos.CorreoElectronico = correoelectronico
		movimientos.Telefono = telefono
		movimientos.Trafico = trafico
		movimientos.TipoDeServicio = tipodeservicio
		movimientos.TipoDeEquipoUnidad = tipodeequipounidad
		movimientos.TipoDeModalidad = tipodemodalidad
		movimientos.CiudadEstadoPaisOrigen = ciudadestadopaisorigen
		movimientos.CiudadEstadoPaisDestino = ciudadestadopaisdestino
		movimientos.Licitacion = licitacion
		movimientos.Target = target
		movimientos.VolumenMensual = volumenmensual
		movimientos.VolumenAnual = volumenanual
		movimientos.TarifaDeCompra = tarifadecompra
		movimientos.Profit = profit
		movimientos.TarifaDeVenta = tarifadeventa
		movimientos.PorcentajeDeRentabilidad = porcentajederentabilidad
		movimientos.IngresoPotencialMensual = ingresopotencialmensual
		movimientos.IngresoPotencialAnual = ingresopotencialanual
		movimientos.UltimoStatus = ultimostatus
		movimientos.NumeroDeRutasOfertadas = numeroderutasofertadas
		movimientos.NoCotizacionEnFowardComentariosOP = nocotizacionenfowardcomentariosop
		movimientos.ComentariosDelStatus = comentariodelstatus
		movimientos.MovimientosRechazados = movimientosrechazados
		movimientos.ComentariosDelRechazo = comentariosdelrechazo

		arregloMovimientos = append(arregloMovimientos, movimientos)

	}

	plantillas.ExecuteTemplate(w, "index", arregloMovimientos)
}

// funcion de direccionamiento a la pagina crear.tmpl
func crear(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola")
	plantillas.ExecuteTemplate(w, "crear", nil)
}

// funcion de insertar en la base de datos
func insertar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		//movid := r.FormValue("movid")
		vendedor := r.FormValue("vendedor")
		fecha := r.FormValue("fecha")
		semana := r.FormValue("semana")
		nombredelcliente := r.FormValue("nombredelcliente")
		tipodecliente := r.FormValue("tipodecliente")
		tamañodelcliente := r.FormValue("tamañodelcliente")
		bussinesconection := r.FormValue("bussinesconection")
		tipodeempresa := r.FormValue("tipodeempresa")
		tipodeindustria := r.FormValue("tipodeindustria")
		tipodemercancia := r.FormValue("tipodemercancia")
		tipodeventa := r.FormValue("tipodeventa")
		servicio := r.FormValue("servicio")
		contacto := r.FormValue("contacto")
		correoelectronico := r.FormValue("correoelectronico")
		telefono := r.FormValue("telefono")
		trafico := r.FormValue("trafico")
		tipodeservicio := r.FormValue("tipodeservicio")
		tipodeequipounidad := r.FormValue("tipodeequipounidad")
		tipodemodalidad := r.FormValue("tipodemodalidad")
		ciudadestadopaisorigen := r.FormValue("ciudadestadopaisorigen")
		ciudadestadopaisdestino := r.FormValue("ciudadestadopaisdestino")
		licitacion := r.FormValue("licitacion")
		target := r.FormValue("target")
		volumenmensual := r.FormValue("volumenmensual")
		volumenanual := r.FormValue("volumenanual")
		tarifadecompra := r.FormValue("tarifadecompra")
		profit := r.FormValue("profit")
		tarifadeventa := r.FormValue("tarifadeventa")
		porcentajederentabilidad := r.FormValue("porcentajederentabilidad")
		ingresopotencialmensual := r.FormValue("ingresopotencialmensual")
		ingresopotencialanual := r.FormValue("ingresopotencialanual")
		ultimostatus := r.FormValue("ultimostatus")
		numeroderutasofertadas := r.FormValue("numeroderutasofertadas")
		nocotizacionenfowardcomentariosop := r.FormValue("nocotizacionenfowardcomentariosop")
		comentariosdelstatus := r.FormValue("comentariosdelstatus")
		movimientosrechazados := r.FormValue("movimientosrechazados")
		comentariosdelrechazo := r.FormValue("comentariosdelrechazo")

		conexionEstablecida := conexionBD()

		insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO pipeline (Vendedor,Fecha,Semana,NombreDelCliente,TipoDeCliente,TamañoDelCliente,BussinesConection,TipoDeEmpresa,TipoDeIndustria,TipoDeMercancia,TipoDeVenta,Servicio,Contacto,CorreoElectronico,Telefono,Trafico,TipoDeServicio,TipoDeEquipoUnidad,TipoDeModalidad,CiudadEstadoPaisOrigen,CiudadEstadoPaisDestino,Licitacion,Target,VolumenMensual,VolumenAnual,TarifaDeCompra,Profit,TarifaDeVenta,PorcentajeDeRentabilidad,IngresoPotencialMensual,IngresoPotencialAnual,UltimoStatus,NumeroDeRutasOfertadas,NoCotizacionEnFowardComentariosOP,ComentariosDelStatus,MovimientosRechazados,ComentariosDelRechazo) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);")
		if err != nil {
			panic(err.Error())
		}
		insertarRegistros.Exec(vendedor, fecha, semana, nombredelcliente, tipodecliente, tamañodelcliente, bussinesconection, tipodeempresa, tipodeindustria, tipodemercancia, tipodeventa, servicio, contacto, correoelectronico, telefono, trafico, tipodeservicio, tipodeequipounidad, tipodemodalidad, ciudadestadopaisorigen, ciudadestadopaisdestino, licitacion, target, volumenmensual, volumenanual, tarifadecompra, profit, tarifadeventa, porcentajederentabilidad, ingresopotencialmensual, ingresopotencialanual, ultimostatus, numeroderutasofertadas, nocotizacionenfowardcomentariosop, comentariosdelstatus, movimientosrechazados, comentariosdelrechazo)

		http.Redirect(w, r, "/", 301)
	}
}

// Editar la base de datos y actualizar
func Actualizar(w http.ResponseWriter, r *http.Request) {
	conexionEstablecida := conexionBD()

	if r.Method == "POST" {

		id := r.FormValue("movid")
		vendedor := r.FormValue("vendedor")
		fecha := r.FormValue("fecha")
		semana := r.FormValue("semana")
		nombredelcliente := r.FormValue("nombredelcliente")
		tipodecliente := r.FormValue("tipodecliente")
		tamañodelcliente := r.FormValue("tamañodelcliente")
		bussinesconection := r.FormValue("bussinesconection")
		tipodeempresa := r.FormValue("tipodeempresa")
		tipodeindustria := r.FormValue("tipodeindustria")
		tipodemercancia := r.FormValue("tipodemercancia")
		tipodeventa := r.FormValue("tipodeventa")
		servicio := r.FormValue("servicio")
		contacto := r.FormValue("contacto")
		correoelectronico := r.FormValue("correoelectronico")
		telefono := r.FormValue("telefono")
		trafico := r.FormValue("trafico")
		tipodeservicio := r.FormValue("tipodeservicio")
		tipodeequipounidad := r.FormValue("tipodeequipounidad")
		tipodemodalidad := r.FormValue("tipodemodalidad")
		ciudadestadopaisorigen := r.FormValue("ciudadestadopaisorigen")
		ciudadestadopaisdestino := r.FormValue("ciudadestadopaisdestino")
		licitacion := r.FormValue("licitacion")
		target := r.FormValue("target")
		volumenmensual := r.FormValue("volumenmensual")
		volumenanual := r.FormValue("volumenanual")
		tarifadecompra := r.FormValue("tarifadecompra")
		profit := r.FormValue("profit")
		tarifadeventa := r.FormValue("tarifadeventa")
		porcentajederentabilidad := r.FormValue("porcentajederentabilidad")
		ingresopotencialmensual := r.FormValue("ingresopotencialmensual")
		ingresopotencialanual := r.FormValue("ingresopotencialanual")
		ultimostatus := r.FormValue("ultimostatus")
		numeroderutasofertadas := r.FormValue("numeroderutasofertadas")
		nocotizacionenfowardcomentariosop := r.FormValue("nocotizacionenfowardcomentariosop")
		comentariosdelstatus := r.FormValue("comentariosdelstatus")
		movimientosrechazados := r.FormValue("movimientosrechazados")
		comentariosdelrechazo := r.FormValue("comentariosdelrechazo")

		modificarRegistros, err := conexionEstablecida.Prepare("UPDATE pipeline SET Vendedor=?, Fecha=?, Semana=?, NombreDelCliente=?, TipoDeCliente=?,TamañoDelCliente=?, BussinesConection=?, TipoDeEmpresa=?, TipoDeIndustria=?, TipoDeMercancia=?, TipoDeVenta=?, Servicio=?, Contacto=?, CorreoElectronico=?, Telefono=?, Trafico=?, TipoDeServicio=?, TipoDeEquipoUnidad=?, TipoDeModalidad=?, CiudadEstadoPaisOrigen=?, CiudadEstadoPaisDestino=?, Licitacion=?, Target=?, VolumenMensual=?, VolumenAnual=?, TarifaDeCompra=?, Profit=?, TarifaDeVenta=?,PorcentajeDeRentabilidad=?, IngresoPotencialMensual=?, IngresoPotencialAnual=?, UltimoStatus=?, NumeroDeRutasOfertadas=?, NoCotizacionEnFowardComentariosOP=?, ComentariosDelStatus=?, MovimientosRechazados=?, ComentariosDelRechazo=? WHERE MovId=?;")
		if err != nil {
			panic(err.Error())
		}
		modificarRegistros.Exec(vendedor, fecha, semana, nombredelcliente, tipodecliente, tamañodelcliente, bussinesconection, tipodeempresa, tipodeindustria, tipodemercancia, tipodeventa, servicio, contacto, correoelectronico, telefono, trafico, tipodeservicio, tipodeequipounidad, tipodemodalidad, ciudadestadopaisorigen, ciudadestadopaisdestino, licitacion, target, volumenmensual, volumenanual, tarifadecompra, profit, tarifadeventa, porcentajederentabilidad, ingresopotencialmensual, ingresopotencialanual, ultimostatus, numeroderutasofertadas, nocotizacionenfowardcomentariosop, comentariosdelstatus, movimientosrechazados, comentariosdelrechazo, id)

		http.Redirect(w, r, "/", 301)
	}
}
