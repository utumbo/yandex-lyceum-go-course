package main

import "os"

func createFile(filename string, data []byte) error {
	f, err := os.Create(filename) // Создание файла с помощью покета os
	if err != nil {
		return err
	}

	_, err = f.Write(data) // Запись данных в файл
	if err != nil {
		f.Close()
		return err
	}

	err = f.Close() // Закрытие файла
	if err != nil {
		return err
	}

	return nil
}

func main(){
	
}
