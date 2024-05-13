package main

import "fmt"

type lista struct {
	pri, ult *nodo
}

type nodo struct {
	val int
	sig *nodo
}

func main() {
}

func New() lista {
	l := lista{}
	return l
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

func (l *lista) AgregarAdelante(n int) {
	var nue nodo
	nue.val = n
	nue.sig = nil

	if l.EstaVacia() {
		l.pri = &nue
		l.ult = &nue
	} else {
		nue.sig = l.pri
		l.pri = &nue
	}
}

func (l *lista) AgregarAtras(n int) {
	var nue nodo
	nue.val = n
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

func (l *lista) eliminarValor(n int) {
	if l.EstaVacia() {
		fmt.Errorf("No se puede eliminar un elemento de una lista vacia.")
		fmt.Println("No se puede eliminar un elemento de una lista vacia")
	}

	ok := false
	ant := l.pri

	for e := l.pri; e != nil && ok == false; e = e.sig {
		if e.val == n {
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
