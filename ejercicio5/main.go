/*
Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan haber
muchos más animales que refugiar.

Por perro necesitan 10 kg de alimento
Por gato 5 kg
Por cada Hamster 250 gramos.
Por Tarántula 150 gramos.


Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal
especificado y que retorne una función y un error (en caso que no exista el animal)

Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal
especificado.

*/

package main

import (
	"errors"
	"fmt"
)

const (
	perro = "perro"
	gato  = "gato"
)

func main() {
	animalPerro, err1 := Animal(perro)
	animalGato, err2 := Animal(gato)

	if err1 == nil && err2 == nil {
		var cantidad float64
		cantidad += animalPerro(5)
		cantidad += animalGato(8)

		fmt.Printf("La cantidad de alimento a comprar es: %.2f\n", cantidad)
		return
	}

	fmt.Printf("%s", err1)
}

func calculoAlimentoPerro(cantidad int) float64 {
	return float64(cantidad * 10)
}

func calculoAlimentoGato(cantidad int) float64 {
	return float64(cantidad * 5)
}

func Animal(animal string) (func(cantidadAnimales int) float64, error) {

	switch animal {
	case "perro":
		return calculoAlimentoPerro, nil
	case "gato":
		return calculoAlimentoGato, nil
	default:
		return nil, errors.New("Animal Invalido")
	}
}
