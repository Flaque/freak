package fs

import (
	"encoding/csv"
	"os"
	"runtime"
)

func HomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func createOrGetFile(name string) (*os.File, error) {
	var file *os.File
	var err error
	if fileExists(name) {
		file, err = os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	} else {
		file, err = os.Create(name)
	}

	if err != nil {
		return nil, err
	}

	return file, nil
}

// WriteToCSV records a row to the file path. It'll create the file if it doesn't exist
func WriteToCSV(data []string, path string) error {

	// Create freak file if not exists, otherwise read
	file, err := createOrGetFile(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write csv line at the end
	err = writer.Write(data)
	if err != nil {
		return err
	}

	return nil
}
