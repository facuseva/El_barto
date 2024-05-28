package Structs

import "fmt"

type Puntos struct {
	cardio       uint
	flexibilidad uint
	resistencia  uint
	fuerza       uint
}

func NewPunto(cardio uint, flex uint, resis uint, fuerza uint) *Puntos {
	return &Puntos{cardio: cardio, flexibilidad: flex, resistencia: resis, fuerza: fuerza}
}
func (p *Puntos) Get_cardio() uint {
	return p.cardio
}
func (p *Puntos) Set_cardio(value uint) {
	p.cardio = value
}

func (p *Puntos) Get_flexibilidad() uint {
	return p.flexibilidad
}
func (p *Puntos) Set_flexibilidad(value uint) {
	p.flexibilidad = value
}
func (p *Puntos) Get_resistencia() uint {
	return p.resistencia
}
func (p *Puntos) Set_resistencia(value uint) {
	p.resistencia = value
}
func (p *Puntos) Get_fuerza() uint {
	return p.fuerza
}
func (p *Puntos) Set_fuerza(value uint) {
	p.fuerza = value
}

func (p Puntos) String() string {
	return fmt.Sprintf("\n\tCardio: %d \n", p.cardio) +
		fmt.Sprintf("\tFlexibilidad: %d \n", p.flexibilidad) +
		fmt.Sprintf("\tResistencia: %d \n", p.resistencia) +
		fmt.Sprintf("\tFuerza: %d", p.fuerza)
}

func (p *Puntos) ImprimirEnLinea() string {
	return fmt.Sprintf("Cardio:$%d$", p.cardio) +
		fmt.Sprintf("Flexibilidad:$%d$", p.flexibilidad) +
		fmt.Sprintf("Resistencia:$%d$", p.resistencia) +
		fmt.Sprintf("Fuerza:$%d$", p.fuerza)

}

func (p *Puntos) Listar_puntos() map[string]uint {
	puntos := map[string]uint{
		"Cardio":       p.cardio,
		"Flexibilidad": p.flexibilidad,
		"Resistencia":  p.resistencia,
		"Fuerza":       p.fuerza,
	}
	return puntos
}
