package utils;

import (
	"os"
)

func WriteFile(fileName string, data string) error {
    f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer f.Close()

    _, err = f.WriteString(data)
    return err
}