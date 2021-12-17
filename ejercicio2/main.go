/*

Un colegio de Buenos Aires necesita calcular el promedio (por alumno) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y
devuelva el promedio y un error en caso que uno de los números ingresados sea negativo

*/

package main

import (
	"errors"
	"fmt"
)

func main() {

	promedio, err := calcularPromedio(7, 8, 3, 2, 10, 9.5)

	if err == nil {
		fmt.Printf("El promedio es: %.2f\n", promedio)
		return
	}

	fmt.Printf("%v\n", err)

}

func calcularPromedio(notas ...float64) (float64, error) {
	var sumaNota float64

	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("ERROR!! Nota negativa")
		}
		sumaNota += nota
	}

	return sumaNota / float64(len(notas)), nil
}
