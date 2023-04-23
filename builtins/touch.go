package builtins

import (
	"fmt"
	"os"
)

func TouchFile(args ...string) error {
	switch len(args) {
	case 1:
		filename := args[0]
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()
		return nil

	default:
		return fmt.Errorf("%w: expected zero or one arguments (directory)", ErrInvalidArgCount)
	}

}
