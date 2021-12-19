/*

Dentro de nuestras estructuras podemos definir etiquetas o anotaciones que hagan referencia a cada uno
de los campos que aparecen luego de declarar el tipo de dato.
*/

package main

import (
	"encoding/json" // convierte una estructura a JSON
	"fmt"
	"reflect" // nos proporciona funcionalidades para poder obtener info de los obj en tiempo de ejecucion
)

/*
type Persona struct {
	PrimerNombre string `json:"primer_nombre"`
	Apellido     string `json:"apellido"`
	Telefono     string `json:"telefono"`
	Direccion    string `json:"direccion"`
}
*/

type Persona struct {
	PrimerNombre string `bd:"primer_nombre"`
	Apellido     string `bd:"apellido"`
	Telefono     string `bd:"telefono"`
	Direccion    string `bd:"direccion"`
}

func main() {

	p := Persona{PrimerNombre: "Agustin", Apellido: "Batistelli", Telefono: "12345678", Direccion: "Calle Falsa 123"}

	// funcion Marshal, nos devuelve los bytes de la representacion de nuestra estructura codificada en
	// json y un error en caso que haya habido algun problema al realizar la conversion

	miJSON, err := json.Marshal(p)

	fmt.Println(string(miJSON))
	fmt.Println(err)

	persona := Persona{}
	p2 := reflect.TypeOf(persona)

	fmt.Println("Type: ", p2.Name())
	fmt.Println("Kind: ", p2.Kind())

	// NumField: devuelve el numero de campos de nuestra estructura
	// Field: devuelve el campo de nuestra estructura, pasandole como parametro el indice

	for i := 0; i < p2.NumField(); i++ {
		field := p2.Field(i)
		tag := field.Tag.Get("bd") // Si queremos acceder al valod de la etiqueta definida

		fmt.Println("FIELD: ", field)
		fmt.Println("TAG: ", tag)
	}
}
