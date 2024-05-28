package Almacenamiento

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func LeerUsuarios() ([][]string, error) {
	f, err := os.Open("Almacenamiento/ArchivosCSV/usuarios.csv")
	if err != nil {
		fmt.Println("error abriendo el archivo: ", err)
	}
	defer f.Close() //cierro el archivo
	r := csv.NewReader(f)
	r.Comma = ','         //indico que esta separado por comas
	r.Comment = '#'       //indico los comentarios
	r.FieldsPerRecord = 4 //indico las columnas

	//el read devuelve un ->[][]string

	rawData, _ := r.ReadAll()

	return rawData, nil
	//var rawData [][]string

}

func GuardarUsuariosCSV(data [][]string) error {
	// Crea un nuevo archivo CSV

	// Abrir el archivo CSV en modo lectura-escritura
	filePath := "Almacenamiento/ArchivosCSV/usuarios.csv"
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

func NuevoUsuario(data [][]string) error {
	// Ruta del archivo CSV
	filePath := "Almacenamiento/ArchivosCSV/usuarios.csv"

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
	file.Close()

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

func EliminarLineaEspecifica(filename string, lineToDelete int) error {
	// Abrir el archivo CSV existente
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Crear un lector CSV asociado al archivo
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// Verificar si la línea a eliminar está dentro del rango válido
	if lineToDelete < 0 || lineToDelete >= len(lines) {
		return fmt.Errorf("índice de línea fuera de rango")
	}

	// Eliminar la línea especificada
	lines = append(lines[:lineToDelete], lines[lineToDelete+1:]...)

	// Volver al principio del archivo
	file.Seek(0, 0)

	// Crear un escritor CSV asociado al archivo
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escribir las líneas actualizadas de vuelta al archivo CSV
	err = writer.WriteAll(lines)
	if err != nil {
		return err
	}

	return nil
}
