package main

import "os"

func createFile(filename string, data []byte) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close() // Отложенный вызов функции закрытия файла

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}
