package main

type nodo struct {
	b   bloque
	sig *nodo
}

type lista struct {
	pri *nodo
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

}
