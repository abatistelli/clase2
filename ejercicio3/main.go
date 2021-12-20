/*

Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas
trabajadas por mes y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y
la categoría, y que devuelva su salario.

*/

package main

import "fmt"

func main() {

	fmt.Printf("El salario es: %.2f\n", calcularSalario(6000, "B"))

}

func calcularSalario(minTrabajados float64, categoria string) float64 {
	horas := minTrabajados / 60

	switch categoria {
	case "A":
		salario := horas * 3
		salario += salario * 0.5
		return salario
	case "B":
		salario := horas * 1.5
		salario += salario * 0.2
		return salario
	case "C":
		salario := horas
		return salario
	default:
		return 0
	}
}
