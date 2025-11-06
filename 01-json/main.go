package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Error string

func (e Error) Error() string { return string(e) }

const (
	errWrongKey     Error = "wrong key"
	errNoData       Error = "no data"
	errInvalidIndex Error = "invalid index"
)

func readFile(fileName string) ([]byte, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile: %w", err)
	}

	return content, nil
}

func processMap(data map[string]any, keys []string) (any, error) {
	key := keys[0]
	value, ok := data[key]
	if !ok {
		return nil, fmt.Errorf("%w: %s", errWrongKey, key)
	}

	if len(keys) == 1 {
		return value, nil
	}

	switch v := value.(type) {
	case map[string]any:
		data, err := processMap(v, keys[1:])
		return data, err
	case []any:
		data, err := processSlice(v, keys[1:])
		return data, err
	default:
		return nil, fmt.Errorf("%w: %s", errNoData, path(keys))
	}
}

func processSlice(data []any, keys []string) (any, error) {
	key, err := strconv.Atoi(keys[0])
	if err != nil {
		return nil, fmt.Errorf("strconv.Atoi: %w", err)
	}

	if key < 0 || key >= len(data) {
		return nil, fmt.Errorf("%w: %v", errInvalidIndex, key)
	}

	value := data[key]

	if len(keys) == 1 {
		return value, nil
	}

	switch v := value.(type) {
	case map[string]any:
		data, err := processMap(v, keys[1:])
		return data, err
	case []any:
		data, err := processSlice(v, keys[1:])
		return data, err
	default:
		return nil, fmt.Errorf("%w: %s", errNoData, path(keys))
	}
}

func path(keys []string) string {
	return strings.Join(keys, "/")
}

func main() {
	fileName := flag.String("f", "", "Имя JSON файла для обработки (обязательно)")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("Ошибка: необходимо указать имя файла с помощью флага -f")
		flag.Usage()
		os.Exit(1)
	}

	keys := flag.Args()
	if len(keys) == 0 {
		fmt.Println("Ошибка: необходимо указать хотя бы один ключ")
		os.Exit(1)
	}

	content, err := readFile(*fileName)
	if err != nil {
		fmt.Printf("Ошибка чтения файла: %v\n", err)
		os.Exit(1)
	}

	var data map[string]any
	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Printf("Ошибка разбора JSON: %v\n", err)
		os.Exit(1)
	}

	value, err := processMap(data, keys)
	if err != nil {
		fmt.Printf("Ошибка обработки данных: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Данные по пути '%s': %v\n", path(keys), value)
}
