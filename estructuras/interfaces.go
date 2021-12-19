/*
Una interfaz es una forma de definir metodo los cuales deben ser utilizados pero no los define

Las interfaces son utilizadas para brindar modularidad al lenguaje, nos permiten implementar el mismo
comportamiento a diferentes objetos

*/

package main

import (
	"fmt"
	"math"
)

const (
	rectType   = "RECT"
	circleType = "CIRCLE"
)

type Geometry interface {
	area() float64
	perim() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func details(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

type Rect struct {
	width, height float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func newCircle(values float64) Circle {
	return Circle{radius: values}
}

func newGeometry(geoType string, values ...float64) Geometry {
	switch geoType {
	case rectType:
		return Rect{width: values[0], height: values[1]}

	case circleType:
		return Circle{radius: values[0]}
	}

	return nil
}

func main() {
	//r := Rect{width: 3, height: 4}
	//c := Circle{radius: 5}
	//details(c)
	//details(r)

	//c := newCircle(2)
	//fmt.Println(c.area())
	//fmt.Println(c.perim())

	r := newGeometry(rectType, 2, 3)
	fmt.Println(r.area())
	fmt.Println(r.perim())

	c := newGeometry(circleType, 2)
	fmt.Println(c.area())
	fmt.Println(c.perim())

	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

}

/*
EXTRA

The "empty interface"

Es un tipo de interface especial que no tiene metodos especificados.

Una interface vacia puede almacenar valores de cualquier tipo porque las interfaces
implementan por lo menos cero metodos

Type assertion (asercion de tipos)

Proveen acceso al tipo exacto de variables de una interface. Si el tipo de datos esta presente en la
interface, entonces recupera el tipo de dato real que esta siendo albergado por la interface

*/
