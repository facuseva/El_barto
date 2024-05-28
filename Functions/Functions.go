package Functions

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/untref-ayp2/TP-2024-el_barto/Almacenamiento"
	"github.com/untref-ayp2/TP-2024-el_barto/Structs"
	st "github.com/untref-ayp2/TP-2024-el_barto/Structs"
)

// Iniciar e importar variables y listas necesarias para el funcionamiento del sistema
func Iniciar() ([4][][]string, error) {
	aux := [4][][]string{}
	flag4, err4 := Almacenamiento.LeerUsuarios()
	if err4 != nil {
		e := errors.New("error al obtener los datos del archivo de usuarios")
		return aux, e
	}
	flag1, err1 := Almacenamiento.LeerEjecicios()
	if err1 != nil {
		e := errors.New("error al obtener los datos del archivo de ejercicios")
		return aux, e
	}
	flag2 := [][]string{}
	flag3, err3 := Almacenamiento.LeerRutinas()
	if err3 != nil {
		e := errors.New("error al obtener los datos del archivo de rutinas")
		return aux, e
	}
	data := [4][][]string{flag1, flag2, flag3, flag4}
	return data, nil
}

// Obtencion de listados guardados en csv
func ObtenerEjercicios(ejercicios [][]string) *[]st.Ejercicio {
	result := []st.Ejercicio{}
	aux := st.NewEjercicio()
	for i := 0; i < len(ejercicios); i++ {
		aux.Set_titulo(ejercicios[i][0])
		aux.Set_descripcion(ejercicios[i][1])
		aux.Set_dificultad(ejercicios[i][3])
		auxiliar, _ := strconv.ParseInt(ejercicios[i][6], 10, 64)
		aux.Set_duracion(uint(auxiliar))
		auxiliar, _ = strconv.ParseInt(ejercicios[i][4], 10, 64)
		aux.Set_calorias_quemadas(uint(auxiliar))
		aux.Set_Puntos(ObtenerPuntos(ejercicios[i][2]))
		auxiliar, _ = strconv.ParseInt(ejercicios[i][7], 10, 64)
		aux.Set_id(int(auxiliar))
		aux.Set_etiquetas()
		result = append(result, *aux)
	}
	return &result
}

// Obtencion de listados guardados en csv
func ObtenerRutinas(rutinas [][]string, sys *Structs.System) *[]st.Rutina {
	result := []st.Rutina{}
	for i := 0; i < len(rutinas); i++ {
		strsplit := strings.Split(rutinas[i][6], "$")
		var ejercicios []st.Ejercicio
		for j := 0; j < len(strsplit); j++ {

			auxiliar, _ := strconv.Atoi(strsplit[j])
			ejercicio := BuscarEjercicioID(auxiliar, sys)
			if ejercicio != nil {
				ejercicios = append(ejercicios, *ejercicio)
			}

		}
		aux := st.NewRutinaConParametros(ejercicios...)
		aux.Set_titulo(rutinas[i][0])
		aux.Set_descripcion(rutinas[i][1])
		obtenerDuracion, _ := strconv.Atoi(rutinas[i][2])
		aux.Set_duracion(uint(obtenerDuracion))
		aux.Set_dificultad(rutinas[i][3])
		aux.ContarCalorias()
		aux.CalcularPuntos()
		//aux.Set_etiquetas()
		sys.AgregarRutina(aux)
		result = append(result, *aux)
	}
	return &result
}

func ObtenerPuntos(texto string) *st.Puntos {
	puntosSt := strings.Split(texto, "$")
	cardio, _ := strconv.Atoi(puntosSt[1])
	flexividad, _ := strconv.Atoi(puntosSt[3])
	resistencia, _ := strconv.Atoi(puntosSt[5])
	fuerza, _ := strconv.Atoi(puntosSt[7])
	puntos := st.NewPunto(uint(cardio), uint(flexividad), uint(resistencia), uint(fuerza))
	return puntos
}

// Busqueda
func BuscarEjercicioID(id int, sys *st.System) *st.Ejercicio {
	aux := *sys.Get_Ejercicios()
	for i, ej := range aux {
		if ej.Get_id() == id {
			return &aux[i]
		}
	}
	return nil
}

// Busca un ejercicio por el titulo
func BuscarEjercicioTitulo(titulo string, sys *st.System) []*st.Ejercicio {
	var resultados []*st.Ejercicio
	aux := *sys.Get_Ejercicios()
	for i := 0; i < len(aux); i++ {
		if strings.Contains(aux[i].Get_titulo(), titulo) {
			resultados = append(resultados, &aux[i])
		} else if strings.Contains(strings.ToUpper(aux[i].Get_titulo()), strings.ToUpper(titulo)) {
			resultados = append(resultados, &aux[i])
		} else if strings.Contains(strings.ToLower(aux[i].Get_titulo()), strings.ToLower(titulo)) {
			resultados = append(resultados, &aux[i])
		}
	}
	if len(resultados) == 0 {
		return nil
	}
	return resultados
}

func Add_Rutina(sys *Structs.System, rutina *Structs.Rutina) {
	sys.AgregarRutina(rutina)
}

func Scanner() string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		return input
	}
	return ""
}

//Miscellaneous:

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func WaitForKeyPress() {
	var input string
	if runtime.GOOS == "windows" {
		// Limpiar la pantalla en Windows
		cmd := exec.Command("cmd", "/c", "pause")
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Run()
	} else {
		// Leer entrada en otros sistemas operativos
		fmt.Scanln(&input)
	}
}

func BuscarRutinaID(id int, sys *st.System) *st.Rutina {
	aux := *sys.Get_Rutinas()
	return &aux[id]
}

func BuscarRutinaTitulo(titulo string, sys *st.System) []*st.Rutina {
	var resultados []*st.Rutina
	aux := *sys.Get_Rutinas()
	for i := 0; i < len(aux); i++ {
		if strings.Contains(aux[i].Get_titulo(), titulo) {
			resultados = append(resultados, &aux[i])
		} else if strings.Contains(strings.ToUpper(aux[i].Get_titulo()), strings.ToUpper(titulo)) {
			resultados = append(resultados, &aux[i])
		} else if strings.Contains(strings.ToLower(aux[i].Get_titulo()), strings.ToLower(titulo)) {
			resultados = append(resultados, &aux[i])
		}
	}
	if len(resultados) == 0 {
		return nil
	}
	return resultados
}

func OrdenarEjerciciosPorTiempo(sys *st.System) *[]st.Ejercicio {

	ejercicios := sys.Get_Ejercicios()
	for i := 0; i < len(*ejercicios)-1; i++ {
		for j := 0; j < len(*ejercicios)-i-1; j++ {
			if (*ejercicios)[j].Get_duracion() > (*ejercicios)[j+1].Get_duracion() {
				(*ejercicios)[j], (*ejercicios)[j+1] = (*ejercicios)[j+1], (*ejercicios)[j]

			}
		}

	}
	return ejercicios
}
