package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("dummyCommit")
}

func AddString(fileName string, stringToAdd string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
		return err
	}

	defer file.Close()
	if _, err = file.WriteString(stringToAdd); err != nil {
		log.Fatalf("failed writing to file: %s", err)
		return err
	}

	return nil
}
