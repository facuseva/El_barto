package Almacenamiento

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/untref-ayp2/TP-2024-el_barto/Structs"
)

func LeerEjecicios() ([][]string, error) {
	f, err := os.Open("Almacenamiento/ArchivosCSV/ejercicios.csv")
	if err != nil {
		fmt.Println("error abriendo el archivo: ", err)
	}
	defer f.Close() //cierro el archivo
	r := csv.NewReader(f)
	r.Comma = ','         //indico que esta separado por comas
	r.Comment = '#'       //indico los comentarios
	r.FieldsPerRecord = 8 //indico las columnas

	//el read devuelve un ->[][]string

	rawData, _ := r.ReadAll()

	//var rawData [][]string
	return rawData, nil
}

func GuardarEjerciciosCSV(data [][]string) error {
	// Crea un nuevo archivo CSV
	filePath := "Almacenamiento/ArchivosCSV/ejercicios.csv"

	// Abrir el archivo CSV en modo de lectura
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Error al abrir el archivo: %v", err)
	}
	defer file.Close()

	// Crear un lector CSV para leer el archivo
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error al leer el archivo CSV: %v", err)
	}

	// Agregar los nuevos datos a la matriz de líneas

	lines = append(lines, data...)

	// Cerrar el archivo antes de escribir
	//file.Close()

	// Abrir el archivo CSV en modo de escritura para escribir los datos actualizados
	file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Error al abrir el archivo: %v", err)
	}
	defer file.Close()

	// Crear un escritor CSV para escribir en el archivo
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escribir todas las líneas de nuevo en el archivo, incluida la nueva fila
	for _, fila := range lines {
		if err := writer.Write(fila); err != nil {
			log.Fatalf("Error al escribir en el archivo CSV: %v", err)
		}
	}

	return nil
}
func NuevoEjercicio(data [][]string) error {
	// Crea un nuevo archivo CSV

	// Abrir el archivo CSV en modo lectura-escritura
	filePath := "Almacenamiento/ArchivosCSV/ejercicios.csv"
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Error al abrir el archivo: %v", err)
	}
	defer file.Close()

	// Crear un lector CSV para leer el archivo
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error al leer el archivo CSV: %v", err)
	}

	// Modificar la línea deseada (por ejemplo, la línea 2)
	lineIndex := 0 // Índice 1 para la segunda línea (los índices comienzan desde 0)
	if lineIndex >= len(lines) {
		log.Fatalf("Índice de línea fuera de rango")
	}

	// Modificar la línea específica
	lines[lineIndex] = data[0] // Modificar el primer campo de la línea

	// Volver al inicio del archivo para escribir los cambios
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatalf("Error al volver al inicio del archivo: %v", err)
	}

	// Crear un escritor CSV para escribir en el archivo
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escribir las líneas modificadas en el archivo
	err = writer.WriteAll(lines)
	if err != nil {
		log.Fatalf("Error al escribir en el archivo CSV: %v", err)
	}
	return nil
}

func ModificarEjercicioEnCsv(id int, e Structs.Ejercicio) {
	lineas, _ := LeerEjecicios()
	var aux [][]string
	temple := [][]string{{e.Get_titulo(), e.Get_descripcion(), e.ImprimirPuntosEnLinea(), e.Get_dificultad(), strconv.Itoa(int(e.Get_calorias_quemadas())), e.Get_etiquetas(), strconv.Itoa(int(e.Get_duracion())), strconv.Itoa(e.Get_id())}}
	aux = append(aux, lineas[:id]...)
	aux = append(aux, temple[0])
	if len(lineas) > id {
		aux = append(aux, lineas[id+1:]...)
	}
	filePath := "Almacenamiento/ArchivosCSV/ejercicios.csv"

	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Error al abrir el archivo: %v", err)
	}
	defer file.Close()

	// Crear un escritor CSV para escribir en el archivo
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escribir las líneas modificadas en el archivo
	err = writer.WriteAll(aux)
	if err != nil {
		log.Fatalf("Error al escribir en el archivo CSV: %v", err)
	}

}

func EliminarEjercicioEnCsv(id int, e Structs.Ejercicio) {
	lineas, _ := LeerEjecicios()
	var aux [][]string

	aux = append(aux, lineas[:id]...)
	aux = append(aux, lineas[id+1:]...)

	filePath := "Almacenamiento/ArchivosCSV/ejercicios.csv"

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Error al abrir el archivo: %v\n", err)
		return
	}
	defer file.Close()

	// Truncar el archivo CSV (borra todo su contenido)
	err = file.Truncate(0)
	if err != nil {
		fmt.Printf("Error al truncar el archivo: %v\n", err)
		return
	}

	file, err = os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Error al abrir el archivo: %v", err)
	}
	defer file.Close()

	// Crear un escritor CSV para escribir en el archivo
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escribir las líneas modificadas en el archivo
	err = writer.WriteAll(aux)
	if err != nil {
		log.Fatalf("Error al escribir en el archivo CSV: %v", err)
	}

}
