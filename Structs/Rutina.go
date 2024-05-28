package Structs

import (
	"errors"
	"fmt"
	"strconv"
)

type Rutina struct {
	titulo           string
	descripcion      string
	duracion         uint
	dificultad       string
	caloriasQuemadas uint
	etiquetas        string
	ejercicios       *[]Ejercicio
	puntos           *Puntos
}

func NewRutina() *Rutina {
	return &Rutina{}
}

func NewRutinaConParametros(ejercicio ...Ejercicio) *Rutina {
	rutina := NewRutina()
	if rutina.ejercicios == nil {
		var ejercicioss []Ejercicio
		ejercicioss = append(ejercicioss, ejercicio...)
		rutina.Set_ejercicios(&ejercicioss)
	} else {
		*rutina.ejercicios = append(*rutina.ejercicios, ejercicio...)
	}
	rutina.puntos = NewPunto(0, 0, 0, 0)
	return rutina
}

func (r *Rutina) Set_descripcion(desc string) {
	r.descripcion = desc
}

func (r *Rutina) Get_descripcion() string {
	return r.descripcion
}

func (r *Rutina) Get_ejercicios() *[]Ejercicio {
	return r.ejercicios
}

func (r *Rutina) Set_ejercicios(ejercicios *[]Ejercicio) {
	r.ejercicios = ejercicios
}

// Agrega varios ejercicios a la rutina
func (r *Rutina) AgregarEjercicios(ejercicios *[]Ejercicio) {
	for _, ejercicio := range *ejercicios {
		r.AgregarEjercicio(&ejercicio)
	}
}

func (r *Rutina) CreacionAutomagica1(ejercicios *[]Ejercicio, tiempo int, tipoEjercicio string, dificultadDeEj string) {
	for _, ejercicio := range *ejercicios {
		if tiempo >= int(ejercicio.Get_duracion()) && ejercicio.Get_etiquetas() == tipoEjercicio && ejercicio.Get_dificultad() == dificultadDeEj {

			r.AgregarEjercicio(&ejercicio)
			tiempo -= int(ejercicio.Get_duracion())
		}
	}
}

// Agrega UN ejercicio a la rutina
func (r *Rutina) AgregarEjercicio(ejercicio *Ejercicio) {
	*r.ejercicios = append(*r.ejercicios, *ejercicio)
}

func (r *Rutina) Get_titulo() string {
	return r.titulo
}

func (r *Rutina) Set_titulo(titulo string) error {
	if titulo != "" {
		r.titulo = titulo
		return nil
	}
	return errors.New("el titulo no es valido")
}
func (r *Rutina) Get_duracion() uint {
	return r.duracion
}

func (r *Rutina) Set_duracion(d uint) error {
	if d != 0 {
		r.duracion = d
	}
	return errors.New("El valor de la duracion no es valido")

}

// Calcula automatiamente la duracion sumando la duracion de los ejercicios en la rutina
func (r *Rutina) CalcularDuracion() uint {
	for _, ejercicio := range *r.ejercicios {
		if r.duracion != 0 {
			r.duracion += ejercicio.Get_duracion()
		} else {
			r.duracion = ejercicio.Get_duracion()
		}
	}
	return r.duracion
}

func (r *Rutina) Get_dificultad() string {
	return r.dificultad
}

func (r *Rutina) Set_dificultad(dificultad string) error {
	if dificultad != "" {
		r.dificultad = dificultad
		return nil
	}
	return errors.New("la difucultadad ingresada no es valida")
}

// Calcula la dificulta tomando la cantidad de veces que aparece la difucltad en los ejercicios
func (r *Rutina) CalcularDificultad() {
	var baja int
	var media int
	var alta int
	for _, ejercicio := range *r.ejercicios {
		if ejercicio.dificultad != "" {
			switch ejercicio.Get_dificultad() {
			case "baja":
				baja++
			case "media":
				media++
			case "alta":
				alta++
			}
		}

	}
	if baja > alta && baja > media {
		r.Set_dificultad("baja")
		return
	}
	if media > baja && media > alta {
		r.Set_dificultad("media")
		return
	}
	r.Set_dificultad("alta")

}

// Calcula las calorias quemadas de la rutina a partir de las calorias quemadas de cada ejercicio
func (r *Rutina) ContarCalorias() uint {
	r.caloriasQuemadas = 0
	for _, elemento := range *r.ejercicios {
		r.caloriasQuemadas += elemento.Get_calorias_quemadas()
	}
	return r.caloriasQuemadas
}

func (r *Rutina) Get_Ejercicios() *[]Ejercicio {
	return r.ejercicios
}

// Calcula los puntos de la rutina a partir de los puntos de los ejercicios
func (r *Rutina) CalcularPuntos() *Puntos {
	for _, valor := range *r.ejercicios {
		if r.Get_puntos() != nil {
			r.puntos.cardio += valor.puntos.cardio
			r.puntos.fuerza += valor.puntos.fuerza
			r.puntos.flexibilidad += valor.puntos.flexibilidad
			r.puntos.resistencia += valor.puntos.resistencia
		} else {
			r.puntos = NewPunto(valor.puntos.cardio, valor.puntos.flexibilidad, valor.puntos.resistencia, valor.puntos.fuerza)
		}

	}
	return r.puntos
}

func (r *Rutina) Get_puntos() *Puntos {
	return r.puntos
}

func (r *Rutina) Get_calorias_quemadas() uint {
	return r.caloriasQuemadas
}

func (r *Rutina) String() string {
	return fmt.Sprintf("Titulo:	%s \n", r.titulo) +
		fmt.Sprintf("Ejercicios: %v \n ", r.ejercicios) +
		fmt.Sprintf("Descripcion: %s \n ", r.descripcion) +
		fmt.Sprintf("Duracion: %d \n ", r.duracion) +
		fmt.Sprintf("Dificultad: %s \n ", r.dificultad) +
		fmt.Sprintf("Calorias quemadas %d \n", r.caloriasQuemadas) +
		fmt.Sprintf("Etiquetas: %v \n ", (r.etiquetas)) +
		fmt.Sprintf("Puntos: %d \n ", r.puntos)

}

// Elimina un ejercicio mediante el id pasado por parametro de la rutina
func (r *Rutina) EliminarEjercicioDeRutina(id int) {
	*r.ejercicios = append((*r.ejercicios)[:id], (*r.ejercicios)[id+1:]...)
}

// Forma de imprimir los id  de los ejercicios para que funcione el lector de csv
func (r *Rutina) ImprimirIdEjerciciosEnLinea() string {
	str := ""
	if r.ejercicios != nil {
		for i, ejercicio := range *r.ejercicios {

			if i == 0 {
				str += strconv.Itoa(ejercicio.Get_id())
			} else {
				str += "$" + strconv.Itoa(ejercicio.Get_id())
			}

		}
	}
	return str
}

func (r *Rutina) Get_etiquetas() string {
	return r.etiquetas
}

func (r *Rutina) Set_etiquetas(e string) {
	r.etiquetas = e
}

func (r *Rutina) Calcular_etiqueta() {
	var cardio int
	var flex int
	var fuerza int
	var res int
	for _, ejercicio := range *r.Get_Ejercicios() {
		switch ejercicio.Get_etiquetas() {
		case "Cardio":
			cardio++
		case "Fuerza":
			fuerza++
		case "Resistencia":
			res++
		case "Flexibilidad":
			flex++
		}

	}
	if flex > cardio && flex > fuerza && flex > res {
		r.Set_etiquetas("Flexibilidad")

	} else if cardio > flex && cardio > fuerza && cardio > res {
		r.Set_etiquetas("Cardio")
	} else if fuerza > cardio && fuerza > flex && fuerza > res {
		r.Set_etiquetas("Fuerza")
	} else {
		r.Set_etiquetas("Resistencia")
	}
}
