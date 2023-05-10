package builtins

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
)
var(
	ErrArgCount =errors.New("invalid argument count")
)

func ListDirectory(args ...string) error {
	var dir string
	switch len(args) {
	case 0: // list current directory
		dir = "."
	case 1:
		dir = args[0]
	default:
		return fmt.Errorf("%w: expected zero or one arguments (directory)", ErrArgCount)
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		fmt.Println(filepath.Join(dir,file.Name()))
	}
	return nil
}