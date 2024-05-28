package Almacenamiento

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/untref-ayp2/TP-2024-el_barto/Structs"
)

func LeerRutinas() ([][]string, error) {
	f, err := os.Open("Almacenamiento/ArchivosCSV/rutinas.csv")
	if err != nil {
		fmt.Println("error abriendo el archivo: ", err)
	}
	defer f.Close() //cierro el archivo
	r := csv.NewReader(f)
	r.Comma = ','         //indico que esta separado por comas
	r.Comment = '#'       //indico los comentarios
	r.FieldsPerRecord = 7 //indico las columnas

	//el read devuelve un ->[][]string

	rawData, _ := r.ReadAll()

	//var rawData [][]string
	return rawData, nil
}

func GuardarRutinasCSV(data [][]string) error {
	// Crea un nuevo archivo CSV
	filePath := "Almacenamiento/ArchivosCSV/rutinas.csv"

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

func ModificarRutinaEnCsv(id int, r Structs.Rutina) {
	lineas, _ := LeerRutinas()
	var aux [][]string
	r.Calcular_etiqueta()
	temple := [][]string{{r.Get_titulo(), r.Get_descripcion(), strconv.Itoa(int(r.Get_duracion())), r.Get_dificultad(), strconv.Itoa(int(r.Get_calorias_quemadas())), r.Get_etiquetas(), r.ImprimirIdEjerciciosEnLinea()}}
	fmt.Println(temple)
	aux = append(aux, lineas[:id]...)
	aux = append(aux, temple[0])
	if len(lineas) > id {
		aux = append(aux, lineas[id+1:]...)
	}
	filePath := "Almacenamiento/ArchivosCSV/rutinas.csv"

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
