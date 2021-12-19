// Una estructura es una coleccion de campos de datos

package main

import "fmt"

type Persona struct {
	Nombre    string
	Genero    string
	Edad      int
	Profesion string
	Peso      float64
	Gustos    Preferencias
}

type Preferencias struct {
	Comidas   string
	Peliculas string
	Series    string
	Animes    string
	Deportes  string
}

func main() {

	//  Diferentes formas de definir una estructura

	//p1 := Persona{"Emma", "Femenino", 33, "Developer", 59, Preferencias{"pollo", "titanic", "", "", ""}}

	p2 := Persona{
		Nombre:    "Pepe",
		Genero:    "Hombre",
		Edad:      29,
		Profesion: "Ingeniero",
		Peso:      80,
		Gustos: Preferencias{
			Comidas:   "asado, pollo",
			Peliculas: "coco",
			Animes:    "shinfeki no kyojin",
		},
	}

	fmt.Printf("El peso de %s es %.2f\n", p2.Nombre, p2.Peso)

	p2.Peso = 85

	fmt.Printf("El peso de %s es %.2f\n", p2.Nombre, p2.Peso)

	var p3 Persona
	p3.Nombre = "Ulises"
	p3.Edad = 15

	fmt.Printf("Persona 3:\n %+v\n", p3)

}
