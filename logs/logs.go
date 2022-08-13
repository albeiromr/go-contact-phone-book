package logs

import (
	"fmt"
	"os"
	"path"
	"log"
)

//crea  el archivo de logs en el sistema de archivos
func CreateLogsFile() *os.File {
	actualPath, _ := os.Getwd()
	logFile := path.Join(actualPath, "/logs/logs.log")

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}

	return file
}

//registra un nuevo log en el archivo de logs
func CreateLog(id string){
	file := CreateLogsFile()
	showError := log.New(file, "showError: ", log.Ldate|log.Ltime|log.Llongfile)
	showError.Printf("el registro (%v) no existe en la base de datos", id)
	defer file.Close()
}
