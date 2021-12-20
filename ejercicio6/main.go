/*
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para
imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables
Nombre, Apellido, DNI, Fecha y que tenga un método detalle
*/

package main

import "fmt"

type Fecha struct {
	dia, mes, anio int
}

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      string
	FecIni   Fecha
}

func (a Alumno) details() {
	fmt.Println("Nombre: 	", a.Nombre)
	fmt.Println("Apellido:  ", a.Apellido)
	fmt.Println("DNI: 		", a.DNI)
	fmt.Println("Fecha: 	", a.FecIni)
}

func main() {

	alumno1 := Alumno{"Alumno1", "Meli1", "12312312", Fecha{13, 12, 2021}}

	alumno2 := Alumno{
		Nombre:   "Alumno2",
		Apellido: "Meli2",
		FecIni:   Fecha{1, 1, 2022},
	}

	fmt.Println("Alumno 1")
	alumno1.details()

	fmt.Println("Alumno 2")
	alumno2.details()

}
