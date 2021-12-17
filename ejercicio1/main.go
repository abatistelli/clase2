/*
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más de $150.000
se le descontará además un 10%.
*/

package main

import "fmt"

func main() {

	fmt.Printf("%.2f\n", calcularImpuesto(70000))

}

func calcularImpuesto(salario float64) float64 {
	if salario > 50000 && salario <= 150000 {
		salario -= salario * 0.17
		return salario
	} else if salario > 150000 {
		salario -= salario * 0.17
		salario -= salario * 0.10
		return salario
	} else {
		return salario
	}
}
