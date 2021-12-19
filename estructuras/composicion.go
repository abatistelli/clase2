package main

import "fmt"

type Vehiculo struct {
	km     float64
	tiempo float64
}

func (v Vehiculo) detalle() {
	fmt.Printf("Km: %.2f\nTiempo: %.2f\n", v.km, v.tiempo)
}

type Auto struct {
	v Vehiculo
}

func (a *Auto) correr(minutos float64) {
	a.v.tiempo = minutos / 60
	a.v.km = a.v.tiempo * 100
}

func (a *Auto) detalle() {
	fmt.Println("\nV:\tAuto")
	a.v.detalle()
}

type Moto struct {
	v Vehiculo
}

func (m *Moto) correr(minutos float64) {
	m.v.tiempo = minutos / 60
	m.v.km = m.v.tiempo * 80
}

func (m *Moto) detalle() {
	fmt.Println("\nV:\tMoto")
	m.v.detalle()
}

func main() {
	auto := Auto{}
	auto.correr(360)
	auto.detalle()

	moto := Moto{}
	moto.correr(360)
	moto.detalle()
}
