package filefactory

import (
	"errors"
	"fmt"
	"os"
	"path"
)

const (
	LogFile = ".log"
	CsvFile = ".csv"
)

// fabrica para generar archivos
func UseFileFactory(fileName string, fileExtension string) error {

	filePath, err := generateFilePath(fileName, fileExtension)
	if err != nil {
		panic(err)
	}

	fileCreationError := createFile(filePath)
	if fileCreationError != nil {
		panic(fileCreationError)
	}

	return nil

}

// crea un archivo en el sistema con el parámetro que se le entregue
func createFile(filePath string) error {
	logFile, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer logFile.Close()
	return nil
}

// crea el path para creación de nuevos archivos
func generateFilePath(fileName string, fileExtension string) (filepath string, err error) {
	switch fileExtension {
	case LogFile:

		currentPath, error := os.Getwd()
		if error != nil {
			return "", error
		}
		completeFileName := fmt.Sprintf("/logs/%v%v", fileName, LogFile)
		logFilePAth := path.Join(currentPath, completeFileName)
		return logFilePAth, nil

	case CsvFile:

		currentPath, error := os.Getwd()
		if error != nil {
			return "", error
		}
		completeFileName := fmt.Sprintf("/database/%v%v", fileName, CsvFile)
		CsvFilePath := path.Join(currentPath, completeFileName)
		return CsvFilePath, nil

	default:
		errorMessage := fmt.Sprintf("la extención %v no es valida, debe crear un archivo (.log) para el manejo de logs o un (.csv) para la base de datos de la aplicación", fileExtension)
		defaultError := errors.New(errorMessage)
		return "", defaultError
	}
}
