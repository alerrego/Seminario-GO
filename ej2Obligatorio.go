package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type nodoBlockChain struct {
	b   bloque
	sig *nodoBlockChain
}

type listaBlockChain struct {
	pri, ult              *nodoBlockChain
	hashInicio, hashFinal string
}

type nodoBilletera struct {
	data billetera
	sig  *nodoBilletera
}

type listaBilleteras struct {
	pri *nodoBilletera
}

type billetera struct {
	nombre   string
	apellido string
	id       int
	monto    float64
}

type transaccion struct {
	monto       float64
	idEnviador  int
	idRecibidor int
	fecha       string
}

type bloque struct {
	hash       string
	hashPrevio string
	data       transaccion
	fecha      string
}

func main() {
	var opc int
	lBilleteras := NewListaBilletera()
	lBlockChain := NewListaBlockChain()
	id := 1
	idTransaccion := 1

	fmt.Println("MENU: ")
	fmt.Println("Ingrese 1 para crear una nueva billetera: ")
	fmt.Println("Ingrese 2 para crear una transaccion: ")
	fmt.Println("Ingrese 3 para validar la cadena: ")
	fmt.Println("Ingrese 4 para obtener el saldo de un usuario recorriendo la cadena: ")
	fmt.Println("Ingrese 5 para recorrer la cadena: ")
	fmt.Println("Ingrese 6 para recorrer las billeteras: ")
	fmt.Println("Ingrese 7 para alterar una cadena: ")
	fmt.Println("Ingrese 0 para finalizar")

	fmt.Scanln(&opc)

	for opc != 0 {

		switch opc {
		case 1:
			lBilleteras.CrearBilletera(&id)
		case 2:
			lBlockChain.CrearTransaccion(&idTransaccion, lBilleteras)
		case 3:
			lBlockChain.ValidarCadena()
		case 4:
			var dniSaldo int
			fmt.Println("Ingrese el id del usuario a obtener el saldo: ")
			fmt.Scanln(&dniSaldo)
			total := lBlockChain.ObtenerSaldo(dniSaldo)
			fmt.Println("El saldo obtenido por el usuario con id: ", dniSaldo, " fue: ", total)
		case 5:
			lBlockChain.RecorrerBlockChain()
		case 6:
			lBilleteras.RecorrerBilleteras()
		case 7:
			var nodo int
			fmt.Println("Ingrese el nodo que quiere alterar: ")
			fmt.Scanln(&nodo)
			lBlockChain.AlterarHash(nodo)
		}

		fmt.Println("MENU: ")
		fmt.Println("Ingrese 1 para crear una nueva billetera: ")
		fmt.Println("Ingrese 2 para crear una transaccion: ")
		fmt.Println("Ingrese 3 para validar la cadena: ")
		fmt.Println("Ingrese 4 para obtener el saldo de un usuario recorriendo la cadena: ")
		fmt.Println("Ingrese 5 para recorrer la cadena: ")
		fmt.Println("Ingrese 6 para recorrer las billeteras: ")
		fmt.Println("Ingrese 7 para alterar una cadena: ")
		fmt.Println("Ingrese 0 para finalizar")

		fmt.Scanln(&opc)
	}
}

func (l *listaBlockChain) AlterarHash(nodo int) {
	ok := false
	var i int
	for e := l.pri; e != nil && ok != true; e = e.sig {
		if i == nodo {
			e.b.hash = "ALTERADO"
			ok = true
		}
		i++
	}
	if ok == false {
		fmt.Println("No existe el nodo ", nodo, " en la BlockChain.")
	}
}
func (l *listaBilleteras) RecorrerBilleteras() {
	for e := l.pri; e != nil; e = e.sig {
		fmt.Println("La billetera pertenece a: ", e.data.nombre, " ", e.data.apellido, " con el ID: ", e.data.id, " y tiene el saldo de: ", e.data.monto)
	}
}
func (l *listaBlockChain) RecorrerBlockChain() {
	for e := l.pri; e != nil; e = e.sig {
		fmt.Println("La transaccion con hash: ", e.b.hash, " hecha del id: ", e.b.data.idEnviador, " a ", e.b.data.idRecibidor, " fue de: ", e.b.data.monto, " en la fecha: ", e.b.fecha)
	}
}

func (l *listaBlockChain) ObtenerSaldo(dniSaldo int) float64 {
	var total float64

	for e := l.pri; e != nil; e = e.sig {
		if e.b.data.idRecibidor == dniSaldo {
			total += e.b.data.monto
		}
	}

	return total
}

func (l *listaBlockChain) ValidarCadena() {
	ant := l.pri
	sig := l.pri.sig
	var i int

	for e := l.pri; e != nil; e = e.sig {
		//CASO INICIO
		if e == l.pri {
			if e.b.hash != l.hashInicio {
				fmt.Println("Se eliminará el bloque", i, " debido a que su hash ha sido alterado.")
				l.pri = l.pri.sig
				//ME ASEGURO QUE AL BORRAR NO QUEDE CON LA LISTA VACIA(PQ ERA EL UNICO ELEM)
				if l.pri != nil {
					l.hashInicio = l.pri.b.hash
					l.pri.b.hashPrevio = l.pri.b.hash
				}
			}
			//CASO MEDIO
		} else if sig != nil {
			if e.b.hash != sig.b.hashPrevio {
				fmt.Println("Se eliminará el bloque ", i, " debido a que su hash ha sido alterado.")
				ant.sig = sig
				sig.b.hashPrevio = ant.b.hash
			}
			//CASO FINAL
		} else {
			if e.b.hash != l.hashFinal {
				fmt.Println("Se eliminará el bloque ", i, " debido a que su hash ha sido alterado.")
				ant.sig = nil
				l.hashFinal = ant.b.hash
			}
		}
		if sig != nil {
			i++
			ant = e
			sig = sig.sig
		}
	}
}

func (l *listaBlockChain) CrearTransaccion(id *int, lBilleteras listaBilleteras) {
	var t transaccion
	fmt.Println("Ingrese su id: ")
	fmt.Scanln(&t.idEnviador)
	fmt.Println("Ingrese el id a enviar la transferencia: ")
	fmt.Scanln(&t.idRecibidor)
	fmt.Println("Ingrese el monto a transferir: ")
	fmt.Scanln(&t.monto)

	if lBilleteras.NoTieneSaldoSuficiente(t.idEnviador, t.monto) {
		fmt.Println("La billetera del usuario con ID: ", t.idEnviador, " no tiene el suficiente saldo para realizar la transaccion de: ", t.monto)
		return
	}

	t.fecha = time.Now().String()

	var b bloque

	b.data = t
	b.fecha = time.Now().String()

	if l.pri == nil {
		hash := CrearHash(id)
		l.hashInicio = hash
		l.hashFinal = hash
		b.hash = hash
		b.hashPrevio = hash
	} else {
		b.hashPrevio = l.UltimoHash()
		b.hash = CrearHash(id)
	}

	l.AgregarBloqueAtras(b)

	lBilleteras.ActualizarBilleteras(t.idEnviador, t.idRecibidor, t.monto)
}
func (l *listaBilleteras) NoTieneSaldoSuficiente(id int, monto float64) bool {
	ok := false
	for e := l.pri; e != nil && ok != true; e = e.sig {
		if e.data.id == id {
			if e.data.monto < monto {
				ok = true
			}
		}
	}
	return ok
}
func (l *listaBilleteras) ActualizarBilleteras(idEnviador int, idRecibidor int, monto float64) {
	for e := l.pri; e != nil; e = e.sig {
		if e.data.id == idEnviador {
			e.data.monto -= monto
		} else if e.data.id == idRecibidor {
			e.data.monto += monto
		}
	}
}
func (l *listaBlockChain) AgregarBloqueAtras(b bloque) {
	var nue nodoBlockChain
	nue.b = b
	nue.sig = nil

	if l.pri == nil {
		l.pri = &nue
		l.ult = &nue
	} else {
		l.ult.sig = &nue
		l.ult = &nue
		l.hashFinal = nue.b.hash
	}

}

func CrearHash(id *int) string {
	hasheador := sha256.New() //CREO EL OBJ PARA TENER LOS METODOS Y PODER CALCULAR EL HASH

	idToStr := strconv.Itoa(*id) //CONVIERTO EL ID A STRING YA QUE ESTA EN INT

	hasheador.Write([]byte(idToStr)) //AGREGO EL ID EN EL HASHEADOR EN FORMA DE BYTES

	hashBytes := hasheador.Sum(nil) //SE CALCULA EL HASH (NIL ES PQ NO QUEREMOS AGREGAR BYTES ADICIONALES)

	hashString := hex.EncodeToString(hashBytes) //RECONVIERTO EL HASH DE BYTES A HEXA PARA MAYOR LEGIBILIDAD (NO ES NECESARIO)

	return hashString
}

func (l *listaBlockChain) UltimoHash() string {
	var ultimoHash string
	ultimoHash = l.ult.b.hash
	return ultimoHash
}

func (l *listaBilleteras) InsertarAdelante(b billetera) {
	var nue nodoBilletera
	nue.data = b
	nue.sig = nil

	if l.pri == nil {
		l.pri = &nue
	} else {
		nue.sig = l.pri
		l.pri = &nue
	}
}

func (l *listaBilleteras) CrearBilletera(id *int) {
	var b billetera

	fmt.Println("Ingrese su nombre: ")
	fmt.Scanln(&b.nombre)
	fmt.Println("Ingrese su apellido: ")
	fmt.Scanln(&b.apellido)
	fmt.Println("Ingrese el monto a depositar en la billetera: ")
	fmt.Scanln(&b.monto)
	b.id = *id
	*id++ //SUMO UNO ASI NO REPITO EL ID DEL USUARIO
	l.InsertarAdelante(b)
	fmt.Println("Tu ID es: ", b.id)
}

func NewListaBilletera() listaBilleteras {
	var l listaBilleteras
	return l
}

func NewListaBlockChain() listaBlockChain {
	var l listaBlockChain
	return l
}
