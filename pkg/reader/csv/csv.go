package csv

import (
	"encoding/csv"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func ReadCsvFile(path string) (map[int][]string, error) {
	file, err := os.Open("assets/vacancy.csv")
	if err != nil {
		log.Error("Ошибка при открытии файла:", err)
		return nil, err
	}

	defer file.Close()

	// Создаем новый читатель CSV
	reader := csv.NewReader(file)

	// Инициализируем map для хранения данных
	data := make(map[int][]string)

	// Читаем данные из файла построчно
	id := 1
	for {
		// Читаем очередную строку из файла
		record, err := reader.Read()
		if err != nil {
			// Если достигнут конец файла, прерываем цикл
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Ошибка при чтении строки:", err)
			continue
		}

		if record[0] == "" {
			continue
		}
		data[id] = record
		id++
	}

	return data, nil
}
