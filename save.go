package main

import "os"

func SaveFile(saveFileName string, scannedFileNames []string) error {
	file, err := os.Create(saveFileName)

	if err != nil {
		return ErrorCreateFile
	}
	defer file.Close()

	for _, name := range scannedFileNames {
		_, err := file.WriteString(name + "\n")
		if err != nil {
			return ErrorWriteFile
		}
	}

	return nil
}
