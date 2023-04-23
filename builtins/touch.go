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
	switch len(args) {
	case 1:
		filename := args[1]
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()
		return nil

	default:
		return fmt.Errorf("%w: expected zero or one arguments (directory)", ErrInvalidCount)
	}

}
