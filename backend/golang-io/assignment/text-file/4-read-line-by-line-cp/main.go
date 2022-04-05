package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Hello World")
}

func ScanToArray(arr *[]string, fileName string) error {

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		*arr = append(*arr, scanner.Text())
	}

	return nil
}

func ScanToMap(dataMap map[string]string, fileName string) error {

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		dataMap[data] = data
	}

	return nil
}
