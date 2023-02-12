package utils

import (
	"bufio"
	"os"
)

func OverrideFile(content, path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	writer.WriteString(content)
	writer.Flush()

	file.Close()
	return nil
}
