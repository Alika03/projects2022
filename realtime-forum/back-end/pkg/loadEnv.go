package pkg

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// LoadEnv extract data from env file and load in operating storage
func LoadEnv(path string) {
	dataBytes, err := extractData(path)
	if err != nil {
		log.Fatal(err)
	}

	storage, err := splitDataIntoMap(dataBytes)
	if err != nil {
		log.Fatal(err)
	}

	if err = loadDataInOS(storage); err != nil {
		log.Fatal(err)
	}
}

func extractData(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = file.Close()
	}()

	return io.ReadAll(file)
}

func loadDataInOS(storage map[string]string) error {
	for key, val := range storage {
		if err := os.Setenv(key, val); err != nil {
			return err
		}
	}

	return nil
}

func splitDataIntoMap(dataBytes []byte) (map[string]string, error) {
	result := make(map[string]string)

	buff := bytes.NewBuffer(dataBytes)
	scan := bufio.NewScanner(buff)

	strLines := make([]string, 0)
	for scan.Scan() {
		if len(scan.Bytes()) == 0 {
			continue
		}
		strLines = append(strLines, scan.Text())
	}

	for _, strLine := range strLines {
		splData := strings.Split(strLine, "=")
		if len(splData) != 2 {
			return nil, fmt.Errorf("not correct data: %v", strLine)
		}
		result[splData[0]] = splData[1]
	}

	return result, nil
}
