package Structs

import (
	"fmt"
)

type System struct {
	usuario *Usuario
	rutinas *[]Rutina

	ejercicios *[]Ejercicio
	puntos     *Puntos
}

func NewSystem() *System {
	return &System{}
}
func (s *System) ToString() string {
	aux := fmt.Sprintf("Usuario: ", s.usuario,
		"\nRutinas: ", s.rutinas,
		"\nEjercicios: ", s.ejercicios,
		"\nPuntos: ", s.puntos)
	return aux
}

func (s *System) Get_Usuario() *Usuario {
	return s.usuario
}

func (s *System) Get_Rutinas() *[]Rutina {
	return s.rutinas
}

func (s *System) Get_Ejercicios() *[]Ejercicio {
	return s.ejercicios
}

func (s *System) Get_Puntos() *Puntos {
	return s.puntos
}

func (s *System) Set_Usuario(usuario *Usuario) {
	s.usuario = usuario
}

func (s *System) Set_Rutinas(rutinas *[]Rutina) {
	s.rutinas = rutinas
}

func (s *System) Set_Ejercicios(ejercicios *[]Ejercicio) {
	s.ejercicios = ejercicios
}

func (s *System) Set_Puntos(puntos *Puntos) {
	s.puntos = puntos
}

func (s *System) AgregarEjercicio(ejercicio *Ejercicio) {
	*s.ejercicios = append(*s.ejercicios, *ejercicio)

}

func (s *System) AgregarRutina(rutina *Rutina) {
	if s.rutinas != nil {
		*s.rutinas = append(*s.rutinas, *rutina)
	} else {
		// SE HACE ASIIIIIII !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		var aux []Rutina
		aux = append(aux, *rutina)
		s.rutinas = &aux
	}
}
