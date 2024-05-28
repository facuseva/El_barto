package main

import (
	"fmt"
	"strconv"

	"github.com/untref-ayp2/TP-2024-el_barto/Functions"
	"github.com/untref-ayp2/TP-2024-el_barto/Interfaz"
	"github.com/untref-ayp2/TP-2024-el_barto/Structs"
)

func main() {
	data, err := Functions.Iniciar() // Inicia el sistema cargando todos los datos de todos los archivos en la variable data
	system := Structs.NewSystem()
	if err != nil {
		fmt.Println(err)
	} else {
		if len(data[3]) == 0 {
			// se crea un usuario nuevo para gestionar
			system.Set_Usuario(Interfaz.CrearUsuario())
		} else {
			// se carga un usuario existente para gestionar
			Nombre := data[3][0][0]
			aux, _ := strconv.Atoi(data[3][0][3])
			Edad := int(aux)
			Altura, _ := strconv.Atoi(data[3][0][2])
			Peso, _ := strconv.Atoi(data[3][0][1])
			system.Set_Usuario(Interfaz.CrearUsuarioParametros(Nombre, uint(Edad), uint(Altura), uint(Peso)))
		}
		on := true
		for on {
			Functions.ClearScreen()
			system.Set_Ejercicios(Functions.ObtenerEjercicios(data[0]))
			system.Set_Rutinas(Functions.ObtenerRutinas(data[2], system))
			fmt.Println("=====================================")
			fmt.Println("Bienvenido ", system.Get_Usuario().Get_Nome(), " al sistema de ejercicios")
			fmt.Println("=====================================")
			fmt.Println("1. Gestionar usuario")
			fmt.Println("2. Gestionar rutina")
			fmt.Println("3. Gestionar ejercicios")
			fmt.Println("4. Salir")
			fmt.Println("=====================================")
			fmt.Print("Ingrese una opcion: ")
			var opcion int
			fmt.Scanf("%d ", &opcion)
			switch opcion {
			case 1:
				Interfaz.GestionarUsuario(system)
			case 2:
				Interfaz.GestionarRutinas(system)
			case 3:
				Interfaz.GestionarEjercicios(system)
			case 4:
				on = false
			default:
				fmt.Printf("Error: Opcion %d no reconocida\n", opcion)
			}
			Functions.ClearScreen()
		}

	}

}
