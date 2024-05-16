package main

import "fmt"

type OptinumSlice []Elemento

type Elemento struct {
	numero      int
	ocurrencias int
}

func main() {
	s := []int{1, 1, 1, 1, 4, 1, 1}
	o := New(s)
	fmt.Println("Optinum Slice: ", o)

	o = Insert(o, 5, -1)
	o = Insert(o, 6, 1)
	o = Insert(o, 6, 0)
	o = Insert(o, 6, 5)
	o = Insert(o, 8, 1000)

	fmt.Println("Optinum Slice despues del insertar: ", o)

	e := SliceArray(o)
	fmt.Println("Slice creado a partir del optinum Slice: ", e)
}

func SliceArray(o OptinumSlice) []int {
	var slice []int
	for _, e := range o {
		for i := 0; i < e.ocurrencias; i++ {
			slice = append(slice, e.numero)
		}
	}
	return slice
}

func Insert(o OptinumSlice, nro int, pos int) OptinumSlice {
	auxPos := pos
	if pos < 0 {
		fmt.Println("La posicion", pos, " que ingreso es negativa, no se podrá insertar el numero ", nro, ", tome una posicion mayor o igual a 0 y menor igual que ", LenSlice(o))
	} else {
		for i, e := range o {
			if e.ocurrencias > auxPos {
				if e.numero == nro {
					o[i].ocurrencias++ //SI ES EL MISMO NRO SOLO LE SUMO LA CANTIDAD DE OCURRENCIAS
				} else {
					if pos == 0 {
						secondPart := o[i:]

						var newElement OptinumSlice
						newElement = append(newElement, Elemento{numero: nro, ocurrencias: 1})

						// Agregar el nuevo elemento en la posición deseada
						o = append(newElement, secondPart...)
					} else {
						var auxOcurrencias int
						auxOcurrencias = e.ocurrencias - auxPos //RESTO LA CANTIDAD DE OCURRENCIAS - POS ASI TENGO LA CANTIDAD PARA EL SEGUNDO CAMBIO SI ES NECESARIO
						o[i].ocurrencias -= auxOcurrencias      //LE RESTO A LAS OCURRENCIAS LA CANTIDAD QUE VA A TENER AHORA EL ELEMENTO

						ok := false
						firstPart := o[:i+1]
						if o[i].ocurrencias == 0 {
							firstPart = o[:i]
							ok = true
						}

						secondPart := o[i+1:]

						var newElement OptinumSlice
						newElement = append(newElement, Elemento{numero: nro, ocurrencias: 1})

						// Agregar el nuevo elemento en la posición deseada
						o = append(firstPart, append(newElement, secondPart...)...)

						if auxOcurrencias > 0 {
							if ok == true {
								firstPart = o[:i+1]
								secondPart = o[i+1:]
							} else {
								firstPart = o[:i+2]
								secondPart = o[i+2:]
							}

							newElement[0] = Elemento{numero: e.numero, ocurrencias: auxOcurrencias}

							o = append(firstPart, append(newElement, secondPart...)...)

						}
					}
				}
				fmt.Println("Se agrego el numero ", nro, " en la pos ", pos)
				return o
			} else {
				auxPos -= e.ocurrencias
			}
		}
		if auxPos >= 0 {
			fmt.Println("Su posicion estaba fuera de rango, por ende se agregara el numero ", nro, " al final en la posicion: ", LenSlice(o))
			o = append(o, Elemento{numero: nro, ocurrencias: 1})

		}
	}
	return o
}
func LastElement(o OptinumSlice) int {
	if !IsEmpty(o) {
		return o[Len(o)-1].numero //-1 PQ LAS POSICIONES DEL OPTINUM ARRANCAN EN 0 Y LEN DEVUELVE LA CANT DE ELEMENTOS DESDE 1
	} else {
		fmt.Println("Estas queriendo averiguar la primer posicion de un Optimus Slice vacio.")
		return -1
	}
}

func FrontElement(o OptinumSlice) int {
	if !IsEmpty(o) {
		return o[0].numero
	} else {
		fmt.Println("Estas queriendo averiguar la primer posicion de un Optimus Slice vacio.")
		return -1
	}
}

func LenSlice(o OptinumSlice) int {
	s := SliceArray(o)
	return len(s)
}

func Len(o OptinumSlice) int {
	return len(o)
}

func IsEmpty(o OptinumSlice) bool {
	return len(o) == 0
}

func New(s []int) OptinumSlice {
	var optinumSlice OptinumSlice
	posOptinum := 0

	for i, val := range s {
		if i == 0 {
			optinumSlice = append(optinumSlice, Elemento{numero: val, ocurrencias: 0})
		}
		if val != optinumSlice[posOptinum].numero {
			optinumSlice = append(optinumSlice, Elemento{numero: val, ocurrencias: 1})
			posOptinum++
		} else {
			optinumSlice[posOptinum].ocurrencias++
		}
	}

	return optinumSlice
}
