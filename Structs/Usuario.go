package Structs

import (
	"errors"
	"fmt"
)

type Usuario struct {
	nombre string
	rutina *[]Rutina
	peso   uint
	altura uint
	puntos *Puntos
	edad   uint
}

func NewUsuario(nombre string, peso uint, altura uint, edad uint) *Usuario {
	return &Usuario{nombre, nil, peso, altura, nil, edad}
}

func (u *Usuario) ToString() string {
	return fmt.Sprintf("Nombre: ", u.nombre, "\n Peso: ", u.peso, "\n Altura: ", u.altura, "\n Edad: ", u.edad)
}

// Getters
func (u *Usuario) Get_Nome() string {
	return u.nombre
}
func (u *Usuario) Get_Peso() uint {
	return u.peso
}
func (u *Usuario) Get_Altura() uint {
	return u.altura
}
func (u *Usuario) Get_Edad() uint {
	return u.edad
}
func (u *Usuario) Get_Rutina() *[]Rutina {
	return u.rutina
}
func (u *Usuario) Get_Puntos() *Puntos {
	return u.puntos
}

// Setters
func (u *Usuario) Set_nombre(nombre string) error {
	if u.nombre != "" {
		u.nombre = nombre
		return nil
	}
	return errors.New("el nombre ingresado no es válido")
}
func (u *Usuario) Set_Peso(peso uint) {
	u.peso = peso
}
func (u *Usuario) Set_Altura(altura uint) {
	u.altura = altura
}
func (u *Usuario) Set_Edad(edad uint) {
	u.edad = edad
}
func (u *Usuario) Set_Rutina(rutina *[]Rutina) {
	u.rutina = rutina
}
func (u *Usuario) Set_Puntos(puntos *Puntos) {
	u.puntos = puntos
}
