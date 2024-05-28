package Interfaz

import (
	"fmt"
	"strconv"

	"github.com/untref-ayp2/TP-2024-el_barto/Almacenamiento"
	"github.com/untref-ayp2/TP-2024-el_barto/Functions"
	"github.com/untref-ayp2/TP-2024-el_barto/Structs"
)

// ABMDL RUTINAS
func GestionarRutinas(sys *Structs.System) {
	ok := false
	for !ok {
		//Functions.ClearScreen()
		fmt.Println("=====================================")
		fmt.Println("    Gestionar Rutinas")
		fmt.Println("=====================================")
		fmt.Println("1. Crear Rutina")
		fmt.Println("2. Eliminar Rutina")
		fmt.Println("3. Modificar Rutina")
		fmt.Println("4. Listar Rutinas")
		fmt.Println("5. Buscar rutina")
		fmt.Println("6. Crear rutina auromatica1 ")
		fmt.Println("7. Guardar y volver")
		fmt.Println("=====================================")
		fmt.Print("Ingrese una opcion: ")
		var opcion string
		fmt.Scanf("%s ", &opcion)
		switch opcion {
		case "1":
			CrearRutina(sys)
		case "2":
			EliminarRutina(sys)
		case "3":
			okk := false
			for !okk {

				Functions.ClearScreen()
				ListarRutinas(sys, 0)
				op := Functions.Scanner()
				fmt.Println("op: ", op)
				opp, _ := strconv.Atoi(op)
				if opp >= 0 {
					ModificarRutina(Functions.BuscarRutinaID(opp, sys), sys)
					okk = true
				}
			}
		case "4":
			Functions.ClearScreen()
			ListarRutinas(sys, 0)
			Functions.WaitForKeyPress()
		case "5":
			Functions.ClearScreen()
			BuscarRutina(sys)
		case "6":
			Functions.ClearScreen()
			CrearRutinaAutomatica(sys)
		case "7":
			ok = true
		default:
			fmt.Println("Opcion invalida")
		}
	}
}
func CrearRutina(sys *Structs.System) {
	ok := false
	okk := false
	okkk := false
	var titulo string
	var ejercicios []Structs.Ejercicio
	var descripcion string
	for !ok {
		Functions.ClearScreen()
		fmt.Println("=====================================")
		fmt.Println("	Nueva Rutina")
		fmt.Println("=====================================")
		fmt.Println("Ingrese el titulo:")
		fmt.Scanf("%s ", &titulo)
		if titulo != "" {
			ok = true
		}
	}
	for !okk {
		Functions.ClearScreen()
		fmt.Println("=====================================")
		fmt.Println("	Nueva Rutina")
		fmt.Println("\nTitulo: ", titulo)
		fmt.Println("=====================================")
		fmt.Println("Ingrese la descripcion:")
		descripcion = Functions.Scanner()
		if descripcion != "" {
			okk = true
		}
	}
	for !okkk {
		Functions.ClearScreen()
		fmt.Println("=====================================")
		fmt.Println("	Nueva Rutina")
		fmt.Println("\nTitulo: ", titulo)
		fmt.Println("Descripcion: ", descripcion)
		fmt.Println("Ejercicios: ")
		listarEjercicios(&ejercicios, 0)
		fmt.Println("=====================================")
		fmt.Println("¿Desea ingresar un ejercicio? (Y/N)")
		var opcion string
		fmt.Scanf("%s ", &opcion)
		if opcion == "Y" || opcion == "y" {
			ListarEjercicios(sys, 0)
			var idEjercicio uint
			fmt.Println("Ingrese el id del ejercicio:")
			fmt.Scanf("%d ", &idEjercicio)
			if Functions.BuscarEjercicioID(int(idEjercicio), sys) != nil {
				//ejerciciosListado := *sys.Get_Ejercicios()
				ejercicios = append(ejercicios, *Functions.BuscarEjercicioID(int(idEjercicio), sys))
			} else {
				fmt.Println("El ejercicio no existe")
				Functions.WaitForKeyPress()
			}
		} else {

			okkk = true
		}
	}
	rutina := Structs.NewRutina()
	rutina.Set_titulo(titulo)
	rutina.Set_descripcion(descripcion)
	rutina.Set_ejercicios(&ejercicios)
	rutina.CalcularPuntos()
	rutina.ContarCalorias()
	rutina.CalcularDificultad()
	rutina.Calcular_etiqueta()
	Functions.Add_Rutina(sys, rutina)
	duracion := strconv.Itoa(int(rutina.CalcularDuracion()))
	calorias := strconv.Itoa(int(rutina.Get_calorias_quemadas()))
	temple := [][]string{{titulo, descripcion, duracion, rutina.Get_dificultad(), calorias, rutina.Get_etiquetas(), rutina.ImprimirIdEjerciciosEnLinea()}}
	fmt.Println(temple)
	Almacenamiento.GuardarRutinasCSV(temple)
	sys.AgregarRutina(rutina)

}
func CrearRutinaAutomatica(sys *Structs.System) {
	ok := false

	var titulo string
	var tiempo int
	var descripcion, tipoDeEj, difDeEj string

	for !ok {
		Functions.ClearScreen()
		fmt.Println("=====================================")
		fmt.Println("	Nueva Rutina")
		fmt.Println("=====================================")
		fmt.Println("Ingrese el titulo:")
		titulo = Functions.Scanner()
		if titulo != "" {
			ok = true
		}
	}
	ok = false
	for !ok {
		Functions.ClearScreen()
		fmt.Println("=====================================")
		fmt.Println("	Nueva Rutina")
		fmt.Println("\nTitulo: ", titulo)
		fmt.Println("=====================================")
		fmt.Println("Ingrese la descripcion:")
		descripcion = Functions.Scanner()
		if descripcion != "" {
			ok = true
		}
	}
	ok = false
	for !ok {
		fmt.Println("=====================================")
		fmt.Println("Ingrese el tiempo que quiere que dure la rutina")
		fmt.Scanf("%d ", &tiempo)
		fmt.Print(tiempo)
		if tiempo != 0 {
			ok = true
		}
	}
	ok = false

	for !ok {
		fmt.Println("=====================================")
		fmt.Println("Ingrese la dificultad de la rutina: ")
		fmt.Println("1. Baja")
		fmt.Println("2. Media")
		fmt.Println("3. Alta")
		opcion := Functions.Scanner()
		switch opcion {
		case "1":
			difDeEj = "baja"
			ok = true
		case "2":
			difDeEj = "media"
			ok = true
		case "3":
			difDeEj = "alta"
			ok = true
		default:
			fmt.Println("*************************************")
			fmt.Println("Por favor ingrese una opcion valida")
			fmt.Println("*************************************")
		}

	}

	ok = false
	for !ok {
		Functions.ClearScreen()
		fmt.Println("=====================================")
		fmt.Println("Ingrese el tipo de ejercicios que se agregara a la rutina ")
		tipoDeEj = Functions.Scanner()
		if tipoDeEj == "Fuerza" || tipoDeEj == "Cardio" || tipoDeEj == "Resistencia" || tipoDeEj == "" {
			ok = true
		}
	}
	var ejercicios []Structs.Ejercicio
	rutina := Structs.NewRutina()
	rutina.Set_titulo(titulo)
	rutina.Set_descripcion(descripcion)
	rutina.Set_ejercicios(&ejercicios)
	rutina.CreacionAutomagica1(sys.Get_Ejercicios(), tiempo, tipoDeEj, difDeEj)
	rutina.CalcularPuntos()
	rutina.ContarCalorias()
	rutina.CalcularDificultad()
	rutina.Calcular_etiqueta()
	Functions.Add_Rutina(sys, rutina)
	duracion := strconv.Itoa(int(rutina.CalcularDuracion()))
	calorias := strconv.Itoa(int(rutina.Get_calorias_quemadas()))
	temple := [][]string{{titulo, descripcion, duracion, rutina.Get_dificultad(), calorias, rutina.Get_etiquetas(), rutina.ImprimirIdEjerciciosEnLinea()}}
	fmt.Println(temple)
	Almacenamiento.GuardarRutinasCSV(temple)
	sys.AgregarRutina(rutina)

}
func EliminarRutina(sys *Structs.System) {
}
func ModificarRutina(rutina *Structs.Rutina, sys *Structs.System) {
	ok := false
	for !ok {
		fmt.Println("=====================================")
		fmt.Println("    Modificar Rutina: ")
		fmt.Println("=====================================")
		fmt.Println("Titulo: ", rutina.Get_titulo())
		fmt.Println("Dificultad: ", rutina.Get_dificultad())
		fmt.Println("Duracion: ", rutina.Get_duracion())
		fmt.Println("Calorias Quemadas: ", rutina.Get_calorias_quemadas())
		fmt.Println("Puntos: ", rutina.Get_puntos())
		fmt.Println("Ejercicios: ")
		listarEjercicios(rutina.Get_Ejercicios(), 0)
		fmt.Println("=====================================")
		fmt.Println("1. Modificar titulo")
		fmt.Println("2. Gestionar los ejercicios de la rutina")
		fmt.Println("3. Eliminar rutina")
		fmt.Println("5. Guardar y volver")
		fmt.Println("=====================================")
		var opcion int
		fmt.Scanf("%d ", &opcion)
		switch opcion {
		case 1:
			Functions.ClearScreen()
			fmt.Println("=====================================")
			fmt.Println("Ingresa el nuevo titulo:")
			fmt.Println("=====================================")
			titulo := Functions.Scanner()
			rutina.Set_titulo(titulo)
		case 2:
			GestionarEjerciciosDeRutina(rutina, sys)
		case 3:
			Functions.ClearScreen()
			fmt.Println("=====================================")
			fmt.Println("¿Seguro? (Y/N)")
			fmt.Println("=====================================")
			var opcion string
			fmt.Scanf("%s ", &opcion)
			if opcion == "Y" || opcion == "y" {
				//EliminarRutina()
				ok = true
			}
		case 5:
			ok = true
			for i, elemento := range *sys.Get_Rutinas() {
				if elemento.Get_titulo() == rutina.Get_titulo() {
					Almacenamiento.ModificarRutinaEnCsv(i, *rutina)
				}
			}

		default:
			fmt.Println("Opcion invalida")
		}
	}
}
func BuscarRutina(sys *Structs.System) {
	Functions.ClearScreen()
	fmt.Println("=====================================")
	fmt.Println("BUSCAR RUTINA:")
	fmt.Println("=====================================")
	fmt.Println("Buscar:")
	var opcion string
	fmt.Scanf("%s ", &opcion)
	aux := Functions.BuscarRutinaTitulo(opcion, sys)
	Functions.ClearScreen()
	if aux == nil {
		Functions.ClearScreen()
		fmt.Println("Ejercicio no encontrado")
		Functions.WaitForKeyPress()
	} else if len(aux) == 1 {
		ModificarRutina(aux[0], sys)
	} else {
		fmt.Println("=====================================")
		fmt.Println("SELECCIONE UNA OPCIÓN:")
		fmt.Println("=====================================")
		for i, rutinas := range aux {
			fmt.Printf("%d. %s\n", i, rutinas.Get_titulo())
		}
		var id int
		fmt.Scanf("%d ", &id)
		aux := Functions.BuscarRutinaID(id, sys)
		ModificarRutina(aux, sys)
	}
}
func ListarRutinas(sys *Structs.System, i int) {
	rutinas := sys.Get_Rutinas()
	if len(*rutinas) == 0 {
		fmt.Println("No hay rutinas cargadas")
		return
	}
	fmt.Println("=====================================")
	fmt.Println("Rutina Nro: ", i)
	fmt.Println("Rutina: ", (*rutinas)[i].Get_titulo())
	fmt.Println("Dificultad: ", (*rutinas)[i].Get_dificultad())
	fmt.Println("Duracion: ", (*rutinas)[i].Get_duracion())
	fmt.Println("Calorias Quemadas: ", (*rutinas)[i].Get_calorias_quemadas())
	fmt.Println("Puntos: ", (*rutinas)[i].Get_puntos())
	fmt.Println("Ejercicios: ")
	listarEjercicios((*rutinas)[i].Get_Ejercicios(), 0)
	fmt.Println("=====================================")
	if len(*rutinas) > i+1 {
		ListarRutinas(sys, i+1)
	}
}

// ABMDL EJERCICIOS
func GestionarEjerciciosDeRutina(rutina *Structs.Rutina, sys *Structs.System) {
	ok := false
	for !ok {
		Functions.ClearScreen()
		fmt.Println("================================")
		fmt.Println(" Gestionar ejercicios de rutina")
		fmt.Println("================================")
		listarEjercicios(rutina.Get_Ejercicios(), 0)
		fmt.Println("=====================================")
		fmt.Println("1. Agregar ejercicio")
		fmt.Println("2. Eliminar ejercicio")
		fmt.Println("3. Volver")
		fmt.Println("=====================================")
		op := Functions.Scanner()
		switch op {
		case "1":
			ListarEjercicios(sys, 0)
			var idEjercicio int
			fmt.Println("Ingrese el id del ejercicio:")
			fmt.Scanf("%d ", &idEjercicio)
			ejerciciosListado := *sys.Get_Ejercicios()
			if idEjercicio > len(ejerciciosListado) {
				fmt.Println("Ejercicio no encontrado")
				Functions.WaitForKeyPress()
			} else {
				for i, e := range ejerciciosListado {
					if e.Get_id() == idEjercicio {
						rutina.AgregarEjercicio(&ejerciciosListado[i])
						break
					}
				}
			}
		case "2":
			listarEjercicios(rutina.Get_Ejercicios(), 0)
			var idEjercicio int
			fmt.Println("Ingrese el id del ejercicio:")
			fmt.Scanf("%d ", &idEjercicio)
			ejerciciosListado := *rutina.Get_Ejercicios()
			if idEjercicio > len(ejerciciosListado) {
				fmt.Println("Ejercicio no encontrado")
				Functions.WaitForKeyPress()
			} else {
				rutina.EliminarEjercicioDeRutina(idEjercicio)
			}
		case "3":
			Functions.ClearScreen()
			ok = true
		default:
			fmt.Println("Opcion invalida")
			Functions.WaitForKeyPress()
			Functions.ClearScreen()
		}
	}
}

func ModificarEjercicio(ej *Structs.Ejercicio, sys *Structs.System) {
	ok := false
	for !ok {
		fmt.Println("=====================================")
		fmt.Println("    Modificar Ejercicio")
		fmt.Println("=====================================")
		fmt.Println("Titulo: ", ej.Get_titulo())
		fmt.Println("Descripcion: ", ej.Get_descripcion())
		fmt.Println("Duracion: ", ej.Get_duracion())
		fmt.Println("Puntos: ", ej.Get_puntos())
		fmt.Println("Dificutad: ", ej.Get_dificultad())
		fmt.Println("Calorias Quemadas: ", ej.Get_calorias_quemadas())
		fmt.Println("Etiquetas: ", ej.Get_etiquetas())
		fmt.Println("=====================================")
		fmt.Println("1. Modificar titulo")
		fmt.Println("2. Modificar descripcion")
		fmt.Println("3. Modificar duracion")
		fmt.Println("4. Modificar puntos")
		fmt.Println("5. Modificar dificultad")
		fmt.Println("6. Modificar calorias quemadas")
		fmt.Println("7. Eliminar ejercicio") //a medias
		fmt.Println("8. Guardar y volver")
		fmt.Println("=====================================")
		var opcion int
		fmt.Scanf("%d ", &opcion)
		switch opcion {
		case 1:
			Functions.ClearScreen()
			fmt.Println("=====================================")
			fmt.Println("Ingresa el nuevo titulo:")
			fmt.Println("=====================================")
			var opcion string
			opcion = Functions.Scanner()
			ej.Set_titulo(opcion)
		case 2:
			Functions.ClearScreen()
			fmt.Println("=====================================")
			fmt.Println("Ingresa la nueva descripcion:")
			fmt.Println("=====================================")
			var opcion string
			opcion = Functions.Scanner()
			ej.Set_descripcion(opcion)
		case 3:
			Functions.ClearScreen()
			fmt.Println("=====================================")
			fmt.Println("Ingresa la nueva duracion (en min):")
			fmt.Println("=====================================")
			var opcion int
			fmt.Scanf("%d ", &opcion)
			if opcion >= 0 {
				ej.Set_duracion(uint(opcion))
			} else {
				fmt.Println("El valor ingesado no es correcto")
				Functions.WaitForKeyPress()
			}

		case 4:
			Functions.ClearScreen()
			fmt.Println("=====================================")
			fmt.Println("Elija el tipo de punto que quiere modificar y el nuevo valor:")
			fmt.Println("1. Cardio")
			fmt.Println("2. Flexibilidad")
			fmt.Println("3. Resistencia")
			fmt.Println("4. Fuerza")
			fmt.Println("=====================================")
			var opcion int
			var valor int
			fmt.Scanf("%d ", &opcion)

			fmt.Println("Indique el nuevo valor: ")
			fmt.Scanf("%d ", &valor)
			if valor >= 0 {

				fmt.Println("=====================================")
				switch opcion {
				case 1:
					ej.Set_cardio(uint(valor))
					fmt.Printf("Se actualizo el valor de cardio a " + fmt.Sprint(valor))
				case 2:
					ej.Set_flexibilidad(uint(valor))
					fmt.Printf("Se actualizo el valor de flexibilidad a " + fmt.Sprint(valor))
				case 3:
					ej.Set_resistencia(uint(valor))
					fmt.Printf("Se actualizo el valor de resistencia a " + fmt.Sprint(valor))
				case 4:
					ej.Set_fuerza(uint(valor))
					fmt.Printf("Se actualizo el valor de fuerza a " + fmt.Sprint(valor))
				default:
					{
						fmt.Printf("Opcion invalida\n")
					}
				}
			} else {
				fmt.Println("El valor ingesado no es correcto")
				Functions.WaitForKeyPress()
			}

			fmt.Println("=====================================")
			fmt.Println(ej.Get_puntos())
		case 5:
			Functions.ClearScreen()
			okk := false
			for !okk {
				fmt.Println("Ingrese la nueva dificultad del ejercicio: ")
				fmt.Println("1. Baja")
				fmt.Println("2. Media")
				fmt.Println("3. Alta")
				opcion := Functions.Scanner()
				switch opcion {
				case "1":
					ej.Set_dificultad("baja")
					okk = true
				case "2":
					ej.Set_dificultad("media")
					okk = true
				case "3":
					ej.Set_dificultad("alta")
					okk = true
				default:
					fmt.Println("*************************************")
					fmt.Println("Por favor ingrese una opcion valida")
					fmt.Println("*************************************")
				}
			}
		case 6:
			Functions.ClearScreen()
			fmt.Println("=====================================")
			fmt.Println("Ingresa la nueva cantidad de calorias que se queman:")
			fmt.Println("=====================================")
			var opcion int
			fmt.Scanf("%d ", &opcion)
			if opcion >= 0 {
				ej.Set_calorias_quemadas(uint(opcion))
			} else {
				fmt.Println("El valor ingesado no es correcto")
				Functions.WaitForKeyPress()
			}

		case 7:
			fmt.Println("=====================================")
			fmt.Println("¿Seguro? (Y/N)")
			fmt.Println("=====================================")
			var opcion string
			opcion = Functions.Scanner()
			if opcion == "Y" || opcion == "y" {
				for i, e := range *sys.Get_Ejercicios() {
					if e.Get_id() == ej.Get_id() {
						Almacenamiento.EliminarEjercicioEnCsv(i, *ej)
					}
				}
				ok = true
			}
		case 8:
			ok = true
			for i, e := range *sys.Get_Ejercicios() {
				if e.Get_id() == ej.Get_id() {
					Almacenamiento.ModificarEjercicioEnCsv(i, *ej)
				}
			}

		default:
			fmt.Println("Opcion invalida")
		}
	}
}
func GestionarEjercicios(sys *Structs.System) {
	Functions.ClearScreen()
	ok := false
	for !ok {
		fmt.Println("=====================================")
		fmt.Println("    Gestionar Ejercicios")
		fmt.Println("=====================================")
		fmt.Println("1. Modificar ejercicios")
		fmt.Println("2. Listar ejercicios")
		fmt.Println("3. Agregar ejercicio")
		fmt.Println("4. Buscar ejercicio por titulo")
		fmt.Println("5. Volver")
		fmt.Println("=====================================")
		var opcion int
		fmt.Scanf("%d ", &opcion)
		switch opcion {
		case 1:
			//Modificar ejercicios: Lista los ejercicios imprimiendolos
			//con un ID que es el indice en el arreglo
			Functions.ClearScreen()
			ListarEjercicios(sys, 0)
			fmt.Println("Seleccione la id del ejercicio que quiere modificar: ")
			var id int
			fmt.Scanf("%d ", &id)
			aux := Functions.BuscarEjercicioID(id, sys)
			if aux == nil {
				fmt.Println("El ejercicio no existe")
				Functions.WaitForKeyPress()
			} else {
				ModificarEjercicio(aux, sys)
				leer, _ := Almacenamiento.LeerEjecicios()
				sys.Set_Ejercicios(Functions.ObtenerEjercicios(leer))
				leer, _ = Almacenamiento.LeerRutinas()
				sys.Set_Rutinas(Functions.ObtenerRutinas(leer, sys))
			}
		case 2:
			Functions.ClearScreen()

			ListarEjercicios(sys, 0)
			Functions.WaitForKeyPress()
			Functions.ClearScreen()
		case 3:
			AgregarEjercicio(sys)
		case 4:
			BuscarEjercicio(sys)
		case 5:
			ok = true
		default:
			fmt.Println("Opcion invalida")
		}
	}
}
func AgregarEjercicio(sys *Structs.System) *Structs.Ejercicio {
	//ejercios
	var titulo, descripcion string
	var duracion, caloriasQuemadas int
	ejercicio := Structs.NewEjercicio()
	//puntos
	var cardio, flexibilidad, resistencia, fuerza int
	ok := false
	for !ok {
		Functions.ClearScreen()
		fmt.Println("=====================================")
		fmt.Println("	Nuevo Ejercicio")
		fmt.Println("=====================================")
		fmt.Println("Ingrese el nombre del ejercicio:")
		titulo = Functions.Scanner()
		if titulo != "" {
			ok = true
		}
	}
	ok = false
	for !ok {
		Functions.ClearScreen()
		fmt.Println("=====================================")
		fmt.Println("	Nuevo Ejercicio")
		fmt.Println("=====================================")
		fmt.Println("Nombre: ", titulo)
		fmt.Println("=====================================")
		fmt.Println("Ingrese descripcion del ejercicio:")
		descripcion = Functions.Scanner()
		if descripcion != "" {
			ok = true
		}
	}
	ok = false
	for !ok {
		Functions.ClearScreen()
		fmt.Println("=====================================")
		fmt.Println("	Nuevo Ejercicio")
		fmt.Println("=====================================")
		fmt.Println("Nombre: ", titulo)
		fmt.Println("Descripcion: ", descripcion)
		fmt.Println("=====================================")
		fmt.Println("Ingrese duracion del ejercicio:")
		fmt.Scanf("%d ", &duracion)
		if duracion >= 0 {
			ok = true
		}
	}
	ok = false
	for !ok {
		Functions.ClearScreen()
		fmt.Println("=====================================")
		fmt.Println("	Nuevo Ejercicio")
		fmt.Println("=====================================")
		fmt.Println("Nombre: ", titulo)
		fmt.Println("Descripcion: ", descripcion)
		fmt.Println("Duracion: ", duracion)
		fmt.Println("=====================================")
		fmt.Println("Ingrese la dificultad del ejercicio: (media,baja,alta)")
		okk := false
		for !okk {
			fmt.Println("Ingrese la nueva dificultad del ejercicio: ")
			fmt.Println("1. Baja")
			fmt.Println("2. Media")
			fmt.Println("3. Alta")
			opcion := Functions.Scanner()
			switch opcion {
			case "1":
				ejercicio.Set_dificultad("baja")
				okk = true
			case "2":
				ejercicio.Set_dificultad("media")
				okk = true
			case "3":
				ejercicio.Set_dificultad("alta")
				okk = true
			default:
				fmt.Println("*************************************")
				fmt.Println("Por favor ingrese una opcion valida")
				fmt.Println("*************************************")
			}
		}
		if ejercicio.Get_dificultad() != "" {
			ok = true
		}
	}
	ok = false
	for !ok {
		Functions.ClearScreen()
		fmt.Println("=====================================")
		fmt.Println("	Nuevo Ejercicio")
		fmt.Println("=====================================")
		fmt.Println("Nombre: ", titulo)
		fmt.Println("Descripcion: ", descripcion)
		fmt.Println("Duracion: ", duracion)
		fmt.Println("Dificultad: ", ejercicio.Get_dificultad())
		fmt.Println("Ingrese las calorias quemadas: ")
		fmt.Scanf("%d ", &caloriasQuemadas)
		if caloriasQuemadas >= 0 {
			ok = true
		}
	}
	ok = false
	for !ok {
		Functions.ClearScreen()
		fmt.Println("=====================================")
		fmt.Println("	Nuevo Ejercicio")
		fmt.Println("=====================================")
		fmt.Println("Nombre: ", titulo)
		fmt.Println("Descripcion: ", descripcion)
		fmt.Println("Duracion: ", duracion)
		fmt.Println("Dificultad: ", ejercicio.Get_dificultad())
		fmt.Println("Calorias: ", caloriasQuemadas)
		fmt.Println("Ingrese la cantidad de puntos de cardio del ejercicio:")
		fmt.Scanf("%d ", &cardio)
		if cardio >= 0 {
			ok = true
		}
	}
	ok = false
	for !ok {
		Functions.ClearScreen()
		fmt.Println("=====================================")
		fmt.Println("	Nuevo Ejercicio")
		fmt.Println("=====================================")
		fmt.Println("Nombre: ", titulo)
		fmt.Println("Descripcion: ", descripcion)
		fmt.Println("Duracion: ", duracion)
		fmt.Println("Dificultad: ", ejercicio.Get_dificultad())
		fmt.Println("Calorias: ", caloriasQuemadas)
		fmt.Println("Cardio: ", cardio)
		fmt.Println("=====================================")
		fmt.Println("Ingrese la cantidad de puntos de flexibilidad del ejercicio:")
		fmt.Scanf("%d ", &flexibilidad)
		if flexibilidad >= 0 {
			ok = true
		}
	}
	ok = false
	for !ok {
		Functions.ClearScreen()
		fmt.Println("=====================================")
		fmt.Println("	Nuevo Ejercicio")
		fmt.Println("=====================================")
		fmt.Println("Nombre: ", titulo)
		fmt.Println("Descripcion: ", descripcion)
		fmt.Println("Duracion: ", duracion)
		fmt.Println("Dificultad: ", ejercicio.Get_dificultad())
		fmt.Println("Calorias: ", caloriasQuemadas)
		fmt.Println("Cardio: ", cardio)
		fmt.Println("Flexibilidad: ", flexibilidad)
		fmt.Println("=====================================")
		fmt.Println("Ingrese la cantidad de puntos de resistencia del ejercicio:")
		fmt.Scanf("%d ", &resistencia)
		if resistencia >= 0 {
			ok = true
		}
	}
	ok = false
	for !ok {
		Functions.ClearScreen()
		fmt.Println("=====================================")
		fmt.Println("	Nuevo Ejercicio")
		fmt.Println("=====================================")
		fmt.Println("Nombre: ", titulo)
		fmt.Println("Descripcion: ", descripcion)
		fmt.Println("Duracion: ", duracion)
		fmt.Println("Dificultad: ", ejercicio.Get_dificultad())
		fmt.Println("Calorias: ", caloriasQuemadas)
		fmt.Println("Cardio: ", cardio)
		fmt.Println("Flexibilidad: ", flexibilidad)
		fmt.Println("Resistencia: ", resistencia)
		fmt.Println("=====================================")
		fmt.Println("Ingrese la cantidad de puntos de fuerza del ejercicio:")
		fmt.Scanf("%d ", &fuerza)
		if fuerza >= 0 {
			ok = true
		}
		fmt.Println("Fuerza: ", fuerza)
		fmt.Println("=====================================")
	}
	id := len(*sys.Get_Ejercicios())

	ejercicio.Set_id(len(*sys.Get_Ejercicios()))
	ejercicio.Set_titulo(titulo)
	ejercicio.Set_descripcion(descripcion)
	ejercicio.Set_duracion(uint(duracion))
	ejercicio.Set_calorias_quemadas(uint(caloriasQuemadas))
	ejercicio.Set_cardio(uint(cardio))
	ejercicio.Set_flexibilidad(uint(flexibilidad))
	ejercicio.Set_resistencia(uint(resistencia))
	ejercicio.Set_fuerza(uint(fuerza))
	//total_puntos := cardio + flexibilidad + resistencia + fuerza
	// seteo de etiquetas automatico y devolucion para escritura en el csv
	ejercicio.Set_etiquetas()
	etiquetas := ejercicio.Get_etiquetas()
	fmt.Println("Puntos : ", ejercicio.Get_puntos())
	temple := [][]string{{titulo, descripcion, ejercicio.ImprimirPuntosEnLinea(), ejercicio.Get_dificultad(), strconv.Itoa(caloriasQuemadas), etiquetas, strconv.Itoa(duracion), strconv.Itoa(id)}}
	Almacenamiento.GuardarEjerciciosCSV(temple)
	sys.AgregarEjercicio(ejercicio)
	return ejercicio
}

// listado de ejercicios en Rutinas
func listarEjercicios(ejercicios *[]Structs.Ejercicio, i int) {

	if len(*ejercicios) == 0 {
		return
	}
	fmt.Println("-------------------------------------")
	fmt.Println("ID: ", (*ejercicios)[i].Get_id())
	fmt.Println("Ejercicio: ", (*ejercicios)[i].Get_titulo())
	fmt.Println("Descripcion: ", (*ejercicios)[i].Get_descripcion())
	fmt.Println("Duracion: ", (*ejercicios)[i].Get_duracion())
	fmt.Println("Dificultad: ", (*ejercicios)[i].Get_dificultad())
	fmt.Println("-------------------------------------")
	if len(*ejercicios) > i+1 {
		listarEjercicios(ejercicios, i+1)
	}
}

// reorganizacion de pantalla descriptiva para que se parezca al orden de almacenamiento en csv
func ListarEjercicios(sys *Structs.System, i int) {
	ejercicios := sys.Get_Ejercicios()
	fmt.Println("=====================================")
	fmt.Println("ID: ", (*ejercicios)[i].Get_id())
	fmt.Println("Ejercicio: ", (*ejercicios)[i].Get_titulo())
	fmt.Println("Descripcion: ", (*ejercicios)[i].Get_descripcion())
	fmt.Println("Dificultad: ", (*ejercicios)[i].Get_dificultad())
	fmt.Println("Calorias quemadas: ", (*ejercicios)[i].Get_calorias_quemadas())
	fmt.Println("Etiquetas: ", (*ejercicios)[i].Get_etiquetas())
	fmt.Println("Duracion: ", (*ejercicios)[i].Get_duracion())
	fmt.Println("Puntos: ", (*ejercicios)[i].Get_puntos())
	fmt.Println("=====================================")
	if len(*ejercicios) > i+1 {
		ListarEjercicios(sys, i+1)
	}
}
func BuscarEjercicio(sys *Structs.System) {
	Functions.ClearScreen()
	fmt.Println("=====================================")
	fmt.Println("BUSCAR EJERCICIO:")
	fmt.Println("=====================================")
	fmt.Println("Buscar:")
	var opcion string
	fmt.Scanf("%s ", &opcion)
	aux := Functions.BuscarEjercicioTitulo(opcion, sys)
	Functions.ClearScreen()
	if aux == nil {
		Functions.ClearScreen()
		fmt.Println("Ejercicio no encontrado")
		Functions.WaitForKeyPress()
	} else if len(aux) == 1 {
		ModificarEjercicio(aux[0], sys)
	} else {
		fmt.Println("=====================================")
		fmt.Println("SELECCIONE UNA OPCIÓN:")
		fmt.Println("=====================================")
		for _, ejercicio := range aux {
			fmt.Printf("%d. %s\n", ejercicio.Get_id(), ejercicio.Get_titulo())
		}
		var id int
		fmt.Scanf("%d ", &id)
		aux := Functions.BuscarEjercicioID(id, sys)
		ModificarEjercicio(aux, sys)
	}
}

// ABMDL USUARIOS
func CrearUsuario() *Structs.Usuario {
	var nombre string
	var edad string
	var peso string
	var altura string
	ok := false
	okk := false
	okkk := false
	okkkk := false
	for !ok {
		Functions.ClearScreen()
		fmt.Println("=====================================")
		fmt.Println("	Nuevo Usuario")
		fmt.Println("=====================================")
		fmt.Println("Ingrese su nombre:")
		fmt.Scanf("%s ", &nombre)
		if nombre != "" {
			ok = true
		}
	}
	for !okk {
		Functions.ClearScreen()
		fmt.Println("=====================================")
		fmt.Println("	Nuevo Usuario")
		fmt.Println("=====================================")
		fmt.Println("Nombre: ", nombre)
		fmt.Println("=====================================")
		fmt.Println("Ingrese su edad:")
		fmt.Scanf("%s ", &edad)
		if edad != "" {
			okk = true
		}
	}
	for !okkk {
		Functions.ClearScreen()
		fmt.Println("=====================================")
		fmt.Println("	Nuevo Usuario")
		fmt.Println("=====================================")
		fmt.Println("Nombre: ", nombre)
		fmt.Println("Edad: ", edad)
		fmt.Println("=====================================")
		fmt.Println("Ingrese su altura (en cm):")
		fmt.Scanf("%s ", &altura)
		if altura != "" {
			okkk = true
		}
	}
	for !okkkk {
		Functions.ClearScreen()
		fmt.Println("=====================================")
		fmt.Println("	Nuevo Usuario")
		fmt.Println("=====================================")
		fmt.Println("Nombre: ", nombre)
		fmt.Println("Edad: ", edad)
		fmt.Println("Altura: ", altura)
		fmt.Println("=====================================")
		fmt.Println("Ingrese su peso (en kg redondeando):")
		fmt.Scanf("%s ", &peso)
		if peso != "" {
			okkkk = true
		}
	}
	weight, _ := strconv.Atoi(peso)
	height, _ := strconv.Atoi(altura)
	aux, _ := strconv.Atoi(edad)
	age := uint(aux)
	temple := [][]string{{nombre, peso, altura, edad}}
	Almacenamiento.NuevoUsuario(temple)
	return Structs.NewUsuario(nombre, uint(weight), uint(height), age)
}
func GestionarUsuario(sys *Structs.System) {

	ok := false
	for !ok {
		Functions.ClearScreen()
		fmt.Println("=====================================")
		fmt.Println("    Gestionar Usuario")
		fmt.Println("=====================================")
		fmt.Println("1. Modificar Usuario")
		fmt.Println("2. Eliminar Usuario")
		fmt.Println("4. Volver")
		fmt.Println("=====================================")
		var opcion int
		fmt.Scanf("%d ", &opcion)
		switch opcion {
		case 1:
			ModificarUsuario(sys)
		case 2:
			fmt.Println("=====================================")
			fmt.Println("¿Seguro? (Y/N)")
			fmt.Println("=====================================")
			var opcion string
			fmt.Scanf("%s ", &opcion)
			if opcion == "Y" || opcion == "y" {
				//EliminarUsuario()
				ok = true
			}
		case 3:
			CrearUsuario()
		case 4:
			ok = true
		default:
			fmt.Println("Opcion invalida")
		}
	}
}
func ModificarUsuario(sys *Structs.System) {

	ok := false
	for !ok {
		Functions.ClearScreen()
		fmt.Println("=====================================")
		fmt.Println("    Modificar Usuario")
		fmt.Println("=====================================")
		fmt.Println("Nombre: ", sys.Get_Usuario().Get_Nome())
		fmt.Println("Edad: ", sys.Get_Usuario().Get_Edad())
		fmt.Println("Altura: ", sys.Get_Usuario().Get_Altura())
		fmt.Println("Peso: ", sys.Get_Usuario().Get_Peso())
		fmt.Println("=====================================")
		fmt.Println("1. Modificar nombre")
		fmt.Println("2. Modificar edad")
		fmt.Println("3. Modificar altura")
		fmt.Println("4. Modificar peso")
		fmt.Println("5. Guardar y volver")
		fmt.Println("=====================================")
		var opcion int
		fmt.Scanf("%d ", &opcion)
		switch opcion {
		case 1:
			Functions.ClearScreen()
			fmt.Println("=====================================")
			fmt.Println("Ingresa el nuevo nombre:")
			fmt.Println("=====================================")
			var opcion string
			fmt.Scanf("%s ", &opcion)
			sys.Get_Usuario().Set_nombre(opcion)
		case 2:
			Functions.ClearScreen()
			fmt.Println("=====================================")
			fmt.Println("Ingresa la nueva edad:")
			fmt.Println("=====================================")
			var opcion uint
			fmt.Scanf("%d ", &opcion)
			sys.Get_Usuario().Set_Edad(opcion)
		case 3:
			Functions.ClearScreen()
			fmt.Println("=====================================")
			fmt.Println("Ingresa la nueva altura: (en cm)")
			fmt.Println("=====================================")
			var opcion uint
			fmt.Scanf("%d ", &opcion)
			sys.Get_Usuario().Set_Altura(opcion)
		case 4:
			Functions.ClearScreen()
			fmt.Println("=====================================")
			fmt.Println("Ingresa el nuevo peso:")
			fmt.Println("=====================================")
			var opcion uint
			fmt.Scanf("%d ", &opcion)
			sys.Get_Usuario().Set_Peso(opcion)
		case 5:
			ok = true
			temple := [][]string{{sys.Get_Usuario().Get_Nome(), strconv.Itoa(int(sys.Get_Usuario().Get_Peso())), strconv.Itoa(int(sys.Get_Usuario().Get_Altura())), strconv.Itoa(int(sys.Get_Usuario().Get_Edad()))}}
			Almacenamiento.GuardarUsuariosCSV(temple)
		default:
			fmt.Println("Opcion invalida")
		}
	}
}
func CrearUsuarioParametros(nombre string, edad uint, altura uint, peso uint) *Structs.Usuario {
	return Structs.NewUsuario(nombre, peso, altura, edad)
}
