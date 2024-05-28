package Structs

import (
	"errors"
	"fmt"
)

type Ejercicio struct {
	id               int
	titulo           string
	descripcion      string
	duracion         uint //en minutos
	puntos           *Puntos
	dificultad       string //media,baja,alta
	caloriasQuemadas uint   // en cal
	etiquetas        string
}

func NewEjercicio() *Ejercicio {
	return &Ejercicio{
		puntos: NewPunto(0, 0, 0, 0), // inicializo el map de etiquetas Usando el ID como clave
	}
}

func (e *Ejercicio) Get_id() int {
	return e.id
}

func (e *Ejercicio) Set_id(id int) {
	e.id = id
}

func (e *Ejercicio) Get_titulo() string {
	return e.titulo
}

func (e *Ejercicio) Set_titulo(titulo string) error {
	if titulo != "" {
		e.titulo = titulo
		return nil
	}
	return errors.New("el título ingresado no es válido")
}

func (e *Ejercicio) Get_descripcion() string {
	return e.descripcion
}

func (e *Ejercicio) Set_descripcion(descripcion string) error {
	if descripcion != "" {
		e.descripcion = descripcion
		return nil
	}
	return errors.New("la descripcion ingresada no es válida")
}

func (e *Ejercicio) Get_duracion() uint {
	return e.duracion
}

func (e *Ejercicio) Set_duracion(duracion uint) error {
	if duracion != 0 {
		e.duracion = duracion
		return nil
	}
	return errors.New("la duración ingresada no es válida")
}

func (e *Ejercicio) Get_puntos() Puntos {
	return *e.puntos
}

func (e *Ejercicio) Set_Puntos(puntos *Puntos) error {
	if puntos != nil {
		e.puntos = puntos
		return nil
	}
	return errors.New("los puntos ingresados no son válidos")
}

func (e *Ejercicio) Get_dificultad() string {
	return e.dificultad
}

func (e *Ejercicio) Set_dificultad(dificultad string) error {
	if dificultad == "baja" || dificultad == "media" || dificultad == "alta" {
		e.dificultad = dificultad
		return nil
	}
	return errors.New("la descripcion ingresada no es válida")
}

func (e *Ejercicio) Get_calorias_quemadas() uint {
	return e.caloriasQuemadas
}

func (e *Ejercicio) Set_calorias_quemadas(calorias_quemadas uint) error {
	if calorias_quemadas != 0 {
		e.caloriasQuemadas = calorias_quemadas
		return nil
	}
	return errors.New("las calorias ingresadas no son válidas")
}

func (e *Ejercicio) Set_etiquetas() {
	etiqueta, _ := e.Calcular_etiqueta()
	e.etiquetas = etiqueta
}

// Brinda las etiquetas del ejercicio
func (e *Ejercicio) Get_etiquetas() string {
	return e.etiquetas
}

// calcula en base a los puntajes de Puntos cual es la fortaleza del ejercicio
func (e *Ejercicio) Calcular_etiqueta() (string, error) {
	if e.puntos == nil {
		return "No hay puntos definidos", errors.New("no hay puntos definidos")
	}
	var mayorEtiqueta string
	var mayorPuntaje uint
	for etiqueta, valor := range e.puntos.Listar_puntos() {
		if valor > mayorPuntaje {
			mayorPuntaje = valor
			mayorEtiqueta = etiqueta
		}
	}
	return mayorEtiqueta, nil
}

func (e *Ejercicio) Set_fuerza(value uint) {
	e.puntos.fuerza = value
}

func (e *Ejercicio) Set_cardio(value uint) {
	e.puntos.cardio = value
}

func (e *Ejercicio) Set_flexibilidad(value uint) {
	e.puntos.flexibilidad = value
}

func (e *Ejercicio) Set_resistencia(value uint) {
	e.puntos.resistencia = value
}

func (e *Ejercicio) ImprimirPuntosEnLinea() string {
	return fmt.Sprintf("Cardio:$%d$", e.puntos.cardio) +
		fmt.Sprintf("Flexibilidad:$%d$", e.puntos.flexibilidad) +
		fmt.Sprintf("Resistencia:$%d$", e.puntos.resistencia) +
		fmt.Sprintf("Fuerza:$%d$", e.puntos.fuerza)

}

func (e *Ejercicio) String() string {
	return fmt.Sprintf("Titulo:	%s \n", e.titulo) +
		fmt.Sprintf("Descripcion: %s \n ", e.descripcion) +
		fmt.Sprintf("Duracion: %d \n ", e.duracion) +
		fmt.Sprintf("Dificultad: %s \n ", e.dificultad) +
		fmt.Sprintf("Calorias quemadas %d \n", e.caloriasQuemadas) +
		fmt.Sprintf("Etiquetas: %v \n ", (e.etiquetas)) +
		fmt.Sprintf("Puntos: %d \n ", e.puntos)
}
