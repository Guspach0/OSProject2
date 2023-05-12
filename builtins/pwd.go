package builtins // package declaration

import (
	"fmt"
	"os"
  "path/filepath"
)

func Pwd(options []string) error { // prints the  pathname of the current working directory
	var dir string
	var err error

	if containsOption(options, "-P") { // the pathname printed will not contain symbolic links for -p
		dir, err = os.Getwd()
	} else if containsOption(options, "-L") { // pathname printed may contain symbolic links. 
		dir, err = filepath.EvalSymlinks(".")
	} else {
		dir, err = filepath.Abs(".")
	}
  
	dir, err = os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	_, err = fmt.Println(dir)
	return err
}

/*
func containsOption(options []string, option string) bool {
	for _, opt := range options {
		if opt == option {
			return true
		}
	}
	return false
}
*/
