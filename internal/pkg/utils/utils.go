package utils

import (
	"os"
)

// InitDir creates a directory if it doesn't exist
func InitDir(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(path, 0o755)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
