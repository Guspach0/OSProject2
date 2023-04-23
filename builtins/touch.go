package builtins

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrInvalidCount = errors.New("invalid argument count")
)

func TouchFile(args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("%w: Expected one argument (file name)", ErrInvalidCount)
	}

	filename := args[0]

	// Create a new file with the specified name.
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}
