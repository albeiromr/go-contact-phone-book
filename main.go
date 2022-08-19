package main

import (
	"go-contact-phone-book/file-factory"
)


func main() {

	csvFileError := filefactory.UseFileFactory("database", filefactory.CsvFile)
	if csvFileError != nil {
		panic(csvFileError)
	}

	logsFileError := filefactory.UseFileFactory("logs", filefactory.LogFile)
	if logsFileError != nil {
		panic(logsFileError)
	}

}
