package main

import (
	"fmt"
)

func main() {
	var temperaturas [10]float64

	var temperaturaMaxima float64 = 0
	var temperaturaMinima float64 = 9999

	const pacientes = 3
	var cantBaja int
	var cantMedia int
	var cantAlta int
	var temp float64

	for i := 0; i < pacientes; i++ {
		fmt.Println("Hola paciente nro: ", i, " por favor ingrese su temperatura corporal: ")
		fmt.Scanln(&temp)
		temperaturas[i] = temp
		switch {
		case temperaturas[i] < 36:
			fmt.Println("Su temperatura es baja")
			cantBaja += 1
			if temperaturas[i] < temperaturaMinima {
				temperaturaMinima = temp
			}
			if temperaturas[i] > temperaturaMaxima {
				temperaturaMaxima = temp
			}
		case (temperaturas[i] >= 36) && (temperaturas[i] <= 37.5):
			fmt.Println("Temperatura media: ")
			cantMedia += 1
			if temperaturas[i] < temperaturaMinima {
				temperaturaMinima = temp
			}
			if temperaturas[i] > temperaturaMaxima {
				temperaturaMaxima = temp
			}
		case temperaturas[i] > 37.5:
			fmt.Println("Temperatura alta: ")
			cantAlta += 1
			if temperaturas[i] < temperaturaMinima {
				temperaturaMinima = temp
			}
			if temperaturas[i] > temperaturaMaxima {
				temperaturaMaxima = temp
			}
		}
	}
	fmt.Println("Gente con temp baja: ", cantBaja)
	fmt.Println("El promedio de pacientes con temperatura baja fue de: ", (cantBaja*100)/3)
	fmt.Println("El promedio de pacientes con temperatura media fue de: ", (cantMedia*100)/3)
	fmt.Println("El promedio de pacientes con temperatura alta fue de: ", (cantAlta*100)/3)

	fmt.Println("El promedio entre la temperatura maxima y la temperatura minima es: ", (temperaturaMaxima+temperaturaMinima)/2)
}
