/*

Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones
de los alumnos de un curso, requiriendo calcular los valores mínimo, máximo y promedio de sus
calificaciones.

Se solicita generar una función que indique qué tipo de cálculo se quiere realizar
(mínimo, máximo o promedio) y que devuelva otra función
( y un error en caso que el cálculo no esté definido) que se le puede pasar una cantidad N de enteros y
devuelva el cálculo que se indicó en la función anterior

*/

package main

import (
	"errors"
	"fmt"
)

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func main() {

	//minFunc, err := operacion(minimo)
	//promFunc, err := operacion(promedio)
	maxFunc, err := operacion(maximo)

	if err == nil {
		//valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
		//fmt.Printf("El valor mas chico es: %.2f\n", valorMinimo

		valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Printf("El valor mas grande es: %.2f\n", valorMaximo)

		//valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
		//fmt.Printf("El promedio es: %.2f\n", valorPromedio)

		return
	}

	fmt.Printf("%s\n", err)
}

func buscarMenor(valores ...float64) float64 {
	menor := valores[0]

	for _, valor := range valores {
		if valor < menor {
			menor = valor
		}
	}
	return menor
}

func buscarMayor(valores ...float64) float64 {
	mayor := valores[0]

	for _, valor := range valores {
		if valor > mayor {
			mayor = valor
		}
	}
	return mayor
}

func calcularPromedio(valores ...float64) float64 {
	var sumTotal float64

	for _, valor := range valores {
		sumTotal += valor
	}

	return sumTotal / float64(len(valores))
}

func operacion(op string) (func(notas ...float64) float64, error) {

	switch op {
	case "minimo":
		return buscarMenor, nil
	case "maximo":
		return buscarMayor, nil
	case "promedio":
		return calcularPromedio, nil
	default:
		return nil, errors.New("Operacion Invalida")

	}
}
