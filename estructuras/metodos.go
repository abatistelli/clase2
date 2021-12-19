/*
Go no tiene clases. Sin embargo, se pueden definir metodos en los tipos de datos.
Un metodo es una funcion con un argumento de recpetor especial.
El receptor aparece en su propia lista de argumentos entre la palabra clave func y el nombre del metodos.
*/

package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	radio float64
}

func (c Circulo) area() float64 {
	return math.Pi * c.radio * c.radio
}

func (c Circulo) perim() float64 {
	return 2 * math.Pi * c.radio
}

func (c *Circulo) setRadio(r float64) {
	fmt.Println("Radio Method: ", r)
	c.radio = r
}
func main() {

	c := Circulo{radio: 5}

	fmt.Println(c.area())
	fmt.Println(c.perim())
	c.setRadio(10)

	fmt.Println("Radio Main: ", c.radio)

	fmt.Println(c.area())
	fmt.Println(c.perim())
}
