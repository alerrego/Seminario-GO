package main

import (
	"fmt"
	"strings"
)

type lista struct {
	pri, ult *nodo
}

type nodo struct {
	val ingresante
	sig *nodo
}
type fechaNacimiento struct {
	dia int
	mes int
	ano int
}
type ingresante struct {
	apellido string
	nombre   string
	ciudad   string
	fecha    fechaNacimiento
	titulo   bool
	carrera  string
	dni      int
}

func main() {
	l := NewList()
	l.CargarLista()
	l.RecorrerLista()
	l.mostrarLista()
}

func NewList() lista {
	l := lista{}
	return l
}

func calcularMaximoAno(m map[int]int) {
	var max int
	var AnoMax int
	for clave, val := range m {
		if val > max {
			max = val
			AnoMax = clave
		}
	}
	fmt.Println("La mayoria de ingresantes nacieron en el ano: ", AnoMax)
}

func calcularMaximoCarrera(m map[string]int) {
	var max int
	var CarreraMax string
	for clave, val := range m {
		if val > max {
			max = val
			CarreraMax = clave
		}
	}
	fmt.Println("La carrera con mas inscriptos fue: ", CarreraMax)
}

func (l *lista) RecorrerLista() {
	mapAno := map[int]int{}
	mapCarrera := map[string]int{}
	for e := l.pri; e != nil; e = e.sig {
		if strings.ToLower(e.val.ciudad) == "bariloche" {
			fmt.Println("El estudiante: ", e.val.nombre, " ", e.val.apellido, " nacio en Bariloche.")
		}
		mapAno[e.val.fecha.ano] += 1
		mapCarrera[strings.ToUpper(e.val.carrera)] += 1
		if !e.val.titulo {
			l.eliminarValor(e.val.dni)
		}
	}
	calcularMaximoAno(mapAno)
	calcularMaximoCarrera(mapCarrera)
}

func (l *lista) CargarLista() {
	i := LeerIngresante()
	for i.dni != 0 {
		l.AgregarAtras(i)
		i = LeerIngresante()
	}

}

func LeerIngresante() ingresante {
	var i ingresante
	fmt.Println("Ingrese el dni del ingresante: ")
	fmt.Scanln(&i.dni)
	if i.dni != 0 {
		var aux string
		fmt.Println("Ingrese el nombre: ")
		fmt.Scanln(&i.nombre)
		fmt.Println("Ingrese el apellido: ")
		fmt.Scanln(&i.apellido)
		fmt.Println("Ingrese la ciudad de origen: ")
		fmt.Scanln(&i.ciudad)
		fmt.Println("Ingrese la fecha de nacimiento: ")
		fmt.Println("dia: ")
		fmt.Scanln(&i.fecha.dia)
		fmt.Println("mes: ")
		fmt.Scanln(&i.fecha.mes)
		fmt.Println("ano: ")
		fmt.Scanln(&i.fecha.ano)
		fmt.Println("Ingrese la carrera a la que se inscribio (APU,LI,LS):")
		fmt.Scanln(&i.carrera)
		fmt.Println("Presento el titulo secundario: ")
		fmt.Scanln(&aux)
		if strings.ToLower(aux) == "si" {
			i.titulo = true
		} else {
			i.titulo = false
		}
	}
	return i
}

func (l *lista) EstaVacia() bool {
	var ok bool

	if l.pri != nil {
		ok = false
	} else {
		ok = true
	}

	return ok
}

func (l *lista) AgregarAdelante(i ingresante) {
	var nue nodo
	nue.val = i
	nue.sig = nil

	if l.EstaVacia() {
		l.pri = &nue
		l.ult = &nue
	} else {
		nue.sig = l.pri
		l.pri = &nue
	}
}

func (l *lista) AgregarAtras(i ingresante) {
	var nue nodo
	nue.val = i
	nue.sig = nil

	if l.EstaVacia() {
		l.pri = &nue
		l.ult = &nue
	} else {
		l.ult.sig = &nue
		l.ult = &nue
	}
}

func (l *lista) mostrarLista() {
	var act = l.pri
	i := 1
	for act != nil {
		fmt.Println("El elemento nro ", i, " tiene el valor: ", act.val)
		i++
		act = act.sig
	}
}

func (l *lista) eliminarValor(dni int) {
	if l.EstaVacia() {
		fmt.Errorf("No se puede eliminar un elemento de una lista vacia.")
		fmt.Println("No se puede eliminar un elemento de una lista vacia")
	}

	ok := false
	ant := l.pri

	for e := l.pri; e != nil && ok == false; e = e.sig {
		if e.val.dni == dni {
			if e == l.pri {
				l.pri = l.pri.sig //SI ES EL PRIMERO REUBICO LOS PUNTEROS Y EL 2DO PUNTERO SERA MI PRIMER (SI NO HAY NADA SERA NIL)
				ok = true
				if e == l.ult { //EXISTE LA POSIBILIDAD QUE SEA EL PRIMER Y ULTIMO ELEMENTO POR ESO ACTUALIZO EL ULTIMO SI LO ES.
					l.ult = nil
				}
			} else {
				ant.sig = e.sig //RECONECTO EL ANTERIOR CON EL SIG Y PIERDO EL NODO ACT
				ok = true
				if e == l.ult { //EXISTE LA POSIBILIDAD QUE SEA EL ULTIMO POR ENDE ACTUALIZO EL ULTIMO
					l.ult = nil
				}
			}
		}
		ant = e //ME REUBICO EL ULTIMO SI NO LO USE
	}

	if !ok {
		fmt.Errorf("El elemento no se encontraba en la lista.")
		fmt.Println("No se puede eliminar un elemento que no existe en la lista")
	}
}
