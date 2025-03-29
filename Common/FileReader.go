package Common

import (
	"bufio"
	"log"
	"os"
)

func ReadInput(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Impossible to open file: %s", err)
	}
	defer file.Close()

	var stringList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringList = append(stringList, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		log.Fatalf("Scanner encountered an err: %s", err)
	}

	return stringList
}
