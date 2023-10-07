package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

const (
	TARGET_DIR_ENV_KEY_NAME = "TARGET_DIRECTORY"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	ff, _ := showFiles()
	for _, f := range ff {
		fmt.Println(f.Name())
	}
}

func showFiles() ([]*os.File, error) {
	targetDir := os.Getenv(TARGET_DIR_ENV_KEY_NAME)
	dir, err := os.ReadDir(targetDir)
	if err != nil {
		return nil, err
	}
	var result []*os.File
	for _, f := range dir {
		if f.IsDir() {
			continue
		}
		f, err := os.Open(createFILEPath(targetDir, f.Name()))
		if err != nil {
			return nil, err
		}
		result = append(result, f)
	}
	return result, nil
}

func createFILEPath(filenames ...string) string {
	return strings.Join(filenames, string(os.PathSeparator))
}
